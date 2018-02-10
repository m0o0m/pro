package back

//额度记录列表
type QuotaRecord struct {
	Id          int64   `json:"id"`
	SiteId      string  `json:"site_id"`       //站点id
	SiteIndexId string  `json:"site_index_id"` //站点前台id
	AdminName   string  `json:"adminName"`     //会员账号
	Money       float64 `json:"money"`         //交易额度
	Balance     float64 `json:"balance"`       //站点视讯余额
	CashType    int8    `json:"cashType"`      //1额度转换  2额度加款  3额度扣款  4预借   5业主充值
	Platform    string  `json:"platform"`      //视讯类型名
	DoType      int8    `json:"doType"`        //操作类型 1存入 2 取出
	State       int8    `json:"status"`        //状态1待审核 2正常  3掉单
	CreateTime  int64   `json:"createTime"`    //操作时间
	Remark      string  `json:"remark"`        //备注
}

//额度记录返回列表
type QuotaRecordBack struct {
	SubtotalMoney float64       `json:"subtotalMoney"` //小计
	TotalMoney    float64       `json:"totalMoney"`    //总计
	TotalCount    int           `json:"totalCount"`    //总计数量
	QuotaRecord   []QuotaRecord `json:"quotaRecord"`   //额度记录列表
}
