package schema

import "global"

//会员返佣期数(事件)
type MemberRetreatCommission struct {
	Id          int64  `xorm:"id"`
	SiteId      string `xorm:"site_id"`       //操作站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	StartTime   int64  `xorm:"start_time"`    //开始时间
	EndTime     int64  `xorm:"end_time"`      //结束时间
	Status      int8   `xorm:"status"`        //退佣状态 0未返佣1已返佣
	CreateTime  int64  `xorm:"create_time"`   //创建时间
	DeleteTime  int64  `xorm:"delete_time"`   //删除时间
}

func (*MemberRetreatCommission) TableName() string {
	return global.TablePrefix + "member_retreat_commission"
}
