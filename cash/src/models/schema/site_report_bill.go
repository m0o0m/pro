package schema

import "global"

//注单表
type SiteReportBill struct {
	Id         int64  `xorm:"'id' PK autoincr" json:"id"`    // 主键id
	ReportData string `xorm:"report_data" json:"reportData"` // 报表数据
	SiteId     string `xorm:"site_id" json:"siteId"`         // 站点ID
	Status     int8   `xorm:"status" json:"status"`          // 推送状态 1已下发 2未下发 3删除
	Qishu      string `xorm:"qishu" json:"qishu"`            // 报表期数 年-月
	StartDate  int64  `xorm:"start_date" json:"startDate"`   // 起始时间
	EndDate    int64  `xorm:"end_date" json:"endDate"`       // 结束时间
}

func (*SiteReportBill) TableName() string {
	return global.TablePrefix + "site_report_bill"
}
