package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	Status string `json:"status"`
}

func (s Server) Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthCheck{Status: "ok"})
}

func (s Server) Readiness(ctx *gin.Context) {
	if err := s.Config.Database.DB.Ping(ctx.Request.Context()); err != nil {
		SetErrorResponse(ctx, http.StatusServiceUnavailable, "unable to verify connection to database", err)
		return
	}
	ctx.JSON(http.StatusOK, HealthCheck{Status: "ok"})
}
