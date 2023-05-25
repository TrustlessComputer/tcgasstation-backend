package entity

import (
	"tcgasstation-backend/utils/helpers"

	"go.mongodb.org/mongo-driver/bson"
)

type TokenNetNetwork string

const (
	TokenNetwork_BTC TokenNetNetwork = "bitcoin"
	TokenNetwork_ETH                 = "ethereum"
	// more
)

type TokenType string

const (
	TokenType_Native TokenType = "native"
	TokenType_ERC20            = "erc20"
)

type TcToken struct {
	BaseEntity `bson:",inline"`

	OutChainTokenID string          `bson:"out_chain_token_id" json:"outChainTokenID"`
	Network         TokenNetNetwork `bson:"network" json:"network"`
	Type            TokenType       `bson:"type" json:"type"`
	Decimals        int             `bson:"decimals" json:"decimals"`
	Name            string          `bson:"name" json:"name"`
	Symbol          string          `bson:"symbol" json:"symbol"`

	TcTokenID string `bson:"tc_token_id" json:"tcTokenID"`

	Status int `bson:"status" json:"status"`

	PriceUsd float64 `bson:"price_usd" json:"priceUsd"`

	IsStableCoin bool `bson:"is_stable_coin" json:"isStableCoin"`
}

func (t TcToken) CollectionName() string {
	return "tc_tokens"
}

func (u TcToken) ToBson() (*bson.D, error) {
	return helpers.ToDoc(u)
}
