package function

import (
	"github.com/go-xorm/xorm"
	"github.com/golyu/sql-build"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type MemberCashRecordBean struct{}

//增加一条出入款记录
func (*MemberCashRecordBean) AddNewRecord(newCashRecord *schema.MemberCashRecord) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err = sess.Insert(newCashRecord)
	return
}

//增加一条出入款记录(sess 事务)
func (*MemberCashRecordBean) AddNewRecordSess(newCashRecord *schema.MemberCashRecord, sessArgs ...*xorm.Session) (count int64, err error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	return sess.Insert(newCashRecord)
}

//添加多条现金记录
func (m *MemberCashRecordBean) AddCashRecordMulti(cashRecords []*schema.MemberCashRecord, sessArgs ...*xorm.Session) (int64, error) {
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	count, err := sess.Table(new(schema.MemberCashRecord).TableName()).InsertMulti(&cashRecords)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//查询现金记录表
func (m *MemberCashRecordBean) GetCashRecordList(this *input.MemberCashRecord, listParams *global.ListParams) (
	info *back.CashRecordAllBack, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info = new(back.CashRecordAllBack)
	memberCashRecord := new(schema.MemberCashRecord)
	var data []back.GetCashRecordList
	sess.Where("site_id = ?", this.SiteId)
	if len(this.Account) != 0 {
		sess.Where("user_name like?", "%"+this.Account+"%")
	}
	if this.SourceType != 0 {
		sess.Where("source_type = ?", this.SourceType)
	}
	if this.ClientType != 0 {
		sess.Where("client_type = ?", this.ClientType)
	}
	if this.ClientType != 0 {
		sess.Where("client_type = ?", this.ClientType)
	}
	if len(this.OrderId) != 0 {
		sess.Where("trade_no = ?", this.OrderId)
	}
	loc, _ := time.LoadLocation("Local")
	if this.StartTime != "" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", this.StartTime, loc)
		if err == nil {
			sess.Where(memberCashRecord.TableName()+".create_time>=?", t.Unix())
		}
	}
	if this.EndTime != "" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", this.EndTime, loc)
		if err == nil {
			sess.Where(memberCashRecord.TableName()+".create_time<=?", t.Unix())
		}
	}

	if this.PageSize == 50 || this.PageSize == 100 || this.PageSize == 200 {
		listParams.PageSize = this.PageSize
	}
	conds := sess.Conds()
	listParams.Make(sess)
	err = sess.Table(memberCashRecord.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, count, err
	}
	count, err = sess.Table(memberCashRecord.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, count, err
	}
	//总计  金额、
	sess.Select("SUM(sales_member_cash_record.balance) as a," +
		"COUNT(sales_member_cash_record.id) as b")
	var cashTotal back.CashRecordListTotalBack
	_, err = sess.Table(memberCashRecord.TableName()).
		Where("site_id=?", this.SiteId).Get(&cashTotal)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, count, err
	}
	var i float64
	if len(data) > 0 {
		for _, v := range data {
			i = i + v.Balance
		}
		cashTotal.SmallTotalMoney = i
		info.CashRecordListTotalBack = cashTotal
	}
	info.GetCashRecordList = data
	return
}

//得到一段时间内的现金记录(出入款账目汇总)
func (m *MemberCashRecordBean) GetCashList(siteId string, times *global.Times) (cashCollect []*back.CashCollectDetails, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Table(new(schema.SiteCashRecord).TableName())
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	times.Make("day_time", sess)
	err = sess.Find(&cashCollect)
	return
}

//出款账目汇总-给予优惠 (包括三个入款优惠加上注册优惠)
func (m *MemberCashRecordBean) GetDiscountCount(siteId string, sTime, eTime int64) (*back.CashCollectDetails, error) {
	cashDiscount := back.CashCollectDetails{}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//SELECT
	//count(*)     AS people_num,
	//	sum(t.money) AS money,
	//	sum(num)     AS num
	//FROM (SELECT
	//sum(dis_balance) AS money,
	//	count(*)            num
	//FROM sales_member_cash_record
	//GROUP BY member_id) t;
	subSql, err := sqlBuild.Select(new(schema.MemberCashRecord).TableName()).
		Column("sum(dis_balance) AS money,count(*) num").
		Where(siteId, "site_id").
		Where(sTime, "create_time>=").
		Where(eTime, "create_time<=").
		GroupBy("member_id").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashDiscount, err
	}
	sql, err := sqlBuild.Select("(" + subSql + ") t").
		Column("count(*) people_num,sum(t.num) num,sum(t.money) money").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashDiscount, err
	}
	_, err = sess.SQL(sql).Get(&cashDiscount)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashDiscount, err
	}
	return &cashDiscount, err
}

//wap查询交易记录
func (m *MemberCashRecordBean) GetWapCashRecord(siteId, siteIndexId, account string, this *input.WapMemberCashRecord, times *global.Times, listparams *global.ListParams) (data []schema.MemberCashRecord, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	cashrecord := new(schema.MemberCashRecord)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("user_name=?", account)
	sess.OrderBy("id desc")
	//if member.SourceType == 14 || member.SourceType == 15 || member.SourceType == 17
	//member.SourceType = member.SourceOneType
	switch this.SourceType {
	case 99:
	case 14, 15, 17:
		sess.Where("source_type=?", this.SourceOneType)
	default:
		sess.Where("source_type=?", this.SourceType)
	}
	if this.OrderNum != 0 {
		sess.Where("trade_no=?", this.OrderNum)
	}
	if times != nil {
		times.Make("create_time", sess)
	}
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(cashrecord.TableName()).Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(cashrecord.TableName()).Where(conds).Count()
	return
}

//获取入款监控数据
func (m *MemberCashRecordBean) GetWapCompanyRecord(siteId, siteIndexId, account string, memberId int64, this *input.WapMemberCashRecord, times *global.Times, listparams *global.ListParams) (data []schema.MemberCompanyIncome, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	cashrecord := new(schema.MemberCompanyIncome)
	sess.Where("site_id=?", siteId)
	sess.Where("site_index_id=?", siteIndexId)
	sess.Where("account=?", account)
	sess.Where("member_id=?", memberId)
	sess.OrderBy("id desc")
	if this.OrderNum != 0 {
		sess.Where("order_num=?", this.OrderNum)
	}
	if times != nil {
		times.Make("create_time", sess)
	}
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(cashrecord.TableName()).Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(cashrecord.TableName()).Where(conds).Count()
	return
}

//批量取消现金报表
func (*MemberCashRecordBean) DelMemberCashRecord(ids []int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mcr := new(schema.MemberCashRecord)
	count, err := sess.In("id", ids).Delete(mcr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//批量删除现金报表
func (*MemberCashRecordBean) PutMemberCashRecord(ids []int64) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mcr := new(schema.MemberCashRecord)
	mcr.DeleteTime = global.GetCurrentTime()
	count, err := sess.Cols("delete_time").In("id", ids).Update(mcr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
