package entity

import (
	"tcgasstation-backend/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
)

type CronJobManager struct {
	BaseEntity `bson:",inline"`

	JobKey string `bson:"job_key"` // the config key of the job server, ex: MAKETPLACE_CRONTAB_START,MINT_NFT_BTC_START

	Group string `bson:"group"` // the name of group, ex: Marketplace, Inscribe, MintBtc, CrawlData ...

	JobName string `bson:"job_name"` // the name of the job: ex: waitForBalance, checkTxMint ....

	Schedule string `bson:"schedule"` // the timer of the job, ex: @every 5s, 0 0 * * *, ....Need split jobs

	Enabled bool `bson:"enabled"` // in the job function, the job will not run when it is not enabled, need some code to do it.

	Description string `bson:"description"`

	LastStatus string `bson:"last_status"` // last time for running/pause

	FunctionName string `bson:"function_name"`

	WebHook string `bson:"webhook"` // the webhook internal link, ex: http://backend/webhook/marketplace/crontMintNft. Currently it be not use.
}

func (u CronJobManager) CollectionName() string {
	return "cron_job_managers"
}

func (u CronJobManager) ToBson() (*bson.D, error) {
	return helpers.ToDoc(u)
}

type CronJobManagerLogs struct {
	BaseEntity  `bson:",inline"`
	RecordID    string      `bson:"record_id"`
	Table       string      `bson:"table"`
	Name        string      `bson:"name"`
	Status      interface{} `bson:"status"`
	RequestMsg  interface{} `bson:"request_msg"`
	ResponseMsg interface{} `bson:"response_msg"`
}

func (u CronJobManagerLogs) CollectionName() string {
	return "cron_job_manager_logs"
}

func (u CronJobManagerLogs) ToBson() (*bson.D, error) {
	return helpers.ToDoc(u)
}
