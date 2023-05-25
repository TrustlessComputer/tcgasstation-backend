package slack

import "github.com/slack-go/slack"

type Config struct {
	Token         string `yaml:"token" json:"token"`
	ChannelLogs   string `yaml:"channelLogs" json:"channel_logs"`
	ChannelOrders string `yaml:"channelOrders" json:"channel_orders"`
	Env           string `yaml:"env" json:"env"`
}

type SlackData struct {
	ChannelName string           `json:"channel_name"`
	Data        slack.Attachment `json:"data"`
}
