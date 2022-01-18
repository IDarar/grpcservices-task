package services

import (
	"context"
	"testing"

	"github.com/IDarar/grpcservices/s1/domain"
	"github.com/IDarar/grpcservices/s1/services/mock_services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name  string
		input domain.SearchResp
		want  domain.SearchResp
	}{
		{
			name: "sorted input",
			input: domain.SearchResp{[]domain.Item{
				{Title: "A"},
				{Title: "B"},
				{Title: "C"},
			}},
			want: domain.SearchResp{[]domain.Item{
				{Title: "A"},
				{Title: "B"},
				{Title: "C"},
			}},
		},
		{
			name: "unsorted input",
			input: domain.SearchResp{[]domain.Item{
				{Title: "B"},
				{Title: "D"},
				{Title: "A"},
				{Title: "C"},
			}},
			want: domain.SearchResp{[]domain.Item{
				{Title: "A"},
				{Title: "B"},
				{Title: "C"},
				{Title: "D"},
			}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)

			defer c.Finish()
			s2 := mock_services.NewMockS2(c)

			s2.EXPECT().Search(context.TODO(), "123").Return(tc.input, nil)

			services := NewSearchService(zap.L().Sugar(), s2)

			res, err := services.Search(context.Background(), "123")
			require.NoError(t, err)

			require.Equal(t, res, tc.want)

		})
	}

}
