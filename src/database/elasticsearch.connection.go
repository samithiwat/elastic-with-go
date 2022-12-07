package database

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/samithiwat/elastic-with-go/src/config"
	"os"
)

func InitElasticDefaultClient(isDebug bool) (*elasticsearch.Client, error) {
	conf, err := config.LoadElasticsearchConfig()
	if err != nil {
		return nil, err
	}

	esConf := elasticsearch.Config{
		Addresses: []string{conf.Host},
		Username:  conf.Username,
		Password:  conf.Password,
	}

	if isDebug {
		esConf.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	client, err := elasticsearch.NewClient(esConf)

	if _, err := client.Info(); err != nil {
		return nil, err
	}

	return client, nil
}

func InitElasticTypedClient(isDebug bool) (*elasticsearch.TypedClient, error) {
	conf, err := config.LoadElasticsearchConfig()
	if err != nil {
		return nil, err
	}

	esConf := elasticsearch.Config{
		Addresses: []string{conf.Host},
		Username:  conf.Username,
		Password:  conf.Password,
	}

	if isDebug {
		esConf.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	return elasticsearch.NewTypedClient(esConf)
}
