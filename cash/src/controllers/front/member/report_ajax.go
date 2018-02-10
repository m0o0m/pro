package member

import (
	"github.com/labstack/echo"
	"global"
	"models/back"
	"strconv"
	"time"
)

type StatisticsAjax struct {
	BetReport []WeekReport
}

type WeekReport struct {
	back.ThisWeekBetReport
	Data []back.ThisWeekBetReport
}

//本周七天报表数据
func (m *StatisticsAjax) GetThisReportAjax(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	account := member.Account
	siteId := member.Site
	siteIndexId := member.SiteIndex
	m.BetReport = m.GetReportData(siteId, siteIndexId, account, 1)
	return ctx.JSON(200, global.ReplyItem(m.BetReport))
}

//上周七天报表数据
func (m *StatisticsAjax) GetLastReportAjax(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	account := member.Account
	siteId := member.Site
	siteIndexId := member.SiteIndex
	m.BetReport = m.GetReportData(siteId, siteIndexId, account, 2)
	return ctx.JSON(200, global.ReplyItem(m.BetReport))
}

//查询报表数据
func (m *StatisticsAjax) GetReportData(siteId, siteIndexId, account string, state int) (Data []WeekReport) {
	then := m.GetWeekDays(state) //获取上周一周时间日期
	loc, _ := time.LoadLocation("Local")
	time1, _ := time.ParseInLocation("2006-01-02 15:04:05", then[0]+" 00:00:00", loc)
	stime := time1.Unix()
	time2, _ := time.ParseInLocation("2006-01-02 15:04:05", then[6]+" 23:59:59", loc)
	etime := time2.Unix()
	//查询本周报表每天总计
	reportCount, err := betreport.GetThisWeekReportCount(siteId, siteIndexId, account, stime, etime)
	//查询本周具体游戏详细报表
	reportData, err := betreport.GetThisWeekReport(siteId, siteIndexId, account, stime, etime)
	Data = make([]WeekReport, len(then))
	if err != nil || reportCount == nil {
		for k, v := range then {
			Data[k].DayTime = v
		}
	} else {
		for k1, v1 := range then {
			date, _ := time.ParseInLocation("2006-01-02 15:04:05", v1+" 00:00:00", loc)
			udate := date.Unix()
			Data[k1].DayTime = v1
			for _, val := range reportCount {
				if val.DayTime == strconv.Itoa(int(udate)) {
					Data[k1].BetAll = val.BetAll
					Data[k1].BetValid = val.BetValid
					Data[k1].Win = val.Win
					Data[k1].Num = val.Num
					Data[k1].WinNum = val.WinNum
					Data[k1].BetAll = val.BetAll
				}
			}
			for _, vdata := range reportData {
				if vdata.DayTime == strconv.Itoa(int(udate)) {
					vdata.DayTime = v1
					Data[k1].Data = append(Data[k1].Data, vdata)
				}
			}
		}
	}
	return Data
}

//获取本周上周的七天日期
func (m *StatisticsAjax) GetWeekDays(state int) (days []string) {
	wday := time.Now().Weekday()
	switch wday {
	case time.Monday:
		days = m.GetWeek(1, state)
	case time.Tuesday:
		days = m.GetWeek(2, state)
	case time.Wednesday:
		days = m.GetWeek(3, state)
	case time.Thursday:
		days = m.GetWeek(4, state)
	case time.Friday:
		days = m.GetWeek(5, state)
	case time.Saturday:
		days = m.GetWeek(6, state)
	case time.Sunday:
		days = m.GetWeek(7, state)
	}
	return
}

func (m *StatisticsAjax) GetWeek(num, state int) (days []string) {
	var day string
	if state == 1 { //本周
		for i := 1; i <= 7; i++ {
			day = time.Now().AddDate(0, 0, i-num).Format("2006-01-02") + m.GetServerWeek(i, state)
			days = append(days, day)
		}
	} else if state == 2 { //上周
		for j := 7; j >= 1; j-- {
			day = time.Now().AddDate(0, 0, 1-j-num).Format("2006-01-02") + m.GetServerWeek(j, state)
			days = append(days, day)
		}
	}
	return
}

//获取一周七天
func (m *StatisticsAjax) GetServerWeek(num, state int) (days string) {
	if state == 1 {
		switch num {
		case 1:
			days = "(周一)"
		case 2:
			days = "(周二)"
		case 3:
			days = "(周三)"
		case 4:
			days = "(周四)"
		case 5:
			days = "(周五)"
		case 6:
			days = "(周六)"
		case 7:
			days = "(周日)"
		}
	} else {
		switch num {
		case 7:
			days = "(周一)"
		case 6:
			days = "(周二)"
		case 5:
			days = "(周三)"
		case 4:
			days = "(周四)"
		case 3:
			days = "(周五)"
		case 2:
			days = "(周六)"
		case 1:
			days = "(周日)"
		}
	}
	return
}
