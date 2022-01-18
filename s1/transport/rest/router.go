package rest

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	SearchService SearchService
}

func NewHandler(searchService SearchService) *Handler {
	return &Handler{
		SearchService: searchService,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	r := router.Group("/")
	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()
		fmt.Println("request took ", time.Since(s))

	})
	r.GET("/search", h.search)

}

func newResponse(c *gin.Context, statusCode int, err error) {
	c.AbortWithError(statusCode, err)
}
