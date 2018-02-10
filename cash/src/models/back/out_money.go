package back

//出款列表
type OutMoneyList struct {
	Id              int64   `json:"id"`                             //主键id
	SiteId          string  `json:"siteId"`                         //站点id
	SiteIndexId     string  `json:"siteIndexId"`                    //站点前台id
	AgencyAccount   string  `json:"agencyAccount"`                  //经销商账号
	MemberAccount   string  `xorm:"user_name" json:"memberAccount"` //会员账号
	IsFirst         int8    `json:"isFirst"`                        //是否首次出款
	OutwardNum      float64 `json:"outwardNum"`                     //提现额度
	Charge          float64 `json:"charge"`                         //手续费
	FavourableMoney float64 `json:"favourableMoney"`                //优惠金额
	ExpeneseMoney   float64 `json:"expeneseMoney"`                  //行政费
	OutwardMoney    float64 `json:"outwardMoney"`                   //实际出款金额
	Balance         float64 `json:"balance"`                        //账户余额
	FavourableOut   int8    `json:"favourableOut"`                  //是否优惠扣除0不是1是
	OutTime         int64   `json:"outTime"`                        //操作时间
	OutStatus       int8    `json:"outStatus"`                      //出款状态5待审核1已出款2预备出款3取消出款4拒绝出款
	DoAgencyAccount string  `json:"doAgencyAccount"`                //操作者账号
	AgencyId        int64   `json:"agencyId"`                       //操作人id
	Remark          string  `json:"remark"`                         //备注
	OutRemark       string  `json:"outRemark"`                      //出款备注（操作人填写）
	IsAutoOut       int8    `json:"isAutoOut"`                      //是否出款
	IsUnderhair     int8    `json:"isUnderhair"`                    //是否下发
	Createtime      int64   `xorm:"create_time" json:"create_time"` //操作时间
	//todo 下面的字段不是很清楚
	Audit         int8 `json:"audit"`         //稽核
	ThirdCheckout int8 `json:"thirdCheckout"` //三方下发
}

type OutMoneyBackList struct {
	TotalCount               int64          `json:"totalCount"`               //总计笔数
	TotalCharge              float64        `json:"totalCharge"`              //总计手续费
	TotalFavourableMoney     float64        `json:"totalFavourable_money"`    //总计优惠金额
	TotalExpeneseMoney       float64        `json:"totalExpeneseMoney"`       //总计行政费
	TotalOutwardMoney        float64        `json:"totalOutwardMoney"`        //总计实际总出款额
	PageTotalCount           int            `json:"pageTotalCount"`           //小计笔数
	PageTotalCharge          float64        `json:"pageTotalCharge"`          //小计手续费
	PageTotalFavourableMoney float64        `json:"pageTotalFavourableMoney"` //小计优惠金额
	PageTotalExpeneseMoney   float64        `json:"pageTotalExpeneseMoney"`   //小计行政费
	PageTotalOutwardMoney    float64        `json:"pageTotalOutwardMoney"`    //小计实际总出款额
	OutMoneyList             []OutMoneyList //出款列表
}

//当前出款手续费
type OutCharge struct {
	OutMoney float64 `json:"out_money"` //出款手续费
}
