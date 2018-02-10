package back

//报表统计 - 入款统计
type CashCountList struct {
	ID             int64   `xorm:"id" json:"id" `                           //id
	SiteId         string  `xorm:"site_id" json:"siteId" `                  //站点
	FirstAgencyId  int64   `xorm:"first_agency_id" json:"firstAgencyId"`    //股东id
	SecondAgencyId int64   `xorm:"second_agency_id" json:"secondAgencyId" ` //总代id
	AgencyId       int64   `xorm:"agency_id" json:"agencyId" `              //代理id
	MemberId       int64   `xorm:"member_id" json:"memberId" `              //会员id
	Account        string  `xorm:"account" json:"account" `                 //账号
	Num            int64   `xorm:"num" json:"num" `                         //笔数
	CashMoney      float64 `xorm:"cash_money" json:"cashMoney" `            //入款金额
	DayType        string  `xorm:"day_type" json:"dayType" `                //统计的的开始时间
	IntoStyle      int8    `xorm:"into_style" json:"intoStyle"`             //入款方式
}

//报表统计 - 入款统计
type CashCountTotal struct {
	SubNum       int64            `xorm:"-"  json:"subNum" `            //小计笔数
	SubCashMoney float64          `xorm:"-" json:"subCashMoney" `       //小计金额
	Num          int64            `xorm:"num"  json:"num" `             //总计笔数
	CashMoney    float64          `xorm:"cash_money" json:"cashMoney" ` //总计金额
	Content      *[]CashCountList `xorm:"-"  json:"content" `           //详情
}
