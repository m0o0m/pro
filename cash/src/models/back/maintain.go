package back

// 维护信息
type MaintainData struct {
	ID         int64  `xorm:"id"`         // 主键id
	MType      int    `xorm:"mtype"`      // 维护类型 1:全网 2:整线 3:单站
	CType      string `xorm:"ctype"`      // 终端 1:pc,2:wap... | 彩票,视讯,电子...
	LindId     string `xorm:"line_id"`    // 线路
	SiteId     string `xorm:"site_id"`    // 站点
	ProductId  string `xorm:"product_id"` // 维护项目 产品ID
	StartTime  int64  `xorm:"starttime"`  // 维护开始时间
	EndTime    int64  `xorm:"endtime"`    // 维护结束时间
	Remark     string `xorm:"remark"`     // 备注
	AddTime    int64  `xorm:"addtime"`    // 添加时间
	UpdateTime int64  `xorm:"updatetime"` // 更新时间
	StartDate  string `xorm:"-"`          // 维护开始日期
	EndDate    string `xorm:"-"`          // 维护结束日期
}

// 维护主页 返回格式
type MaintainIndexRes struct {
	Data  map[int]*map[string]*map[int]*MaintainData
	List  map[int]*map[string]*map[string]*map[int]*MaintainData
	Sites []SiteSiteIndexBack
}
type MaintainDataRes struct {
	ID         int64
	MType      int
	CType      string
	LindId     string
	SiteId     string
	ProductId  string
	StartTime  int64
	EndTime    int64
	Remark     string
	AddTime    int64
	UpdateTime int64
	StartDate  string
	EndDate    string
	LindIds    map[string]LineIds
}
type LineIds struct {
	ID         int64
	MType      int
	CType      string
	LindId     string
	SiteId     string
	ProductId  string
	StartTime  int64
	EndTime    int64
	Remark     string
	AddTime    int64
	UpdateTime int64
	StartDate  string
	EndDate    string
	SiteIds    map[string]MaintainData
}

// 维护状态 返回格式
type IsMaintainRes struct {
	Return    int    `json:"return"`    // 返回值 1:未维护 2:维护中
	Remark    string `json:"remark"`    // 备注
	StartTime int64  `json:"starttime"` // 维护开始时间
	EndTime   int64  `json:"endtime"`   // 维护结束时间
}
