package function

import (
	"global"
	"models/schema"
)

type MemberCashCountBean struct{}

//根据Id获取一条数据
func (*MemberCashCountBean) GetInfo(id int64) (info schema.MemberCashCount, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Where("member_id=?", id).Get(&info)
	return
}

//修改
func (*MemberCashCountBean) UpdateRecord(memberInfo *schema.MemberCashCount) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("member_id=?", memberInfo.Member)
	sess.Cols("deposit_count,deposit_num,deposit_max,draw_num,draw_count,draw_max,spread_money")
	count, err = sess.Update(memberInfo)
	return
}

//确定入款时候修改数据[]
func (*MemberCashCountBean) ChangeData(id int64, money float64) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newCashCountBean := new(MemberCashCountBean)
	newCashCount := new(schema.MemberCashCount)
	info, flag, err := newCashCountBean.GetInfo(id)
	if err != nil {
		return 0, err
	}
	if !flag {
		return 0, err
	}
	newCashCount.Member = info.Member
	newCashCount.DepositCount = info.DepositCount + 1
	newCashCount.DepositNum = info.DepositNum + money
	if money > info.DepositMax {
		newCashCount.DepositMax = money
	} else {
		newCashCount.DepositMax = info.DepositMax
	}
	newCashCount.DrawNum = info.DrawNum
	newCashCount.DrawCount = info.DrawCount
	newCashCount.DrawMax = info.DrawMax //最大取款数
	newCashCount.SpreadMoney = info.SpreadMoney
	count, err = newCashCountBean.UpdateRecord(newCashCount)
	return
}

//确定出款的时候修改数据
func (*MemberCashCountBean) ConfirmOutMoneyChangeData(id int64, money float64) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newCashCountBean := new(MemberCashCountBean)
	newCashCount := new(schema.MemberCashCount)
	info, flag, err := newCashCountBean.GetInfo(id)
	if err != nil {
		return 0, err
	}
	if !flag {
		return 0, err
	}

	if info.DrawMax < money {
		newCashCount.DrawMax = money
	}
	newCashCount.Member = id
	newCashCount.DrawNum = info.DrawNum + 1
	newCashCount.DrawCount = info.DrawCount + money
	count, err = newCashCountBean.ChangeOutMoney(newCashCount)
	return
}

//更改会员现金统计表的出款
func (*MemberCashCountBean) ChangeOutMoney(newCash *schema.MemberCashCount) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Where("member_id=?", newCash.Member)
	sess.Cols("draw_num,draw_count,draw_max")
	count, err = sess.Update(newCash)
	return
}
