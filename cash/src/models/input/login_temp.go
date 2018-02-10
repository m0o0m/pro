package input

type LoginAdd struct {
	Title   string `json:"title"`   //标题
	Type    int8   `json:"type"`    //类别1会员注册  2代理注册 3试玩注册
	Content string `json:"content"` //会员注册上面的文案
	Color   string `json:"color"`   //标题颜色属性
}
type LoginUpdate struct {
	Id      int64  `json:"id"`      //注册模板id
	Title   string `json:"title"`   //标题
	Type    int8   `json:"type"`    //类别1会员注册  2代理注册 3试玩注册
	Content string `json:"content"` //会员注册上面的文案
	Color   string `json:"color"`   //标题颜色属性
}
type LoginStatus struct {
	Id     int64 `json:"id"`     //注册模板id
	Status int8  `json:"status"` //注册模板状态
}
type StatusData struct {
	State int8 `json:"state"` //状态1启用 2停用
}

//注册模板详情
type LoginRegListDetailIn struct {
	Id int64 `query:"id" valid:"Min(1);ErrorCode(30041)"` //注册模板id
}
