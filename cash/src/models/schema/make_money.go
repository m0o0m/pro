package schema

import "global"

//会员出款管理
type MakeMoney struct {
	Id              int64   `xorm:"id"`
	SiteId          string  `xorm:"site_id"`             //操作站点id
	SiteIndexId     string  `xorm:"site_index_id"`       //站点前台id
	MemberId        int64   `xorm:"member_id"`           //会员id
	UserName        string  `xorm:"user_name"`           //会员账号
	LevelId         string  `xorm:"level_id"`            //会员所属层级
	AgencyId        int64   `xorm:"agency_id"`           //会员所属代理id
	AgencyAccount   string  `xorm:"agency_account"`      //会员所属代理账号
	IsFirst         int8    `xorm:"is_first"`            //是否首次出款,0不是1是
	OutwardNum      float64 `xorm:"outward_num"`         //提出金额
	Charge          float64 `xorm:"charge"`              //手续费
	FavourableMoney float64 `xorm:"favourable_money"`    //优惠金额
	ExpeneseMoney   float64 `xorm:"expenese_money"`      //行政费
	OutwardMoney    float64 `xorm:"outward_money"`       //实际出款金额
	Balance         float64 `xorm:"balance"`             //账户余额
	FavourableOut   int     `xorm:"favourable_out"`      //是否优惠扣除0不是1是
	OutStatus       int     `xorm:"out_status"`          //出款状态1已出款2预备出款3取消出款4拒绝出款5待审核
	Remark          string  `xorm:"remark"`              //会员备注
	OutRemark       string  `xorm:"out_remark"`          //出款备注,管理员填写
	DoAgencyId      int64   `xorm:"do_agency_id"`        //操作人id(agency表主键)
	DoAgencyAccount string  `xorm:"do_agency_account"`   //操作人账号
	CreateTime      int64   `xorm:"create_time created"` //提出时间
	DoUrl           string  `xorm:"do_url"`              //会员提交网址
	ClientType      int     `xorm:"client_type"`         //客户端类型1pc 2wap 3android 4ios
	OutTime         int64   `xorm:"out_time"`            //出款时间
	IsAutoOut       int8    `xorm:"is_auto_out"`         //是否自动出款
	IsUnderhair     int8    `xorm:"is_underhair"`        //是否下发
	OrderNum        string  `xorm:"order_num"`           //订单号
	//下面两个字段不确定
	Audit         int8 `xorm:"audit"`          //稽核
	ThirdCheckout int8 `xorm:"third_checkout"` //三方下发
}

func (*MakeMoney) TableName() string {
	return global.TablePrefix + "make_money"
}
