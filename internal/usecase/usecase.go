package usecase

import (
	"tcgasstation-backend/external/quicknode"
	"tcgasstation-backend/internal/repository"
	"tcgasstation-backend/utils/config"
	discordclient "tcgasstation-backend/utils/discord"
	"tcgasstation-backend/utils/eth"
	"tcgasstation-backend/utils/global"
	"tcgasstation-backend/utils/googlecloud"
	"tcgasstation-backend/utils/oauth2service"
	"tcgasstation-backend/utils/redis"
	"tcgasstation-backend/utils/slack"
)

type Usecase struct {
	Repo          repository.Repository
	Config        *config.Config
	QuickNode     *quicknode.QuickNode
	Cache         redis.IRedisCache
	Auth2         *oauth2service.Auth2
	Storage       googlecloud.IGcstorage
	DiscordClient *discordclient.Client

	TcClient, EthClient *eth.Client

	Slack slack.Slack
}

func NewUsecase(global *global.Global, r repository.Repository) (*Usecase, error) {
	u := new(Usecase)
	u.Repo = r
	u.QuickNode = global.QuickNode
	u.Cache = global.Cache
	u.Storage = global.GCS
	u.Auth2 = global.Auth2
	u.DiscordClient = global.DiscordClient

	u.Slack = global.Slack

	u.TcClient = global.TcClient
	u.EthClient = global.EthClient

	u.Config = global.Conf

	return u, nil
}

func (u *Usecase) Version() string {
	return "bridges-API Server - version 1"
}
