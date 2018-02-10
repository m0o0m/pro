package function

import (
	"github.com/go-xorm/xorm"
	"global"
	"models/schema"
)

//稽核日志
type MemeberAuditLogBean struct {
}

//添加稽核日志
func (m *MemeberAuditLogBean) InsertMulti(auditLogs []*schema.MemberAuditLog, sessArgs ...*xorm.Session) (int64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	if len(auditLogs) > 0 {

	}
	num, err := sess.Table(new(schema.MemberAuditLog).TableName()).
		InsertMulti(&auditLogs)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return num, err
	}
	return num, err
}
