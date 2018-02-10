package schema

import "global"

//自助返水查询
type MemberRetreatWaterSelf struct {
	Id          int64   `xorm:"id PK autoincr"`
	SiteId      string  `xorm:"site_id"`       //操作站点id
	SiteIndexId string  `xorm:"site_index_id"` //站点前台id
	AdminUser   string  `xorm:"admin_user"`    //操作者
	MemberId    string  `xorm:"member_id"`     //会员id
	Account     string  `xorm:"account"`       //会员账号
	OrderNum    int64   `xorm:"order_num"`     //订单号，反水时间，年月日
	CreateTime  int64   `xorm:"create_time"`   //创建时间就是反水时间
	Betting     float64 `xorm:"betting"`       //总有效投注
	Money       float64 `xorm:"money"`         //反水金额
}

func (*MemberRetreatWaterSelf) TableName() string {
	return global.TablePrefix + "member_retreat_water_self"
}
