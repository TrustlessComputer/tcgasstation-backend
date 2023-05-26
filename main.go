package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	discordclient "tcgasstation-backend/utils/discord"
	"tcgasstation-backend/utils/eth"
	"tcgasstation-backend/utils/slack"
	"time"

	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"tcgasstation-backend/external/quicknode"
	"tcgasstation-backend/internal/delivery"
	"tcgasstation-backend/internal/delivery/crontabManager"
	httpHandler "tcgasstation-backend/internal/delivery/http"
	"tcgasstation-backend/internal/repository"
	"tcgasstation-backend/internal/usecase"
	"tcgasstation-backend/utils/config"
	"tcgasstation-backend/utils/connections"
	"tcgasstation-backend/utils/global"
	"tcgasstation-backend/utils/googlecloud"
	_logger "tcgasstation-backend/utils/logger"
	"tcgasstation-backend/utils/oauth2service"
	"tcgasstation-backend/utils/redis"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	migrate "github.com/xakep666/mongo-migrate"
)

var logger _logger.Ilogger
var mongoConnection connections.IConnection
var conf *config.Config

func init() {
	logger = _logger.NewLogger(true)

	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	mongoCnn := fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority", c.Databases.Mongo.Scheme, c.Databases.Mongo.User, c.Databases.Mongo.Pass, c.Databases.Mongo.Host)
	mongoDbConnection, err := connections.NewMongo(mongoCnn)
	if err != nil {
		logger.AtLog().Logger.Error("Cannot connect mongoDB ", zap.Error(err))
		panic(err)
	}

	conf = c
	mongoConnection = mongoDbConnection
}

//	@title			tcDAPP APIs
//	@version		1.0.0
//	@description	This is a sample server tcgasstation-backend server.

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/tcgasstation-backend/v1
func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	// log.Println("init sentry ...")
	// sentry.InitSentry(conf)
	startServer()
}

func startServer() {
	logger.AtLog().Logger.Info("starting server ...")
	cache, _ := redis.NewRedisCache(conf.Redis)
	r := mux.NewRouter()
	gcs, err := googlecloud.NewDataGCStorage(*conf)

	slack := slack.NewSlack(conf.Slack)

	qn := quicknode.NewQuickNode(conf, cache)
	dcl := discordclient.NewClient()

	// init tc client
	tcClientWrap, err := ethclient.Dial(conf.BlockchainConfig.TCEndpoint)
	if err != nil {
		_logger.AtLog.Logger.Error("error initializing tcClient service", zap.Error(err))
		return
	}
	tcClient := eth.NewClient(tcClientWrap)

	// init eth client
	ethClientWrap, err := ethclient.Dial(conf.BlockchainConfig.ETHEndpoint)
	if err != nil {
		_logger.AtLog.Logger.Error("error initializing ethClientWrap service", zap.Error(err))
		return
	}
	ethClient := eth.NewClient(ethClientWrap)

	auth2Service := oauth2service.NewAuth2()
	g := global.Global{
		MuxRouter:     r,
		Conf:          conf,
		DBConnection:  mongoConnection,
		Cache:         cache,
		GCS:           gcs,
		QuickNode:     qn,
		Auth2:         auth2Service,
		DiscordClient: dcl,
		Slack:         *slack,
		TcClient:      tcClient,
		EthClient:     ethClient,
	}

	repo, err := repository.NewRepository(&g)
	if err != nil {
		logger.AtLog().Logger.Error("Cannot init repository", zap.Error(err))
		return
	}

	// migration
	migrate.SetDatabase(repo.DB)
	if migrateErr := migrate.Up(-1); migrateErr != nil {
		logger.AtLog().Error("migrate failed", zap.Error(err))
	}

	uc, err := usecase.NewUsecase(&g, *repo)
	if err != nil {
		logger.AtLog().Error("LoadUsecases - Cannot init usecase", zap.Error(err))
		return
	}

	servers := make(map[string]delivery.AddedServer)
	// api fixed run:
	h, _ := httpHandler.NewHandler(&g, *uc)
	servers["http"] = delivery.AddedServer{
		Server:  h,
		Enabled: true,
	}

	//var wait time.Duration
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// start a group cron:
	if len(conf.CronTabList) > 0 {
		for _, cronKey := range conf.CronTabList {
			fmt.Printf("%s is running... \n", cronKey)
			crontabManager.NewCrontabManager(cronKey, &g, *uc).StartServer()
		}
	}

	// uc.TestSendNotify()
	// Run our server in a goroutine so that it doesn't block.
	for name, server := range servers {
		if server.Enabled {
			if server.Server != nil {
				go server.Server.StartServer()
			}
			logger.AtLog().Logger.Info(fmt.Sprintf("%s is enabled", name))
		} else {
			logger.AtLog().Logger.Info(fmt.Sprintf("%s is disabled", name))
		}
	}

	// Block until we receive our signal.
	<-c
	wait := time.Second
	// // Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// // Doesn't block if no connections, but will otherwise wait
	// // until the timeout deadline.
	// err := srv.Shutdown(ctx)
	// if err != nil {
	// 	logger.AtLog().Logger.Error("httpDelivery.StartServer - Server can not shutdown", err)
	// 	return
	// }
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	<-ctx.Done() //if your application should wait for other services
	// to finalize based on context cancellation xxx.
	logger.AtLog().Logger.Warn("httpDelivery.StartServer - server is shutting down!!!")
	tracer.Stop()
	os.Exit(0)

}
