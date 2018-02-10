package schema

import "global"

//电子游戏类型
type MgGameType struct {
	Id   int64  `xorm:"id PK autoincr" json:"id"`
	Type string `xorm:"type" json:"type"`
}

func (*MgGameType) TableName() string {
	return global.TablePrefix + "mg_game_type"
}
