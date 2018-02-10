//存款
package wdeposit

import (
	"controllers"
	"global"
	"models/input"
	"models/schema"
	"net/http"

	"fmt"
	"github.com/labstack/echo"
)

type WapDeposit struct {
	controllers.BaseController
}

//线上存款 -微信
func (w *WapDeposit) WechatDeposit(ctx echo.Context) error {
	//验证存款金额字段
	deposit := new(input.WapDeposit)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//根据代理Id获取会员的代理信息
	agency, have, err := agencyBean.GetAgency(member.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//sales_site_pay_set查看会员充值是否有优惠
	paySet, have, err := sitePaySetBean.GetPaySetInfo(member.SiteId, member.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//查看该站点下是否开启线上入款优惠
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	if paySet.OnlineIsDeposit == 1 {
		onlineEntryRecord.IsDiscount = 2
	}
	//查看会员是否是首次存款
	have, err = onlineEntryRecordBean.IsFirst(member.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//是否首次存款
	if have {
		onlineEntryRecord.IsFirstDeposit = 1
	} else {
		onlineEntryRecord.IsFirstDeposit = 2
	}
	//赋值线上入款(默认未成功)
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1        //未支付
	onlineEntryRecord.SourceDeposit = 2 //wap端
	onlineEntryRecord.AmountDeposit = deposit.Money
	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}
	//请求第三方支付接口
	client := &http.Client{}
	//访问的支付路由
	url := "http://www.baidu.com"
	//提交请求
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(10219, ctx))
	}
	//更新会员出入款纪录
	memberCashCount := new(schema.MemberCashCount)
	memberCashCount.Member = member.Id
	memberCashCount.DepositNum = deposit.Money
	//更新会员余额
	member.Balance += deposit.Money
	//更新入款状态
	onlineEntryRecord.Status = 2 //已经支付
	//添加现金流水
	cash := new(schema.MemberCashRecord)
	cash.SiteId = member.SiteId
	cash.SiteIndexId = member.SiteIndexId
	cash.MemberId = member.Id
	cash.UserName = member.Account
	cash.AgencyId = agency.Id
	cash.AgencyAccount = agency.Account
	//类型:线上入款
	cash.SourceType = 2
	//存入
	cash.Type = 1
	cash.Balance = deposit.Money
	//客户端类型 wap
	cash.ClientType = 2
	//更新会员稽核纪录TODO
	aduit := new(schema.MemberAudit)
	//站点下是否开启线上入款优惠(2开启优惠)
	if paySet.LineIsDeposit == 2 {
		//查看站点设置是否是首次存款才有优惠
		if (paySet.OnlineIsDepositDiscount == 1 && !have) || paySet.OnlineIsDepositDiscount == 2 {
			//判断存款金额是否达到存款优惠标准
			if paySet.LineDiscountStandard <= deposit.Money {
				//获取会员优惠金额
				if deposit.Money*paySet.OnlineDiscountPercent < paySet.OnlineDepositMin {
					aduit.DepositMoney = paySet.OnlineDepositMin
				} else if deposit.Money*paySet.OnlineDiscountPercent > paySet.OnlineDepositMax {
					aduit.DepositMoney = paySet.OnlineDepositMax
				} else {
					aduit.DepositMoney = deposit.Money * paySet.OnlineDiscountPercent
				}
			}
		}
		//是否开启了常态稽核 1:开启
		if paySet.OnlineIsNormalAudit == 1 {
			//常态稽核
			aduit.NormalMoney = deposit.Money * paySet.OnlineNormalAuditPercent
		}
		//是否开启了综合性稽核 1:开启
		if paySet.OnlineIsMultipleAudit == 1 {
			//综合性稽核
			aduit.MultipleMoney = (deposit.Money + aduit.DepositMoney) * float64(paySet.OnlineMultipleAuditTimes)
		}
		aduit.MemberId = member.Id
		aduit.Account = member.Account
		aduit.DepositMoney = deposit.Money
		aduit.SiteId = member.SiteId
		aduit.SiteIndexId = member.SiteIndexId
	}
	//设置线上入款的优惠金额
	onlineDepositBean.OnlineRecord = new(schema.OnlineEntryRecord)
	onlineDepositBean.OnlineRecord.DepositDiscount = aduit.DepositMoney
	//更新会员余额
	member.Balance += aduit.DepositMoney
	onlineDepositBean.Audit = aduit
	onlineDepositBean.Cash = memberCashCount
	onlineDepositBean.Member = member
	onlineDepositBean.OnlineRecord = onlineEntryRecord
	onlineDepositBean.MemberRecord = cash
	err = onlineEntryRecordBean.Update(onlineDepositBean)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//线上存款 -网银
func (w *WapDeposit) BankDeposit(ctx echo.Context) error {
	//验证存款金额和入款银行字段
	deposit := new(input.WapDepositBank)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//判断银行Id是否存在
	have, err = bankCardBean.DepositById(deposit.Bank)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10221, ctx))
	}
	//根据代理Id获取会员的代理信息
	agency, have, err := agencyBean.GetAgency(member.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//sales_site_pay_set查看会员充值是否有优惠
	paySet, have, err := sitePaySetBean.GetPaySetInfo(member.SiteId, member.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//查看该站点下是否开启线上入款优惠
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	if paySet.OnlineIsDeposit == 1 {
		onlineEntryRecord.IsDiscount = 2
	}
	//查看会员是否是首次存款
	have, err = onlineEntryRecordBean.IsFirst(member.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//是否首次存款
	if have {
		onlineEntryRecord.IsFirstDeposit = 1
	} else {
		onlineEntryRecord.IsFirstDeposit = 2
	}
	//赋值线上入款(默认未成功)
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1        //未支付
	onlineEntryRecord.SourceDeposit = 2 //wap端
	onlineEntryRecord.AmountDeposit = deposit.Money
	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}
	//请求第三方支付接口
	client := &http.Client{}
	//访问的支付路由
	url := "http://www.baidu.com"
	//提交请求
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(10219, ctx))
	}
	//更新会员出入款纪录
	memberCashCount := new(schema.MemberCashCount)
	memberCashCount.Member = member.Id
	memberCashCount.DepositNum = deposit.Money
	//更新会员余额
	member.Balance += deposit.Money
	//更新入款状态
	onlineEntryRecord.Status = 2 //已经支付
	//添加现金流水
	cash := new(schema.MemberCashRecord)
	cash.SiteId = member.SiteId
	cash.SiteIndexId = member.SiteIndexId
	cash.MemberId = member.Id
	cash.UserName = member.Account
	cash.AgencyId = agency.Id
	cash.AgencyAccount = agency.Account
	//类型:线上入款
	cash.SourceType = 2
	//存入
	cash.Type = 1
	cash.Balance = deposit.Money
	//客户端类型 wap
	cash.ClientType = 2
	//更新会员稽核纪录TODO
	aduit := new(schema.MemberAudit)
	//站点下是否开启线上入款优惠(2开启优惠)
	if paySet.LineIsDeposit == 2 {
		//查看站点设置是否是首次存款才有优惠
		if (paySet.OnlineIsDepositDiscount == 1 && !have) || paySet.OnlineIsDepositDiscount == 2 {
			//判断存款金额是否达到存款优惠标准
			if paySet.LineDiscountStandard <= deposit.Money {
				//获取会员优惠金额
				if deposit.Money*paySet.OnlineDiscountPercent < paySet.OnlineDepositMin {
					aduit.DepositMoney = paySet.OnlineDepositMin
				} else if deposit.Money*paySet.OnlineDiscountPercent > paySet.OnlineDepositMax {
					aduit.DepositMoney = paySet.OnlineDepositMax
				} else {
					aduit.DepositMoney = deposit.Money * paySet.OnlineDiscountPercent
				}
			}
		}
		//是否开启了常态稽核 1:开启
		if paySet.OnlineIsNormalAudit == 1 {
			//常态稽核
			aduit.NormalMoney = deposit.Money * paySet.OnlineNormalAuditPercent
		}
		//是否开启了综合性稽核 1:开启
		if paySet.OnlineIsMultipleAudit == 1 {
			//综合性稽核
			aduit.MultipleMoney = (deposit.Money + aduit.DepositMoney) * float64(paySet.OnlineMultipleAuditTimes)
		}
		aduit.MemberId = member.Id
		aduit.Account = member.Account
		aduit.DepositMoney = deposit.Money
		aduit.SiteId = member.SiteId
		aduit.SiteIndexId = member.SiteIndexId
	}
	Onli := new(schema.OnlineEntryRecord)
	//设置线上入款的优惠金额
	Onli.DepositDiscount = aduit.DepositMoney
	onlineDepositBean.OnlineRecord = Onli
	//更新会员余额
	member.Balance += aduit.DepositMoney
	onlineDepositBean.Audit = aduit
	onlineDepositBean.Cash = memberCashCount
	onlineDepositBean.Member = member
	onlineDepositBean.OnlineRecord = onlineEntryRecord
	onlineDepositBean.MemberRecord = cash
	err = onlineEntryRecordBean.Update(onlineDepositBean)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//线上存款 -点卡
func (w *WapDeposit) CardDeposit(ctx echo.Context) error {
	//验证存款金额,点卡金额,入款银行字段
	deposit := new(input.WapDepositCard)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断会员Id是否存在
	fmt.Println(memberRedis.Id)
	member, have, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//判断银行类型Id是否存在 todo:点卡类型暂不做判断
	//have, err = bankCardBean.DepositById(deposit.Bank)
	//if err != nil {
	//	global.GlobalLogger.Error("err:", err)
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//if !have {
	//	return ctx.JSON(200, global.ReplyError(10221, ctx))
	//}
	//根据代理Id获取会员的代理信息
	agency, have, err := agencyBean.GetAgency(member.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//sales_site_pay_set查看会员充值是否有优惠
	paySet, have, err := sitePaySetBean.GetPaySetInfo(member.SiteId, member.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		fmt.Println("未找到对应站点的支付设定")
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//查看该站点下是否开启线上入款优惠
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	if paySet.OnlineIsDeposit == 1 {
		onlineEntryRecord.IsDiscount = 2
	}
	//查看会员是否是首次存款
	have, err = onlineEntryRecordBean.IsFirst(member.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//是否首次存款
	if have {
		onlineEntryRecord.IsFirstDeposit = 1
	} else {
		onlineEntryRecord.IsFirstDeposit = 2
	}
	//赋值线上入款(默认未成功)
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1        //未支付
	onlineEntryRecord.SourceDeposit = 2 //wap端
	onlineEntryRecord.AmountDeposit = deposit.Money
	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}
	//请求第三方支付接口
	client := &http.Client{}
	//访问的支付路由
	url := "http://www.baidu.com"
	//提交请求
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(10219, ctx))
	}
	//更新会员出入款纪录
	memberCashCount := new(schema.MemberCashCount)
	memberCashCount.Member = member.Id
	memberCashCount.DepositNum = deposit.Money
	//更新会员余额
	member.Balance += deposit.Money
	//更新入款状态
	onlineEntryRecord.Status = 2 //已经支付
	//添加现金流水
	cash := new(schema.MemberCashRecord)
	cash.SiteId = member.SiteId
	cash.SiteIndexId = member.SiteIndexId
	cash.MemberId = member.Id
	cash.UserName = member.Account
	cash.AgencyId = agency.Id
	cash.AgencyAccount = agency.Account
	//类型:线上入款
	cash.SourceType = 2
	//存入
	cash.Type = 1
	cash.Balance = deposit.Money
	//客户端类型 wap
	cash.ClientType = 2
	//更新会员稽核纪录
	aduit := new(schema.MemberAudit)
	//站点下是否开启线上入款优惠(2开启优惠)
	if paySet.LineIsDeposit == 2 {
		//查看站点设置是否是首次存款才有优惠
		if (paySet.OnlineIsDepositDiscount == 1 && !have) || paySet.OnlineIsDepositDiscount == 2 {
			//判断存款金额是否达到存款优惠标准
			if paySet.LineDiscountStandard <= deposit.Money {
				//获取会员优惠金额
				if deposit.Money*paySet.OnlineDiscountPercent < paySet.OnlineDepositMin {
					aduit.DepositMoney = paySet.OnlineDepositMin
				} else if deposit.Money*paySet.OnlineDiscountPercent > paySet.OnlineDepositMax {
					aduit.DepositMoney = paySet.OnlineDepositMax
				} else {
					aduit.DepositMoney = deposit.Money * paySet.OnlineDiscountPercent
				}
			}
		}
		//是否开启了常态稽核 1:开启
		if paySet.OnlineIsNormalAudit == 1 {
			//常态稽核
			aduit.NormalMoney = deposit.Money * paySet.OnlineNormalAuditPercent
		}
		//是否开启了综合性稽核 1:开启
		if paySet.OnlineIsMultipleAudit == 1 {
			//综合性稽核
			aduit.MultipleMoney = (deposit.Money + aduit.DepositMoney) * float64(paySet.OnlineMultipleAuditTimes)
		}
		aduit.MemberId = member.Id
		aduit.Account = member.Account
		aduit.DepositMoney = deposit.Money
		aduit.SiteId = member.SiteId
		aduit.SiteIndexId = member.SiteIndexId
	}
	//设置线上入款的优惠金额
	//var onlineDepositBean.OnlineRecord *schema.OnlieEntryRecord
	onlineDepositBean.OnlineRecord = new(schema.OnlineEntryRecord)
	onlineDepositBean.OnlineRecord.DepositDiscount = aduit.DepositMoney
	//更新会员余额
	member.Balance += aduit.DepositMoney
	onlineDepositBean.Audit = aduit
	onlineDepositBean.Cash = memberCashCount
	onlineDepositBean.Member = member
	onlineDepositBean.OnlineRecord = onlineEntryRecord
	onlineDepositBean.MemberRecord = cash
	err = onlineEntryRecordBean.Update(onlineDepositBean)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//公司存款 -微信
func (w *WapDeposit) WechatCompanyDeposit(ctx echo.Context) error {
	//验证存款金额,转账账号,转账时间
	deposit := new(input.WapDepositCardIn)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//根据代理Id获取会员的代理信息
	agency, have, err := agencyBean.GetAgency(member.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//根据存款优惠类型判断站点下是否开启线上入款优惠
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	//赋值线上入款(默认未成功)
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1        //未支付
	onlineEntryRecord.SourceDeposit = 2 //wap端
	onlineEntryRecord.AmountDeposit = deposit.Money
	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}
	//请求第三方支付接口
	client := &http.Client{}
	//访问的支付路由
	url := "http://www.baidu.com"
	//提交请求
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(10219, ctx))
	}
	//根据存款优惠类型做稽核TODO

	//更新会员稽核纪录
	aduit := new(schema.MemberAudit)
	//更新会员出入款纪录
	memberCashCount := new(schema.MemberCashCount)
	memberCashCount.Member = member.Id
	memberCashCount.DepositNum = deposit.Money
	//更新会员余额
	member.Balance += deposit.Money
	//更新入款状态
	onlineEntryRecord.Status = 2 //已经支付
	//添加现金流水
	cash := new(schema.MemberCashRecord)
	cash.SiteId = member.SiteId
	cash.SiteIndexId = member.SiteIndexId
	cash.MemberId = member.Id
	cash.UserName = member.Account
	cash.AgencyId = agency.Id
	cash.AgencyAccount = agency.Account
	//类型:线上入款
	cash.SourceType = 2
	//存入
	cash.Type = 1
	cash.Balance = deposit.Money
	//客户端类型 wap
	cash.ClientType = 2
	//设置线上入款的优惠金额
	onlineDepositBean.OnlineRecord.DepositDiscount = aduit.DepositMoney
	//更新会员余额
	member.Balance += aduit.DepositMoney
	onlineDepositBean.Audit = aduit
	onlineDepositBean.Cash = memberCashCount
	onlineDepositBean.Member = member
	onlineDepositBean.OnlineRecord = onlineEntryRecord
	onlineDepositBean.MemberRecord = cash
	err = onlineEntryRecordBean.Update(onlineDepositBean)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//公司存款 -网银
func (w *WapDeposit) BankCompanyDeposit(ctx echo.Context) error {
	//验证存款金额,存款人姓名,转账方式,转账账号,转账时间,卡类型
	deposit := new(input.WapCompanyDepositBank)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//判断银行Id是否存在
	have, err = bankCardBean.DepositById(deposit.Bank)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10221, ctx))
	}
	//根据代理Id获取会员的代理信息
	agency, have, err := agencyBean.GetAgency(member.ThirdAgencyId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//根据存款优惠类型判断站点下是否开启线上入款优惠
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	//赋值线上入款(默认未成功)
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1        //未支付
	onlineEntryRecord.SourceDeposit = 2 //wap端
	onlineEntryRecord.AmountDeposit = deposit.Money
	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}
	//请求第三方支付接口
	client := &http.Client{}
	//访问的支付路由
	url := "http://www.baidu.com"
	//提交请求
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(10219, ctx))
	}
	//根据存款优惠类型做稽核TODO
	//更新会员稽核纪录
	aduit := new(schema.MemberAudit)
	//更新会员出入款纪录
	memberCashCount := new(schema.MemberCashCount)
	memberCashCount.Member = member.Id
	memberCashCount.DepositNum = deposit.Money
	//更新会员余额
	member.Balance += deposit.Money
	//更新入款状态
	onlineEntryRecord.Status = 2 //已经支付
	//添加现金流水
	cash := new(schema.MemberCashRecord)
	cash.SiteId = member.SiteId
	cash.SiteIndexId = member.SiteIndexId
	cash.MemberId = member.Id
	cash.UserName = member.Account
	cash.AgencyId = agency.Id
	cash.AgencyAccount = agency.Account
	//类型:线上入款
	cash.SourceType = 2
	//存入
	cash.Type = 1
	cash.Balance = deposit.Money
	//客户端类型 wap
	cash.ClientType = 2
	//设置线上入款的优惠金额
	onlineDepositBean.OnlineRecord.DepositDiscount = aduit.DepositMoney
	//更新会员余额
	member.Balance += aduit.DepositMoney
	onlineDepositBean.Audit = aduit
	onlineDepositBean.Cash = memberCashCount
	onlineDepositBean.Member = member
	onlineDepositBean.OnlineRecord = onlineEntryRecord
	onlineDepositBean.MemberRecord = cash
	err = onlineEntryRecordBean.Update(onlineDepositBean)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)

}
