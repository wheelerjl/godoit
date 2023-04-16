package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/wheelerjl/godoit/internal/variables"
)

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

func (b BotClient) SendComplexNotification(userID, msg string) error {
	ch, err := b.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}

	b.Session.ChannelMessageDelete(ch.ID, ch.LastMessageID)

	message := discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Image of Gopher",
			Description: "Todo",
			Color:       100,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://miro.medium.com/v2/resize:fit:1000/0*YISbBYJg5hkJGcQd.png",
			},
		},
	}
	if _, err := b.Session.ChannelMessageSendComplex(ch.ID, &message); err != nil {
		return err
	}
	return nil
}
