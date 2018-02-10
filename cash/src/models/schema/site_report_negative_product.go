package schema

import "global"

//站点报表负数记录对应商品分类比例
type SiteReportNegativeProduct struct {
	NegativeId int64   `xorm:"negative_id"` //报表负数记录id
	ProductId  int64   `xorm:"product_id"`  //商品分类id
	ReportWin  float64 `xorm:"report_win"`  //报表盈利数字
}

func (*SiteReportNegativeProduct) TableName() string {
	return global.TablePrefix + "site_report_negative_product"
}
