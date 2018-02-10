package back

//额度统计列表
type QuotaCountListBack struct {
	Num        int64   `xorm:"a" json:"num"`
	Proportion float64 `xorm:"proportion" json:"proportion"`
	Platform   string  `xorm:"platform" json:"platform"`
	Money      float64 `xorm:"b" json:"money"`
}

//额度统计最终列表
type AllQuotaCountListBack struct {
	TurnNum      int64   `json:"outNumber"`     //转出笔数
	RetreatNum   int64   `json:"inNumber"`      //转入笔数
	Proportion   float64 `json:"proportionate"` //抽成比例
	Platform     string  `json:"type"`          //类型名称
	TurnMoney    float64 `json:"vedioOut"`      //转出金额
	RetreatMoney float64 `json:"vedioIn"`       //转入金额
	WalletAdd    float64 `json:"walletAdd"`     //钱包增加额度
	WalletReduce float64 `json:"walletReduce"`  //钱包减少额度
	ResultMoney  float64 `json:"resultMoney"`   //结果额度
	ResultRatio  float64 `json:"resultRatio"`   //结果比例
}

//总列表
type AllList struct {
	Data  []AllQuotaCountListBack `json:"data"`
	Total QuotaTotal              `json:"total"`
}

//总计
type QuotaTotal struct {
	TurnMoneyTotal    float64 `json:"turnMoneyToatl"`    //转出总计
	RetreatMoneyTotal float64 `json:"retreatMoney"`      //转入总计
	TurnNumTotal      int64   `json:"turnNumTotal"`      //转出次数总计
	RetreatNumTotal   int64   `json:"retreatNumTotal"`   //转入次数总计
	WalletAddTotal    float64 `json:"walletAddTotal"`    //钱包增加额度
	WalletReduceTotal float64 `json:"walletReduceTotal"` //钱包减少额度
	ResultMoneyTotal  float64 `json:"resultMoneyTotal"`  //结果额度
	ResultRatioTotal  float64 `json:"resultRatioTotal"`  //结果比例
}

//充值记录
type QuotaReBack struct {
	Id          int64   `xorm:"id" json:"id"`
	SiteId      string  `xorm:"site_id" json:"siteId"`            //站点ID
	SiteIndexId string  `xorm:"site_index_id" json:"siteIndexId"` //多站点ID
	OrderNum    string  `xorm:"order_num" json:"orderNum"`        // 订单号
	Money       float64 `xorm:"money" json:"money"`               // 交易额度
	Type        int8    `xorm:"type" json:"type"`                 // 支付方式  1第三方入款，2公司入款
	UpdataTime  int64   `xorm:"update_time" json:"updataTime"`    // 支付时间
	State       int8    `xorm:"state" json:"state"`               // 状态1未支付2支付
	Remark      string  `xorm:"remark" json:"remark"`             // 备注
}

//视讯平台
type GetPlatform struct {
	Id       int64  `json:"id"`
	Platform string `json:"platform"` //视讯
}
