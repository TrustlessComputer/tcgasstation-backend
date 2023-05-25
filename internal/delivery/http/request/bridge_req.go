package request

type GenerateDepositAddressReq struct {
	TCAddress string `json:"tcAddress"`
	PayType   string `json:"payType"`
	TcAmount  string `json:"tcAmount"`
}
