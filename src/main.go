package main

import (
	"context"
	"fmt"
	"github.com/samithiwat/elastic-with-go/src/pb"
	searchRepo "github.com/samithiwat/elastic-with-go/src/repository/search"
	courseSearchRepo "github.com/samithiwat/elastic-with-go/src/repository/search/course"
	courseSrv "github.com/samithiwat/elastic-with-go/src/service/search/course"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/samithiwat/elastic-with-go/src/config"
	"github.com/samithiwat/elastic-with-go/src/database"
	cacheRepo "github.com/samithiwat/elastic-with-go/src/repository/cache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		sig := <-s

		log.Info().
			Str("service", "graceful shutdown").
			Msgf("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Error().
				Str("service", "graceful shutdown").
				Msgf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Error().
						Str("service", "graceful shutdown").
						Err(err).
						Msgf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("%v was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}

func main() {
	conf, err := config.LoadAppConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "search").
			Msg("Failed to start service")
	}

	esClient, err := database.InitElasticClient(conf.App.Debug)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "search").
			Msg("Failed to init elasticsearch client")
	}

	redisClient, err := database.InitRedisConnect()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "search").
			Msg("Failed to init redis client")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "search").
			Msg("Failed to start service")
	}
	defer lis.Close()

	cacheRepository := cacheRepo.NewRepository(redisClient)

	searchRepository := searchRepo.NewRepository(esClient)

	courseSearchRepository := courseSearchRepo.NewRepository(searchRepository)

	courseService := courseSrv.NewService(courseSearchRepository, cacheRepository, conf.App.CacheTTL)

	grpcServer := grpc.NewServer()

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	pb.RegisterSearchServiceServer(grpcServer, courseService)

	reflection.Register(grpcServer)
	go func() {
		log.Info().
			Str("service", "search").
			Msgf("Newbie search starting at port %v", conf.App.Port)

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal().
				Err(err).
				Str("service", "search").
				Msg("Failed to start service")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"cache": func(ctx context.Context) error {
			return redisClient.Close()
		},
		"server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})

	<-wait

	grpcServer.GracefulStop()

	log.Info().
		Str("service", "search").
		Msg("End of Program")
}
