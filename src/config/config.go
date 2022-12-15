package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type App struct {
	Port        int  `mapstructure:"port"`
	Debug       bool `mapstructure:"debug"`
	CacheTTL    int  `mapstructure:"cache_ttl"`
	MaxFileSize int  `mapstructure:"max_file_size"`
}

type AppConfig struct {
	App App `mapstructure:"app"`
}

func LoadAppConfig() (config *AppConfig, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := AppConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf, err
}
