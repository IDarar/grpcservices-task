package transport

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IDarar/grpcservices/s2/domain"
	p "github.com/IDarar/grpcservices/search_services"

	"github.com/IDarar/grpcservices/s2/config"
	"google.golang.org/grpc"
)

type SearchService interface {
	Search(ctx context.Context, keyword string) (domain.SearchResult, error)
}

type SearchServer struct {
	Service SearchService
}

func RunServer(cfg config.GRPCConfig, services SearchService) error {
	s := &SearchServer{Service: services}

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return fmt.Errorf("listen failed: %w", err)
	}

	grpcSrv := grpc.NewServer()

	p.RegisterSearchServer(grpcSrv, s)

	sigCh := make(chan os.Signal, 1)

	//get signal from os to gracefully shutdown server
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		<-sigCh
		grpcSrv.GracefulStop()
		wg.Done()
	}()

	fmt.Println("Starting server ...")
	if err := grpcSrv.Serve(lis); err != nil {
		return fmt.Errorf("failed serving: %w", err)
	}

	wg.Wait()

	return nil
}

func (s *SearchServer) Search(ctx context.Context, req *p.SearchReq) (*p.SearchResp, error) {
	res, err := s.Service.Search(ctx, req.Keyword)
	if err != nil {
		return nil, err
	}

	//convert internal struct to protobuf
	resp := &p.SearchResp{}
	for i := range res.Items {
		resp.Items = append(resp.Items, &p.Item{Link: res.Items[i].Link, Title: res.Items[i].Title})
	}

	return resp, nil
}
