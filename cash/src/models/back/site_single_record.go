package back

//掉单列表
type SiteSingleRecord struct {
	Id       int64   `json:"id"`       //编号
	Username string  `json:"username"` //会员账号
	Money    float64 `json:"money"`    //交易额度
	Ctype    int64   `json:"ctype"`    //转出方
	Vtype    int64   `json:"vtype"`    //转入方
	DoTime   int64   `json:"doTime"`   //掉单时间
	Remark   string  `json:"remark"`   //备注
	Type     int8    `json:"type"`     //1表示掉单审核中，2表示审核通过，3无效申请
	Platform string  `json:"platform"` //交易平台名称
}

//掉单列表
type SiteSingleRecordBack struct {
	Id       int64   `json:"id"`       //编号
	Username string  `json:"username"` //会员账号
	Money    float64 `json:"money"`    //交易额度
	Ctype    string  `json:"ctype"`    //转出方
	Vtype    string  `json:"vtype"`    //转入方
	DoTime   int64   `json:"doTime"`   //掉单时间
	Remark   string  `json:"remark"`   //备注
	Type     int8    `json:"type"`     //交易别 1表示掉单审核中，2表示审核通过，3无效申请
}

//掉单返回列表
type SiteSingleRecordsBack struct {
	SubtotalMoney        float64                `json:"subtotalMoney"` //小计
	TotalMoney           float64                `json:"totalMoney"`    //总计
	TotalCount           int                    `json:"totalCount"`    //总计数量
	SiteSingleRecordBack []SiteSingleRecordBack `json:"siteSingleRecordBack"`
}
