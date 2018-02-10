package function

import (
	"github.com/go-xorm/xorm"
	"global"
	"models/back"
	"models/schema"
)

type MemberRebateRecordBean struct {
}

//根据会员id查询对应的返佣记录
func (m *MemberRebateRecordBean) GetRecordInfoInMemberId(ids []int) (rebateRecords []schema.MemberRebateRecord, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err = sess.Table(new(schema.MemberRebateRecord).TableName()).
		In("member_id", ids).
		Cols("member_id", "count(rebate) as rebate").
		Where("status = ?", 1). //已返佣的
		Find(&rebateRecords)
	return
}

//根据返佣记录id查询返佣记录
func (m *MemberRebateRecordBean) GetRebateRecordById(ids []int64) (rebateRecords []schema.MemberRebateRecord, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	err = sess.Table(new(schema.MemberRebateRecord).TableName()).
		In("id", ids).
		Where("status = ?", 1). //已返佣的
		Find(&rebateRecords)
	return
}

//返佣冲销
func (m *MemberRebateRecordBean) WriteoffRebateRecordById(ids []int64, sessArgs ...*xorm.Session) (
	int64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	num, err := sess.Table(new(schema.MemberRebateRecord).TableName()).
		In("id", ids).
		Where("status = ?", 1). //已返佣的
		Update(map[string]interface{}{"status": 2})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}

//根据期数id查询会员返佣记录
func (m *MemberRebateRecordBean) GetRebateRecord(commissionId int) (
	[]back.RebateDetail, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//根据期数id查询出返佣记录
	var rebateRecords []back.RebateDetail
	rebateRecord := &schema.MemberRebateRecord{}
	member := &schema.Member{}
	agency := &schema.Agency{}
	err := sess.Table(rebateRecord.TableName()).
		Select(rebateRecord.TableName()+".id, "+member.TableName()+".account as account, "+agency.TableName()+".account as agency, betting as all_bet, rebate, "+rebateRecord.TableName()+".status").
		Join("INNER", member.TableName(), rebateRecord.TableName()+".member_id = "+member.TableName()+".id").
		Join("INNER", agency.TableName(), member.TableName()+".third_agency_id = "+agency.TableName()+".id").
		Where("periods_id = ?", commissionId).
		Find(&rebateRecords)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rebateRecords, err
	}
	return rebateRecords, err
}
