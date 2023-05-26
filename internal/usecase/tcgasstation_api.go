package usecase

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
	"tcgasstation-backend/internal/delivery/http/request"
	"tcgasstation-backend/internal/delivery/http/response"
	"tcgasstation-backend/internal/entity"
	"tcgasstation-backend/utils"
	"tcgasstation-backend/utils/btc"
	"tcgasstation-backend/utils/encrypt"
	"tcgasstation-backend/utils/eth"
	"tcgasstation-backend/utils/logger"
	"time"

	"go.uber.org/zap"
)

const (
	NETWORK_BTC = "btc"
	NETWORK_ETH = "eth"
)

func (u *Usecase) GenerateDepositAddress(data *request.GenerateDepositAddressReq) (*response.GenerateDepositAddressResp, error) {

	var privateKey, privateKeyEnCrypt, receiveAddress string
	var err error

	keyToEncrypt := os.Getenv("SECRET_KEY")

	if len(keyToEncrypt) == 0 {
		return nil, errors.New("payType invalid")
	}

	if !eth.ValidateAddress(data.TCAddress) {
		return nil, errors.New("tcAddress invalid")
	}

	tcAmountFloat := big.NewFloat(data.TcAmount)
	tcAmountFloat.Mul(tcAmountFloat, new(big.Float).SetFloat64(math.Pow10(18)))

	tcAmount := new(big.Int)
	tcAmountFloat.Int(tcAmount)

	// todo check max 100TC?
	maxAmountIntWei, _ := big.NewInt(0).SetString(big.NewInt(0).Mul(new(big.Int).SetInt64(MAX_TC_TO_BUY), new(big.Int).SetUint64(uint64(math.Pow10(18)))).String(), 10)

	fmt.Println("tcAmount float: ", tcAmountFloat)
	fmt.Println("tcAmount int: ", tcAmount)
	fmt.Println("maxAmountIntWei: ", maxAmountIntWei)

	if tcAmount.Cmp(maxAmountIntWei) > 0 {
		return nil, errors.New("max 100 TC")
	}

	// check from db:
	buyedTc := u.CheckTotalBuy(data.TCAddress)
	totalBuy := big.NewInt(0).Add(buyedTc, tcAmount)
	if totalBuy.Cmp(maxAmountIntWei) > 0 {

		leftAmount := big.NewInt(0).Sub(maxAmountIntWei, buyedTc)
		return nil, errors.New(fmt.Sprintf("Only %.4f TC left to buy.", float64(leftAmount.Uint64())/1e18))
	}

	feeRateCurrent, err := u.getFeeRateFromChain()

	if err != nil {
		return nil, err
	}
	fastestFee := feeRateCurrent.FastestFee

	// tc btc price big int:
	tcPriceFloat := big.NewFloat(TC_BTC_PRICE)
	tcPriceFloat.Mul(tcPriceFloat, new(big.Float).SetFloat64(math.Pow10(8)))

	tcPrice := new(big.Int)
	tcPriceFloat.Int(tcPrice)

	feeInfos, err := u.calBuyTcFeeInfo(tcPrice.Int64(), int64(agvFileSize), int64(fastestFee), 0, 0)
	if err != nil {
		logger.AtLog.Logger.Error("u.calMintFeeInfo.Err", zap.Error(err))
		return nil, err
	}

	fmt.Println("feeInfos: ", feeInfos)

	switch data.PayType {
	case NETWORK_BTC:
		privateKey, _, receiveAddress, err = btc.GenerateAddressSegwit()
		if err != nil {
			logger.AtLog.Logger.Error("u.ApiCreateNewGM.GenerateAddressSegwit", zap.Error(err))
			return nil, err
		}
		privateKeyEnCrypt, err = encrypt.EncryptToString(privateKey, keyToEncrypt)
		if err != nil {
			logger.AtLog.Logger.Error("u.CreateMintReceiveAddress.Encrypt", zap.Error(err))
			return nil, err
		}
	case NETWORK_ETH:

		privateKey, _, receiveAddress, err = eth.GenerateAddress()
		if err != nil {
			logger.AtLog.Logger.Error("GenerateDepositAddress.ethClient.GenerateAddress", zap.Error(err))
			return nil, err
		}
		privateKeyEnCrypt, err = encrypt.EncryptToString(privateKey, keyToEncrypt)
		if err != nil {
			logger.AtLog.Logger.Error("u.GenerateDepositAddress.Encrypt", zap.Error(err))
			return nil, err
		}

	}

	totalPaymentFloat, _ := big.NewFloat(0).SetString(feeInfos[data.PayType].TcPrice)
	totalPaymentFloat = totalPaymentFloat.Mul(totalPaymentFloat, new(big.Float).SetFloat64(data.TcAmount))

	fmt.Println("payment amount by tc amount: ", totalPaymentFloat)

	totalPaymentInt := new(big.Int)
	totalPaymentFloat.Int(totalPaymentInt)

	totalPaymentInt = big.NewInt(0).Add(totalPaymentInt, feeInfos[data.PayType].NetworkFeeBigInt)

	fmt.Println("payment amount + network by tc amount: ", totalPaymentFloat)

	if len(privateKeyEnCrypt) > 0 {

		expiredTime := utils.INSCRIBE_TIMEOUT
		if u.Config.ENV == "develop" {
			expiredTime = 1
		}
		if data.PayType == utils.NETWORK_ETH {
			expiredTime = 2 // just 1h for checking eth balance
		}

		expiredAt := time.Now().Add(time.Hour * time.Duration(expiredTime))

		newDeposit := &entity.TcGasStation{
			TcAddress:      strings.ToLower(data.TCAddress),
			PayType:        data.PayType,
			Status:         0, // pending
			ExpiredAt:      expiredAt,
			ReceiveAddress: receiveAddress, // temp address for the user send to
			PrivateKey:     privateKeyEnCrypt,
			PaymentFee:     feeInfos[data.PayType].NetworkFee, // fee by payType
			PaymentAmount:  totalPaymentInt.String(),          // fee by payType
			FeeInfo:        feeInfos,
			TcAmount:       tcAmount.String(),
		}
		err = u.Repo.InsertTcGasStation(newDeposit)
		if err != nil {
			logger.AtLog.Logger.Error("u.GenerateDepositAddress.InsertTcGasStation", zap.Error(err))
			return nil, err
		}
		return &response.GenerateDepositAddressResp{
			TCAddress:     newDeposit.TcAddress,
			TcAmount:      newDeposit.TcAmount,
			Address:       newDeposit.ReceiveAddress,
			PaymentFee:    feeInfos[data.PayType].NetworkFee, // fee by payType
			PaymentAmount: totalPaymentInt.String(),          // fee by payType
			ExpiredAt:     &newDeposit.ExpiredAt,
			FeeInfos:      feeInfos,
			PayType:       data.PayType,
		}, nil
	}

	return nil, errors.New("payType invalid")
}

func (u *Usecase) HistoryTcGasStation(address string) ([]*entity.TcGasStation, error) {
	return u.Repo.FindByTcAddress(address)
}

func (u Usecase) calBuyTcFeeInfo(mintBtcPrice, fileSize, feeRate int64, btcRate, ethRate float64) (map[string]entity.BuyTcFeeInfo, error) {

	//fmt.Println("fileSize, feeRate: ", fileSize, feeRate)

	listBuyTcFeeInfo := make(map[string]entity.BuyTcFeeInfo)

	tcPrice := big.NewInt(0)
	feeSendFund := big.NewInt(utils.FEE_BTC_SEND_AGV)

	feeSendTc := big.NewInt(0)

	totalAmountToMint := big.NewInt(0)
	netWorkFee := big.NewInt(0)

	var err error

	// cal min price:
	tcPrice = tcPrice.SetUint64(uint64(mintBtcPrice))

	if fileSize > 0 {

		calNetworkFee := int64(fileSize/4) * feeRate
		// fee mint:
		feeSendTc = big.NewInt(calNetworkFee)

	}

	fmt.Println("feeSendTc: ", feeSendTc)

	if btcRate <= 0 {
		btcRate, err = GetExternalPrice("BTC", "USDT")
		if err != nil {
			logger.AtLog.Logger.Error("getExternalPrice", zap.Error(err))
			return nil, err
		}

		ethRate, err = GetExternalPrice("ETH", "USDT")
		if err != nil {
			logger.AtLog.Logger.Error("GetExternalPrice", zap.Error(err))
			return nil, err
		}
	}

	fmt.Println("btcRate, ethRate", btcRate, ethRate)

	// total amount by BTC:
	netWorkFee = netWorkFee.Add(feeSendTc, feeSendFund) // + feeSendTc	+ feeSendFund

	totalAmountToMint = totalAmountToMint.Add(tcPrice, netWorkFee) // tcPrice, netWorkFee

	listBuyTcFeeInfo["btc"] = entity.BuyTcFeeInfo{

		TcPrice: tcPrice.String(),

		FeeRate: int(feeRate),

		InscribeFee: feeSendTc.String(),
		NetworkFee:  netWorkFee.String(),
		TotalAmount: totalAmountToMint.String(),
		SendFundFee: feeSendFund.String(),

		TcPriceBigInt:     tcPrice,
		InscribeFeeInt:    feeSendTc,
		SendFundFeeBigInt: feeSendFund,
		NetworkFeeBigInt:  netWorkFee,
		TotalAmountBigInt: totalAmountToMint,

		EthPrice: ethRate,
		BtcPrice: btcRate,

		Decimal: 8,
	}

	//fmt.Println("feeInfos[btc].TcPriceBigIn1", listBuyTcFeeInfo["btc"].TcPriceBigInt)

	// 1. convert mint price btc to eth  ==========
	// tcPriceByEth, _, _, err := u.convertBTCToETHWithPriceEthBtc(fmt.Sprintf("%f", float64(tcPrice.Uint64())/1e8), btcRate, ethRate)
	// if err != nil {
	// 	logger.AtLog.Logger.Error("calBuyTcFeeInfo.convertBTCToETHWithPriceEthBtc", zap.Error(err))
	// 	return nil, err
	// }

	// 1. set mint price by eth
	// tcPriceEth, ok := big.NewInt(0).SetString(tcPriceByEth, 10)
	// if !ok {
	// 	err = errors.New("can not set tcPriceByEth")
	// 	logger.AtLog.Logger.Error("u.calBuyTcFeeInfo.Set(TcPriceByEth)", zap.Error(err))
	// 	return nil, err
	// }

	tcPriceByEthFloat := big.NewFloat(TC_ETH_PRICE)
	tcPriceByEthFloat.Mul(tcPriceByEthFloat, new(big.Float).SetFloat64(math.Pow10(18)))

	tcPriceEth := new(big.Int)
	tcPriceByEthFloat.Int(tcPriceEth)

	// 2. convert mint fee btc to eth  ==========
	feeMintNftByEth, _, _, err := u.convertBTCToETHWithPriceEthBtc(fmt.Sprintf("%f", float64(feeSendTc.Uint64())/1e8), btcRate, ethRate)
	if err != nil {
		logger.AtLog.Logger.Error("calBuyTcFeeInfo.convertBTCToETHWithPriceEthBtc", zap.Error(err))
		return nil, err
	}
	// 2. set mint fee by eth
	feeMintNftEth, ok := big.NewInt(0).SetString(feeMintNftByEth, 10)
	if !ok {
		err = errors.New("can not set feeMintNftByEth")
		logger.AtLog.Logger.Error("u.calBuyTcFeeInfo.Set(feeMintNftByEth)", zap.Error(err))
		return nil, err
	}

	// 3. fee send master by eth:
	feeSendFundEth := big.NewInt(utils.FEE_ETH_SEND_MASTER * 1e18)

	// total amount by ETH:
	netWorkFeeEth := big.NewInt(0).Add(feeMintNftEth, feeSendFundEth) // + inscribe fee	+ feeSendFundEth

	totalAmountToMintEth := big.NewInt(0).Add(tcPriceEth, netWorkFeeEth) // tcPrice, netWorkFee

	listBuyTcFeeInfo["eth"] = entity.BuyTcFeeInfo{
		TcPrice:     tcPriceEth.String(),
		FeeRate:     int(feeRate),
		InscribeFee: feeMintNftEth.String(),
		NetworkFee:  netWorkFeeEth.String(),
		TotalAmount: totalAmountToMintEth.String(),
		SendFundFee: feeSendFundEth.String(),

		TcPriceBigInt:     tcPriceEth,
		InscribeFeeInt:    feeMintNftEth,
		SendFundFeeBigInt: feeSendFundEth,

		NetworkFeeBigInt:  netWorkFeeEth,
		TotalAmountBigInt: totalAmountToMintEth,

		EthPrice: ethRate,
		BtcPrice: btcRate,

		Decimal: 18,
	}

	return listBuyTcFeeInfo, err
}

func (u Usecase) convertBTCToETHWithPriceEthBtc(amount string, btcPrice, ethPrice float64) (string, float64, float64, error) {

	//amount = "0.1"
	powIntput := math.Pow10(8)
	powIntputBig := new(big.Float)
	powIntputBig.SetFloat64(powIntput)
	amountMintBTC, _ := big.NewFloat(0).SetString(amount)
	amountMintBTC.Mul(amountMintBTC, powIntputBig)
	// if err != nil {
	// 	logger.AtLog.Logger.Error("strconv.ParseFloat", zap.Error(err))
	// 	return "", err
	// }

	_ = amountMintBTC
	btcToETH := btcPrice / ethPrice
	// btcToETH := 14.27 // remove hardcode, why tri hardcode this??

	rate := new(big.Float)
	rate.SetFloat64(btcToETH)
	amountMintBTC.Mul(amountMintBTC, rate)

	pow := math.Pow10(10)
	powBig := new(big.Float)
	powBig.SetFloat64(pow)

	amountMintBTC.Mul(amountMintBTC, powBig)
	result := new(big.Int)
	amountMintBTC.Int(result)

	logger.AtLog.Logger.Info("convertBTCToETH", zap.String("amount", amount), zap.Float64("btcPrice", btcPrice), zap.Float64("ethPrice", ethPrice))
	return result.String(), btcPrice, ethPrice, nil
}

func (u *Usecase) CheckTotalBuy(address string) *big.Int {
	totalBuy := big.NewInt(0)
	mintReadyList, _ := u.Repo.FindByTcAddress(address)

	if mintReadyList == nil {
		return totalBuy
	}

	for _, mItem := range mintReadyList {

		if mItem.IsConfirm {
			tcAmount, _ := big.NewInt(0).SetString(mItem.TcAmount, 10)
			totalBuy = big.NewInt(0).Add(totalBuy, tcAmount)

		} else if mItem.Status == entity.StatusTcGasStation_Pending || mItem.Status == entity.StatusTcGasStation_WaitForConfirm {
			if time.Since(mItem.ExpiredAt) < 1*time.Second {
				tcAmount, _ := big.NewInt(0).SetString(mItem.TcAmount, 10)
				totalBuy = big.NewInt(0).Add(totalBuy, tcAmount)
			}
		}
	}
	return totalBuy
}
