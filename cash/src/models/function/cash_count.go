package function

import (
	"errors"
	"github.com/golyu/sql-build"
	"global"
	"models/back"
	"models/schema"
	"sync"
	"time"
)

type CashCountBean struct {
}

var cashCountSwitch = global.NewSwitch() //定时任务开关控制
var memberBean = new(MemberBean)

//聚合查询入款数据
func (m *CashCountBean) GetAllCash(sTime, eTime int64) (cashReports []*schema.CashCountReport, err error) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	var companyReports, manualReports, onlineReports []*schema.CashCountReport
	var companyReportErr, manualReportErr, onlineReportErr error
	currentTime := global.GetCurrentTime()
	tm := time.Unix(sTime, 0)
	timeTag := tm.Format("2006-01-02")                                //天标志
	day, _ := time.ParseInLocation("2006-01-02", timeTag, time.Local) //统计哪天
	dayTime := day.Unix()
	go func() {
		defer wg.Done()
		// DESCRIPTION:公司入款
		companyReports, companyReportErr = m.GetCompanyCash(sTime, eTime, currentTime)
		if companyReportErr != nil {
			return
		}
		// DESCRIPTION:附加会员信息
		companyReportErr = m.AppendMember(companyReports, timeTag, dayTime, currentTime)
	}()
	go func() {
		// DESCRIPTION:线上人工入款
		defer wg.Done()
		manualReports, manualReportErr = m.GetManualCash(sTime, eTime)
		if manualReportErr != nil {
			return
		}
		// DESCRIPTION:附加会员信息
		manualReportErr = m.AppendMember(companyReports, timeTag, dayTime, currentTime)
	}()
	go func() {
		// DESCRIPTION:线上三方入款
		defer wg.Done()
		onlineReports, onlineReportErr = m.GetOnlineCash(sTime, eTime)
		if onlineReportErr != nil {
			return
		}
		// DESCRIPTION:附加会员信息
		onlineReportErr = m.AppendMember(companyReports, timeTag, dayTime, currentTime)
	}()
	wg.Wait()
	if onlineReportErr != nil {
		err = onlineReportErr
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	} else if manualReportErr != nil {
		err = manualReportErr
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	} else if companyReportErr != nil {
		err = companyReportErr
		global.GlobalLogger.Error("err:%s", err.Error())
		return cashReports, err
	} else if manualReportErr != nil {
		err = manualReportErr
		return cashReports, err
	} else if companyReportErr != nil {
		err = companyReportErr
		return cashReports, err
	}
	cashReports = append(cashReports, companyReports...)
	cashReports = append(cashReports, manualReports...)
	cashReports = append(cashReports, onlineReports...)
	return cashReports, err
}

//给入款信息附加上会员信息
func (m *CashCountBean) AppendMember(cashReports []*schema.CashCountReport, timeTag string, dayTime, current int64) (err error) {
	var memberIds []int64
	cashReportMap := make(map[int64]*schema.CashCountReport)
	for i := range cashReports {
		cashReports[i].DayType = timeTag
		cashReports[i].DayTime = dayTime
		cashReports[i].DoTime = current
		cashReportMap[cashReports[i].MemberId] = cashReports[i] //memberId作key
		memberIds = append(memberIds, cashReports[i].MemberId)
	}
	// DESCRIPTION:会员资料
	var members []schema.Member
	members, err = memberBean.GetMemberByIds(memberIds)
	if err != nil {
		return
	}
	if len(memberIds) != len(members) {
		return errors.New("Member number and member ids does not correspond")
	}
	for _, member := range members {
		cashReport, ok := cashReportMap[member.Id]
		if ok {
			cashReport.AgentId = member.ThirdAgencyId
			cashReport.SecondAgencyId = member.SecondAgencyId
			cashReport.FirstAgencyId = member.FirstAgencyId
		} else {
			return errors.New("Member number and member ids does not correspond")
		}
	}
	return
}

//统计人工入款数据 (0)
func (m *CashCountBean) GetManualCash(sTime, eTime int64) (cashReports []*schema.CashCountReport, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	manualAccess := new(schema.ManualAccess)
	err = sess.Table(manualAccess.TableName()).
		Select("site_id,site_index_id,1 AS into_style,member_id,account,COUNT(*) AS num,SUM(money) AS cash_money").
		Where("deposit_type = 1"). //人工存款
		Where("access_type = 1").  //存款
		Where("create_time >= ?", sTime).
		Where("create_time <= ?", eTime).
		GroupBy("member_id").
		Find(&cashReports)
	return
}

//统计公司入款数据 (1)
func (m *CashCountBean) GetCompanyCash(sTime, eTime, currentTime int64) (cashReports []*schema.CashCountReport, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	memberCompanyIncome := new(schema.MemberCompanyIncome)
	err = sess.Table(memberCompanyIncome.TableName()).
		Select("site_id,site_index_id,1 AS into_style,member_id,account,COUNT(*) AS num,SUM(deposit_money) AS cash_money").
		Where("status = 1").
		Where("update_time >= ?", sTime).
		Where("update_time <= ?", eTime).
		GroupBy("member_id").
		Find(&cashReports)
	return
}

//统计线上入款数据 (2)
func (m *CashCountBean) GetOnlineCash(sTime, eTime int64) (cashReports []*schema.CashCountReport, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	err = sess.Table(onlineEntryRecord.TableName()).
		Select("site_id,site_index_id,1 AS into_style,member_id,member_account as account,COUNT(*) as num,SUM(amount_deposit) AS cash_money").
		Where("status = 2").
		Where("third_pay_time >= ?", sTime).
		Where("third_pay_time <= ?", eTime).
		GroupBy("member_id").
		Find(&cashReports)
	return
}

//将统计数据插入到每日入款统计表 (2)
func (m *CashCountBean) InsertOrUpdate(cashReports *[]*schema.CashCountReport, mutexs ...*sync.Mutex) (err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sql, err := sqlBuild.Insert(new(schema.CashCountReport).TableName()).
		Values(cashReports).
		OrUpdate().
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	if len(mutexs) > 0 {
		mutexs[0].Lock()
		defer mutexs[0].Unlock()
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
func (m *CashCountBean) timingSwitch(stopChan chan interface{}) {
	t := time.NewTicker(5 * time.Minute) //5分钟统计一次
	//t := time.NewTicker(2 * time.Second) //2秒统计一次
	go func() {
		for {
			select {
			case <-t.C:
				//统计数据,并将统计的数据插入到discount_report
				sTime, eTime := global.GetToday()
				allReports, err := m.GetAllCash(sTime, eTime)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
				}
				mutex := global.GetSiteCashCountReportMutex("all")
				err = m.InsertOrUpdate(&allReports, mutex)
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
func (m *CashCountBean) TimingSwitch(open int64) int64 {
	switch open {
	case cashCountSwitch.IsOpen():
		return 71025
	case 1:
		cashCountSwitch.Open(m.timingSwitch)
	case 2:
		cashCountSwitch.Close()
	default:
		panic("error:open not in (1,2)")
	}
	return 0
}

//获取入款统计列表
func (m *CashCountBean) GetList(siteId string, memberIds []int64, intoStyle int64, times *global.Times, params *global.ListParams) (
	cashCountTotal back.CashCountTotal, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	cashReportSchema := new(schema.CashCountReport)
	sess.Table(cashReportSchema.TableName())
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	if len(memberIds) > 0 {
		sess.In("member_id", memberIds)
	}
	if intoStyle > 0 { //入款方式过滤
		sess.Where("into_style = ?", intoStyle)
	}
	times.Make("day_time", sess) //
	conds := sess.Conds()
	params.Make(sess) //分页
	// DESCRIPTION:查询当前页数据
	var cashCountList []back.CashCountList
	sess.Select("id,site_id,first_agency_id,second_agency_id,agent_id ,member_id,account,num,cash_money,into_style,day_type")
	sess.GroupBy("id,site_id,first_agency_id,second_agency_id,agent_id ,member_id,account,num,cash_money,into_style,day_type")
	err = sess.Find(&cashCountList)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	// DESCRIPTION:查询统计信息
	b, err := sess.Table(cashReportSchema.TableName()).
		Where(conds).
		Select("COUNT(*) as num,SUM(cash_money) as cash_money").
		Get(&cashCountTotal)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cashCountTotal, err
	}
	if !b {
		err = errors.New("find 0 row")
		return cashCountTotal, err
	}
	cashCountTotal.Content = &cashCountList
	return cashCountTotal, err
}

//入款账目汇总
func (m *CashCountBean) CashCollectCount(siteId string, times *global.Times) (cashCollect []*back.CashCollectDetails, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Table(new(schema.CashCountReport).TableName())
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	times.Make("day_time", sess)
	sess.Select("into_style as source_type,sum(cash_money) as money,count(member_id) as people_num,sum(num) as num")
	sess.Where("into_style < ?", 3) //0人工存入1公司入款2线上入款
	sess.GroupBy("into_style")
	err = sess.Find(&cashCollect)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return cashCollect, err
	}
	return cashCollect, err
}
