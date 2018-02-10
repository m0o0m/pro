package input

type VideoPlay struct {
	VType  string `query:"vType"`  //游戏类型
	GameId string `query:"gameId"` //具体游戏ID
}

//跳转游戏所需内容
type VideoUserData struct {
	SiteId   string
	IndexId  string
	UserName string
	Platform string //游戏平台名称
	AgentId  int64
	UaId     int64  //总代id
	ShId     int64  //股东id
	Media    string //设备类型  wap pc app
	GameID   string //子游戏id
	IP       string //ip
	Lang     string //语言类型
	Cur      string //货币类型
	Limit    string //限额
	Domain   string //登陆的域名
	IsSw     bool   //是否是试玩
}
