package back

type CsGame struct {
	Id       int64  `xorm:"'id'" json:"id"`
	CsName   string `xorm:"'cs_name' default('ull')" json:"cs_name"`       //彩票名字
	CsType   string `xorm:"'cs_type' default('ull')" json:"cs_type"`       //彩票类型
	CsState  int64  `xorm:"'cs_state' default(1)" json:"cs_state"`         //开关，默认开启
	CsLxType string `xorm:"'cs_lx_type' default('ull')" json:"cs_lx_type"` //彩票分类
	CsHref   string `xorm:"'cs_href' default('ull')" json:"cs_href"`       //卡司彩票链接
	CsSort   int64  `xorm:"'cs_sort' default(ull)" json:"cs_sort"`         //排序字段
	Hot      int64  `xorm:"'hot' default(ull)" json:"hot"`                 //卡司彩票类型
	CsColumn string `xorm:"'cs_column' default('ull')" json:"cs_column"`   //栏目
}

type FcGameList struct {
	GameName string `json:"game_name"` //游戏名
	GameType string `json:"game_type"` //游戏类型
}
