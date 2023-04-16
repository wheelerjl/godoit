package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) SendDiscordNotification(ctx *gin.Context) {
	userID, ok := ctx.GetQuery("userId")
	if !ok {
		ErrorResponse(ctx, http.StatusBadRequest, "missing userID query parameter", nil)
		return
	}
	subjects, err := s.Config.Database.GetSubjects(ctx.Request.Context())
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "unable to get subjects", err)
		return
	}
	if err := s.Config.DiscordBot.SendNotification(userID, fmt.Sprintf("subject count: %d", len(subjects))); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "unable to send notification", err)
	}
	ctx.Status(http.StatusNoContent)
}