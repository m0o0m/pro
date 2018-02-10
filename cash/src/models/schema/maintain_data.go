package schema

import "global"

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
}

func (*MaintainData) TableName() string {
	return global.TablePrefix + "maintain_data"
}
