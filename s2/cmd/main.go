package main

import (
	"log"

	"github.com/IDarar/grpcservices/s2/config"
	"github.com/IDarar/grpcservices/s2/pkg/googlesearch"
	"github.com/IDarar/grpcservices/s2/services"
	"go.uber.org/zap"

	transport "github.com/IDarar/grpcservices/s2/transport/grpcapi"
)

func main() {
	cfg, err := config.Init("config.json")
	if err != nil {
		log.Fatal(err)
	}

	s, err := googlesearch.NewGoogleSearchAPICLient(cfg.GoogleSearch.ApiKey, cfg.GoogleSearch.Cx)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()

	services := services.NewSearchServie(s, sugar)

	err = transport.RunServer(cfg.GRPC, services)
	if err != nil {
		log.Fatal(err)
	}
}
