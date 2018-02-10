package back

//存取款列表
type MemberCashRecord struct {
	Id            int64   `json:"id"`
	UserName      string  `json:"userName"`      // 会员账号
	AfterBalance  float64 `json:"afterBalance"`  //会员表余额
	Balance       float64 `json:"balance"`       // 交易金额
	DisBalance    float64 `json:"disBalance" `   //优惠金额
	Type          int8    `json:"type"`          //交易别1  存入    2 取出
	SourceType    int8    `json:"sourceType"`    // 数据来源类型0人工存入1公司入款2线上入款3人工取出4线上取款5出款6注册优惠7下单8额度转换
	Remark        string  `json:"remark"`        // 备注
	AgencyAccount string  `json:"agencyAccount"` // 操作人账号
	CreateTime    int64   `json:"createTime"`    //时间
}

//存取款返回列表
type MemberCashRecordBack struct {
	TotalMoney       float64            `json:"totalMoney"`    //总计金额
	SubTotalMoney    float64            `json:"subTotalMoney"` //小计金额
	TotalCount       int                `json:"totalCount"`    //总计笔数
	MemberCashRecord []MemberCashRecord `json:"memberCashRecord"`
}

//账目汇总
type CompanyIntoMoney struct {
	Money     float64 `json:"money"`              //收入/支出金额
	PeopleNum int64   `xorm:"-" json:"peopleNum"` //交易人数
	Count     int64   `xorm:"-" json:"count"`     //交易笔数
}

//汇总
type Summary struct {
	ProfitLoss    float64 `json:"profitLoss"`    //实际盈亏（元）
	DepositMoney  float64 `json:"depositMoney"`  //入款总额（元）
	DepositPeople int64   `json:"depositPeople"` //入款人数（人）
	DepositCount  int64   `json:"depositCount"`  //入款次数（次）
	PayoutMoney   float64 `json:"payoutMoney"`   //出款总额（元）
	PayoutPeople  int64   `json:"payoutPeople"`  //出款人数（人）
	PayoutCount   int64   `json:"payoutCount"`   //出款次数（次）
}
