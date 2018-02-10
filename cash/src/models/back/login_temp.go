package back

type LoginTemp struct {
	Id      int64  `json:"id"`      //注册模板id
	Title   string `json:"title"`   //标题
	Type    int8   `json:"type"`    //类别1会员注册  2代理注册 3试玩注册
	Content string `json:"content"` //会员注册上面的文案
	Color   string `json:"color"`   //标题颜色属性
	State   int8   `json:"state"`   //状态1启用 2停用
}

//注册模板详情
type LoginRegDetailBack struct {
	Id      int64  ` xorm:"id" json:"id"`          //注册模板id
	Title   string `xorm:"title" json:"title"`     //标题
	Type    int8   `xorm:"type" json:"type"`       //类别1会员注册  2代理注册 3试玩注册
	Content string `xorm:"content" json:"content"` //会员注册上面的文案
	Color   string `xorm:"color" json:"color"`     //标题颜色属性
}
