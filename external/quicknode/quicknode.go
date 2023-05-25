package quicknode

import (
	"encoding/json"
	"tcgasstation-backend/internal/usecase/structure"
	"tcgasstation-backend/utils/config"
	"tcgasstation-backend/utils/helpers"
	"tcgasstation-backend/utils/logger"
	"tcgasstation-backend/utils/redis"

	"go.uber.org/zap"
)

type QuickNode struct {
	conf      *config.Config
	serverURL string
	cache     redis.IRedisCache
}

func NewQuickNode(conf *config.Config, cache redis.IRedisCache) *QuickNode {
	return &QuickNode{
		conf:      conf,
		serverURL: conf.QuickNode,
		cache:     cache,
	}
}

func (q QuickNode) AddressBalance(walletAddress string) ([]WalletAddressBalanceResp, error) {
	headers := make(map[string]string)
	reqBody := RequestData{
		Method: "qn_addressBalance",
		Params: []string{
			walletAddress,
		},
	}

	data, _, _, err := helpers.HttpRequest(q.serverURL, "POST", headers, reqBody)
	if err != nil {
		return nil, err
	}

	resp := []WalletAddressBalanceResp{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q QuickNode) GetListUTXOs(address string, minConfirm, maxConfirm uint64) ([]*structure.UTXO, error) {
	utxoList, err := q.AddressBalance(address)
	if err != nil {
		logger.AtLog.Error("GetListUTXOs", "Get list utxo error", zap.Any("error", err))
		return nil, err
	}

	result := []*structure.UTXO{}

	// get current block height to check confirmation blocks
	currentHeight, err := q.GetBlockCountFromQuickNode()
	if err != nil {
		logger.AtLog.Error("GetListUTXOs", "Get current btc block height error", zap.Any("error", err))
		return nil, err
	}
	for _, utxo := range utxoList {
		confirmationBlocks := currentHeight - uint64(utxo.Height) + 1
		if confirmationBlocks < minConfirm || confirmationBlocks > maxConfirm {
			continue
		}

		result = append(result, &structure.UTXO{
			TxHash:    utxo.Hash,
			TxOutputN: utxo.Index,
			Value:     uint64(utxo.Value),
		})

	}
	return result, nil
}

func (q QuickNode) GetBlockCountFromQuickNode() (uint64, error) {
	type getBlockCountResp struct {
		Result uint64 `json:"result"`
		Error  error  `json:"error"`
	}

	headers := make(map[string]string)
	reqBody := RequestData{
		Method: "getblockcount",
		Params: []string{},
	}

	data, _, _, err := helpers.HttpRequest(q.serverURL, "POST", headers, reqBody)
	if err != nil {
		return 0, err
	}

	resp := getBlockCountResp{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return 0, err
	}

	return resp.Result, nil
}
