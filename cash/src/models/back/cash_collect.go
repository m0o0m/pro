package back

import "fmt"

//出入款汇总详情
type CashCollectDetails struct {
	SourceType int64   `xorm:"source_type" json:"sourceType" ` //
	Money      float64 `xorm:"money" json:"money"`             //收入/支出金额
	PeopleNum  int64   `xorm:"people_num" json:"peopleNum"`    //交易人数
	Num        int64   `xorm:"num" json:"num"`                 //交易笔数
}

//钱(笔)(人)
func (m *CashCollectDetails) GetString() string {
	money := fmt.Sprintf("%10.2f", m.Money)
	num := fmt.Sprintf("%d", m.Num)
	people := fmt.Sprintf("%d", m.PeopleNum)
	return money + "(" + num + "笔)(" + people + "人"
}

//出入款汇总
type CashCollect struct {
	IncomeSum  float64               `json:"incomeSum" `  //总收入
	OutlaySum  float64               `json:"outlaySum" `  //总支出
	ProfitLoss float64               `json:"profitLoss" ` //实际盈亏 -
	Total      float64               `json:"total" `      //账目统计 +
	Details    []*CashCollectDetails `json:"details" `    //详情
}
