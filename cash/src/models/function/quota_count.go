package function

import (
	"fmt"
	"global"
	"math"
	"models/back"
	"models/input"
	"models/schema"
)

type QuotaCountBean struct{}

//额度统计列表
func (*QuotaCountBean) QuotaCountList(this *input.QuotaCountList, times *global.Times) (*back.AllList, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sL := new(schema.SiteLevel)
	list := new(back.AllList)
	var sLl schema.SiteLevel
	if this.SiteId == "" {
		return nil, nil
	}
	has, err := sess.Table(sL.TableName()).
		Where("site_level like?", "%"+this.SiteId+"%").Get(&sLl)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	if !has {
		return nil, err
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, err
	}
	sess.Where("sales_member_balance_conversion.site_id=?", this.SiteId)

	if this.Account != "" {
		sess.Where("sales_member_balance_conversion.account=?", this.Account)
	}
	mbc := new(schema.MemberBalanceConversion)
	slP := new(schema.SiteLevelPlatform)
	pla := new(schema.Platform)
	//根据时间段查询
	times.Make("sales_member_balance_conversion.create_time", sess)
	sess.Where(slP.TableName()+".level_id=?", sLl.Id)
	conds := sess.Conds()

	var mbc_income []back.QuotaCountListBack
	//转入
	sess.Select("count(sales_member_balance_conversion.from_type) as a,SUM(sales_member_balance_conversion.money) as b,sales_platform.platform,sales_site_level_platform.proportion")
	err = sess.Table(mbc.TableName()).
		Join("LEFT", slP.TableName(),
			mbc.TableName()+".from_type ="+slP.TableName()+".platform_id").
		Join("LEFT", pla.TableName(), mbc.TableName()+".from_type ="+
			pla.TableName()+".id").
		Where("sales_member_balance_conversion.from_type = sales_site_level_platform.platform_id!=?", 0).
		GroupBy("sales_platform.platform,sales_site_level_platform.proportion").
		Find(&mbc_income)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	var da back.AllQuotaCountListBack
	var dataA []back.AllQuotaCountListBack
	if len(mbc_income) > 0 {
		for _, v := range mbc_income {
			da.RetreatNum = v.Num
			da.Platform = v.Platform
			da.RetreatMoney = v.Money
			da.Proportion = v.Proportion
			da.TurnMoney = 0
			da.TurnNum = 0
			dataA = append(dataA, da)
		}
	}

	//转出
	var mbc_out []back.QuotaCountListBack
	sess.Select("count(sales_member_balance_conversion.from_type) as a,SUM(sales_member_balance_conversion.money) as b,sales_platform.platform,sales_site_level_platform.proportion")
	err = sess.Table(mbc.TableName()).
		Join("LEFT", slP.TableName(),
			mbc.TableName()+".for_type ="+slP.TableName()+".platform_id").
		Join("LEFT", pla.TableName(), mbc.TableName()+".for_type ="+
			pla.TableName()+".id").
		Where("sales_member_balance_conversion.for_type = sales_site_level_platform.platform_id!=?", 0).
		Where(conds).
		GroupBy("sales_platform.platform,sales_site_level_platform.proportion").
		Find(&mbc_out)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	if len(dataA) > 0 {
		for k, v := range dataA {
			if len(mbc_out) > 0 {
				for _, vv := range mbc_out {
					if vv.Platform == v.Platform {
						dataA[k].TurnNum = vv.Num
						dataA[k].TurnMoney = vv.Money
					} else {
						da.RetreatNum = 0
						da.RetreatMoney = 0
						da.ResultMoney = 0
						da.RetreatNum = 0
						da.Proportion = vv.Proportion
						da.Platform = vv.Platform
						da.TurnMoney = vv.Money
						da.TurnNum = vv.Num
						dataA = append(dataA, da)
					}
				}
			}

		}
	} else {
		if len(mbc_out) > 0 {
			for _, hg := range mbc_out {
				da.RetreatNum = 0
				da.RetreatMoney = 0
				da.ResultMoney = 0
				da.RetreatNum = 0
				da.Proportion = hg.Proportion
				da.Platform = hg.Platform
				da.TurnMoney = hg.Money
				da.TurnNum = hg.Num
				dataA = append(dataA, da)
			}
		}
	}
	var total back.QuotaTotal
	if len(dataA) > 0 {
		for k, v := range dataA {
			if v.Proportion != 0 {
				dataA[k].WalletAdd = Round(v.Proportion/100*v.TurnMoney, 2)
				dataA[k].WalletReduce = Round(v.Proportion/100*v.RetreatMoney, 2)

			} else {
				dataA[k].WalletAdd = Round(10/100*v.TurnMoney, 2)
				dataA[k].WalletReduce = Round(10/100*v.ResultMoney, 2)
			}
			dataA[k].ResultMoney = v.TurnMoney - v.RetreatMoney
			dataA[k].ResultRatio = dataA[k].WalletAdd - dataA[k].WalletReduce
			total.TurnMoneyTotal = total.TurnMoneyTotal + v.TurnMoney
			total.RetreatMoneyTotal = total.RetreatMoneyTotal + v.RetreatMoney
			total.TurnNumTotal = total.TurnNumTotal + v.TurnNum
			total.RetreatNumTotal = total.RetreatNumTotal + v.RetreatNum
			total.ResultMoneyTotal = total.ResultMoneyTotal + dataA[k].ResultMoney
			total.ResultRatioTotal = total.ResultRatioTotal + dataA[k].ResultRatio
			total.WalletAddTotal = total.WalletAddTotal + dataA[k].WalletAdd
			total.WalletReduceTotal = total.WalletReduceTotal + dataA[k].WalletReduce
		}
	}
	list.Data = dataA
	list.Total = total
	return list, err
}

//额度记录列表
func (*QuotaCountBean) QuotaRecordList(this *input.QuotaRecordList, params *global.ListParams, times *global.Times) (
	back.QuotaRecordBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var list back.QuotaRecordBack
	siteCashRecord := new(schema.SiteCashRecord)
	//判断并组合where条件
	if this.SiteId != "" {
		sess.Where("site_id = ?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id = ?", this.SiteIndexId)
	}
	if this.AdminName != "" {
		sess.Where("admin_name = ?", this.AdminName)
	}
	if this.State != 0 {
		sess.Where("state = ?", this.State)
	}
	if this.DoType != 0 {
		sess.Where("do_type = ?", this.DoType)
	}
	if this.CashType != 0 {
		sess.Where("cash_type = ?", this.CashType)
	}
	if this.VdType != 0 {
		sess.Where("vd_type = ?", this.VdType)
	}
	//把where条件取出来保存，否则后面任意一次操作都会将where条件清空，不能用同样的条件做多次查询
	conds := sess.Conds()
	qrd := make([]back.QuotaRecord, 0)
	err := sess.Table(siteCashRecord.TableName()).Select("money").Find(&qrd)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, 0, err
	}
	var totalMoney float64 //总计金额
	for k := range qrd {
		totalMoney += qrd[k].Money
	}
	//获得分页记录
	params.Make(sess)
	times.Make("create_time", sess)
	qr := make([]back.QuotaRecord, 0)
	platform := new(schema.Platform)
	sql := fmt.Sprintf("%s.id = %s.vd_type", platform.TableName(), siteCashRecord.TableName())
	err = sess.Table(siteCashRecord.TableName()).Join("LEFT", platform.TableName(), sql).Where(conds).Find(&qr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, 0, err
	}
	var subTotalMoney float64 //小计金额
	for k := range qr {
		subTotalMoney += qr[k].Money
	}
	//获得符合条件的记录数
	count, err := sess.Table(siteCashRecord.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, count, err
	}
	list.TotalCount = len(qrd)
	list.TotalMoney = totalMoney
	list.SubtotalMoney = subTotalMoney
	list.QuotaRecord = qr
	return list, count, err
}

//额度充值记录
func (*QuotaCountBean) QuotaRecList(this *input.QuotaRecord, lisparms *global.ListParams, times *global.Times) (
	[]back.QuotaReBack, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.QuotaReBack
	if this.Type != 0 {
		sess.Where("type=?", this.Type)
	}
	if this.Status != 0 {
		sess.Where("state=?", this.Status)
	}
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	//根据时间段查询
	times.Make("update_time", sess)
	conds := sess.Conds()
	lisparms.Make(sess)
	spr := new(schema.SitePayRecord)
	err := sess.Table(spr.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
	}
	count, err := sess.Table(spr.TableName()).Where(conds).Count()
	err = sess.Table(spr.TableName()).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	count, err = sess.Table(spr.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//修改额度充值记录状态
func (m *QuotaCountBean) QuotaRecordUpdate(sitePayRecord *input.SitePayRecordUpdate) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	spr := new(schema.SitePayRecord)
	spr.AdminUser = sitePayRecord.AdminUser
	spr.State = sitePayRecord.State
	if sitePayRecord.Remark != "" {
		spr.Remark = sitePayRecord.Remark
		sess.Cols("remark")
	}
	count, err := sess.ID(sitePayRecord.Id).Cols("admin_user,state").Update(spr)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//入款管理-添加或修改第三方或者银行卡
func (m *QuotaCountBean) AddOrUpdatePayName(add *input.SitePayNameAdd) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sitePayNameSchema := new(schema.SitePayName)
	sitePayNameSchema.PayName = add.PayName
	sitePayNameSchema.PayType = add.PayType
	sitePayNameSchema.State = add.State
	sitePayNameSchema.PayId = add.PayId
	sitePayNameSchema.PayKey = add.PayKey
	sitePayNameSchema.FUrl = add.FUrl
	sitePayNameSchema.Vircarddoin = add.Vircarddoin
	sitePayNameSchema.TerminalId = add.TerminalId
	sitePayNameSchema.Type = add.Type
	sitePayNameSchema.MyName = add.MyName
	sitePayNameSchema.Address = add.Address
	sitePayNameSchema.Lid = add.Lid
	if add.Id > 0 {
		return sess.ID(add.Id).Update(sitePayNameSchema)
	}
	count, err := sess.InsertOne(sitePayNameSchema)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}

//入款管理 - 查询第三方或者银行卡
func (m *QuotaCountBean) GetPayNameList(req *input.SitePayNameList) (list []*back.SitePayNameList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sitePayNameSchema := new(schema.SitePayName)
	if req.Type != 0 {
		sess.Where(sitePayNameSchema.TableName()+".type = ?", req.Type)
	}
	if req.State != 0 {
		sess.Where(sitePayNameSchema.TableName()+".state = ?", req.State)
	}
	sL := new(schema.SiteLevel)
	pT := new(schema.PaidType)
	err = sess.Table(sitePayNameSchema.TableName()).
		Join("LEFT", sL.TableName(), sitePayNameSchema.TableName()+".lid="+
			sL.TableName()+".id").
		Join("LEFT", pT.TableName(), sitePayNameSchema.TableName()+".pay_type="+
			pT.TableName()+".id").
		Find(&list)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return list, err
	}
	return list, err
}

//保留两位小数
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

//获取视讯平台下拉框
func (*QuotaCountBean) GetPlatform() ([]back.GetPlatform, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	p := new(schema.Platform)
	var platform []back.GetPlatform
	err := sess.Table(p.TableName()).Find(&platform)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return platform, err
	}
	return platform, err
}

//获取支付类型下拉框
func (*QuotaCountBean) ThirdTypeDrop() ([]back.ThirdTypeDropBack, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	pt := new(schema.PaidType)
	var data []back.ThirdTypeDropBack
	err := sess.Table(pt.TableName()).Where("type_status=?", 1).Find(&data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	return data, err
}

func (*QuotaCountBean) BankCardRechargeAdd(this *input.BankCardRecharge) (count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sitePayRecord := new(schema.SitePayRecord)

	sitePayRecord.SiteId = this.SiteId
	sitePayRecord.State = 1
	sitePayRecord.OrderNum = this.OrderNum
	sitePayRecord.AdminUser = this.AdminUser
	sitePayRecord.Money = float64(this.Money)
	sitePayRecord.Type = this.Type
	sitePayRecord.DoTime = global.GetCurrentTime()
	sitePayRecord.Remark = fmt.Sprintf("%s站点预存款%d，转入账号%s", this.SiteId, this.Money, this.PayCard)
	count, err = sess.Table(sitePayRecord.TableName()).Insert(sitePayRecord)
	return
}
