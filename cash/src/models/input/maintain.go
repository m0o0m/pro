package input

// 维护信息
type MaintainData struct {
	ID         int64  `json:"id"`         // 主键id
	MType      int    `json:"mtype"`      // 维护类型 1:全网 2:整线 3:单站
	CType      string `json:"ctype"`      // 终端 1:pc,2:wap... | 彩票,视讯,电子...
	LindId     string `json:"line_id"`    // 线路
	SiteId     string `json:"site_id"`    // 站点
	ProductId  string `json:"product_id"` // 维护项目 产品ID
	StartTime  int64  `json:"starttime"`  // 维护开始时间
	EndTime    int64  `json:"endtime"`    // 维护结束时间
	Remark     string `json:"remark"`     // 备注
	AddTime    int64  `json:"addtime"`    // 添加时间
	UpdateTime int64  `json:"updatetime"` // 更新时间
	StartDate  string `json:"-"`          // 维护开始日期
	EndDate    string `json:"-"`          // 维护结束日期
}
