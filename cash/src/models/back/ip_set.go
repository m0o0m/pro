package back

//ip开关列表
type IpSetList struct {
	Id      int64  `xorm:"id" json:"id"`
	IpStart string `xorm:"ip_start" json:"ip_start"` // 起始ip
	IpEnd   string `xorm:"ip_end" json:"ip_end"`     // 结束ip
	Type    string `xorm:"type" json:"type"`         // 控制类型：1为客户后阳台，2为代理后台，3为前台，4为wap端
	State   int8   `xorm:"state" json:"state"`       // 状态：1为启用，2为停用
	Remark  string `xorm:"remark" json:"remark"`     // 备注
}

//ip白名单列表
type WhiteListBack struct {
	Id     int64  `xorm:"id" json:"id"`          //
	SiteId string `xorm:"site_id" json:"siteId"` // 站点id
	Ip     string `xorm:"ip" json:"ip"`          //ip白名单:英文逗号分隔  填写多个
	State  int8   `xorm:"state" json:"state"`    // 状态：1为启用，2为停用
	Remark string `xorm:"remark" json:"remark"`  // 备注
}
