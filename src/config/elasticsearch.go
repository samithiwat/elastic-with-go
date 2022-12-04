package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Elasticsearch struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type ElasticsearchConfig struct {
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch"`
}

func LoadElasticsearchConfig() (config *Elasticsearch, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("elasticsearch")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := ElasticsearchConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf.Elasticsearch, err
}
