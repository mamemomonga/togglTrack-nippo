package slackutil

import (
	"log"

	"github.com/slack-go/slack"
)

type SlackUtil struct {
	api *slack.Client
}

func New(token string) *SlackUtil {
	t := &SlackUtil{}
	t.api = slack.New(token, slack.OptionDebug(false))
	return t
}

func (t *SlackUtil) PostSimple(msg string, channel string) {
	attachment := slack.Attachment{
		Text: msg,
	}
	channelID, timestamp, err := t.api.PostMessage(
		channel,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	log.Printf("debug: Message successfully sent to channel %s at %s", channelID, timestamp)
}
