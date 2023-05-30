package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tcgasstation-backend/utils/eth"
	"tcgasstation-backend/utils/logger"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const MAX_TC_TO_BUY = 100

const MinBTCConfirmation = 3
const MinTCConfirmation = 6

const MinETHConfirmation = 15

const TIME_TO_REPLACE_BY_FEE = 20 // 20 min

const TCTokenDecimal = 18

var DEPOSIT_LAST_SCANNED_BTC_BLOCK_HEIGHT = "deposit_last_scanned_btc_block_height"

const TOTAL_BYTES_1_TOKEN = 1301

const MINT_BTC_FEE = 10000 // sat

const INC_FEE_RATE = 30

const GAS_LIMIT_WITHDRAW_ETH = 100000

const TC_ETH_PRICE = 0.0069
const TC_BTC_PRICE = 0.000472167

const agvFileSize = 570

type FeeRates struct {
	FastestFee  int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee     int `json:"hourFee"`
	EconomyFee  int `json:"economyFee"`
	MinimumFee  int `json:"minimumFee"`
}

func (u Usecase) getFeeRateFromChain() (*FeeRates, error) {

	response, err := http.Get("https://mempool.space/api/v1/fees/recommended")

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println("responseData", string(responseData))

	feeRateObj := &FeeRates{}

	err = json.Unmarshal(responseData, &feeRateObj)
	if err != nil {
		logger.AtLog.Logger.Error("err", zap.Error(err))
		return nil, err
	}
	return feeRateObj, nil

}

func (u Usecase) ReplaceByFeeTcWithBtcTx(tx string, feeRate int) (string, error) {

	var resp struct {
		Result string `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	payloadStr := fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"method": "eth_replaceTxWithTargetFeeRate",
			"params": [
				"%s",%d
			],
			"id": 1
		}`, tx, feeRate)

	payload := strings.NewReader(payloadStr)

	fmt.Println("payloadStr: ", payloadStr)

	client := &http.Client{}
	req, err := http.NewRequest("POST", u.Config.BlockchainConfig.TCEndpoint, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// fmt.Println("body", string(body))

	/*
			{
		    "jsonrpc": "2.0",
		    "id": 1,
		    "result": "tx"
		}
	*/

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}
	if len(resp.Result) == 0 && resp.Error != nil {
		// error:
		return "", errors.New(resp.Error.Message)
	}

	// inscribe ok now:

	return resp.Result, nil
}

func (u Usecase) SendSlackInfo(prefix string, tcAddress string, txID string, processTxID string, amount string, funcName string, id string, isOrderSlack bool) {
	msg := ""
	if processTxID == "" {
		msg = fmt.Sprintf("%v \nTCAddress: %v\nTxID: %v\nAmount: %v", prefix, tcAddress, txID, amount)
	} else {
		msg = fmt.Sprintf("%v \nTCAddress: %v\nTxID: %v\nProcessTxID: %v\nAmount: %v", prefix, tcAddress, txID, processTxID, amount)
	}

	if isOrderSlack {
		u.sendOrderSlack(id, msg, funcName, "")
	} else {
		u.sendLogSlack(id, msg, funcName, "")
	}
}

// dec is 1e18 or 1e8
func getAmountFormatStr(amount string, dec float64) string {
	amtBN, _ := new(big.Int).SetString(amount, 10)
	return fmt.Sprintf("%.4f", float64(float64(amtBN.Uint64()))/dec)
}

func convertBTCtoNanopBTC(amount float64) uint64 {
	return uint64(amount*1e9 + 0.5)
}

func switchAmountWithDecimals(amount string, fromDecimal, toDecimal int) (*big.Int, error) {
	amountFrom, ok := big.NewInt(0).SetString(amount, 10)
	if !ok {
		return nil, errors.New("can not convert amount to bigInt")
	}
	fmt.Println("amount from:", amountFrom.String())
	fmt.Println("decimal from:", fromDecimal)
	fmt.Println("decimal to:", toDecimal)

	var amountTo *big.Int
	newDecimal := toDecimal - fromDecimal
	fmt.Println("newDecimal:", newDecimal)

	if newDecimal >= 0 {
		newAmountDecimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(newDecimal)), nil)
		fmt.Println("newAmountDecimal:", newAmountDecimal)
		amountTo = new(big.Int).Mul(amountFrom, newAmountDecimal)
		fmt.Println("amountConvert:", amountTo)
	} else {
		newDecimal = fromDecimal - toDecimal
		newAmountDecimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(newDecimal)), nil)
		fmt.Println("newAmountDecimal:", newAmountDecimal)
		amountTo = new(big.Int).Div(amountFrom, newAmountDecimal)
		fmt.Println("amountConvert:", amountTo)
	}
	return amountTo, nil
}

func (u Usecase) sendLogSlack(ids, funcName, requestMsgStr, messageStr string) {
	channelID := u.Config.Slack.ChannelLogs
	preText := fmt.Sprintf("[App: %s][recordIDs %s] - %s", "tcGasStation", ids, requestMsgStr)
	if _, _, err := u.Slack.SendMessageToSlackWithChannel(channelID, preText, funcName, messageStr); err != nil {
		fmt.Println("s.Slack.SendMessageToSlack err", err)
		return
	}
	fmt.Println("s.Slack.SendMessageToSlack Success")
}

// alert success deposit/withdraw:
func (u Usecase) sendOrderSlack(ids, funcName, requestMsgStr, messageStr string) {
	channelID := u.Config.Slack.ChannelOrders
	preText := fmt.Sprintf("[App: %s][recordIDs %s] - %s", "tcGasStation new Order", ids, requestMsgStr)
	if _, _, err := u.Slack.SendMessageToSlackWithChannel(channelID, preText, funcName, messageStr); err != nil {
		fmt.Println("s.Slack.SendMessageToSlack err", err)
		return
	}
	fmt.Println("s.Slack.SendMessageToSlack Success")

}

func GetExternalPrice(tokenSymbol, toSymbol string) (float64, error) {

	if len(toSymbol) == 0 {
		toSymbol = "USDT"
	}

	binanceAPI := os.Getenv("BINANCE_API")
	if binanceAPI == "" {
		binanceAPI = "https://api.binance.us"
	}
	binancePriceURL := fmt.Sprintf("%v/api/v3/ticker/price?symbol=", binanceAPI)
	var price struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	var jsonErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return 0, nil
	}
	tk := strings.ToUpper(tokenSymbol)
	fullURL := binancePriceURL + tk + toSymbol
	fmt.Println("fullURL: ", fullURL)
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &price)
	if err != nil {
		err = json.Unmarshal(body, &jsonErr)
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	value, err := strconv.ParseFloat(price.Price, 32)
	if err != nil {
		log.Println("getExternalPrice", tokenSymbol, err)
		return 0, err
	}
	return value, nil
}

func (u *Usecase) convertToNewCoinWithPrice(fromAmount string, fromDecimal, toDecimal int, fromPrice, toPrice float64) (*big.Int, float64, float64, error) {

	//amount = "0.1"
	powIntput := math.Pow10(fromDecimal)
	powIntputBig := new(big.Float)
	powIntputBig = powIntputBig.SetFloat64(powIntput)
	amountMintBTC, _ := big.NewFloat(0).SetString(fromAmount)
	amountMintBTC = amountMintBTC.Mul(amountMintBTC, powIntputBig)

	btcToETH := fromPrice / toPrice

	fmt.Println("fromPrice: ", fromPrice)
	fmt.Println("toPrice: ", toPrice)
	fmt.Println("rate btcToETH: ", btcToETH)

	rate := new(big.Float)
	rate.SetFloat64(btcToETH)
	amountMintBTC = amountMintBTC.Mul(amountMintBTC, rate)

	fmt.Println("amountMintBTC*rate: ", amountMintBTC.String())

	pow := math.Pow10(toDecimal - fromDecimal)
	powBig := new(big.Float)
	powBig.SetFloat64(pow)

	fmt.Println("powBig ", powBig.String())

	amountMintBTC = amountMintBTC.Mul(amountMintBTC, powBig)
	result := new(big.Int)
	amountMintBTC.Int(result)

	logger.AtLog.Logger.Info("convertToNewCoinWithPrice", zap.String("amount", fromAmount), zap.Float64("fromPrice", fromPrice), zap.Float64("toPrice", toPrice))
	return result, fromPrice, toPrice, nil

}

func (u *Usecase) EstFeeDepositBtc(fastestFee int) (int, error) {

	if fastestFee == 0 {
		feeRateCurrent, err := u.getFeeRateFromChain()

		if err != nil {
			return 0, err
		}
		fastestFee = feeRateCurrent.FastestFee
	}

	feeBtc := TOTAL_BYTES_1_TOKEN * fastestFee / 4
	feeBtc = feeBtc + feeBtc*INC_FEE_RATE/100

	if feeBtc < MINT_BTC_FEE {
		feeBtc = MINT_BTC_FEE
	}
	return feeBtc, nil

}

func (u *Usecase) EstFeeDepositEth(depositFeeByBtc int) (*big.Int, error) {

	fmt.Println("EstFeeDepositEth.depositFeeByBtc: ", depositFeeByBtc)

	btcRate, err := GetExternalPrice("BTC", "USDT")
	if err != nil {
		logger.AtLog.Error("GetExternalPrice(BTC)", zap.Error(err))
		return nil, err
	}

	toRate, err := GetExternalPrice(strings.ToUpper("ETH"), "USDT")
	if err != nil {
		logger.AtLog.Error(fmt.Sprintf("GetExternalPrice(%s)", "TC"), zap.Error(err))
		return nil, err
	}

	fmt.Println("toRate: ", toRate)

	mintPriceByTo, _, _, err := u.convertToNewCoinWithPrice(fmt.Sprintf("%f", float64(depositFeeByBtc)/1e8), 8, 18, btcRate, toRate)

	fmt.Println("mintPriceByTo: ", mintPriceByTo)

	return mintPriceByTo, nil
}

// format number:
func bigIntString(balance *big.Int, decimals int64) string {
	amount := bigIntFloat(balance, decimals)
	deci := fmt.Sprintf("%%0.%vf", decimals)
	return clean(fmt.Sprintf(deci, amount))
}
func bigIntStringWithDec(balance *big.Int, decimals, dec int64) string {
	amount := bigIntFloat(balance, decimals)
	deci := fmt.Sprintf("%%0.%vf", dec)
	return clean(fmt.Sprintf(deci, amount))
}

func bigIntFloat(balance *big.Int, decimals int64) *big.Float {
	if balance.Sign() == 0 {
		return big.NewFloat(0)
	}
	bal := big.NewFloat(0)
	bal.SetInt(balance)
	pow := bigPow(10, decimals)
	p := big.NewFloat(0)
	p.SetInt(pow)
	bal.Quo(bal, p)
	return bal
}

func bigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}

func clean(newNum string) string {
	stringBytes := bytes.TrimRight([]byte(newNum), "0")
	newNum = string(stringBytes)
	if stringBytes[len(stringBytes)-1] == 46 {
		newNum += "0"
	}
	if stringBytes[0] == 46 {
		newNum = "0" + newNum
	}
	return newNum
}

func CheckTxAndBlockConfirmed(client *eth.Client, txHash string, minConfirmationBlocks uint64,
) (bool, error) {

	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	receipt, err := client.TransactionReceipt(context, common.HexToHash(txHash))

	if err != nil {
		return false, err
	}

	// tx failed:
	if receipt.Status == 0 {
		return false, nil
	}

	// status: 1. check confirm:
	latestBlockHeight, err := client.GetClient().BlockNumber(context)

	if err != nil {
		return false, err
	}

	if big.NewInt(int64(latestBlockHeight)).Cmp(big.NewInt(0).Add(receipt.BlockNumber, big.NewInt(int64(minConfirmationBlocks)))) == -1 {
		return false, fmt.Errorf("it needs %v confirmation blocks for the process, the requested block (%v) but the latest block (%v)", minConfirmationBlocks, receipt.BlockNumber.Uint64(), latestBlockHeight)
	}
	// check block hash string:
	headerByNumber, err := client.BlockByNumber(context, receipt.BlockNumber)
	if err != nil {
		return false, err
	}
	if headerByNumber.Hash().String() != receipt.BlockHash.String() {
		return false, fmt.Errorf("the requested evm BlockHash %s is being on fork branch", receipt.BlockHash.String())
	}

	return true, nil
}

func (u *Usecase) estFeeDepositEth() (*big.Int, error) {

	depositFeeByBtc, err := u.EstFeeDepositBtc(0)
	if err != nil {
		return nil, err
	}

	mintPriceByTo, err := u.EstFeeDepositEth(depositFeeByBtc)
	if err != nil {
		return nil, err
	}

	return mintPriceByTo, nil
}
