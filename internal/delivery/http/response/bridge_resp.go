package response

import (
	"time"
)

type GenerateDepositAddressResp struct {
	TCAddress     string      `json:"tcAddress"`
	Address       string      `json:"address"`
	PayType       string      `json:"payType"`
	PaymentFee    string      `json:"paymentFee"`    // by pay type
	PaymentAmount string      `json:"paymentAmount"` // by pay type
	TcAmount      string      `json:"tcAmount"`      // buy amount from user
	FeeInfos      interface{} `json:"feeInfos"`
	ExpiredAt     *time.Time  `json:"expiredAt"`
}
