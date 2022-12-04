package database

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/samithiwat/elastic-with-go/src/config"
	"os"
)

func InitElasticClient(isDebug bool) (*elasticsearch.TypedClient, error) {
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
