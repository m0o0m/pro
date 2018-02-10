package back

//视讯ip白名单
type GameWhiteList struct {
	Id      int    `json:"id"`      //序号
	Ip      string `json:"ip"`      //ip
	Remarks string `json:"remarks"` //备注
}
