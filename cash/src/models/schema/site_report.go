package schema

import "global"

//注单表
type SiteReport struct {
	Id         int64   `xorm:"'id' PK autoincr" json:"id"`     // 主键id
	SiteId     string  `xorm:"site_id" json:"site_id"`         // 站点ID
	Num        int64   `xorm:"num" json:"num"`                 // 总笔数
	BetAll     float64 `xorm:"bet_all" json:"bet_all"`         // 总计打码
	BetValid   float64 `xorm:"bet_valid" json:"bet_valid"`     // 有效打码
	Win        float64 `xorm:"win" json:"win"`                 // 结果
	WinNum     int64   `xorm:"win_num" json:"win_num"`         // 赢的笔数
	Jack       float64 `xorm:"jack" json:"jack"`               // 彩金
	VType      string  `xorm:"v_type" json:"v_type"`           // 游戏类型
	DayTime    int64   `xorm:"day_time" json:"day_time"`       // 标示统计哪天
	CreateTime int64   `xorm:"create_time" json:"create_time"` // 添加时间
}

func (*SiteReport) TableName() string {
	return global.TablePrefix + "site_report"
}
