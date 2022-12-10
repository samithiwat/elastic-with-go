package main

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/_example_apps/constant"
	"github.com/samithiwat/elastic-with-go/_example_apps/domain/entity/chula-course/course"
	"github.com/samithiwat/elastic-with-go/_example_apps/utils"
	"os"
	"sync"
)

var (
	// SubscribedTopics topic name in convention as <application>.<service>.<method>.<status>
	SubscribedTopics = []string{"cugetreg.scraper.scrape.pending", "cugetreg.backend.create.pending"}
)

func main() {
	client := InitElasticsearchClient()

	go CourseEventListener(client)

	var wait chan struct{}

	<-wait
}

func InitElasticsearchClient() *elasticsearch.Client {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "admin",
		Logger: &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	})

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error while init elastic client")
	}

	{
		if _, err := client.Info(); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while connect to elastic client")
		}
	}

	return client
}

func InsertData(client *elasticsearch.Client, courseList *[]*course.Course) {
	var courseDocList []*course.CourseDoc
	buf := bytes.Buffer{}
	currentBatch := 0
	nCourseList := len(*courseList)

	var wg sync.WaitGroup

	for pos, c := range *courseList {
		courseDoc := &course.CourseDoc{
			AbbrName:     c.AbbrName,
			CourseNo:     c.CourseNo,
			CourseNameTh: c.CourseNameTh,
			CourseNameEn: c.CourseNameEn,
			CourseDescTh: c.CourseDescTh,
			CourseDescEn: c.CourseDescEn,
			RawData:      c,
		}
		courseDocList = append(courseDocList, courseDoc)

		if err := utils.AppendDocToBuffer(c.ID.OID, courseDoc, &buf); err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while create request body")
		}

		currentBatch = pos / constant.DocPerBatch
		if pos == nCourseList-1 {
			currentBatch++
		}

		if pos > 0 && pos%constant.DocPerBatch == 0 || pos == nCourseList-1 {
			go func() {
				wg.Add(1)
				defer wg.Done()

				res, err := client.Bulk(bytes.NewReader(buf.Bytes()), client.Bulk.WithIndex(constant.IndexName))
				if err != nil {
					log.Fatal().
						Err(err).
						Msg("Error while create data to elasticsearch database")
				}

				if res.IsError() {
					raw := map[string]interface{}{}

					if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
						log.Fatal().Msgf("Failure to to parse response body: %s", err)
					}

					log.Printf("  Error: [%d] %s: %s",
						res.StatusCode,
						raw["error"].(map[string]interface{})["type"],
						raw["error"].(map[string]interface{})["reason"],
					)

				}

				resMap := map[string]interface{}{}

				if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
					log.Fatal().
						Err(err).
						Msgf("Failure to to parse response body: %s", err)
				}

				for _, item := range resMap["items"].([]interface{}) {
					status := item.(map[string]interface{})["index"].(map[string]interface{})["status"].(float64)

					if status > 201 {
						resErr := item.(map[string]interface{})["index"].(map[string]interface{})["error"].(map[string]interface{})

						log.Printf("  Error: [%.0f]: %s: %s",
							status,
							resErr["type"],
							resErr["reason"],
						)
					}
				}

				buf.Reset()
			}()
		}
	}

	wg.Wait()
}

func CourseEventListener(client *elasticsearch.Client) {
	conn, err := amqp.DialConfig("amqp://guest:guest@localhost:5672/", amqp.Config{
		Vhost: "newbie",
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal().
			Msg("Failed to open a channel")
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(
		constant.ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create exchange")
	}

	queue, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create queue")
	}

	for _, topic := range SubscribedTopics {
		log.Info().
			Str("queue_name", queue.Name).
			Str("exchange", constant.ExchangeName).
			Str("topic", topic).
			Msg("Binding queue")

		if err := ch.QueueBind(
			queue.Name,
			topic,
			constant.ExchangeName,
			false,
			nil,
		); err != nil {
			log.Fatal().
				Err(err).
				Str("queue_name", queue.Name).
				Str("exchange", constant.ExchangeName).
				Str("topic", topic).
				Msg("Failed to binding queue")
		}
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Fatal().
			Msg("Failed to registered consumer")
	}

	go func() {
		for d := range msgs {
			log.Info().
				Str("topic", d.RoutingKey).
				Msg("Received message")

			var courseList []*course.Course

			if err := json.Unmarshal(d.Body, &courseList); err != nil {
				log.Fatal().
					Err(err).
					Msg("Error while parsing course to struct")
			}

			InsertData(client, &courseList)

			log.Info().
				Str("topic", d.RoutingKey).
				Msg("Successfully insert data to elasticsearch database")
		}
	}()

	log.Info().
		Msg("Ready to accept message from RabbitMQ")

	var forever chan struct{}

	<-forever
}
