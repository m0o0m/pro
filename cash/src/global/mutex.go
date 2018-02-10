package global

import "sync"

//单站点定时统计插入的锁
var siteCashCountReportMutexs sync.Map

//得到站点入款统计插入的单站点锁
func GetSiteCashCountReportMutex(siteId string) *sync.Mutex {
	temp, _ := siteCashCountReportMutexs.LoadOrStore(siteId, new(sync.Mutex))
	mutex, _ := temp.(*sync.Mutex)
	return mutex
}
