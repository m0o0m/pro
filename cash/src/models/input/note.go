package input

//电子游戏查询
type VdGameList struct {
	Name string `query:"name"` //游戏名字
	Type string `query:"type"` //游戏类型
}

//电子游戏添加参数
type VdGameAdd struct {
	Type   string `json:"type" valid:"Required;ErrorCode(60301)"`       //游戏类型
	Name   string `json:"name" valid:"Required;ErrorCode(60302)"`       //名字
	Topid  int8   `json:"topid" valid:"Range(1,5);ErrorCode(60305)"`    //'父类型id 1 SLOTS，2 TABLE GAMES，3 VIDEO POKER，5 Others'
	Itemid int8   `json:"itemid" valid:"Range(11,31);ErrorCode(60306)"` //'子类型id 11 3 Reel Slots，12 5 Reel Slots，13 Bonus Screen ，14 Others 21 BlackJack 22 OtherCasinoGames 23 OtherTableGames 24 Others 25 Poker 26 Roulette 31 VIDEO POKER'
	Gameid string `json:"gameid" valid:"Required;ErrorCode(60303)"`     //'游戏id'
	Image  string `json:"image" valid:"Required;ErrorCode(60304)"`      //'图片'
}

//电子游戏修改参数
type VdGameUpdate struct {
	Id        int64 `json:"id" valid:"Required;ErrorCode(60310)"`
	Status    int8  `json:"status"`    //1正常，2不可用，3维护
	Recommend int8  `json:"recommend"` //推荐度（优先）
	Ckr       int8  `json:"ckr"`       //推荐度
	IsSw      int8  `json:"is_sw"`     //试玩线路开启为1 关闭为2
	IsZs      int8  `json:"is_zs"`     //正式线路开启为1 关闭为2
}

//电子游戏修改参数（修改内容）
type VdGameContentUpdate struct {
	Id     int64  `json:"id" valid:"Required;ErrorCode(60310)"`
	Type   string `json:"type" valid:"Required;ErrorCode(60301)"`       //游戏类型
	Name   string `json:"name" valid:"Required;ErrorCode(60302)"`       //名字
	Topid  int8   `json:"topid" valid:"Range(1,5);ErrorCode(60305)"`    //'父类型id 1 SLOTS，2 TABLE GAMES，3 VIDEO POKER，5 Others'
	Itemid int8   `json:"itemid" valid:"Range(11,31);ErrorCode(60306)"` //'子类型id 11 3 Reel Slots，12 5 Reel Slots，13 Bonus Screen ，14 Others 21 BlackJack 22 OtherCasinoGames 23 OtherTableGames 24 Others 25 Poker 26 Roulette 31 VIDEO POKER'
	Gameid string `json:"gameid" valid:"Required;ErrorCode(60303)"`     //'游戏id'
	Image  string `json:"image" valid:"Required;ErrorCode(60304)"`      //'图片'
}

type VdGameStatusUpdate struct {
	SiteId      string `json:"site_id" valid:"MaxSize(4);ErrorCode(50058)"`       //站点id
	SiteIndexId string `json:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	GameId      int64  `json:"game_id" valid:"Required;ErrorCode(60310)"`         //游戏id
}

//添加游戏视讯类型
type VdGameTypeAdd struct {
	Type string `query:"type" valid:"Required;ErrorCode(60301)"` //游戏类型
}
