package discord

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/wheelerjl/godoit/internal/database"
	"github.com/wheelerjl/godoit/internal/variables"
)

const embedColorBlue = 100

type EmbedData struct {
	Subject  database.Subject
	Activity database.Activity
}

type BotClient struct {
	Session *discordgo.Session
}

func NewDiscordBotClient(variables variables.Variables) (client BotClient, err error) {
	client.Session, err = discordgo.New(fmt.Sprintf("Bot %s", variables.DiscordToken))
	if err != nil {
		return client, err
	}
	return client, nil
}

func (b BotClient) SendNotification(userID, msg string) error {
	ch, err := b.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}

	if err := b.Session.ChannelMessageDelete(ch.ID, ch.LastMessageID); err != nil {
		return err
	}

	if _, err := b.Session.ChannelMessageSend(ch.ID, msg); err != nil {
		return err
	}
	return nil
}

func (b BotClient) SendComplexNotification(userID string, data []EmbedData) error {
	ch, err := b.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}

	b.Session.ChannelMessageDelete(ch.ID, ch.LastMessageID)

	var embeds []*discordgo.MessageEmbed
	for _, embed := range data {
		newEmbed := discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       embed.Activity.Name,
			Description: embed.Activity.Description,
			Timestamp:   embed.Activity.StartTime.Format(time.RFC3339),
			Color:       embedColorBlue,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: embed.Subject.ImageURL,
			},
		}
		embeds = append(embeds, &newEmbed)
	}
	message := &discordgo.MessageSend{
		Content: "Activities",
		Embeds:  embeds,
	}
	if _, err := b.Session.ChannelMessageSendComplex(ch.ID, message); err != nil {
		return err
	}
	return nil
}
