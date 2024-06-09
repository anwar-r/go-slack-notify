package main

import (
	"github.com/slack-go/slack"
)

func SendSlackNotification(apiToken string, channelID string, username string, message string, iconEmojiType string, emojiURL string) error {
	// Initialize a new Slack client with the provided API token
	api := slack.New(apiToken)

	// Set up the message parameters
	params := slack.PostMessageParameters{
		IconEmoji: iconEmojiType,
		Username:  username,
	}

	// Attachments for the message
	attachment := slack.Attachment{
		Text:     message,
		ThumbURL: emojiURL,
	}

	// Create an options array with attachments
	options := []slack.MsgOption{
		slack.MsgOptionPostMessageParameters(params),
		slack.MsgOptionAttachments(attachment),
	}

	// Post the message to the specified Slack channel
	_, _, err := api.PostMessage(channelID, options...)
	if err != nil {
		return err
	}
	return nil
}
