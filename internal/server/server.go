package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return &Server{router: r}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
