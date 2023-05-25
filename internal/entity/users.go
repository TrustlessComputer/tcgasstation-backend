package entity

import (
	"tcgasstation-backend/utils"
	"time"
)

const (
	WalletType_BTC_PRVKEY = "btc_prvkey"
)

type Users struct {
	BaseEntity              `bson:",inline"`
	Message                 string     `bson:"message"`
	WalletAddress           string     `bson:"wallet_address" json:"wallet_address,omitempty"`                         // eth wallet define user in platform by connect wallet and sign
	WalletAddressPayment    string     `bson:"wallet_address_payment" json:"wallet_address_payment,omitempty"`         // eth wallet artist receive royalty
	WalletAddressBTC        string     `bson:"wallet_address_btc" json:"wallet_address_btc,omitempty"`                 // btc wallet artist receive royalty
	WalletAddressBTCTaproot string     `bson:"wallet_address_btc_taproot" json:"wallet_address_btc_taproot,omitempty"` // btc wallet receive minted nft
	DisplayName             string     `bson:"display_name" json:"display_name,omitempty"`
	Bio                     string     `bson:"bio" json:"bio,omitempty"`
	Avatar                  string     `bson:"avatar" json:"avatar"`
	WalletType              string     `bson:"wallet_type" json:"wallet_type"`
	LastLoggedinAt          *time.Time `bson:"last_loggedin_at" json:"last_loggedin_at"`
}

func (t *Users) CollectionName() string {
	return utils.COLLECTION_USERS
}
