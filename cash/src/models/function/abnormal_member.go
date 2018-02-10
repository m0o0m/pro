package function

import (
	"database/sql"
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type AbnormalMemberBean struct{}

//异常会员查询
func (*AbnormalMemberBean) GetAbnormalMemberList(this *input.AbnormalMemberList, listParam *global.ListParams) (data []back.AbnormalMemberList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	abnormalMember := new(schema.Member)
	memberBank := new(schema.MemberBank)
	sess.Where("remark is not null")
	if this.Key != "" {
		sess.Where("remark like ?", "%"+this.Key+"%")
	}
	if this.Type == 1 { //已处理
		sess.Where("remark like ?", "%PK系统检测异常会员%")
	} else if this.Type == 2 { //未处理
		sess.Where("remark not like ?", "%PK系统检测异常会员%")
	}
	listParam.Make(sess)
	conds := sess.Conds()
	sql := fmt.Sprintf("%s.id = %s.member_id", abnormalMember.TableName(), memberBank.TableName())
	err = sess.Table(abnormalMember.TableName()).
		Select(abnormalMember.TableName()+".site_id,"+abnormalMember.TableName()+".id,"+abnormalMember.TableName()+".account,"+memberBank.TableName()+".card,"+memberBank.TableName()+".card_name,"+memberBank.TableName()+".card_address").
		Join("LEFT", memberBank.TableName(), sql).
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(abnormalMember.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//异常会员处理
func (*AbnormalMemberBean) AbnormalMemberSet(this *input.AbnormalMemberSet) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	abnormalMember := new(schema.Member)
	var siteAccounts string
	for i, p := range this.AbnormalMembers {
		if i == len(this.AbnormalMembers)-1 {
			siteAccounts = siteAccounts + "'" + p.SiteId + p.Account + "'"
		} else {
			siteAccounts = siteAccounts + "'" + p.SiteId + p.Account + "',"
		}
	}
	var result sql.Result
	result, err = sess.Exec("update " + abnormalMember.TableName() + " set remark=concat(remark,'PK系统检测异常会员') where concat(site_id,account) in(" + siteAccounts + ")")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	count, err = result.RowsAffected()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
