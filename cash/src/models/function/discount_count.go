package function

import (
	"errors"
	"fmt"
	"github.com/golyu/sql-build"
	"global"
	"models/back"
	"models/schema"
	"time"
)

type DiscountCountBean struct {
}

var discountCountSwitch = global.NewSwitch() //定时任务开关控制

//查询出一天返水数据 (1)
func (m *DiscountCountBean) GetRetreatWater(sTime, eTime int64) (discountReports []*schema.DiscountReport, err error) {
	sql := "SELECT t1.site_id, t1.site_index_id, member_id, t2.third_agency_id" +
		" AS agent_id, t2.second_agency_id AS ua_id, t2.first_agency_id AS sh_id," +
		" t1.account, count(*) AS num, sum(rebate_water) AS discount_money FROM" +
		" sales_member_retreat_water_record t1 JOIN sales_member t2 ON t1.member_id" +
		" = t2.id WHERE t1.status = 1"
	sql += " And start_time >= " + fmt.Sprintf("%d", sTime)
	sql += " And start_time <= " + fmt.Sprintf("%d", eTime)
	sql += " GROUP BY t1.site_id, t1.site_index_id, member_id,t1.account"
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err = sess.SQL(sql).Find(&discountReports)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return discountReports, err
	}
	tm := time.Unix(sTime, 0)
	timeTag := tm.Format("2006-01-02 15:04:05")                                //天标志
	day, _ := time.ParseInLocation("2006-01-02 15:04:05", timeTag, time.Local) //统计哪天
	dayTime := day.Unix()
	doTime := time.Now().Unix() //当前时间
	for i := range discountReports {
		discountReports[i].DayTime = dayTime
		discountReports[i].DayType = timeTag
		discountReports[i].DoTime = doTime
	}
	return discountReports, err
}

//插入到每日优惠统计表 (2)
func (m *DiscountCountBean) InsertOrUpdate(discountReports *[]*schema.DiscountReport) (err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sql, err := sqlBuild.Insert(new(schema.DiscountReport).TableName()).
		Values(discountReports).
		OrUpdate().
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	_, err = sess.Exec(sql)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	return err
}

// 5分钟统计一次
//为避免数据积压,下一次任务的开始时间和这一次任务的结束时间作为时间间隔的条件
func (m *DiscountCountBean) timingSwitch(stopChan chan interface{}) {
	t := time.NewTicker(5 * time.Minute) //5分钟统计一次
	//t := time.NewTicker(2 * time.Second) //2秒统计一次
	go func() {
		for {
			select {
			case <-t.C:
				//统计数据,并将统计的数据插入到discount_report
				sTime, eTime := global.GetToday()
				discountReports, err := m.GetRetreatWater(sTime, eTime)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
				}
				err = m.InsertOrUpdate(&discountReports)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
				}
			case <-stopChan:
				return
			}
		}
	}()
	return
}

//开启或者关闭定时统计
func (m *DiscountCountBean) TimingSwitch(open int64) int64 {
	switch open {
	case discountCountSwitch.IsOpen():
		return 71025
	case 1:
		discountCountSwitch.Open(m.timingSwitch)
	case 2:
		discountCountSwitch.Close()
	default:
		panic("error:open not in (1,2)")
	}
	return 0
}

//重新统计
func (m *DiscountCountBean) Recount() {
	//控制器端做
}

//优惠统计列表
func (m *DiscountCountBean) GetList(siteId string, memberIds []int64, times *global.Times, params *global.ListParams) (
	retreatWaterRecordTotal back.DiscountCountTotal, err error) {
	sess := global.GetXorm().NewSession()
	discountReportSchema := new(schema.DiscountReport)
	sess.Table(discountReportSchema.TableName())
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	if len(memberIds) > 0 {
		sess.In("member_id", memberIds)
	}
	times.Make("day_time", sess) //时间段
	conds := sess.Conds()
	params.Make(sess) //分页
	// DESCRIPTION:查询当前页
	var retreatWaterRecordList []back.DiscountCountList
	sess.Select("id,site_id,agent_id ,member_id,account,num,discount_money,day_type")
	sess.GroupBy("member_id,id,site_id,agent_id ,account,num,discount_money,day_type")
	err = sess.Find(&retreatWaterRecordList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return retreatWaterRecordTotal, err
	}
	// DESCRIPTION:查询统计信息
	b, err := sess.Table(discountReportSchema.TableName()).
		Where(conds).
		Select("COUNT(*) as num,SUM(discount_money) as discount_money").
		Get(&retreatWaterRecordTotal)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return retreatWaterRecordTotal, err
	}
	if !b {
		err = errors.New("find 0 row")
		return retreatWaterRecordTotal, err
	}
	retreatWaterRecordTotal.Content = &retreatWaterRecordList
	return retreatWaterRecordTotal, err
}

//查询出款账目汇总的统计信息--给予返水 (总额度,总笔数,总人数)
func (m *DiscountCountBean) GetDiscountCount(siteId string, times *global.Times) (*back.CashCollectDetails, error) {
	cashCollect := back.CashCollectDetails{}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Table(new(schema.DiscountReport).TableName())
	if siteId != "" {
		sess.Where("site_id =?", siteId)
	}
	times.Make("day_time", sess)
	sess.Select("sum(discount_money) as money,count(member_id) as people_num,sum(num) as num")
	_, err := sess.Get(&cashCollect)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashCollect, err
	}
	return &cashCollect, err
}
