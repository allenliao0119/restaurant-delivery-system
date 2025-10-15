package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xorm.io/xorm"
)

type Server struct {
	router *gin.Engine
	engine *xorm.Engine
}

func New(engine *xorm.Engine) *Server {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return &Server{
		router: r,
		engine: engine,
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
