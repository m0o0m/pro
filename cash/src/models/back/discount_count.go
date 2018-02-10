package back

//报表统计 - 优惠统计
type DiscountCountList struct {
	ID            int64   `xorm:"id" json:"id" `                        //id
	SiteId        string  `xorm:"site_id" json:"siteId" `               //站点
	AgencyId      int64   `xorm:"agency_id" json:"agencyId" `           //代理id
	MemberId      int64   `xorm:"member_id" json:"memberId" `           //会员id
	Account       string  `xorm:"account" json:"account" `              //账号
	Num           int64   `xorm:"num" json:"num" `                      //笔数
	DiscountMoney float64 `xorm:"discount_money" json:"discountMoney" ` //优惠金额
	DayType       string  `xorm:"day_type" json:"dayType" `             //统计的的开始时间
}

//报表统计 - 优惠统计
type DiscountCountTotal struct {
	SubNum           int64                `xorm:"-"  json:"subNum" `                     //小计笔数
	SubDiscountMoney float64              `xorm:"-"  json:"subDiscountMoney" `           //小计金额
	Num              int64                `xorm:"num"  json:"num" `                      //总计笔数
	DiscountMoney    float64              `xorm:"discount_money"  json:"discountMoney" ` //总计金额
	Content          *[]DiscountCountList `xorm:"-"  json:"content" `                    //详情
}
