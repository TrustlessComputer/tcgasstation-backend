package response

import (
	"time"
)

type GenerateDepositAddressResp struct {
	TCAddress string      `json:"tcAddress"`
	Address   string      `json:"address"`
	EstFee    string      `json:"estFee"`
	FeeInfos  interface{} `json:"feeInfos"`
	ExpiredAt *time.Time  `json:"expiredAt"`
}
