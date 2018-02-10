package schema

import "global"

//开户人及下属账号
type Agency struct {
	Id              int64   `xorm:"'id' PK autoincr"`      //主键id
	SiteId          string  `xorm:"site_id"`               //站点id
	SiteIndexId     string  `xorm:"site_index_id"`         //站点前台id
	ParentId        int64   `xorm:"parent_id"`             //所属上级
	RoleId          int64   `xorm:"role_id"`               //所属角色id
	Account         string  `xorm:"account"`               //登录账号
	LoginKey        string  `xorm:"login_key"`             //登录之后的token
	IsLogin         int8    `xorm:"is_login"`              //是否在线
	Password        string  `xorm:"password"`              //登录密码
	OperatePassword string  `xorm:"operate_password"`      //操作密码
	LoginErrCount   int64   `xorm:"login_err_count"`       //登录错误次数
	LoginCount      int64   `xorm:"login_count"`           //登录次数
	LoginIp         string  `xorm:"login_ip"`              //登录IP
	LoginTime       int64   `xorm:"login_time"`            //登录时间
	LastLoginIp     string  `xorm:"last_login_ip"`         //上次登录IP
	LastLoginTime   int64   `xorm:"last_login_time"`       //上次登录时间
	Username        string  `xorm:"username"`              //代理名称
	Remark          string  `xorm:"remark"`                //备注
	Level           int8    `xorm:"level"`                 //账号等级(1:开户人2:股东3:总代理4:代理)
	IsSub           int8    `xorm:"is_sub"`                //是否子账号 1是2否
	IsDefault       int8    `xorm:"is_default"`            //是否等级里默认账号 1是2否
	VideoBalance    float64 `xorm:"video_balance"`         //视讯余额
	Status          int8    `xorm:"status"`                //状态 1正常2禁用
	CreateTime      int64   `xorm:"'create_time' created"` //创建时间
	DeleteTime      int64   `xorm:"delete_time"`           //软删除时间(为0表示未删除)
}

func (*Agency) TableName() string {
	return global.TablePrefix + "agency"
}
