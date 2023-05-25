package global

import (
	"tcgasstation-backend/external/quicknode"
	"tcgasstation-backend/utils/config"
	_pConnection "tcgasstation-backend/utils/connections"
	discordclient "tcgasstation-backend/utils/discord"
	"tcgasstation-backend/utils/eth"
	"tcgasstation-backend/utils/googlecloud"
	"tcgasstation-backend/utils/oauth2service"
	"tcgasstation-backend/utils/redis"
	"tcgasstation-backend/utils/slack"

	"github.com/gorilla/mux"
)

type Global struct {
	Conf             *config.Config
	MuxRouter        *mux.Router
	DBConnection     _pConnection.IConnection
	GCS              googlecloud.IGcstorage
	S3Adapter        googlecloud.S3Adapter
	Cache            redis.IRedisCache
	CacheAuthService redis.IRedisCache
	QuickNode        *quicknode.QuickNode
	Auth2            *oauth2service.Auth2
	DiscordClient    *discordclient.Client

	Slack slack.Slack

	TcClient, EthClient *eth.Client
}
