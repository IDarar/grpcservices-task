package rest

import (
	"context"

	"github.com/IDarar/grpcservices/s1/domain"
	"github.com/gin-gonic/gin"
)

type SearchService interface {
	Search(ctx context.Context, keyword string) (domain.SearchResp, error)
}

func (h *Handler) search(c *gin.Context) {
	keyword := c.Query("keyword")

	resp, err := h.SearchService.Search(c.Request.Context(), keyword)
	if err != nil {
		newResponse(c, 500, err)
		return
	}

	c.JSON(200, resp)
}
