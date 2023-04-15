package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wheelerjl/godoit/internal/config"
)

type Server struct {
	config config.Config
	router *gin.Engine
}

func NewServer(conf config.Config) Server {
	s := Server{
		config: conf,
	}
	if s.config.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.router = gin.Default()
	s.router.GET("/readiness", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	s.router.GET("/liveness", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Create a config route to print environment variables if applicable
	if s.config.Debug && s.config.Env == "dev" {
		s.router.GET("/config", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, s.config)
		})
	}
	return s
}

func (s Server) Start() error {
	return s.router.Run(fmt.Sprintf(":%d", s.config.Port))
}
