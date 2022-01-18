package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/IDarar/grpcservices/s1/config"
	"github.com/IDarar/grpcservices/s1/s2client"
	"github.com/IDarar/grpcservices/s1/services"
	"github.com/IDarar/grpcservices/s1/transport/rest"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Init("config.json")
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()

	s2, err := s2client.NewS2Client(cfg)
	if err != nil {
		log.Fatal(err)
	}

	services := services.NewSearchService(sugar, s2)

	h := rest.NewHandler(services)

	srv := rest.NewServer(h.Init(), cfg.HTTP.Port)

	//run server in goroutine to not block execution
	//and be able to listen system signals
	go func() {
		fmt.Println("Starting server ...")
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	//register channel on signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	//shutdown server and return
	if err := srv.Stop(context.Background()); err != nil {
		return
	}

	sugar.Info("shutted  server ... ")
}
