package s2client

import (
	"context"
	"fmt"

	"github.com/IDarar/grpcservices/s1/config"
	"github.com/IDarar/grpcservices/s1/domain"
	"github.com/IDarar/grpcservices/search_services"
	"google.golang.org/grpc"
)

type S2Client struct {
	Client search_services.SearchClient
}

func NewS2Client(cfg config.Config) (*S2Client, error) {
	conn, err := grpc.Dial(cfg.S2.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to dial s2 grpc server: %w", err)
	}

	c := search_services.NewSearchClient(conn)

	return &S2Client{
		Client: c,
	}, nil
}

func (s *S2Client) Search(ctx context.Context, keyword string) (domain.SearchResp, error) {
	res, err := s.Client.Search(ctx, &search_services.SearchReq{Keyword: keyword})
	if err != nil {
		return domain.SearchResp{}, fmt.Errorf("failed to call s2 search method: %w", err)
	}

	return convertToDomain(res), nil
}

func convertToDomain(res *search_services.SearchResp) domain.SearchResp {
	sResp := domain.SearchResp{}

	for i := range res.Items {
		sResp.Items = append(sResp.Items, domain.Item{Title: res.Items[i].Title, Link: res.Items[i].Link})
	}

	return sResp
}
