package schema

import "global"

//电子游戏
type MgGame struct {
	Id        int64  `xorm:"id PK autoincr" json:"id"`
	Topid     int8   `xorm:"topid" json:"topid"`   //'父类型id 1 SLOTS，2 TABLE GAMES，3 VIDEO POKER，5 Others'
	Itemid    int8   `xorm:"itemid" json:"itemid"` //'子类型id 11 3 Reel Slots，12 5 Reel Slots，13 Bonus Screen ，14 Others 21 BlackJack 22 OtherCasinoGames 23 OtherTableGames 24 Others 25 Poker 26 Roulette 31 VIDEO POKER'
	Gameid    string `xorm:"gameid" json:"gameid"` //'游戏id'
	Name      string `xorm:"name" json:"name"`     //'名字'
	Image     string `xorm:"image" json:"image"`   //'图片'
	Status    int8   `xorm:"status" json:"status"` //'1正常，2不可用，3维护'
	Type      string `xorm:"type" json:"type"`
	Recommend int8   `xorm:"recommend" json:"recommend"` //'推荐度'
	IsSw      int8   `xorm:"is_sw" json:"is_sw"`         //'试玩线路开启为1 关闭为2'
	IsZs      int8   `xorm:"is_zs" json:"is_zs"`         //'正式线路开启为1 关闭为2'
}

func (*MgGame) TableName() string {
	return global.TablePrefix + "mg_game"
}
