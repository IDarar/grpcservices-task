package services

import (
	"context"
	"fmt"
	"sort"

	"github.com/IDarar/grpcservices/s1/domain"
	"go.uber.org/zap"
)

//Methods represent functions of s2. Interface is needed to wrap grpc client
type S2 interface {
	Search(ctx context.Context, keyword string) (domain.SearchResp, error)
}

type SearchService struct {
	s2 S2
	l  *zap.SugaredLogger
}

func NewSearchService(l *zap.SugaredLogger, s2 S2) *SearchService {
	return &SearchService{s2: s2, l: l}
}

func (s *SearchService) Search(ctx context.Context, keyword string) (domain.SearchResp, error) {
	res, err := s.s2.Search(ctx, keyword)
	if err != nil {
		s.l.Error("failed to call s2 search: ", err)
		return res, fmt.Errorf("failed to call s2 search: %w", err)
	}

	//sort result before return
	sort.Slice(res.Items, func(i, j int) bool {
		return res.Items[i].Title < res.Items[j].Title
	})

	return res, nil
}
