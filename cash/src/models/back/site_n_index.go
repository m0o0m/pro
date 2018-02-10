package back

//n_index.html页面需要的站点数据
type NIndexData struct {
	Notice    string              //公告内容
	FlashList []*NIndexFlashImage //轮播图
	LogoUrl   string              //logo地址
	SiteName  string              //站点名称

	Header interface{} //页头
	Footer interface{} //页尾
}

//egame.html页面需要的站点数据
type EgameData struct {
	Notice    string        //公告内容
	LogoUrl   string        //logo地址
	SiteName  string        //站点名称
	GameTitle []string      //电子导航
	GameData  EgameDataInfo //电子信息
}

type NIndexFlashImage struct {
	ImgUrl  string //图片路径
	ImgLink string //链接地址
}

//sports.html页面需要的站点数据
type SportsData struct {
	Notice      string   //公告内容
	LogoUrl     string   //logo地址
	SiteName    string   //站点名称
	SportsOrder []string //体育平台排序
}

//livetop.html页面需要的站点数据
type LiveTopData struct {
	Notice    string   //公告内容
	LogoUrl   string   //logo地址
	SiteName  string   //站点名称
	LiveOrder []string //视讯平台排序
}

//livetop.html页面需要的站点数据
type LiveTopAjaxData struct {
	LiveId    int           //模版选择
	LiveOrder []LiveTopName //视讯平台排序
}

type LiveTopName struct {
	ProductName string `json:"product_name"` //商品名
	VType       string `json:"v_type"`       //游戏类型
}

//lottery.html页面需要的站点数据
type LotteryData struct {
	Notice       string   //公告内容
	LogoUrl      string   //logo地址
	SiteName     string   //站点名称
	LotteryOrder []string //彩票平台排序

	Header interface{} //页头
	Footer interface{} //页尾
	CdnUrl string
}

//wapview.html页面需要的站点数据
type WapviewData struct {
	WapDomain  string //手机域名
	SiteName   string //站点名称
	IosUrl     string //ios下载地址
	AndroidUrl string //安卓下载地址
}

//平台游戏数据
type EgameDataInfo struct {
	Count int      `json:"count"` //游戏总数
	Data  []MgGame `json:"data"`
	Wh    int8     `json:"wh"`   // 1 维护 2  不维护
	Type  string   `json:"type"` // 类型
}

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

//youhui.html页面需要的站点数据
type YouhuiData struct {
	Notice   string           //公告内容
	LogoUrl  string           //logo地址
	SiteName string           //站点名称
	YhTitle  []YouhuiDataInfo //优惠分类信息
	YhData   []SiteActivity
}

type SiteActivity struct {
	Id          int64  `xorm:"'id' PK autoincr" json:"id"`         //id
	TopId       int64  `xorm:"top_id" json:"top_id"`               //上级栏目
	SiteId      string `xorm:"site_id" json:"site_id"`             //站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	Title       string `xorm:"title" json:"title"`                 //标题
	Content     string `xorm:"content" json:"content"`             //内容
	Img         string `xorm:"img" json:"img"`                     //标题图片路径
	State       int8   `xorm:"state" json:"state"`                 //状态 1启用  2关闭
	Sort        int64  `xorm:"sort" json:"sort"`                   //排序
	From        int8   `xorm:"from" json:"from"`                   //'1-PC 2-WAP
	Itype       int64  `xorm:"itype" json:"itype"`                 //类型代码
	TypeName    string `xorm:"type_name" json:"type_name"`         //类型名称
	AddTime     int64  `xorm:"add_time" json:"add_time"`           //操作时间
	DeleteTime  int64  `xorm:"delete_time" json:"delete_time"`     //删除时间   0未删除
}

//平台游戏数据
type YouhuiDataInfo struct {
	Id    int64  `json:"id"`    //游戏总数
	Title string `json:"title"` //游戏总数
}
