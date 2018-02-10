package schema

import "global"

//EG彩票游戏
type EgGames struct {
	Id       int64  `xorm:"'id' PK autoincr notnull" json:"id"`
	EgName   string `xorm:"'eg_name' default('ull')" json:"eg_name"`       //彩票名字
	EgType   string `xorm:"'eg_type' default('ull')" json:"eg_type"`       //彩票类型
	EgState  int64  `xorm:"'eg_state' default(1)" json:"eg_state"`         //开关，默认开启
	EgLxType string `xorm:"'eg_lx_type' default('ull')" json:"eg_lx_type"` //彩票分类
	EgHref   string `xorm:"'eg_href' default('ull')" json:"eg_href"`       //eg彩票链接
	EgSort   int64  `xorm:"'eg_sort' default(ull)" json:"eg_sort"`
	Hot      int64  `xorm:"'hot' default(0)" json:"hot"` //eg热门彩票
}

func (*EgGames) TableName() string {
	return global.TablePrefix + "eg_games"
}
