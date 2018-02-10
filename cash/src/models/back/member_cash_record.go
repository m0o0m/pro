package back

//现金记录列表返回
type GetCashRecordList struct {
	Id           int64   `json:"id"`
	SiteId       string  `json:"siteId"`          //操作站点id
	UserName     string  `json:"userName"`        //会员账号
	AgencyId     int64   `json:"agencyId"`        //会员所属代理id
	SourceType   int64   `json:"sourceType"`      //数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换9优惠返水10自助返水11会员返佣
	Type         int64   `json:"type"`            //1.存入2.取出
	TradeNo      string  `json:"tradeNo"`         //下单单号
	Balance      float64 `json:"balance"`         //金额
	DisBalance   float64 `json:"disBalance"`      //优惠金额
	AfterBalance float64 `json:"afterBalance"`    //操作后余额
	Remark       string  `json:"remark"`          //备注
	ClientType   int64   `json:"clientType"`      //客户端类型1pc 2wap 3android 4ios
	CreateTime   int64   `json:"createTime"`      //添加时间
	DeleteTime   int64   `json:"deleteTime"`      //删除时间（0为正常  否则就是隐藏）
	Status       int8    `xorm:"-" json:"status"` //状态1正常2不正常
}

//总计，小计
type CashRecordListTotalBack struct {
	GrandTotal      int64   `xorm:"b" json:"grandTotal"`      //总计笔数
	GrandTotalMoney float64 `xorm:"a" json:"grandTotalMoney"` //总计金额
	SmallTotalMoney float64 `xorm:"-" json:"smallTotalMoney"` //小计
}

//现金记录返回
type CashRecordAllBack struct {
	GetCashRecordList       []GetCashRecordList     `json:"getCashRecordList"`
	CashRecordListTotalBack CashRecordListTotalBack `json:"cashRecordListTotalBack"`
}
