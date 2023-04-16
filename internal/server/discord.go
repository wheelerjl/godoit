package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wheelerjl/godoit/internal/discord"
)

func (s Server) SendDiscordNotification(ctx *gin.Context) {
	userID, ok := ctx.GetQuery("assigned_user_id")
	if !ok {
		SetErrorResponse(ctx, http.StatusBadRequest, "missing assigned_user_id query parameter", nil)
		return
	}
	activities, err := s.Config.Database.GetActivities(ctx.Request.Context(), userID)
	if err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to get activities", err)
		return
	}
	subjects, err := s.Config.Database.GetSubjects(ctx.Request.Context())
	if err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to get subjects", err)
		return
	}

	var embeds []discord.EmbedData
	for _, activity := range activities {
		// Discord limits the max amount of embeds to 10
		if len(embeds) == 10 {
			continue
		}
		newEmbed := discord.EmbedData{
			Activity: activity,
		}
		for _, subject := range subjects {
			if activity.SubjectID == subject.SubjectID {
				newEmbed.Subject = subject
				continue
			}

		}
		embeds = append(embeds, newEmbed)
	}
	if err := s.Config.DiscordBot.SendComplexNotification(userID, embeds); err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to send notification", err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
