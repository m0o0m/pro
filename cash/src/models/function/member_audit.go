package function

import (
	"github.com/go-xorm/xorm"
	"global"
	"models/schema"
)

type MemberAuditBean struct{}

//根据条件批量修改稽核状态和时间
func (*MemberAuditBean) ChangeAuditSAT(newRecord *schema.MemberAudit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err = sess.Where("site_id=?", newRecord.SiteId).Where("site_index_id=?", newRecord.SiteIndexId).Where("member_id=?", newRecord.MemberId).Where("status=?", 2).Cols("end_time,status").Update(newRecord)
	return
}

//增加一条稽核记录
func (*MemberAuditBean) AddNewRecord(newRecord *schema.MemberAudit) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err = sess.Insert(newRecord)
	return
}

//添加多条稽核记录
func (m *MemberAuditBean) InsertAuditMulti(audits []*schema.MemberAudit, sessArgs ...*xorm.Session) (int64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	return sess.Table(new(schema.MemberAudit).TableName()).InsertMulti(&audits)
}

//修改多条稽核记录为已经处理 通过稽核时间和会员id
func (*MemberAuditBean) OverAudit(beginTime int64, memberIds []int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	num, err := sess.Table(new(schema.MemberAudit).TableName()).
		Where("begin_time = ?", beginTime).
		Where("status = ?", 1).
		In("member_id", memberIds).
		Update(map[string]interface{}{"status": 2})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}
