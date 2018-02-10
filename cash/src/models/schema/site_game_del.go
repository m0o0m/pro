package schema

import "global"

//电子游戏
type SiteGameDel struct {
	SiteId      string `xorm:"site_id" json:"site_id"`             //站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	GameId      int64  `xorm:"game_id" json:"game_id"`             //游戏id
}

func (*SiteGameDel) TableName() string {
	return global.TablePrefix + "site_game_del"
}
