//出入款查询[admin]
package function

import (
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type BankInOutBean struct{}

//公司入款
func (*BankInOutBean) DepositByCompany(this *input.InDeposit, listparam *global.ListParams, times *global.Times) (data []back.InDepositBack,
	count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mci := new(schema.MemberCompanyIncome)
	agency := new(schema.Agency)
	if this.Account != "" {
		sess.Where(mci.TableName()+".account=?", this.Account)
	}
	if this.OrderNum != "" {
		sess.Where(mci.TableName()+".order_num=?", this.OrderNum)
	}
	if this.SiteId != "" {
		sess.Where(mci.TableName()+".site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where(mci.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	if this.Equipment != 0 {
		sess.Where(mci.TableName()+".client_type=?", this.Equipment)
	}
	//根据时间段查询
	times.Make("sales_member_company_income.update_time", sess)
	listparam.Make(sess)
	conds := sess.Conds()
	err = sess.Table(mci.TableName()).
		Join("LEFT", agency.TableName(), mci.TableName()+".do_agency_id ="+agency.TableName()+".id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(mci.TableName()).
		Join("LEFT", agency.TableName(), mci.TableName()+".do_agency_id ="+agency.TableName()+".id").
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//线上入款
func (*BankInOutBean) DepositByOnline(this *input.InDeposit, listparam *global.ListParams, times *global.Times) (data []back.InDepositOnlineBack,
	count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mci := new(schema.OnlineEntryRecord)
	agency := new(schema.Agency)
	member := new(schema.Member)
	third := new(schema.OnlineIncomeThird)
	if this.Account != "" {
		sess.Where(mci.TableName()+".member_account=?", this.Account)
	}
	if this.OrderNum != "" {
		sess.Where(mci.TableName()+".third_order_number=?", this.OrderNum)
	}
	if this.Equipment != 0 {
		sess.Where(mci.TableName()+".source_deposit=?", this.Equipment)
	}
	//根据时间段查询
	times.Make("sales_online_entry_record.create_time", sess)
	conds := sess.Conds()
	listparam.Make(sess)
	err = sess.Table(mci.TableName()).
		Join("LEFT", agency.TableName(), mci.TableName()+".operate_id ="+agency.TableName()+".id").
		Join("LEFT", member.TableName(), mci.TableName()+".member_id="+member.TableName()+".id").
		Join("LEFT", third.TableName(), mci.TableName()+".third_id="+third.TableName()+".id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(mci.TableName()).
		Join("LEFT", agency.TableName(), mci.TableName()+".operate_id ="+agency.TableName()+".id").
		Join("LEFT", member.TableName(), mci.TableName()+".member_id="+member.TableName()+".id").
		Join("LEFT", third.TableName(), mci.TableName()+".third_id="+third.TableName()+".id").
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//出款查询
func (*BankInOutBean) OutDepositSearch(this *input.OutDeposit, listparams *global.ListParams, times *global.Times) (data []back.OutDepositBack, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	out := new(schema.MakeMoney)
	audit := new(schema.MemberAudit)
	agency := new(schema.Agency)
	if this.Account != "" {
		sess.Where(out.TableName()+".user_name=?", this.Account)
	}
	if this.Equipment != 0 {
		sess.Where(out.TableName()+".client_type=?", this.Equipment)
	}
	if this.SiteId != "" {
		sess.Where(out.TableName()+".site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where(out.TableName()+".site_index_id=?", this.SiteIndexId)
	}
	times.Make("sales_make_money.out_time", sess)
	conds := sess.Conds()
	listparams.Make(sess)
	err = sess.Table(out.TableName()).
		Join("LEFT", audit.TableName(), out.TableName()+".create_time="+audit.TableName()+".end_time").
		Join("LEFT", agency.TableName(), out.TableName()+".do_agency_id="+agency.TableName()+".id").
		Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(out.TableName()).
		Join("LEFT", audit.TableName(), out.TableName()+".create_time="+audit.TableName()+".end_time").
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//得到出款信息
func (m *BankInOutBean) GetOutList(siteId string, times *global.Times) (money []*schema.MakeMoney, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Table(new(schema.MakeMoney).TableName())
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	times.Make("out_time", sess)
	err = sess.Find(&money)
	return
}
