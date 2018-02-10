package schema

import "global"

//视讯ip白名单
type GameWhiteList struct {
	Id      int    `xorm:"id"`      //序号
	Ip      string `xorm:"ip"`      //ip
	Remarks string `xorm:"remarks"` //备注
}

func (*GameWhiteList) TableName() string {
	return global.TablePrefix + "game_whitelist"
}
