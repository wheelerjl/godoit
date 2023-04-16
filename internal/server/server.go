package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wheelerjl/godoit/internal/database"
	"github.com/wheelerjl/godoit/internal/discord"
	"github.com/wheelerjl/godoit/internal/variables"
)

type DefaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Config struct {
	Variables  variables.Variables
	Database   database.Client
	DiscordBot discord.BotClient
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
	// Create a config route to print environment variables if applicable
	if s.Config.Variables.Debug && s.Config.Variables.Env == "dev" {
		s.Router.GET("/config", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, s.Config.Variables)
		})
	}

	s.Router.NoRoute(s.RouteNotFound)
	s.Router.GET("/readiness", s.Readiness)
	s.Router.GET("/liveness", s.Liveness)
	s.Router.POST("/discord", s.SendDiscordNotification)

	subjectGroup := s.Router.Group("/subjects")
	subjectGroup.POST("", s.AddSubject)
	subjectGroup.GET("", s.GetSubjects)
	subjectGroup.GET("/:id", s.GetSubject)
	subjectGroup.DELETE("/:id", s.RemoveSubject)

	activityGroup := s.Router.Group("/activities")
	activityGroup.POST("", s.AddActivity)
	activityGroup.GET("", s.GetActivities)
	activityGroup.GET("/:id", s.GetActivity)
	activityGroup.DELETE("/:id", s.RemoveActivity)
}
