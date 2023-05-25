package entity

import (
	"math/big"
	"tcgasstation-backend/utils/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type StatusTcGasStation int

const (
	// for deposit+withdraw
	StatusTcGasStation_Pending        StatusTcGasStation = iota // 0: wait for the payment
	StatusTcGasStation_WaitForConfirm                           // 1: wait for confirm
	StatusTcGasStation_TxConfirmed                              // 2: tx confirmed

	// deposit:
	StatusTcGasStation_SubmittedTransfer // 3: submitted transaction on Tc chain (will have tc tx)
	StatusTcGasStation_InscribedBtc      // 4: trigger on btc (will have btc tx)
	StatusTcGasStation_Success           // 5: mint success (tc tx confirmed)

	StatusTcGasStation_TimedOut // 6:

	StatusTcGasStation_NeedToRefund // 7:
	StatusTcGasStation_Refunding    // 8:
	StatusTcGasStation_Refunded     // 9:

	StatusTcGasStation_Invalid // 10:

)

type TcGasStation struct {
	BaseEntity `bson:",inline"`

	TcAddress string `bson:"tc_address" json:"tcAddress"` // user tc address

	Status StatusTcGasStation `bson:"status" json:"status"`

	ExpiredAt time.Time `bson:"expired_at"`

	PayType string `bson:"pay_type" json:"payType"`

	ReceiveAddress string `bson:"receiveAddress"` // address generated to receive coin from users.
	PrivateKey     string `bson:"privateKey"`     // private key of the receive wallet.

	AmountInTC string `bson:"amount_tc" json:"amountTc"` // buy amount for user (in tc chain, dec 18)
	Fee        string `bson:"fee" json:"fee"`            // deposit/withdraw fee decimal 18

	EstFee string `bson:"est_fee" json:"estFee"` // est fee decimal 18

	AmountReceivedBuy string `bson:"amount_received_buy" json:"amountReceivedBuy"` // by btc, or eth ...
	TxTcProcessBuy    string `bson:"tx_tc_process_buy" json:"txTcProcessBuy"`      // tx TC process buy on tc chain
	TxBtcProcessBuy   string `bson:"tx_btc_process_buy" json:"txBtcProcessBuy"`    // tx btc inscribe to process buy

	FeeInfo interface{} `bson:"fee_info" json:"feeInfo"` // some note...

	Note string `bson:"note" json:"note"` // some note...
}

func (t TcGasStation) CollectionName() string {
	return "tc_gas_stations"
}

func (u TcGasStation) ToBson() (*bson.D, error) {
	return helpers.ToDoc(u)
}

type FeeInfo struct {
	FeeBtc   int    `json:"feeBtc"`
	FeeToken string `json:"feeToken"`
	FeeRate  int    `json:"feeRate"`
}

type BuyTcFeeInfo struct {

	//string
	TcPrice string `json:"tcPrice"`

	FeeRate int `json:"feeRate"`

	NetworkFee  string `json:"networkFee"`
	InscribeFee string `json:"inscribeFee"`
	SendFundFee string `json:"sendFundFee"`
	TotalAmount string `json:"totalAmount"`

	// big number
	TcPriceBigInt    *big.Int `json:"-"`
	NetworkFeeBigInt *big.Int `json:"-"`

	InscribeFeeInt    *big.Int `json:"-"`
	SendFundFeeBigInt *big.Int `json:"-"`

	TotalAmountBigInt *big.Int `json:"-"`

	EthPrice float64 `json:"ethPrice"`
	BtcPrice float64 `json:"btcPrice"`
	Decimal  int     `json:"decimal"`
}
