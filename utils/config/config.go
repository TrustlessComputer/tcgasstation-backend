package config

import (
	"context"
	"os"
	"regexp"
	"strconv"
	"tcgasstation-backend/utils/slack"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/joho/godotenv"
)

const DevEnv = "develop"
const MainEnv = "production"

type BlockchainConfig struct {
	ETHEndpoint string
	TCEndpoint  string
}

type BitcoinParams struct {
	FirstScannedBTCBlkHeight uint64

	MasterPubKeys         [][]byte
	GeneralMultisigWallet string
	NumRequiredSigs       int
	TotalSigs             int
	MinDepositAmount      uint64
	MinWithdrawAmount     uint64
	DepositFee            uint64

	ChainParam *chaincfg.Params
	InputSize  int
	OutputSize int
	MaxTxSize  int
	MaxFeeRate int
}

var BitcoinParamsRegtest = &BitcoinParams{
	FirstScannedBTCBlkHeight: uint64(1),
	MasterPubKeys: [][]byte{
		[]byte{0x2, 0x80, 0x88, 0xd3, 0x6f, 0xc8, 0xc5, 0x5c, 0x77, 0x99, 0x83, 0xa5, 0xc7, 0xd7, 0xee, 0xdf, 0x12, 0x4f, 0xe2, 0xb5, 0x91, 0xb, 0xc9, 0x2d, 0x69, 0xdd, 0x4a, 0x91, 0xc, 0xc5, 0xf4, 0xf9, 0x22},
		[]byte{0x2, 0x56, 0xa1, 0x55, 0xa4, 0xad, 0x18, 0x31, 0x2d, 0xeb, 0xc2, 0xd3, 0x6f, 0xd8, 0x83, 0xc6, 0x45, 0x62, 0x2f, 0xb6, 0x8d, 0xb0, 0x29, 0x23, 0xa8, 0x92, 0xe1, 0xfe, 0x1d, 0x77, 0xc0, 0xd6, 0x1d},
		[]byte{0x3, 0x34, 0x38, 0x3a, 0x21, 0xa, 0x87, 0x84, 0x72, 0xc1, 0xa4, 0x92, 0xe7, 0x52, 0x9b, 0xbd, 0xdc, 0x83, 0x8f, 0xb7, 0x1f, 0x43, 0x5b, 0x73, 0xa1, 0xc0, 0x88, 0x6f, 0xa, 0xce, 0xab, 0x83, 0x79},
		[]byte{0x2, 0x63, 0x30, 0x9, 0x2b, 0x16, 0x81, 0x9d, 0x1b, 0xe3, 0x7d, 0x15, 0xa3, 0x3d, 0xd, 0xde, 0xd7, 0x6b, 0xdc, 0x50, 0x1c, 0xdc, 0x1, 0x7, 0x30, 0x3e, 0x17, 0x68, 0x65, 0x55, 0x32, 0x4e, 0x49},
		[]byte{0x3, 0x9c, 0x47, 0x10, 0x29, 0xc8, 0xdb, 0xef, 0xec, 0xcd, 0x46, 0xe7, 0x1f, 0x75, 0x85, 0x9e, 0xfd, 0x6e, 0x77, 0x25, 0x8, 0x1a, 0x70, 0x1c, 0xcb, 0x60, 0xd5, 0xf6, 0x87, 0x2e, 0xc0, 0x79, 0x80},
		[]byte{0x2, 0x98, 0x48, 0x3c, 0x14, 0x72, 0xc0, 0xf, 0xfa, 0x23, 0xdc, 0x90, 0x7a, 0xfd, 0xc4, 0x5, 0x5f, 0x6f, 0xcf, 0x78, 0x50, 0xe4, 0x32, 0xfb, 0x62, 0xd, 0x32, 0xc7, 0xd0, 0xa, 0x2a, 0x10, 0x9a},
		[]byte{0x3, 0xff, 0x89, 0x4b, 0x62, 0x1, 0x5c, 0x8, 0x78, 0x3f, 0x8a, 0x75, 0x94, 0x92, 0x6e, 0x11, 0xa2, 0xdb, 0x4d, 0xeb, 0x93, 0x30, 0x8d, 0x30, 0xaa, 0x5e, 0xfb, 0x8a, 0x22, 0xfb, 0xa2, 0xa4, 0x84},
	},
	GeneralMultisigWallet: "bcrt1qgh5mj4uy023euwr2r9saey3r80laz5xfy9l898n95x640dg4dkmqqc0dva",
	NumRequiredSigs:       5,
	TotalSigs:             7,
	MinDepositAmount:      uint64(40000),
	MinWithdrawAmount:     uint64(30000),
	DepositFee:            uint64(10000),

	ChainParam: &chaincfg.RegressionNetParams,
	InputSize:  130,
	OutputSize: 43,
	MaxTxSize:  51200, // 50 KB
	// MaxFeeRate: 30,
}

var BitcoinParamsMaintest = &BitcoinParams{
	FirstScannedBTCBlkHeight: uint64(787496),
	MasterPubKeys: [][]byte{
		[]byte{0x2, 0xe, 0x8, 0xa, 0xe3, 0xcf, 0xf5, 0x1d, 0xc3, 0xc0, 0x83, 0xaf, 0xa9, 0x24, 0x71, 0x9c, 0x2f, 0xca, 0x62, 0x89, 0x74, 0x70, 0xb4, 0x8b, 0x9, 0x51, 0x3, 0x6f, 0x32, 0x9e, 0xdb, 0x5f, 0xe7},
		[]byte{0x3, 0xc2, 0x3c, 0x3d, 0x6f, 0x83, 0xbe, 0xc9, 0x56, 0xde, 0x6a, 0x54, 0x90, 0xac, 0x2d, 0xe7, 0xee, 0x5c, 0xf8, 0x63, 0x22, 0x84, 0x9c, 0x61, 0xed, 0x62, 0x5b, 0x69, 0x8f, 0x4a, 0x4a, 0xee, 0x49},
		[]byte{0x3, 0x3b, 0x75, 0x80, 0x77, 0x8f, 0x4d, 0x2e, 0x46, 0x20, 0x6a, 0xd5, 0x32, 0x66, 0x18, 0xb6, 0xd6, 0x4c, 0x46, 0x1a, 0xe, 0x47, 0xb2, 0x5a, 0x77, 0xad, 0x72, 0xdc, 0x56, 0x4e, 0xa6, 0xca, 0xdb},
		[]byte{0x2, 0xfc, 0x81, 0x32, 0xba, 0xb2, 0x85, 0x71, 0x82, 0x3f, 0x82, 0x3d, 0x74, 0xe5, 0xd4, 0xa2, 0xff, 0xcb, 0xb, 0xe7, 0x2c, 0x49, 0x63, 0xb3, 0x73, 0x75, 0xf7, 0xc4, 0x41, 0xf9, 0x3e, 0xda, 0x96},
		[]byte{0x2, 0xb2, 0x14, 0xd8, 0x19, 0x77, 0x31, 0x59, 0xa3, 0xae, 0x9c, 0x30, 0xf9, 0x85, 0xa4, 0xe0, 0x56, 0x1d, 0x98, 0x9d, 0xf6, 0x27, 0xfb, 0xbd, 0xd6, 0x9d, 0x33, 0xdd, 0xa7, 0x25, 0x38, 0x35, 0xb5},
		[]byte{0x3, 0xb8, 0xc6, 0xf2, 0x80, 0x9d, 0xe5, 0xd, 0x6a, 0x43, 0x57, 0xb9, 0xac, 0xce, 0xaa, 0xd, 0x8, 0xda, 0xd4, 0x75, 0xd9, 0x6a, 0xbf, 0x70, 0x14, 0xe7, 0x2a, 0xeb, 0x68, 0xe0, 0xb1, 0xa4, 0xb4},
		[]byte{0x2, 0xee, 0x76, 0x26, 0xa3, 0x4f, 0xd, 0xb7, 0x57, 0x21, 0xa4, 0x44, 0x8c, 0xa8, 0x6e, 0x59, 0xa1, 0x32, 0x2d, 0xa6, 0xbe, 0xf8, 0x86, 0xbf, 0x64, 0xa5, 0x94, 0xa7, 0xed, 0x20, 0xfd, 0xc2, 0x52},
	},
	GeneralMultisigWallet: "bc1qajyp9ekpepmhftxq8aeps4cv8gjkat00nfk9lplqec7mevhv4z6qxy37z8",
	NumRequiredSigs:       5,
	TotalSigs:             7,
	MinDepositAmount:      uint64(40000),
	MinWithdrawAmount:     uint64(30000),
	DepositFee:            uint64(10000),

	ChainParam: &chaincfg.MainNetParams,
	InputSize:  192,
	OutputSize: 43,
	MaxTxSize:  51200, // 50 KB
	// MaxFeeRate: 150,
}

type Config struct {
	Debug         bool
	StartHTTP     bool
	Context       *Context
	Databases     *Databases
	Redis         RedisConfig
	ENV           string
	ServicePort   string
	QuickNode     string
	BlockStream   string
	NftExplorer   string
	TokenExplorer string
	BFSService    string
	BNSService    string
	Gcs           *GCS

	Slack slack.Config

	BlockchainConfig BlockchainConfig

	BitcoinParams *BitcoinParams

	// list crontab to run:
	CronTabList []string

	GoogleSecretKey string // read it from google
}

type Context struct {
	TimeOut int
}

type Databases struct {
	Postgres *DBConnection
	Mongo    *DBConnection
}

type DBConnection struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	Sslmode string
	Scheme  string
}

type Mongo struct {
	DBConnection
}

type GCS struct {
	ProjectId string
	Bucket    string
	Auth      string
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       string
	ENV      string
}

func NewConfig(filePaths ...string) (*Config, error) {
	if len(filePaths) > 0 {
		godotenv.Load(filePaths[0])
	} else {
		godotenv.Load()
	}
	services := make(map[string]string)
	isDebug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	isStartHTTP, _ := strconv.ParseBool(os.Getenv("START_HTTP"))

	timeOut, err := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	if err != nil {
		panic(err)
	}

	services["og"] = os.Getenv("OG_SERVICE_URL")

	bitcoinParams := &BitcoinParams{}
	env := os.Getenv("ENV")
	switch env {
	case DevEnv:
		{
			bitcoinParams = BitcoinParamsRegtest
		}
	case MainEnv:
		{
			bitcoinParams = BitcoinParamsMaintest

		}
	default:
		panic("Invalid env")
	}

	conf := &Config{
		ENV:       os.Getenv("ENV"),
		StartHTTP: isStartHTTP,
		Context: &Context{
			TimeOut: timeOut,
		},
		Debug:         isDebug,
		ServicePort:   os.Getenv("SERVICE_PORT"),
		QuickNode:     os.Getenv("QUICKNODE_URL"),
		BlockStream:   os.Getenv("BLOCK_STREAM_URL"),
		NftExplorer:   os.Getenv("NFT_EXPLORER_URL"),
		TokenExplorer: os.Getenv("TOKEN_EXPLORER_URL"),
		BFSService:    os.Getenv("BFS_SERVICE_URL"),
		BNSService:    os.Getenv("BNS_SERVICE_URL"),
		Databases: &Databases{
			Mongo: &DBConnection{
				Host:   os.Getenv("MONGO_HOST"),
				Port:   os.Getenv("MONGO_PORT"),
				User:   os.Getenv("MONGO_USER"),
				Pass:   os.Getenv("MONGO_PASSWORD"),
				Name:   os.Getenv("MONGO_DB"),
				Scheme: os.Getenv("MONGO_SCHEME"),
			},
		},
		Redis: RedisConfig{
			Address:  os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       os.Getenv("REDIS_DB"),
			ENV:      os.Getenv("REDIS_ENV"),
		},
		Gcs: &GCS{
			ProjectId: os.Getenv("GCS_PROJECT_ID"),
			Bucket:    os.Getenv("GCS_BUCKET"),
			Auth:      os.Getenv("GCS_AUTH"),
			Endpoint:  os.Getenv("GCS_ENDPOINT"),
			Region:    os.Getenv("GCS_REGION"),
			AccessKey: os.Getenv("GCS_ACCESS_KEY"),
			SecretKey: os.Getenv("GCS_SECRET_KEY"),
		},
		BlockchainConfig: BlockchainConfig{
			ETHEndpoint: os.Getenv("ETH_ENDPOINT"),
			TCEndpoint:  os.Getenv("TC_ENDPOINT"),
		},
		Slack: slack.Config{
			Token:         os.Getenv("SLACK_TOKEN"),
			ChannelLogs:   os.Getenv("SLACK_CHANNEL_LOGS"),
			ChannelOrders: os.Getenv("SLACK_CHANNEL_ORDERS"),
			Env:           os.Getenv("ENV"),
		},
		CronTabList: regexp.MustCompile(`\s*[,;]\s*`).Split(os.Getenv("CRONTAB_LIST"), -1),

		BitcoinParams: bitcoinParams,
	}

	googleSecretKey, err := GetGoogleSecretKey(os.Getenv("SECRET_KEY"))
	if err != nil {
		if env == MainEnv {
			panic("can not GetGoogleSecretKey")
		} else {
			googleSecretKey = os.Getenv("SECRET_KEY")
		}

	}
	conf.GoogleSecretKey = googleSecretKey

	return conf, nil
}

func (c Config) IsMainnet() bool {
	return c.ENV == MainEnv
}
func (c Config) IsDevnet() bool {
	return c.ENV == DevEnv
}

// get google secret:
func GetGoogleSecretKey(name string) (string, error) {

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}