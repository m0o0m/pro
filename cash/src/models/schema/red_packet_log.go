package schema

import "global"

//红包日志表
type RedPacketLog struct {
	Id             int64   `xorm:"'id' PK autoincr"`
	SetId          int64   `xorm:"set_id"`                //红包设置id
	Uuid           string  `xorm:"uuid"`                  //红包的uuid
	Money          float64 `xorm:"money"`                 //红包金额
	MemberId       int64   `xorm:"member_id"`             //用户id
	Account        string  `xorm:"account"`               //用户名
	SiteId         string  `xorm:"site_id"`               //站点id
	SiteIndexId    string  `xorm:"site_index_id"`         //前台id
	CreateIp       string  `xorm:"create_ip"`             //创建ip
	CreateTime     int64   `xorm:"'create_time' created"` //创建时间
	StartTime      int64   `xorm:"start_time"`            //开始时间
	EndTime        int64   `xorm:"end_time"`              //结束时间
	InStartTime    int64   `xorm:"in_start_time"`         //存款起始时间
	InEndTime      int64   `xorm:"in_end_time"`           //存款结束时间
	InSum          float64 `xorm:"in_sum"`                //存款额度
	AuditStartTime int64   `xorm:"audit_start_time"`      //有效打码开始时间
	AuditEndTime   int64   `xorm:"audit_end_time"`        //有效打码结束时间
	BetSum         float64 `xorm:"bet_sum"`               //有效打码量
	MinMoney       float64 `xorm:"min_money"`             //红包最小额度
	LevelId        string  `json:"level_id"`              //可参加活动会员的分组，0为无限制
	Title          string  `xorm:"title"`                 //红包名
	MakeSure       int64   `xorm:"make_sure"`             //是否被抢1,未被抢,2已抢
	PType          int64   `xorm:"p_type"`                //客户端类型0pc 1wap 2android 3ios
	BalanceMoney   float64 `xorm:"balance_money"`
	LevelEs        string  `xorm:"level_es"`
	IsIp           int64   `xorm:"is_ip"` //1为限制ip,2为不限制
	Finish         int64   `xorm:"finish"`
}

func (*RedPacketLog) TableName() string {
	return global.TablePrefix + "red_packet_log"
}
