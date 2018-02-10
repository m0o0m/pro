package schema

import "global"

//pk彩票彩种表
type PkGames struct {
	Id      int64  `xorm:"'id' PK autoincr notnull"`
	Name    string `xorm:"'name' default('ull')"`
	Type    string `xorm:"'type' default('0')"`
	State   int64  `xorm:"'state' default(1)"`
	KyQishu int64  `xorm:"'ky_qishu' notnull default(0)"` //跨月期数
	Ym      int64  `xorm:"'ym' notnull default(0)"`       //年月
	StTime  string `xorm:"'st_time' default(ull)"`        //基准开盘时间，北京时间，已此作为其它期数开盘时间的基准
	StQishu int64  `xorm:"'st_qishu' default(ull)"`       //基准期数
	StQihao int64  `xorm:"'st_qihao' default(ull)"`       //基准期号，当天的第多少期
	StAllqi int64  `xorm:"'st_allqi' default(ull)"`       //基准期数，一天一共多少期
	LType   string `xorm:"'l_type' default('ull')"`
	Sort    int64  `xorm:"'sort' default(0)"` //排序
}

func (*PkGames) TableName() string {
	return global.TablePrefix + "pk_games"
}
