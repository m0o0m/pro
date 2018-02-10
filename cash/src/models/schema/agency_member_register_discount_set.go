package schema

import "global"

//会员注册优惠设定
type AgencyMemberRegisterDiscountSet struct {
	SiteId      string  `xorm:"site_id"`               //站点id
	SiteIndexId string  `xorm:"site_index_id"`         //站点前台id
	AgencyId    int64   `xorm:"agency_id"`             //代理id
	Offer       float64 `xorm:"offer"`                 //加入会员赠送优惠金额
	AddMosaic   int64   `xorm:"add_mosaic"`            //优惠打码倍数
	IsIp        int8    `xorm:"is_ip"`                 //是否限制IP 1:是2:否
	CreateTime  int64   `xorm:"'create_time' created"` //创建时间
	DeleteTime  int64   `xorm:"delete_time"`           //软删除时间(为0表示未删除)
}

func (*AgencyMemberRegisterDiscountSet) TableName() string {
	return global.TablePrefix + "agency_member_register_discount_set"
}
