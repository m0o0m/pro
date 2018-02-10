package schema

import "global"

//交易平台
type Platform struct {
	Id         int64  `xorm:"id"`
	Platform   string `xorm:"platform"`    //平台名称
	Status     int8   `xorm:"status"`      //状态
	DeleteTime int64  `xorm:"delete_time"` //删除时间
}

func (*Platform) TableName() string {
	return global.TablePrefix + "platform"
}
