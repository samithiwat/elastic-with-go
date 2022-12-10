package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type RabbitMQ struct {
	Host  string `mapstructure:"host"`
	VHost string `mapstructure:"vhost"`
}

type RabbitMQConfig struct {
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq"`
}

func LoadRabbitMQConfig() (config *RabbitMQ, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("rabbitmq")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := RabbitMQConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf.RabbitMQ, nil
}
