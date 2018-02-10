package function

import (
	"github.com/golyu/sql-build"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"
)

type MakeMoneyBean struct{}

//出款列表
func (*MakeMoneyBean) OutMoney(this *input.OutMoneyList, times *global.Times, listParams *global.ListParams) (
	back.OutMoneyBackList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	makeMoney := new(schema.MakeMoney)
	var data back.OutMoneyBackList
	//判断并组合where条件
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.Level != "" {
		sess.In("level_id", this.Level)
	}
	if this.OutStatus != 0 {
		sess.Where("out_status = ?", this.OutStatus)
	}
	if this.ClientType != 0 {
		sess.Where("client_type = ?", this.ClientType)
	}
	if this.AgencyAccount != "" {
		sess.Where("agency_account = ?", this.AgencyAccount)
	}
	if this.AgencyId != 0 {
		sess.Where("agency_id=?", this.AgencyId)
	}
	if this.LowerLimit != 0 && this.UpperLimit != 0 {
		sess.Where("outward_money < ?", this.UpperLimit).And("outward_money > ?", this.LowerLimit)
	}
	if this.SelectBy == 1 { //会员账号
		sess.Where("user_name = ?", this.Conditions)
	} else if this.SelectBy == 2 { //操作者账号
		sess.Where("do_agency_account=?", this.Conditions)
	}
	if this.Automatic != 0 {
		sess.Where("is_underhair=?", this.Automatic)
	}
	//根据时间段查询
	times.Make("create_time", sess)
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	outMoneyList := make([]back.OutMoneyList, 0)
	err := sess.Table(makeMoney.TableName()).Find(&outMoneyList)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	var totalCharge float64              //总计手续费
	var totalFavourableMoney float64     //总计优惠金额
	var totalExpeneseMoney float64       //总计行政费
	var totalOutwardMoney float64        //总计实际总出款额
	var pageTotalCharge float64          //小计手续费
	var pageTotalFavourableMoney float64 //小计优惠金额
	var pageTotalExpeneseMoney float64   //小计行政费
	var pageTotalOutwardMoney float64    //小计实际总出款额
	for k := range outMoneyList {
		//总计
		totalCharge += outMoneyList[k].Charge
		totalFavourableMoney += outMoneyList[k].FavourableMoney
		totalExpeneseMoney += outMoneyList[k].ExpeneseMoney
		totalOutwardMoney += outMoneyList[k].OutwardMoney
	}
	//获得分页记录
	listParams.Make(sess)
	oml := make([]back.OutMoneyList, 0)
	//重新传入表名和where条件查询记录
	err = sess.Table(makeMoney.TableName()).Where(conds).OrderBy("create_time DESC").Find(&oml)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	var omlt back.OutMoneyList
	var omlts []back.OutMoneyList
	for k := range oml {
		omlt.Id = oml[k].Id
		omlt.AgencyId = oml[k].AgencyId
		omlt.Charge = oml[k].Charge
		omlt.OutwardMoney = oml[k].OutwardMoney
		omlt.ExpeneseMoney = oml[k].ExpeneseMoney
		omlt.FavourableMoney = oml[k].FavourableMoney
		omlt.AgencyAccount = oml[k].AgencyAccount
		omlt.Remark = oml[k].Remark
		omlt.SiteIndexId = oml[k].SiteIndexId
		omlt.Balance = oml[k].Balance
		omlt.SiteId = oml[k].SiteId
		omlt.DoAgencyAccount = oml[k].DoAgencyAccount
		omlt.MemberAccount = oml[k].MemberAccount
		omlt.FavourableOut = oml[k].FavourableOut
		omlt.IsFirst = oml[k].IsFirst
		omlt.OutRemark = oml[k].OutRemark
		omlt.OutStatus = oml[k].OutStatus
		omlt.IsAutoOut = oml[k].IsAutoOut
		omlt.IsUnderhair = oml[k].IsUnderhair
		omlt.OutTime = oml[k].OutTime
		omlt.OutwardNum = oml[k].OutwardNum
		omlt.Createtime = oml[k].Createtime
		//小计
		pageTotalCharge += oml[k].Charge
		pageTotalFavourableMoney += oml[k].FavourableMoney
		pageTotalExpeneseMoney += oml[k].ExpeneseMoney
		pageTotalOutwardMoney += oml[k].OutwardMoney
		omlts = append(omlts, omlt)
	}
	//获得符合条件的记录数
	count, err := sess.Table(makeMoney.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	data.OutMoneyList = omlts
	data.TotalCount = count //总计笔数
	data.TotalCharge = totalCharge
	data.TotalFavourableMoney = totalFavourableMoney
	data.TotalExpeneseMoney = totalExpeneseMoney
	data.TotalOutwardMoney = totalOutwardMoney
	data.PageTotalCount = len(oml) //小计笔数
	data.PageTotalCharge = pageTotalCharge
	data.PageTotalFavourableMoney = pageTotalFavourableMoney
	data.PageTotalExpeneseMoney = pageTotalExpeneseMoney
	data.PageTotalOutwardMoney = pageTotalOutwardMoney
	return data, err
}

//确定出款
func (*MakeMoneyBean) ConfirmOutMoney(confirm *input.ConfirmOutMoney, agency schema.Agency) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	//更改出款时间和出款状态
	makeMonBean := new(MakeMoneyBean)
	makeMoney := new(schema.MakeMoney)
	newCash := new(MemberCashCountBean)
	makeMoney.OutStatus = 1
	makeMoney.OutTime = time.Now().Unix()
	makeMoney.DoAgencyId = confirm.AgencyId
	makeMoney.DoAgencyAccount = agency.Account
	count, err = sess.Where("id=?", confirm.Id).Cols("out_status,out_time,do_agency_id,do_agency_account").Update(makeMoney)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//修改会员现金统计表
	info, flag, err := makeMonBean.GetInfoById(confirm.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if !flag {
		return 0, err
	}
	count, err = newCash.ConfirmOutMoneyChangeData(info.MemberId, info.OutwardNum)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if count == 0 {
		return count, err
	}
	//稽核状态批量修改为已经稽核
	newAudit := new(schema.MemberAudit)
	newAudit.EndTime = time.Now().Unix()
	newAudit.Status = 1
	newAudit.MemberId = info.MemberId
	newAudit.SiteId = info.SiteId
	newAudit.SiteIndexId = info.SiteIndexId
	memberAudit := new(MemberAuditBean)
	count, err = memberAudit.ChangeAuditSAT(newAudit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return
}

//取消出款
func (*MakeMoneyBean) CancleOutMoney(cancle *input.CancleOutMoney, money float64, agency schema.Agency) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//更改状态
	sess.Begin()
	makeMoneyBean := new(MakeMoneyBean)
	makeMoney := new(schema.MakeMoney)
	makeMoney.OutStatus = 3
	makeMoney.OutTime = time.Now().Unix()
	makeMoney.DoAgencyId = cancle.AgencyId
	makeMoney.OutRemark = cancle.Reason
	makeMoney.DoAgencyAccount = agency.Account
	count, err := sess.Where("id=?", cancle.Id).Cols("out_remark,out_status,out_time,do_agency_id,do_agency_account").Update(makeMoney)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	//根据id获取出款详情
	info, flag, err := makeMoneyBean.GetInfoById(cancle.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if !flag {
		return 0, err
	}
	//更改会员余额,取消出款将这笔钱给他加回去
	memberBean := new(MemberBean)
	count, err = memberBean.ChangeBalance(info.MemberId, money, 1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}

	//会员现金流水
	newCashRecord := new(schema.MemberCashRecord)
	newCashRecord.SiteIndexId = info.SiteIndexId
	newCashRecord.SiteId = info.SiteId
	newCashRecord.MemberId = info.MemberId
	newCashRecord.UserName = info.UserName
	newCashRecord.AgencyId = info.AgencyId
	newCashRecord.AgencyAccount = info.AgencyAccount
	newCashRecord.SourceType = 8
	newCashRecord.TradeNo = ""
	newCashRecord.Type = 1
	newCashRecord.Balance = info.OutwardNum                     //操作金额
	newCashRecord.AfterBalance = info.Balance + info.OutwardNum //操作后的余额
	newCashRecord.Remark = "取消出款添加"
	newCashRecord.ClientType = 0
	newCashRecord.CreateTime = time.Now().Unix()
	memberCashBean := new(MemberCashRecordBean)
	count, err = memberCashBean.AddNewRecord(newCashRecord)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return count, err
}

//拒绝出款
func (*MakeMoneyBean) RefuseOutMoney(refuse *input.RefuseOutMoney, agencyAccount string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	makeMoney := new(schema.MakeMoney)
	makeMoney.OutStatus = 4
	makeMoney.OutRemark = refuse.Reason
	makeMoney.DoAgencyId = refuse.AgencyId
	makeMoney.DoAgencyAccount = agencyAccount
	makeMoney.OutTime = time.Now().Unix()
	count, err := sess.Where("id=?", refuse.Id).
		Cols("out_status,out_time,out_remark,do_agency_id,do_agency_account").
		Update(makeMoney)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//预备出款
func (*MakeMoneyBean) PrepareOutMoney(prepare *input.PrepareOutMoney, agency schema.Agency, user *global.RedisStruct) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	makeMoney := new(schema.MakeMoney)
	makeMoney.OutStatus = 2
	makeMoney.DoAgencyId = prepare.AgencyId
	makeMoney.DoAgencyAccount = agency.Account
	makeMoney.OutTime = time.Now().Unix()
	count, err := sess.Table(makeMoney.TableName()).
		Where("id=?", prepare.Id).
		Where("site_id=?", user.SiteId).
		Where("site_index_id=?", user.SiteIndexId).
		Cols("out_time,out_status,do_agency_id,do_agency_account").
		Update(makeMoney)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//根据id获取会员出款详情
func (*MakeMoneyBean) GetInfoById(id int64) (schema.MakeMoney, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info schema.MakeMoney
	flag, err := sess.Table(info.TableName()).Where("id=?", id).Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, flag, err
	}
	return info, flag, err
}

//根据出款状态和会员id获取看是否存在存款成功的记录
func (*MakeMoneyBean) GetStatus(id int64) (flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(schema.MakeMoney)
	flag, err = sess.Table(info.TableName()).Where("member_id=?", id).Where("out_status=?", 1).Exist()
	return
}

//获取没有操作或者处于待操作状态的记录
func (*MakeMoneyBean) GetOperateRecord(site, siteIndex string) (infolist []schema.MakeMoney, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	info := new(schema.MakeMoney)
	err = sess.Table(info.TableName()).Where("out_status=?", 2).Or("out_status=?", 5).Where("site_id", site).Where("site_index_id", siteIndex).Find(&infolist)
	count, err = sess.Table(info.TableName()).Where("out_status=?", 2).Or("out_status=?", 5).Where("site_id", site).Where("site_index_id", siteIndex).Count()
	return

}

//查询出款账目汇总的统计信息--会员出款 (总额度,总笔数,总人数)
func (m *MakeMoneyBean) GetMakeMoneyCount(siteId string, sTime, eTime int64) (*back.CashCollectDetails, error) {
	cashTake := back.CashCollectDetails{SourceType: 4}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//SELECT
	//count(*)        people_num,
	//	sum(t.num)   AS num,
	//	sum(t.money) AS money
	//FROM (SELECT
	//sum(outward_money) AS money,count(member_id) num
	//FROM sales_make_money
	//GROUP BY member_id) t;

	subSql, err := sqlBuild.Select(new(schema.MakeMoney).TableName()).
		Column("sum(outward_money) AS money,count(member_id) num").
		Where(siteId, "site_id").
		Where(sTime, "out_time>=").
		Where(eTime, "out_time<=").
		GroupBy("member_id").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashTake, err
	}
	sql, err := sqlBuild.Select("(" + subSql + ") t").
		Column("count(*) people_num,sum(t.num) num,sum(t.money) money").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashTake, err
	}
	_, err = sess.SQL(sql).Get(&cashTake)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashTake, err
	}
	return &cashTake, err
}

//查询出款账目汇总的统计信息--出款被扣金额 (总额度,总笔数,总人数)
func (m *MakeMoneyBean) GetTakeOutCount(siteId string, sTime, eTime int64) (*back.CashCollectDetails, error) {
	var cashTakeOut = back.CashCollectDetails{SourceType: 3}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//SELECT count(*) people_num,sum(t.num) num,sum(t.money) money
	//FROM
	//(SELECT sum(outward_num - outward_money) as money,count(member_id) num
	//FROM sales_make_money
	//WHERE outward_num > outward_money
	//GROUP BY member_id) t;
	subSql, err := sqlBuild.Select(new(schema.MakeMoney).TableName()).
		Column("sum(outward_num - outward_money) as money,count(member_id) num").
		Where(siteId, "site_id").
		WhereFunc("outward_money", "outward_num > ").
		Where(sTime, "out_time>=").
		Where(eTime, "out_time<=").
		GroupBy("member_id").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashTakeOut, err
	}
	sql, err := sqlBuild.Select("(" + subSql + ") t").
		Column("count(*) people_num,sum(t.num) num,sum(t.money) money").
		String()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashTakeOut, err
	}
	_, err = sess.SQL(sql).Get(&cashTakeOut)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return &cashTakeOut, err
	}
	return &cashTakeOut, err
}
