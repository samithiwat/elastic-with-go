package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/src/config"
)

func InitRedisConnect() (cache *redis.Client, err error) {
	conf, err := config.LoadRedisConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "auth").
			Msg("Failed to start service")
	}

	cache = redis.NewClient(&redis.Options{
		Addr: conf.Host,
		DB:   conf.DB,
	})

	return
}
