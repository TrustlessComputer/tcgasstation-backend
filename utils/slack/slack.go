package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

type Slack struct {
	client *slack.Client
	token  string
	env    string
}

func NewSlack(cnf Config) *Slack {
	client := slack.New(cnf.Token)

	return &Slack{
		token:  cnf.Token,
		client: client,
		env:    cnf.Env,
	}
}

func (s Slack) postMessage(channelId string, messages ...slack.MsgOption) (string, string, error) {
	return s.client.PostMessage(channelId, messages...)
}

func (s Slack) SendMessageToSlackWithChannel(channelID string, pretext string, title string, text string) (string, string, error) {
	attachment := slack.Attachment{
		Color:   "#1766ff",
		Pretext: fmt.Sprintf("[%s] %s", s.env, pretext),
		Title:   title,
		Text:    text,
	}

	return s.postMessage(channelID, slack.MsgOptionAttachments(attachment))
}
