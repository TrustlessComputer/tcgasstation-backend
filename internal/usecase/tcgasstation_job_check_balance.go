package usecase

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"tcgasstation-backend/internal/entity"
	"tcgasstation-backend/utils"
	"tcgasstation-backend/utils/btc"
	"tcgasstation-backend/utils/logger"
	"tcgasstation-backend/utils/slack"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (u *Usecase) JobTcGasStation_CheckBalance() error {

	list, err := u.Repo.ListTcGasStationPending() // list pending

	if err != nil {
		return err
	}
	if len(list) == 0 {
		return nil
	}

	for _, item := range list {

		// check balance:
		balance := big.NewInt(0)
		confirm := -1

		if item.PayType == utils.NETWORK_ETH {

			// check balance:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			balance, err = u.EthClient.GetBalance(ctx, item.ReceiveAddress)
			fmt.Println("GetBalance eth response: ", balance, err)

			confirm = 1

		} else if item.PayType == utils.NETWORK_BTC {

			balanceQuickNode, err := btc.GetBalanceFromQuickNode(item.ReceiveAddress, u.Config.QuickNode)
			if err != nil {
				err = errors.Wrap(err, "btc.GetBalanceFromQuickNode")
			}
			if balanceQuickNode != nil {
				balance = big.NewInt(int64(balanceQuickNode.Balance))
				// check confirm:
				if len(balanceQuickNode.Txrefs) > 0 {
					var txInfo *btc.QuickNodeTx
					txInfo, err = btc.CheckTxfromQuickNode(balanceQuickNode.Txrefs[0].TxHash, u.Config.QuickNode)
					if err == nil {
						if txInfo != nil {
							confirm = txInfo.Result.Confirmations
						}

					} else {
						// slack
					}
				}
			}
		}
		if err != nil {
			fmt.Printf("Could not GetBalance Bitcoin - with err: %v", err)
			time.Sleep(300 * time.Millisecond)
			continue
		}
		if balance == nil {
			err = errors.New("balance is nil")
			fmt.Println(err)
			continue
		}

		if balance.Uint64() == 0 {
			continue
		}

		// get required amount to check vs temp wallet balance:
		amount, ok := big.NewInt(0).SetString(item.PaymentAmount, 10)
		if !ok {
			err = errors.New("cannot parse amount")
			fmt.Println(err)
			continue
		}

		if amount.Uint64() == 0 {
			err := errors.New("balance is zero")
			fmt.Println(err)
			continue
		}

		// set receive balance:
		item.AmountReceived = amount.String()

		if balance.Uint64() < amount.Uint64() {
			err := fmt.Errorf("Not enough amount %d < %d ", balance.Uint64(), amount.Uint64())
			fmt.Println(err)

			item.Status = entity.StatusTcGasStation_NeedToRefund
			item.ReasonRefund = "Not enough balance."
			u.Repo.UpdateTcGasStation(&item)
			continue
		}

		if confirm == 0 {
			item.Status = entity.StatusTcGasStation_WaitForConfirm
			u.Repo.UpdateTcGasStation(&item)
		}
		if confirm >= 1 {
			// received fund:
			item.Status = entity.StatusTcGasStation_ReceivedFund
			item.IsConfirm = true

			logger.AtLog.Logger.Info(fmt.Sprintf("TcGasStation_Job_CheckBalance.CheckReceiveFund.%s", item.ReceiveAddress), zap.Any("item", item))

		}

		_, err = u.Repo.UpdateTcGasStation(&item)
		if err != nil {
			fmt.Printf("Could not UpdateTcGasStation uuid %s - with err: %v", item.UUID, err)
			continue
		}

	}
	return nil
}

func (u *Usecase) JobTcGasStation_CheckVolumeTC() error {
	listTCBuyed, _ := u.Repo.ListTcGasStationByStatus([]entity.StatusTcGasStation{entity.StatusTcGasStation_Success})

	totalTC := big.NewInt(0)

	totalBTC := big.NewInt(0)
	totalETH := big.NewInt(0)

	requestETH := 0
	requestBTC := 0

	userBuyEth := make(map[string]int)
	userBuyBtc := make(map[string]int)

	for _, item := range listTCBuyed {
		amountTc, _ := big.NewInt(0).SetString(item.TcAmount, 10)
		totalTC = big.NewInt(0).Add(totalTC, amountTc)
		requestBTC += 1

		if item.PayType == NETWORK_BTC {
			amount, _ := big.NewInt(0).SetString(item.PaymentAmount, 10)
			totalBTC = big.NewInt(0).Add(totalBTC, amount)

			if _, ok := userBuyBtc[item.TcAddress]; !ok {
				userBuyBtc[item.TcAddress] += 1
			}

		}
		if item.PayType == NETWORK_ETH {
			amount, _ := big.NewInt(0).SetString(item.PaymentAmount, 10)
			totalETH = big.NewInt(0).Add(totalETH, amount)
			requestETH += 1
		}
		if _, ok := userBuyEth[item.TcAddress]; !ok {
			userBuyEth[item.TcAddress] += 1
		}

	}
	fmt.Println("tc buying total: ", totalTC)
	fmt.Println("ETH total: ", totalETH)
	fmt.Println("BTC total: ", totalBTC)

	channelID := "C059KMQMKQX" // "C05A4BEDFSR"
	preText := fmt.Sprintf("[App: %s]", "TcGasStation monitor")

	message := fmt.Sprintf(fmt.Sprintf("%s TC. \n%s ETH (%d requests/%d users). \n%s BTC (%d requests/%d users).", bigIntStringWithDec(totalTC, 18, 4), bigIntStringWithDec(totalETH, 18, 4), requestETH, len(userBuyEth), bigIntStringWithDec(totalBTC, 8, 4), requestBTC, len(userBuyBtc)))

	slack := slack.NewSlack2(os.Getenv("SLACK_TOKEN2"), u.Config.ENV)

	if _, _, err := slack.SendMessageToSlackWithChannel(channelID, preText, "Gas Station Volume", message); err != nil {
		fmt.Println("s.Slack.SendMessageToSlack err", err)
		return err
	}
	fmt.Println("s.Slack.SendMessageToSlack Success")

	return nil

}
