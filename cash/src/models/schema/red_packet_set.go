package schema

import "global"

//红包活动设定表
type RedPacketSet struct {
	Id             int64   `xorm:"'id' PK autoincr"`
	Title          string  `xorm:"title"`                 //活动标题
	SiteId         string  `xorm:"site_id"`               //站点id
	SiteIndexId    string  `xorm:"site_index_id"`         //前台id
	Description    string  `xorm:"description"`           //活动简介
	MaxCount       int64   `xorm:"max_count"`             //每人最多获奖次数
	StartTime      int64   `xorm:"start_time"`            //活动开始时间
	EndTime        int64   `xorm:"end_time"`              //活动结束时间
	InStartTime    int64   `xorm:"in_start_time"`         //存款起始时间
	InEndTime      int64   `xorm:"in_end_time"`           //存款结束时间
	InSum          float64 `xorm:"in_sum"`                //存款额度
	AuditStartTime int64   `xorm:"audit_start_time"`      //有效打码起始时间
	AuditEndTime   int64   `xorm:"audit_end_time"`        //有效打码结束时间
	BetSum         float64 `xorm:"bet_sum"`               //有效打码量
	EndTitle       string  `xorm:"end_title"`             //结束标题
	EndDescription string  `xorm:"end_description"`       //活动结束简介
	LevelId        string  `xorm:"level_id"`              //可参加活动会员的分组，0为无限制
	TotalMoney     float64 `xorm:"total_money"`           //红包总额
	MinMoney       float64 `xorm:"min_money"`             //红包最小额度
	RedNum         int64   `xorm:"red_num"`               //红包数量
	CreateIp       string  `xorm:"create_ip"`             //创建ip
	CreateUid      int64   `xorm:"create_uid"`            //创建管理员的uid
	CreateTime     int64   `xorm:"'create_time' created"` //活动创建时间
	IsIp           int64   `xorm:"is_ip"`                 //1为限制ip；2为不限制
	StyleId        int64   `xorm:"style_id"`              //红包皮肤,关联红包皮肤表red_packet_set
	IsShow         int64   `xorm:"is_show"`               //1为展示，2为不展示
	AppointMoney   int64   `xorm:"appoint_money"`         //额外领取存款门槛
	RedType        int64   `xorm:"red_type"`              //红包类型，1拼手气，2普通
	Status         int64   `xorm:"status"`                //活动状态 ，1未开始，2进行中，3已经结束，4已删除
	IsGenerate     int64   `xorm:"is_generate"`           //是否生成 1,未生成,2,已生成',
	DepositAchieve float64 `xorm:"deposit_achieve"`       //存款达到
	ReceiveAgain   int8    `xorm:"receive_again"`         //是否能重复领取红包1是2否
}

func (*RedPacketSet) TableName() string {
	return global.TablePrefix + "red_packet_set"
}
