package schema

import "global"

//注册文案模板
type InfoRegStyle struct {
	Id      int64  `xorm:"id"`      //注册模板id
	Title   string `xorm:"title"`   //标题
	Type    int8   `xorm:"type"`    //类别1会员注册  2代理注册 3试玩注册
	Content string `xorm:"content"` //会员注册上面的文案
	Color   string `xorm:"color"`   //标题颜色属性
	State   int8   `xorm:"state"`   //状态1启用 2停用
}

func (*InfoRegStyle) TableName() string {
	return global.TablePrefix + "info_reg_style"
}
