package schema

import "global"

//三方对接验证加密表
type ApiClients struct {
	Id        int64  `xorm:"id"`
	UserId    int64  `xorm:"user_id"`    //客户id
	Name      string `xorm:"name"`       //名称
	Secret    string `xorm:"secret"`     //证书
	Revoked   int8   `xorm:"revoked"`    //1开启，2关闭
	CreatedAt int64  `xorm:"created_at"` //创建时间
	UpdatedAt int64  `xorm:"updated_at"` //更新时间
	SiteId    string `xorm:"site_id"`    //站点id
}

func (*ApiClients) TableName() string {
	return global.TablePrefix + "api_clients"
}
