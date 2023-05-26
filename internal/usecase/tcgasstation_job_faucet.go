package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"tcgasstation-backend/internal/entity"
	"tcgasstation-backend/utils/encrypt"
	"tcgasstation-backend/utils/logger"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (u Usecase) JobTcGasStation_SendTCNow() error {

	if len(os.Getenv("TC_MULTI_CONTRACT")) == 0 {
		err := errors.New("TC_MULTI_CONTRACT empty")
		go u.sendLogSlack("", "TcGasStation_SendTCNow.TC_MULTI_CONTRACT", "empty", err.Error())
		return err
	}
	if len(os.Getenv("PRIVATE_KEY_FEE_TC_WALLET")) == 0 {
		err := errors.New("PRIVATE_KEY_FEE_TC_WALLET empty")
		go u.sendLogSlack("", "TcGasStation_SendTCNow.PRIVATE_KEY_FEE_TC_WALLET", "empty", err.Error())
		return err
	}
	if len(os.Getenv("SECRET_KEY")) == 0 {
		err := errors.New("SECRET_KEY empty")
		go u.sendLogSlack("", "TcGasStation_SendTCNow.SECRET_KEY", "empty", err.Error())
		return err
	}

	// check pending first:
	recordsPending, _ := u.Repo.FindTcGasStationByStatus(entity.StatusTcGasStation_InscribedBtc)
	if len(recordsPending) > 0 {
		u.JobTcGasStation_CheckTx(recordsPending)

	}

	// check pending again:
	recordsPending, _ = u.Repo.FindTcGasStationByStatus(entity.StatusTcGasStation_InscribedBtc)
	if len(recordsPending) > 0 {
		return nil
	}

	feeRate := 20

	feeRateCurrent, err := u.getFeeRateFromChain()
	if err == nil {
		feeRate = feeRateCurrent.FastestFee
		// feeRate += 10
	}

	recordsNeedTrigger, _ := u.Repo.FindTcGasStationByStatus(entity.StatusTcGasStation_SubmittedTransfer) // only have tc tx
	fmt.Println("recordsNeedTrigger len: ", len(recordsNeedTrigger))

	if len(recordsNeedTrigger) > 0 {
		// submit raw data:
		tempItem := recordsNeedTrigger[0]

		// check tx tc first:
		context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		status, err := u.TcClient.GetTransaction(context, tempItem.TxTcProcessBuy)
		fmt.Println("GetTransaction status, err ", tempItem.TxTcProcessBuy, status, err)
		if err == nil {
			if status > 0 {
				// pass:
				_, err = u.Repo.UpdateTcGasStation_Status_ByTxDeposit(tempItem.TxTcProcessBuy, entity.StatusTcGasStation_Success)
				if err != nil {
					go u.sendLogSlack(tempItem.UUID, "JobTcGasStation_CheckTx.UpdateTcGasStation", "UpdateTcGasStation", err.Error())
				}
				go u.sendLogSlack(tempItem.UUID, "JobTcGasStation_CheckTx.UpdateStatusFaucetByTxTc", "Update status success before Re-Trigger: ", tempItem.TxTcProcessBuy)
				return nil
			}

		} else {
			go u.sendLogSlack(tempItem.UUID, "JobTcGasStation_CheckTx.GetTransaction", "CheckTxBefore Re-Trigger: ", err.Error())
		}

		txBtc, err := u.SubmitTCToBtcChain(tempItem.TxTcProcessBuy, feeRate)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("JobTcGasStation_CheckTx.SubmitTCToBtcChain"), zap.Error(err))
			go u.sendLogSlack(tempItem.UUID, "JobTcGasStation_CheckTx.Re-SubmitTCToBtcChain", "call send vs tcTx: "+tempItem.TxTcProcessBuy, err.Error())
			return err
		}
		// update for tx:
		_, err = u.Repo.UpdateTcGasStation_TxBtcProcessBuy_ByTxBuy(tempItem.TxTcProcessBuy, txBtc, entity.StatusTcGasStation_InscribedBtc)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("JobTcGasStation_CheckTx.UpdateFaucetByTxTc"), zap.Error(err))
			go u.sendLogSlack(tempItem.UUID, "JobTcGasStation_CheckTx.Re-SubmitTCToBtcChain.UpdateFaucetByTxTc", "update by tx err: "+tempItem.TxTcProcessBuy+", btcTx:"+txBtc, err.Error())
			return err
		}
		go u.sendLogSlack(tempItem.UUID, "JobTcGasStation_CheckTx.Re-SubmitTCToBtcChain", "okk=>tcTx/btcTx", "https://explorer.trustless.computer/tx/"+tempItem.TxTcProcessBuy+"; https://mempool.space/tx/"+txBtc)
		return nil
	}

	// new send TC now:
	recordToSendTc, _ := u.Repo.FindTcGasStationByStatus(entity.StatusTcGasStation_ReceivedFund)
	fmt.Println("record need send tc: ", len(recordToSendTc))

	if len(recordToSendTc) == 0 {
		return nil
	}

	// send TC:
	destinations := make(map[string]*big.Int)

	var uuids []string

	totalAmount := big.NewInt(0)

	t := 0

	var finalTcGasStationRecords []*entity.TcGasStation

	// get list again:
	for _, item := range recordToSendTc {

		if t >= 200 {
			break
		}
		if _, ok := destinations[item.TcAddress]; ok {
			continue
		}

		amount, ok := big.NewInt(0).SetString(item.TcAmount, 10)
		if !ok {
			continue
		}
		destinations[item.TcAddress] = amount

		totalAmount = big.NewInt(0).Add(totalAmount, amount)
		uuids = append(uuids, item.UUID)
		finalTcGasStationRecords = append(finalTcGasStationRecords, item)

		t += 1
	}

	if len(destinations) == 0 {
		return nil
	}

	uuidStr := strings.Join(uuids, ",")

	fmt.Println("destinations: ", destinations)

	privateKeyDeCrypt, err := encrypt.DecryptToString(os.Getenv("PRIVATE_KEY_FEE_TC_WALLET"), os.Getenv("SECRET_KEY"))
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("GenMintFreeTemAddress.Decrypt.%s.Error", "can decrypt"), zap.Error(err))
		return err
	}

	go u.sendLogSlack(fmt.Sprintf("%d", len(uuids)), "JobTcGasStation_CheckTx.SendMulti", "call send with total amount:", totalAmount.String())

	txID, err := u.TcClient.SendMulti(
		os.Getenv("TC_MULTI_CONTRACT"),
		privateKeyDeCrypt,
		destinations,
		totalAmount,
		0,
	)
	fmt.Println("txID, err ", txID, err)

	if err != nil {
		go u.sendLogSlack(uuidStr, "JobTcGasStation_CheckTx.SendMulti", fmt.Sprintf("call send %s err", totalAmount.String()), err.Error())
		return err
	}

	go u.sendLogSlack(uuidStr, "JobTcGasStation_CheckTx.SendMulti", "ok=> tx", txID)

	// update status 1 first:
	if len(uuids) > 0 {
		_, err = u.Repo.UpdateTcGasStation_Status_TxTcProcessDeposit_ByUuids(uuids, entity.StatusTcGasStation_SubmittedTransfer, txID)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("JobTcGasStation_CheckTx.UpdateFaucetByTxTc"), zap.Error(err))
			go u.sendLogSlack(uuidStr, "JobTcGasStation_SendTCNow.UpdateTcGasStation_Status_TxTcProcessDeposit_ByUuids", "update by tx err: "+txID, err.Error())
			return err
		}
	}

	// submit raw data:
	txBtc, err := u.SubmitTCToBtcChain(txID, feeRate)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("JobTcGasStation_CheckTx.SubmitTCToBtcChain"), zap.Error(err))
		go u.sendLogSlack(uuidStr, "JobTcGasStation_CheckTx.SubmitTCToBtcChain", "call send vs tcTx: "+txID, err.Error())
		return err
	}

	go u.sendLogSlack(uuidStr, "JobTcGasStation_CheckTx.SubmitTCToBtcChain", "okk=>tcTx/btcTx", "https://explorer.trustless.computer/tx/"+txID+"\n https://mempool.space/tx/"+txBtc)
	// update tx by uuids:
	if len(uuids) > 0 {
		_, err = u.Repo.UpdateTcGasStation_TxBtcProcessBuy_ByTxBuy(txID, txBtc, entity.StatusTcGasStation_InscribedBtc)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("JobTcGasStation_CheckTx.UpdateFaucetByTxTc"), zap.Error(err))
			go u.sendLogSlack(uuidStr, "JobTcGasStation_SendTCNow.UpdateTcGasStation_Status_TxTcProcessDeposit_ByUuids", "update by tx err: "+txID, err.Error())
			return err
		}
	}

	fmt.Println("-------done------")

	return nil
}
func (u Usecase) JobTcGasStation_CheckTx(recordsToCheck []*entity.TcGasStation) error {

	mapCheckTxPass := make(map[string]bool)
	mapCheckTxFalse := make(map[string]string)

	// check tc tx:
	for _, item := range recordsToCheck {
		context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		status, err := u.TcClient.GetTransaction(context, item.TxTcProcessBuy)

		fmt.Println("GetTransaction status, err ", item.TxTcProcessBuy, status, err)

		if err == nil {
			if status > 0 {
				// pass:
				mapCheckTxPass[item.TxTcProcessBuy] = true
				item.Status = entity.StatusTcGasStation_Success
				_, err = u.Repo.UpdateTcGasStation(item)
				if err != nil {
					go u.sendLogSlack(item.UUID, "JobTcGasStation_CheckTx.UpdateTcGasStation", "UpdateTcGasStation", err.Error())
				}

			} else {
				mapCheckTxFalse[item.TxTcProcessBuy] = "status != 1"
			}
		} else {
			// if error maybe tx is pending or rejected
			// TODO check timeout to detect tx is rejected or not.
			if strings.Contains(err.Error(), "not found") {
				now := time.Now()
				updatedTime := item.UpdatedAt
				if updatedTime != nil {

					duration := now.Sub(*updatedTime).Minutes()
					if duration >= 30 {
						u.sendLogSlack(item.UUID, "JobTcGasStation_CheckTx", fmt.Sprintf("long time to confirm okk? tcTx: https://explorer.trustless.computer/tx/%s, btcTx: https://mempool.space/tx/%s", item.TxTcProcessBuy, item.TxTcProcessBuy), fmt.Sprintf("%.2f mins ago", duration))
						break
					}
				}
			}

			mapCheckTxFalse[item.TxTcProcessBuy] = "err: " + err.Error()
		}
	}
	if len(mapCheckTxFalse) > 0 {
		var uuids []string
		var errs []string
		for k, v := range mapCheckTxFalse {
			uuids = append(uuids, k)
			errs = append(errs, v)
		}

		go u.sendLogSlack(strings.Join(uuids, ","), "JobTcGasStation_CheckTx.UpdateTcGasStation", "mapCheckTxFalse", strings.Join(errs, ","))
	}
	return nil

}

func (u Usecase) SubmitTCToBtcChain(tx string, feeRate int) (string, error) {

	var resp struct {
		Result string `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	payloadStr := fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"method": "eth_inscribeTxWithTargetFeeRate",
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
