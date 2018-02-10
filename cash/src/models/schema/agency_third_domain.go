package schema

import "global"

//代理推广域名
type AgencyThirdDomain struct {
	Id         int64  `xorm:"'id' PK autoincr"`      //主键id
	AgencyId   int64  `xorm:"agency_id"`             //代理id
	Domain     string `xorm:"domain"`                //推广域名（唯一判断）
	CreateTime int64  `xorm:"'create_time' created"` //添加时间
	DeleteTime int64  `xorm:"delete_time"`           //删除时间
}

func (*AgencyThirdDomain) TableName() string {
	return global.TablePrefix + "agency_third_domain"
}
