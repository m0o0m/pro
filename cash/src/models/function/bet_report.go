package function

import (
	"global"
	"models/back"
	"models/schema"
)

//会员每日打码统计信息
type BetReportBean struct {
}

/**
SELECT
  spread_id,
  product_name,
  t2.id          AS product_id,
  SUM(bet_valid) AS product_bet
FROM `sales_bet_report_account` `t1` LEFT JOIN `sales_product` t2
    ON t1.platform_id = t2.platform_id AND t1.game_type = t2.type_id
  LEFT JOIN `sales_member` t3 ON t1.account = t3.account
WHERE
  t1.site_id = 'aaaa' AND t1.site_index_id = 'a' AND day_time >= 1509465600 AND day_time <= 1511971200 AND spread_id > 0
GROUP BY spread_id,t1.platform_id,t1.game_type;
*/
//根据每日打码统计返回有效打码信息
func (m *BetReportBean) CountValidBet(siteId, siteIndexId string, sTime, eTime int64) (
	betReports []*back.PreRebateRecordProduct, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	betReportAccountSchema := new(schema.BetReportAccount)
	productSchema := new(schema.Product)
	memberSchema := new(schema.Member)
	err = sess.Table(betReportAccountSchema.TableName()).
		Alias("t1").
		Join("LEFT", productSchema.TableName(), "t1.platform_id = "+productSchema.TableName()+".platform_id AND t1.game_type = "+productSchema.TableName()+".type_id").
		Join("LEFT", memberSchema.TableName(), "t1.account = "+memberSchema.TableName()+".account").
		Select("spread_id,t1.v_type,product_name, "+productSchema.TableName()+".id as product_id, SUM(bet_valid) as product_bet").
		Where("t1.site_id = ?", siteId).
		Where("t1.site_index_id = ?", siteIndexId).
		Where("day_time >= ?", sTime).
		Where("day_time <= ?", eTime).
		Where("spread_id > ?", 0).
		GroupBy("spread_id,t1.platform_id,t1.v_type").
		Find(&betReports)
	return
}

//查询本周打码量总计
func (m *BetReportBean) GetThisWeekReportCount(siteId, siteIndexId, account string, stime, etime int64) (WeekReport []back.WeekBetReportCount, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.BetReportAccount)
	err = sess.Table(product.TableName()).
		Select("SUM(bet_all) as bet_all,SUM(bet_valid) as bet_valid,SUM(win) as win,SUM(num) as num,SUM(win_num) as win_num,day_time,v_type").
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("account=?", account).
		Where("day_time >= ?", stime).
		Where("day_time <= ?", etime).
		GroupBy("day_time").
		Find(&WeekReport)
	return
}

//查询本周打码量
func (m *BetReportBean) GetThisWeekReport(siteId, siteIndexId, account string, stime, etime int64) (WeekReport []back.ThisWeekBetReport, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.BetReportAccount)
	err = sess.Table(product.TableName()).
		Select("SUM(bet_all) as bet_all,SUM(bet_valid) as bet_valid,SUM(win) as win,SUM(num) as num,SUM(win_num) as win_num,day_time,v_type").
		Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).
		Where("account=?", account).
		Where("day_time >=?", stime).
		Where("day_time <=?", etime).
		GroupBy("v_type,day_time").
		Find(&WeekReport)
	return
}
