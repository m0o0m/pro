package schema

import "global"

//会员现金统计表
type MemberCashCount struct {
	Member       int64   `xorm:"'member_id' PK "` //会员Id
	DepositCount int64   `xorm:"deposit_count" `  //存款次数
	DepositNum   float64 `xorm:"deposit_num" `    //存款总额
	DepositMax   float64 `xorm:"deposit_max" `    //最大存款数
	DrawNum      int64   `xorm:"draw_num" `       //取款次数
	DrawCount    float64 `xorm:"draw_count" `     //取款总额
	DrawMax      float64 `xorm:"draw_max" `       //最大取款数
	SpreadMoney  float64 `xorm:"spread_money" `   //会员推广获利金额
}

func (*MemberCashCount) TableName() string {
	return global.TablePrefix + "member_cash_count"
}
