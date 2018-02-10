package back

//查询电子列表返回数据
type VdGameList struct {
	Id        int64  `xorm:"id PK autoincr" json:"id"`
	Gameid    string `xorm:"gameid" json:"gameid"`       //'游戏id'
	Name      string `xorm:"name" json:"name"`           //'名字'
	Image     string `xorm:"image" json:"image"`         //'图片'
	Status    int8   `xorm:"status" json:"status"`       //'1正常，0不可用，2维护'
	Recommend int8   `xorm:"recommend" json:"recommend"` //'推荐度'
	IsSw      int8   `xorm:"is_sw" json:"is_sw"`         //'试玩线路开启为1 关闭为2'
	IsZs      int8   `xorm:"is_zs" json:"is_zs"`         //'正式线路开启为1 关闭为2'
}
