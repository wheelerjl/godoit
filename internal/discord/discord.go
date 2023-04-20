package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/wheelerjl/godoit/internal/variables"
)

const embedColorBlue = 100

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

func (b BotClient) SendComplexNotification(userID string, msg discordgo.MessageSend) error {
	ch, err := b.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}

	b.Session.ChannelMessageDelete(ch.ID, ch.LastMessageID)
	if _, err := b.Session.ChannelMessageSendComplex(ch.ID, &msg); err != nil {
		return err
	}
	return nil
}
