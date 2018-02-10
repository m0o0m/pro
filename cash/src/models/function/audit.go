package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type AuditsBean struct{}

//稽核日志记录
func (*AuditsBean) GetAuditLogList(this *input.AuditLogGet, listParam *global.ListParams) (
	[]back.AuditLogList, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.AuditLogList
	auditlog := new(schema.MemberAuditLog)
	sess.Where("site_id=?", this.SiteId)
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where("account=?", this.Account)
	}
	if this.StartTime != "" {
		loc, _ := time.LoadLocation("Local")
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", this.StartTime+" 00:00:00", loc)
		StartTime := t.Unix()
		sess.Where("update_date>=?", StartTime)
	}
	if this.EndTime != "" {
		loc, _ := time.LoadLocation("Local")
		t, _ := time.ParseInLocation("2006-01-02 15:04:05", this.EndTime+" 23:59:59", loc)
		EndTime := t.Unix()
		sess.Where("update_date>=?", EndTime)
	}
	conds := sess.Conds()
	listParam.Make(sess)
	err := sess.Table(auditlog.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	counts, err := sess.Table(auditlog.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, counts, err
	}
	return data, counts, err
}

//查询会员即时稽核
func (*AuditsBean) GetMenmberNowAudit(this *input.AuditNow) (
	[]back.MemberAuditNow, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.MemberAuditNow
	auditlog := new(schema.MemberAuditRecord)
	sess.Where("site_id=?", this.SiteId)

	if this.Account != "" {
		sess.Where("account=?", this.Account)
	}
	err := sess.Table(auditlog.TableName()).Where("status=?", 1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询会员稽核期间的打码情况
func (*AuditsBean) GetMemberValidBet(account, siteId string, sTime, eTime int64) (
	[]back.AuditBet, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.AuditBet
	if account != "" {
		sess.Where("account=?", account)
	}
	sess.Where("site_id=?", siteId)
	sess.Where("day_time>=?", sTime)
	sess.Where("day_time<=?", eTime)
	sess.Select("id,site_id,site_index_id,account,sum(bet_valid) as bet_valid,game_type")
	bra := new(schema.BetReportAccount)
	err := sess.Table(bra.TableName()).
		GroupBy("game_type").Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

//查询会员账号是否存在
func (*AuditsBean) IsMemberAccount(Account, SiteId string) (
	back.AuditMember1, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var member back.AuditMember1
	sess.Where("site_id=?", SiteId)
	sess.Where("account=?", Account)
	m := new(schema.Member)
	hs, err := sess.Table(m.TableName()).Get(&member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return member, hs, err
	}
	return member, hs, err
}

//稽核日志记录(后台管理)
func (*AuditsBean) AuditLogList(this *input.AuditLogAdmin, listParam *global.ListParams, times *global.Times) (data []back.AuditLogAdminList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	auditlog := new(schema.MemberAudit)
	member := new(schema.Member)
	if this.SiteId != "" {
		sess.Where(auditlog.TableName()+".site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where(auditlog.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where(member.TableName()+".account=?", this.Account)
	}
	listParam.Make(sess)
	times.Make("begin_time", sess)
	conds := sess.Conds()
	sql := fmt.Sprintf("%s.member_id = %s.id", auditlog.TableName(), member.TableName())
	sess.Select(fmt.Sprintf("%s.id,%s.site_id,%s.account,%s.begin_time,%s.end_time",
		auditlog.TableName(), auditlog.TableName(), member.TableName(), auditlog.TableName(), auditlog.TableName()))
	err = sess.Table(auditlog.TableName()).Join("LEFT", member.TableName(), sql).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(auditlog.TableName()).Where(conds).Join("LEFT", member.TableName(), sql).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return
}

//稽核记录(后台管理)
func (*AuditsBean) AuditList(this *input.AuditLogAdmin, listParam *global.ListParams, times *global.Times) (data []back.AuditAdminList, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	audit := new(schema.MemberAuditLog)
	member := new(schema.Member)
	if this.SiteId != "" {
		sess.Where(audit.TableName()+".site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where(audit.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	if this.Account != "" {
		sess.Where(member.TableName()+".account=?", this.Account)
	}
	listParam.Make(sess)
	times.Make("update_date", sess)
	conds := sess.Conds()
	sql := fmt.Sprintf("%s.member_id = %s.id", audit.TableName(), member.TableName())
	sess.Select(fmt.Sprintf("%s.id,%s.site_id,%s.account,%s.member_id,%s.update_date,%s.content",
		audit.TableName(), audit.TableName(), member.TableName(), audit.TableName(), audit.TableName(), audit.TableName()))
	err = sess.Table(audit.TableName()).Join("LEFT", member.TableName(), sql).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(audit.TableName()).Where(conds).Join("LEFT", member.TableName(), sql).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//查询会员最后一条稽核记录
func (*AuditsBean) AuditLastOne(siteId, account string) (data back.MemberAuditNow, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	auditlog := new(schema.MemberAuditRecord)
	sess.Where("site_id=?", siteId)
	sess.Where("account=?", account)
	has, err = sess.Table(auditlog.TableName()).OrderBy("id desc").Get(&data)
	return
}
