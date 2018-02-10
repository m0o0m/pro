package global

import (
	"fmt"
	"strconv"
	"time"
)

func GetCurrentTime() int64 {
	//local, _ := time.LoadLocation("Local")	//服务器设置的时间
	local, err := time.LoadLocation("UTC") //北京时间
	if err != nil {
		GlobalLogger.Error("err:%s", err.Error())
	}
	return time.Now().In(local).Unix()
}

//获取北京时间（时间戳）
func GetBeijingtime() int64 {
	local, err := time.LoadLocation("UTC") //北京时间
	if err != nil {
		GlobalLogger.Error("err:%s", err.Error())
	}
	var str = time.Now().In(local).Format("2006-01-02 15:04:05")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	return t.Unix()
}

//格式化的时间(年月日 时分秒)转时间戳
func FormatTime2Timestamp(startTime string, endTime string) (sTime, eTime int64, code int64) {
	sTime, code = FormatTime2Timestamp2(startTime)
	if code > 0 {
		code = 30155
		return
	}
	eTime, code = FormatTime2Timestamp2(endTime)
	if code > 0 {
		code = 30156
		return
	}
	if sTime > eTime {
		code = 30139
		return
	}
	return
}

//格式化的日期(年月日)转时间戳
func FormatDay2Timestamp(startDay string, endDay string) (sTime, eTime int64, code int64) {
	sTime, code = FormatDay2Timestamp2(startDay)
	if code > 0 {
		code = 30155
		return
	}
	eTime, code = FormatDay2Timestamp2(endDay)
	if code > 0 {
		code = 30156
		return
	}
	if sTime > eTime {
		code = 30139
		return
	}
	return
}

//格式化的时间(年月日 时分秒)转时间戳
func FormatTime2Timestamp2(datetime string) (dateTime int64, code int64) {
	loc, _ := time.LoadLocation("Local")
	if len(datetime) != 19 {
		code = 70028
		return
	} else {
		sT, stErr := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
		if stErr != nil {
			code = 70028
			return
		}
		dateTime = sT.Unix()
	}
	return
}

//格式化的日期(年月日)转时间戳
func FormatDay2Timestamp2(day string) (dayTime int64, code int64) {
	loc, _ := time.LoadLocation("Local")
	if len(day) < 10 {
		code = 70015
		return
	} else {
		sT, stErr := time.ParseInLocation("2006-01-02", day, loc)
		if stErr != nil {
			code = 70015
			return
		}
		dayTime = sT.Unix()
	}
	return
}

//格式化的日期(年月)得到本月开始时间和本月结束时间
func FormatMonth2Timestamp(day string) (startDay, endDay int64, code int64) {
	code = 70015
	if len(day) == 7 {
		year, yearErr := strconv.ParseInt(day[0:4], 10, 64)
		if yearErr != nil {
			return
		}
		month, monthErr := strconv.ParseInt(day[5:7], 10, 64)
		if monthErr != nil {
			return
		}
		days := GetCurrentMonthDays(year, month)
		return FormatDay2Timestamp(linkDate(year, month, 1), linkDate(year, month, days))
	}
	return
}

func linkDate(ints ...int64) (date string) {
	for _, v := range ints {
		if date != "" {
			date += "-"
		}
		day := fmt.Sprintf("%d", v)
		if len(day) == 1 {
			day = "0" + day
		}
		date += day
	}
	return
}

//根据年月得到本月天数
func GetCurrentMonthDays(year int64, month int64) int64 {
	if month == 2 {
		if year%4 == 0 {
			return 29
		}
		return 28
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	}
	return 31
}

//得到今天0时0分0秒和23时59分59秒
func GetToday() (int64, int64) {
	timeStr := time.Now().Format("2006-01-02")
	//fmt.Println("timeStr:", timeStr)
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	timeNumber := t.Unix()
	return timeNumber, timeNumber + 86399
}

//得到明天0时0分0秒和23时59分59秒
func GetTomorrow() (int64, int64) {
	s, e := GetToday()
	return s + 86400, e + 86400
}

//传入开始结束时间,返回之间的每一天
func GetEveryDay(begin, end time.Time) (days []time.Time) {
	dd := end.Sub(begin)
	for i := 0; i <= int(dd.Hours()/24); i++ {
		d, _ := time.ParseDuration("-24h")
		day := end.Add(d * time.Duration(i))
		days = append(days, day)
	}
	return
}

//传入开始结束时间,返回之间的每一天的0点和23点59分59秒
func GetEveryDayTimes(begin, end string) (days []*Times, code int64) {
	beginTime, err := time.ParseInLocation("2006-01-02", begin, time.Local)
	if err != nil {
		code = 30155
		return days, code
	}
	endTime, err := time.ParseInLocation("2006-01-02", end, time.Local)
	if err != nil {
		code = 30156
		return days, code
	}
	if beginTime.Unix() > endTime.Unix() {
		code = 30139
		return days, code
	}
	dayTimes := GetEveryDay(beginTime, endTime)
	for _, v := range dayTimes {
		days = append(days, &Times{v.Unix(), v.Unix() + 86399})
	}
	return days, code
}

//得到传入时间戳当天0时0分0秒的时间戳
func GetTimeStart(timeStamp int64) int64 {
	timeStr := time.Unix(timeStamp, 0).Format("2006-01-02")
	//fmt.Println("timeStr:", timeStr)
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	timeNumber := t.Unix()
	return timeNumber
}
