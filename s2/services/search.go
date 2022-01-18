package services

import (
	"context"
	"fmt"

	"github.com/IDarar/grpcservices/s2/domain"
	"go.uber.org/zap"
)

type SearchApi interface {
	SearchReq(ctx context.Context, keyword string) (domain.SearchResult, error)
}

type SearchServie struct {
	SearchApi SearchApi
	l         *zap.SugaredLogger
}

func NewSearchServie(s SearchApi, l *zap.SugaredLogger) *SearchServie {
	return &SearchServie{l: l, SearchApi: s}
}

func (s *SearchServie) Search(ctx context.Context, keyword string) (domain.SearchResult, error) {
	s.l.Info("reqeust for: ", keyword)
	res, err := s.SearchApi.SearchReq(ctx, keyword)
	if err != nil {
		s.l.Error("failed to call search api: ", err)
		return res, fmt.Errorf("failed to call search api: %w", err)
	}

	return res, nil
}
