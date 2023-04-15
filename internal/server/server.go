package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wheelerjl/godoit/internal/database"
	"github.com/wheelerjl/godoit/internal/variables"
)

type Config struct {
	Variables variables.Variables
	Database  database.Client
}

type Server struct {
	Router *gin.Engine
	Config Config
}

func NewServer(config Config) Server {
	s := Server{
		Config: config,
	}
	if s.Config.Variables.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.Router = gin.Default()
	s.SetupRoutes()
	return s
}

func (s Server) Start() error {
	return s.Router.Run(fmt.Sprintf(":%d", s.Config.Variables.Port))
}

func (s Server) SetupRoutes() {
	// Change default handling for if the route isn't found to present a json message
	s.Router.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			gin.H{"status": "error", "message": "route not found"},
		)
	})

	s.Router.GET("/readiness", func(ctx *gin.Context) {
		code := http.StatusOK
		message := gin.H{"status": "ok"}
		if err := s.Config.Database.DB.Ping(); err != nil {
			code = http.StatusServiceUnavailable
			message = gin.H{"status": "error", "message": fmt.Sprintf("unable to verify connection to database: %v", err)}
		}
		ctx.JSON(
			code,
			message,
		)
	})

	s.Router.GET("/liveness", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{"status": "ok"},
		)
	})

	// Create a config route to print environment variables if applicable
	if s.Config.Variables.Debug && s.Config.Variables.Env == "dev" {
		s.Router.GET("/config", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, s.Config.Variables)
		})
	}
}
