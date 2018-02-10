package schema

import "global"

//会员退水期数(事件)
type MemberRetreatWater struct {
	Id          int64   `xorm:"id"`
	SiteId      string  `xorm:"site_id"`       //操作站点id
	SiteIndexId string  `xorm:"site_index_id"` //站点前台id
	AdminUser   string  `xorm:"admin_user"`    //操作者
	LevelId     string  `xorm:"level_id"`      //层级,用逗号分隔
	StartTime   int64   `xorm:"start_time"`    //开始时间
	EndTime     int64   `xorm:"end_time"`      //结束时间
	CreateTime  int64   `xorm:"create_time"`   //创建时间
	Event       string  `xorm:"event"`         //事件
	NoPeopleNum int64   `xorm:"no_people_num"` //冲销人数
	PeopleNum   int64   `xorm:"people_num"`    //退水人数
	Money       float64 `xorm:"money"`         //金额'
	Bet         float64 `xorm:"bet"`           //综合打码倍数
}

func (*MemberRetreatWater) TableName() string {
	return global.TablePrefix + "member_retreat_water"
}
