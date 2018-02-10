package back

//额度转换
type TransferCreditResult struct {
	Result bool `json:"result"`
	Data   struct {
		Code    int64   `json:"code" `    //code码
		Balance float64 `json:"balance" ` //余额
	} `json:"data"`
}

//跳转游戏
type GameResult struct {
	Result bool `json:"result"`
	Data   struct {
		Code     int64  `json:"code" `    //code码
		LoginUrl string `json:"loginUrl"` //余额
	} `json:"data"`
}
