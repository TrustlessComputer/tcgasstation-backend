package request

type UpsertRedisRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ListtokenIDsReq struct {
	InscriptionID []string `json:"inscriptionIDs"`

	SellOrdAddress string            `json:"seller_ord_address"`
	SellerAddress  string            `json:"seller_address"`
	Price          string            `json:"price"`
	PayType        map[string]string `bson:"payType"`
}

type UpdateCronJobStatusRequest struct {
	JobKey string `json:"jobKey"`
	Enable bool   `json:"enable"`
}

type UpdateCronJobStatusByFuncNameRequest struct {
	FuncName string `json:"funcName"`
	Enable   bool   `json:"enable"`
}

type GetCronJobInfoRequest struct {
	JobKey string `json:"jobKey"`
}

type RescanDepositTxs struct {
	From uint64 `json:"from"`
	To   uint64 `json:"to"`
}

type FilterDepositWithdrawReq struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	TokenID string `json:"tokenID"`
}
