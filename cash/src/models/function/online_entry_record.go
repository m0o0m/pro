package function

import (
	"time"

	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
)

type OnlineEntryRecordBean struct{}

type OnlineDeposit struct {
	Audit        *schema.MemberAudit
	Member       *schema.Member
	OnlineRecord *schema.OnlineEntryRecord
	Cash         *schema.MemberCashCount
	MemberRecord *schema.MemberCashRecord
	AuditStatus  int64 //如果总余额小于放宽额度，则把该会员的所有稽核状态更新为已处理
}

//确定一条线上入款
func (*OnlineEntryRecordBean) ConfirmOnlineDeposit(inputDeposit *input.ConfirmDeposit) (
	int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//首先更改入款状态
	sess.Begin()
	entryRecord := new(schema.OnlineEntryRecord)
	entryRecord.Status = 1
	count, err := sess.Where("id=?", inputDeposit.Id).Cols("status").Update(entryRecord)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//获取该条存款详情
	entry := new(OnlineEntryRecordBean)
	info, flag, err := entry.GetInfoById(inputDeposit.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if !flag {
		return count, err
	}
	//根据会员id，修改会员余额
	memberBean := new(MemberBean)
	balance := info.AmountDeposit + info.DepositDiscount + info.OtherDepositDiscount
	count, err = memberBean.ChangeBalance(info.MemberId, balance, 1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}

	//根据会员Id取出会员详情
	members := new(MemberBean)
	member, flags, err := members.GetInfoById(info.MemberId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if !flags {
		return 0, err
	}
	//更改member_cash_count 现金统计记录
	newCashCountBean := new(MemberCashCountBean)
	count, err = newCashCountBean.ChangeData(member.Id, info.AmountDeposit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return count, err
	}
	if count == 0 {
		return count, err
	}
	//更改现金系统记录
	newCashRecord := new(schema.MemberCashRecord)
	newCashRecord.CreateTime = time.Now().Unix()
	newCashRecord.SourceType = 2
	newCashRecord.TradeNo = info.ThirdOrderNumber
	newCashRecord.SiteId = member.SiteId
	newCashRecord.SiteIndexId = member.SiteIndexId
	newCashRecord.MemberId = info.MemberId
	newCashRecord.UserName = info.MemberAccount
	newCashRecord.Balance = info.AmountDeposit
	newCashRecord.AgencyId = info.AgencyId
	newCashRecord.ClientType = int64(info.SourceDeposit)
	newCashRecord.Type = 1
	newCashRecord.AgencyAccount = info.AgencyAccount
	newCashRecord.Remark = info.Remark
	newCashRecord.AfterBalance = info.AmountDeposit + info.DepositDiscount + info.OtherDepositDiscount
	memberCashRecord := new(MemberCashRecordBean)
	count, err = memberCashRecord.AddNewRecord(newCashRecord)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return 0, err
	}
	if count == 0 {
		return count, err
	}
	sess.Commit()
	return count, err
}

//根据线上入款id取出线上入款详情
func (*OnlineEntryRecordBean) GetInfoById(id int64) (info schema.OnlineEntryRecord, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Where("id=?", id).Get(&info)
	return
}

//根据线上入款订单号third_order_number取出线上入款详情
func (*OnlineEntryRecordBean) GetInfoByOrder(order string) (info schema.OnlineEntryRecord, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Where("third_order_number=?", order).Get(&info)
	return
}

//取消一条线上入款
//TODO 自动取消没有完成
func (*OnlineEntryRecordBean) CancleOneDeposit(newDeposit *input.CancleDeposit) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineDeposit := new(schema.OnlineEntryRecord)
	onlineDeposit.Status = 2
	count, err := sess.Where("id=?", newDeposit.Id).
		Cols("status").Update(onlineDeposit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//线上入款列表
func (*OnlineEntryRecordBean) DepositList(depositList *input.OnlineDepositList, times *global.Times, listparam *global.ListParams) (
	back.OnlineDepositBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	deposite := new(schema.OnlineEntryRecord)
	var infolist back.OnlineDepositBack
	agency := new(schema.Agency)
	if depositList.SourceDeposit != 0 { //入款来源
		sess.Where(deposite.TableName()+".source_deposit=?", depositList.SourceDeposit)
	}
	if depositList.ThirdId != "" { //入款商户
		//分割字符串
		thirdId := strings.Split(depositList.ThirdId, ",")
		sess.In(deposite.TableName()+".third_id", thirdId)
	}
	//金额
	if depositList.LowerLimitMoney >= 0 && depositList.UpperLimitMoney != 0 { //最大值最小值都有
		sess.Where(deposite.TableName()+".amount_deposit>?", depositList.LowerLimitMoney).
			And(deposite.TableName()+".amount_deposit<?", depositList.UpperLimitMoney)
	}
	//状态
	if depositList.Status != 0 {
		sess.Where(deposite.TableName()+".status=?", depositList.Status)
	}
	//站点
	if depositList.SiteId != "" {
		sess.Where(deposite.TableName()+".site_id=?", depositList.SiteId)
	}
	if depositList.SiteIndexId != "" {
		sess.Where(deposite.TableName()+".site_index_id=?", depositList.SiteIndexId)
	}
	//代理账号
	if depositList.AgencyId != 0 {
		sess.Where(deposite.TableName()+".agency_id=?", depositList.AgencyId)
	}
	//层级
	if depositList.Level != "" {
		sess.Where(deposite.TableName()+".level=?", depositList.Level)
	}
	if depositList.SelectBy == 1 { //会员账号
		sess.Where(deposite.TableName()+".member_account=?", depositList.Conditions)
	} else if depositList.SelectBy == 2 { //手动订单号模糊搜索
		sess.Where(deposite.TableName()+".third_order_number=?", depositList.Conditions)
	}
	if depositList.IsDiscount == 1 {
		sess.Where(deposite.TableName() + ".is_discount=1")
	} else if depositList.IsDiscount == 2 {
		sess.Where(deposite.TableName() + ".is_discount=2")
	}
	conds := sess.Conds()
	info := make([]back.OnlineListDeposit, 0)
	oit := new(schema.OnlineIncomeThird)
	sql1 := fmt.Sprintf("%s.paid_type = %s.id", deposite.TableName(), oit.TableName())
	sql2 := fmt.Sprintf("%s.operate_id = %s.id", deposite.TableName(), agency.TableName())
	err := sess.Table(deposite.TableName()).
		Join("LEFT", oit.TableName(), sql1).
		Join("LEFT", agency.TableName(), sql2).
		Find(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, err
	}
	//入款列表返回的数据
	var data []back.OnlineListDeposit
	var totalMoney float64               //总计存入金额
	var totalDepositDiscount float64     //总计存款优惠
	var totalOtherDiscount float64       //总计其他优惠
	var totalDeposit float64             //总计存入总金额
	var pageTotalMoney float64           //小计存入金额
	var pageTotalDepositDiscount float64 //小计存款优惠
	var pageTotalOtherDiscount float64   //小计其他优惠
	var pageTotalDeposit float64         //小计存入总金额
	for k := range info {
		//总计
		totalMoney += info[k].AmountDeposit
		totalDepositDiscount += info[k].DepositDiscount
		totalOtherDiscount += info[k].OtherDepositDiscount
		totalDeposit += info[k].AmountDeposit + info[k].DepositDiscount + info[k].OtherDepositDiscount
	}
	//获得分页记录
	listparam.Make(sess)
	times.Make(deposite.TableName()+".create_time", sess)
	onlineListDeposit := make([]back.OnlineListDeposit, 0)
	//重新传入表名和where条件查询记录
	err = sess.Table(deposite.TableName()).
		Join("LEFT", oit.TableName(), sql1).
		Join("LEFT", agency.TableName(), sql2).
		Where(conds).
		OrderBy("sales_online_entry_record.create_time DESC").
		Find(&onlineListDeposit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, err
	}
	var old back.OnlineListDeposit
	for k := range onlineListDeposit {
		old.Id = onlineListDeposit[k].Id
		old.ThirdOrderNumber = onlineListDeposit[k].ThirdOrderNumber
		old.SourceDeposit = onlineListDeposit[k].SourceDeposit
		old.CreateTime = onlineListDeposit[k].CreateTime
		old.AgencyAccount = onlineListDeposit[k].AgencyAccount
		old.MemberAccount = onlineListDeposit[k].MemberAccount
		old.ThirdPayTime = onlineListDeposit[k].ThirdPayTime
		old.AmountDeposit = onlineListDeposit[k].AmountDeposit
		old.Status = onlineListDeposit[k].Status
		old.Title = onlineListDeposit[k].Title
		old.IsFirstDeposit = onlineListDeposit[k].IsFirstDeposit
		old.Account = onlineListDeposit[k].Account
		//小计
		pageTotalMoney += onlineListDeposit[k].AmountDeposit
		pageTotalDepositDiscount += onlineListDeposit[k].DepositDiscount
		pageTotalOtherDiscount += onlineListDeposit[k].OtherDepositDiscount
		pageTotalDeposit += onlineListDeposit[k].AmountDeposit + onlineListDeposit[k].DepositDiscount + onlineListDeposit[k].OtherDepositDiscount
		data = append(data, old)
	}

	//获得符合条件的记录数
	count, err := sess.Table(deposite.TableName()).
		Join("LEFT", oit.TableName(), sql1).
		Join("LEFT", agency.TableName(), sql2).
		Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, err
	}
	infolist.AllData = data
	infolist.TotalMoney = totalMoney
	infolist.TotalDeposit = totalDeposit
	infolist.TotalDepositDiscount = totalDepositDiscount
	infolist.TotalOtherDiscount = totalOtherDiscount
	infolist.PageTotalMoney = pageTotalMoney
	infolist.PageTotalDeposit = pageTotalDeposit
	infolist.PageTotalDepositDiscount = pageTotalDepositDiscount
	infolist.PageTotalOtherDiscount = pageTotalOtherDiscount
	infolist.TotalCount = count
	infolist.PageCount = len(onlineListDeposit)
	return infolist, err
}

//获取某个线上支付设定下面所有的存款记录
func (*OnlineEntryRecordBean) GetRecordByPaidId(selectDis *input.GetInfoSetupDeposit, listParam *global.ListParams) (
	[]back.OnePaidSetupBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var infolist []back.OnePaidSetupBack
	newDepositRecord := new(schema.OnlineEntryRecord)
	if selectDis.PaidSetup > 0 {
		sess.Where("paid_setup_id=?", selectDis.PaidSetup)
	}
	if selectDis.OrderNumber != "" {
		sess.Where("third_order_number=?", selectDis.OrderNumber)
	}
	if selectDis.StartTime != "" {
		if selectDis.EndTime != "" {
			sess.Where("third_pay_time>? and third_pay_time<?", selectDis.StartTime, selectDis.EndTime)
		} else {
			sess.Where("third_pay_time>?", selectDis.StartTime)
		}
	} else {
		if selectDis.EndTime != "" {
			sess.Where("third_pay_time<?", selectDis.EndTime)
		}
	}
	conds := sess.Conds()
	//获得分页记录
	listParam.Make(sess)
	err := sess.Table(newDepositRecord.TableName()).Find(&infolist)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, 0, err
	}
	//获得符合条件的记录数
	count, err := sess.Table(newDepositRecord.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return infolist, count, err
	}
	return infolist, count, err
}

//获取一个线上入款详情
func (*OnlineEntryRecordBean) GetInfo(id int64) (info schema.OnlineEntryRecord, flag bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	flag, err = sess.Table(info.TableName()).Where("id=?", id).Get(&info)
	return
}

//获取拒绝，取消出款原因
func (*OnlineEntryRecordBean) OutRemark(this *input.OutRemark) (back.OutRemark, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var info back.OutRemark
	makeMoney := new(schema.MakeMoney)
	_, err := sess.Table(makeMoney.TableName()).Where("id=?", this.Id).Get(&info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return info, err
	}
	return info, err
}

//监控页面获取没有确认得线上入款
func (*OnlineEntryRecordBean) GetNotConfirm(site, siteIndex string) (infolist []schema.OnlineEntryRecord, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	newOnline := new(schema.OnlineEntryRecord)
	err = sess.Table(newOnline.TableName()).Where("status=?", 2).Where("site_id=?", site).Where("site_index_id=?", siteIndex).Find(&infolist)
	count, err = sess.Table(newOnline.TableName()).Where("status=?", 2).Where("site_id=?", site).Where("site_index_id=?", siteIndex).Count()
	return
}

//添加存款纪录
func (*OnlineEntryRecordBean) Add(this *schema.OnlineEntryRecord) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	count, err := sess.Insert(this)
	return count, err
}

//查看某个会员是否是首次存款
func (*OnlineEntryRecordBean) IsFirst(id int64) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	onlineEntryRecord.Id = id
	have, err := sess.Exist(onlineEntryRecord)
	return have, err
}

//查看该站点下某个会员是否是首次存款
func (*OnlineEntryRecordBean) IsFirstIncome(id int64, siteId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	onlineEntryRecord.Id = id
	onlineEntryRecord.SiteId = siteId
	have, err := sess.Exist(onlineEntryRecord)
	return have, err
}

//会员存款成功后用事务更新多表的数据
func (*OnlineEntryRecordBean) Update(this *OnlineDeposit) (err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//开启事务
	sess.Begin()
	//更新会员余额
	_, err = sess.Where("id=?", this.Member.Id).Cols("balance").Update(this.Member)
	if err != nil {
		defer sess.Rollback()
		return
	}
	//更新公司入款纪录的状态和优惠金额
	_, err = sess.Where("id=?", this.OnlineRecord.Id).Cols("status,third_pay_time").Update(this.OnlineRecord)
	if err != nil {
		defer sess.Rollback()
		return
	}
	//添加会员现金流水记录
	_, err = sess.Insert(this.MemberRecord)
	if err != nil {
		defer sess.Rollback()
		return
	}
	//更新会员现金统计
	//1:获取会员现金纪录,如果没有,则插入,否则更新数据
	if this.Cash.DepositCount == 1 { //没有，插入
		_, err = sess.Insert(this.Cash)
	} else {
		_, err = sess.Where("member_id=?", this.Cash.Member).Update(this.Cash)
	}
	if err != nil {
		defer sess.Rollback()
		return
	}
	//更新会员的上一条稽核纪录的结束时间
	if this.Audit.Id > 0 { //判断上一条稽核记录是否已处理，此条件为未处理
		_, err = sess.Where("id=?", this.Audit.Id).Cols("end_time").Update(&schema.MemberAudit{EndTime: this.Audit.BeginTime})
		if err != nil {
			defer sess.Rollback()
			return
		}
		this.Audit.Id = 0 //把要新插入的稽核记录id清0
	}
	if this.AuditStatus == 2 { //更新该会员所有稽核状态为已处理
		_, err = sess.Where("site_id=?", this.Audit.SiteId).Where("member_id=?", this.Audit.MemberId).Where("status=?", 1).Cols("status").Update(&schema.MemberAudit{Status: 2})
		if err != nil {
			defer sess.Rollback()
			return
		}
	}
	//插入一条新的会员稽核记录
	_, err = sess.Insert(this.Audit)
	if err != nil {
		defer sess.Rollback()
		return
	}
	sess.Commit()
	return
}
