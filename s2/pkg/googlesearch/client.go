package googlesearch

import (
	"context"
	"fmt"

	"github.com/IDarar/grpcservices/s2/domain"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type GoogleSearchAPICLient struct {
	s  *customsearch.Service
	cx string
}

func NewGoogleSearchAPICLient(apiKey string, cx string) (*GoogleSearchAPICLient, error) {
	service, err := customsearch.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to init google api service")
	}

	return &GoogleSearchAPICLient{s: service, cx: cx}, nil
}

//call google search api and returns titles and links of the first 10 results
func (c *GoogleSearchAPICLient) SearchReq(ctx context.Context, keyword string) (domain.SearchResult, error) {
	l := c.s.Cse.List()
	l.Cx(c.cx)

	l.Q(keyword)

	l.Context(ctx)

	res, err := l.Fields(googleapi.Field("items(title,link)")).Do()
	if err != nil {
		return domain.SearchResult{}, fmt.Errorf("failed to call google api: %w", err)
	}

	return convertToDomain(res), nil
}

func convertToDomain(res *customsearch.Search) domain.SearchResult {
	sResp := domain.SearchResult{}

	for i := range res.Items {
		sResp.Items = append(sResp.Items, domain.Item{Title: res.Items[i].Title, Link: res.Items[i].Link})
	}

	return sResp
}
