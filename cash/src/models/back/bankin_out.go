//出入款查询[admin]
package back

//公司入款
type InDepositBack struct {
	Id              int64   `xorm:"id" json:"id"`
	SiteId          string  `xorm:"site_id" json:"siteId"`                   //操作站点id
	SiteIndexId     string  `xorm:"site_index_id" json:"siteIndexId"`        // 站点前台id
	OrderNum        string  `xorm:"order_num" json:"orderNum"`               // 订单号
	Account         string  `xorm:"account" json:"account"`                  //会员账号
	DepositUsername string  `xorm:"deposit_username" json:"depositUsername"` // 存款人名称
	DepositMoney    float64 `xorm:"deposit_money" json:"depositMoney"`       // 存入金额
	DepositDiscount float64 `xorm:"deposit_discount" json:"depositDiscount"` // 存入优惠
	OtherDiscount   float64 `xorm:"other_discount" json:"otherDiscount"`     // 其他优惠
	DepositCount    float64 `xorm:"deposit_count" json:"depositCount"`       // 存入总额
	ClientType      int8    `xorm:"client_type" json:"clientType"`           // 客户端类型 1pc 2wap 3android 4ios
	Status          int8    `xorm:"status" json:"status"`                    // 状态,0未处理1已确认2已取消
	DoAgency        string  `xorm:"username" json:"doAgency"`                // 操作人
	UpdateTime      int64   `xorm:"update_time" json:"updateTime"`           // 操作人操作时间
	DepositMethod   int8    `xorm:"deposit_method" json:"depositMethod"`     //存款类型
}

//线上入款
type InDepositOnlineBack struct {
	Id                   int64   `xorm:"id" json:"id"`
	ThirdOrderNumber     string  `xorm:"third_order_number" json:"orderNum"`          //第三方平台订单号
	SourceDeposit        int     `xorm:"source_deposit" json:"clientType"`            //入款来源(1.pc 2.wap)
	LocalOperateTime     int64   `xorm:"create_time" json:"updateTime"`               //本地平台操作时间(本平台点击跳转到三方平台付账时候的时间,也是前台查询时间的搜索条件)
	AmountDeposit        float64 `xorm:"amount_deposit" json:"depositMoney"`          //存款金额
	MemberAccount        string  `xorm:"member_account" json:"account"`               //存款会员账号
	SiteIndexId          string  `xorm:"site_index_id" json:"siteIndexId"`            //前台站点id
	SiteId               string  `xorm:"site_id" json:"siteId"`                       //站点Id
	Status               int     `xorm:"status" json:"status"`                        //状态(1.未支付2.已经支付3.已取消4已确认)
	DepositDiscount      float64 `xorm:"deposit_discount" json:"depositDiscount"`     //存款优惠金额
	OtherDepositDiscount float64 `xorm:"other_deposit_discount" json:"otherDiscount"` //其他优惠金额
	OperateId            int     `xorm:"operate_id" json:"operateId"`                 //操作者id
	ThirdIdTitle         string  `xorm:"title" json:"thirdId"`                        //第三方支付平台
}

//出款查询
type OutDepositBack struct {
	UserName        string  `xorm:"user_name" json:"userName"`               //会员账号
	ClientType      int     `xorm:"client_type" json:"clientType"`           //客户端类型0pc 1wap 2android 3ios
	OutwardNum      float64 `xorm:"outward_num" json:"outwardNum"`           //提出金额
	Charge          float64 `xorm:"charge" json:"charge"`                    //手续费
	FavourableMoney float64 `xorm:"favourable_money" json:"favourableMoney"` //优惠金额
	ExpeneseMoney   float64 `xorm:"expenese_money" json:"expeneseMoney"`     //行政费
	OutwardMoney    float64 `xorm:"outward_money" json:"outwardMoney"`       //实际出款金额
	Balance         float64 `xorm:"balance" json:"balance"`                  //账户余额
	FavourableOut   int     `xorm:"favourable_out" json:"favourableOut"`     //是否优惠扣除0不是1是
	OutTime         int64   `xorm:"out_time" json:"outTime"`                 //出款时间
	DoAgencyName    string  `xorm:"username" json:"username"`                //操作人
	OutStatus       int     `xorm:"out_status" json:"outStatus"`             //出款状态1已出款2预备出款3取消出款4拒绝出款5待审核
	NormalMoney     float64 `xorm:"normal_money" json:"normalMoney"`         //常态稽核金额
	MultipleMoney   float64 `xorm:"multiple_money" json:"multipleMoney"`     //综合稽核金额
}
