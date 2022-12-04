package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Redis struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RedisConfig struct {
	Redis Redis `mapstructure:"redis"`
}

func LoadRedisConfig() (config *Redis, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("redis")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := RedisConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf.Redis, nil
}
