package schema

import "global"

//视讯额度掉单申请表
type SiteVideoMember struct {
	Id              int64   `xorm:"id"`
	Account         string  `xorm:"account"`             //登录账号
	GUsername       string  `xorm:"g_username"`          //视讯账号
	Password        string  `xorm:"password"`            //登录密码
	Token           string  `xorm:"token"`               //
	Balance         float64 `xorm:"balance"`             //视讯余额
	Platform        string  `xorm:"platform"`            //视讯平台
	MemberId        int64   `xorm:"member_id"`           //会员id
	SiteId          string  `xorm:"site_id"`             //
	AgentId         int64   `xorm:"agent_id"`            //
	IndexId         string  `xorm:"index_id"`            //
	UaId            int64   `xorm:"ua_id"`               //
	ShId            int64   `xorm:"sh_id"`               //
	RegIp           string  `xorm:"reg_ip"`              //
	LastIp          string  `xorm:"last_ip"`             //
	Cur             string  `xorm:"cur"`                 //币种
	CreateTime      int64   `xorm:"create_time created"` //申请时间
	UpdateTime      int64   `xorm:"update_time"`         //操作时间
	LastBalanceTime int64   `xorm:"last_balance_time"`   //最后更新余额时间
	Limit           string  `xorm:"limit"`               //限额
	Status          int8    `xorm:"status"`              //状态 1正常 2停用
}

func (*SiteVideoMember) TableName() string {
	return global.TablePrefix + "site_video_member"
}
