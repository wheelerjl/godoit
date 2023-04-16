package server

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/wheelerjl/godoit/internal/database"
)

type MessageData struct {
	Subject   database.Subject
	Activites []database.Activity
}

const embedColorBlue = 100

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

	message := buildMessage(userID, activities, subjects)
	if err := s.Config.DiscordBot.SendComplexNotification(userID, message); err != nil {
		SetErrorResponse(ctx, http.StatusInternalServerError, "unable to send notification", err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func buildMessage(userID string, activities []database.Activity, subjects []database.Subject) discordgo.MessageSend {
	index := 0
	data := []MessageData{}
	indexMap := make(map[string]int)

	for _, activity := range activities {
		for _, subject := range subjects {
			if subject.SubjectID == activity.SubjectID {
				i, ok := indexMap[subject.SubjectID]
				if !ok {
					newData := MessageData{
						Subject:   subject,
						Activites: []database.Activity{activity},
					}
					data = append(data, newData)
					indexMap[subject.SubjectID] = index
					index++
				} else {
					data[i].Activites = append(data[i].Activites, activity)
				}

				continue
			}
		}
	}

	var embeds []*discordgo.MessageEmbed
	for _, value := range data {
		var fields []*discordgo.MessageEmbedField
		for _, activity := range value.Activites {
			whatField := discordgo.MessageEmbedField{
				Name:   "What",
				Value:  activity.Name,
				Inline: true,
			}
			whenField := discordgo.MessageEmbedField{
				Name:   "When",
				Value:  fmt.Sprintf("<t:%d:d>", activity.StartTime.Unix()),
				Inline: true,
			}
			whereField := discordgo.MessageEmbedField{
				Name:   "Where",
				Value:  activity.Location,
				Inline: true,
			}
			howField := discordgo.MessageEmbedField{
				Name:   "How",
				Value:  activity.Description,
				Inline: false,
			}
			fields = append(fields, &whatField)
			fields = append(fields, &whenField)
			fields = append(fields, &whereField)
			fields = append(fields, &howField)
		}
		newEmbed := discordgo.MessageEmbed{
			Type:  discordgo.EmbedTypeRich,
			Title: fmt.Sprintf("Who: %s", value.Subject.Name),
			Color: embedColorBlue,
			Image: &discordgo.MessageEmbedImage{
				URL: value.Subject.ImageURL,
			},
			Fields: fields,
		}
		embeds = append(embeds, &newEmbed)
	}

	return discordgo.MessageSend{
		Content: fmt.Sprintf("Hey <@%s>, you've got some work to do!", userID),
		Embeds:  embeds,
	}
}
