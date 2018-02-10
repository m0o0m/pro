package member

import (
	"controllers"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"global"
	"io/ioutil"
	"models/back"
	"models/input"
	"models/schema"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type OtherController struct {
	controllers.BaseController
}

//根据商品类型获取商品名（头部导航栏）
func (*OtherController) GetProductName(ctx echo.Context) error {
	var typeId []int64
	ids := ctx.FormValue("type_id")
	s := strings.Split(ids, ",")
	for k := range s {
		id, _ := strconv.Atoi(s[k])
		typeId = append(typeId, int64(id))
	}
	//获取站点信息
	site, flag, err := GetSiteInfo(ctx)
	if err != nil {
		global.GlobalLogger.Error("error%s", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	//根据站点查询哪些商品被剔除
	productIds, err := otherBean.GetProductDel(site.SiteId, site.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data, err := otherBean.GetProductByType(typeId, productIds)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//根据id获取站点支付设定
func (*OtherController) GetPaySetInfoById(ctx echo.Context) error {
	setId := new(input.GetPaySetInfoById)
	code := global.ValidRequestMember(setId, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(setId.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(paySet))
}

//获取层级支付设定
func (*OtherController) PaySet(ctx echo.Context) error {
	payset := new(input.MemberLevelPaySet)
	code := global.ValidRequest(payset, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := memberLevelBean.MemberLevelPatSetOne(payset)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//优惠活动申请提交
func (*OtherController) SelfHelpApply(ctx echo.Context) error {
	apply := new(input.SelfHelpApllyAdd)
	code := global.ValidRequest(apply, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//验证验证码
	codes := ctx.Request().Header.Get("code")
	key, err := getMemberRedis(codes)
	if err == redis.Nil {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if key == "" || strings.ToLower(key) != strings.ToLower(apply.Code) {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	//删除验证码
	err = global.GetRedis().Del(codes).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//根据站点账号取出对应的会员信息
	info, flag, err := memberBean.GetInfoBySite(memberRedis.Site, memberRedis.SiteIndex, apply.Account)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//账号不存在,已经被删除在获取的时候已经被排除出去，或者是该站点下面不存在该账号
	if !flag {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//账号被禁用
	if info.Status == 2 {
		return ctx.JSON(200, global.ReplyError(20002, ctx))
	}

	count, err := selfHelpApplyforBean.AutoDiscountAdd(apply, memberRedis)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.NoContent(204)
}

//公司入款数据
func (*OtherController) GetCompanyData(ctx echo.Context) error {
	companyData := new(back.GetCompanyData)
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
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	companyData.PaySet.OnlineDepositMin = paySet.OnlineDepositMin
	companyData.PaySet.OnlineDepositMax = paySet.OnlineDepositMax
	companyData.PaySet.LineDepositMax = paySet.LineDepositMax
	companyData.PaySet.LineDepositMin = paySet.LineDepositMin
	//支付类型
	infoList, err := paidTypeBean.GetSitePaidTypeData(memberRedis.Site, memberRedis.SiteIndex, memberRedis.LevelId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	companyData.PaidType = make(map[string]int)
	for k, _ := range infoList {
		if infoList[k].TypeStatus == 1 {
			companyData.PaidType[infoList[k].PaidTypeName] = infoList[k].Id
		}
	}
	//收款人信息
	bankInfo, err := memberCompanyIncomeBean.GetSetAccountInfo(&input.SiteId{member.SiteId, member.SiteIndexId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	companyData.OnlineIncomeSet = bankInfo
	//入款银行
	siteBank, err := bankCardBean.GetAllIncomeBank(member.SiteId, member.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	companyData.SiteIncomeBank = siteBank

	//m.Datetime = time.Now().String()[:19]
	//m.UrlLink = siteInfo.UrlLink
	//m.SiteName = siteName
	//m.OrderNum = strings.Replace(time.Now().String()[:10], "-", "", 2) + siteId + siteIndexId + strconv.FormatInt(time.Now().Unix(), 10)
	//m.TradeDate = time.Now().String()[:10]
	return ctx.JSON(200, global.ReplyItem(companyData))
}

//快速充值中心存款数据
func (*OtherController) GetFastIncomeData(ctx echo.Context) error {
	accounts := new(input.GetFastIncomeData)
	code := global.ValidRequest(accounts, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if accounts.Account != accounts.Reaccount {
		return ctx.JSON(200, global.ReplyError(10164, ctx))
	}
	//根据域名查询出站点id
	host := ctx.Request().Host
	siteDomin := strings.Split(host, ":")[0]
	if len(siteDomin) == 0 {
		global.GlobalLogger.Error("Login get host is nil!")
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	siteInfo, flag, err := siteDomainBean.GetSiteInfoByDomian(siteDomin)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//未查询到域名
	if !flag {
		global.GlobalLogger.Error("Login GetSiteInfoByDomian error: %s", siteDomin)
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}

	incomeData := new(back.GetFastIncomeData)

	/*if ctx.Get("member") != nil { //已登录
		memberRedis := ctx.Get("member").(*global.MemberRedisToken)
		incomeData.AccountSelf = memberRedis.Account
	} else {
		incomeData.AccountSelf = ""
	}*/
	//memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	/*//模拟数据
	memberRedis:=new(global.MemberRedisToken)
	memberRedis.Id=5
	memberRedis.SiteIndex="a"
	memberRedis.Site="aaa"
	memberRedis.LevelId="A007"
	memberRedis.Account="sitong5"
	memberRedis.Type="pc"
	memberRedis.ExpirTime=0
	memberRedis.Status=1*/
	//判断会员Id是否存在
	incomeData.Account = accounts.Account
	member, have, err := memberBean.GetInfoBySite(siteInfo.SiteId, siteInfo.SiteIndexId, accounts.Account)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//入款上下限
	incomeData.PaySet.OnlineDepositMin = paySet.OnlineDepositMin
	incomeData.PaySet.OnlineDepositMax = paySet.OnlineDepositMax
	//支付类型（线上+公司）
	//线上支付类型+支付设定
	infoList, err := paidTypeBean.GetFastIncomeData(siteInfo.SiteId, siteInfo.SiteIndexId, member.LevelId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(infoList) == 0 {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	fmt.Println(infoList)
	info := new(back.FastIncomeData)
	incomeData.Income = make(map[int]*back.FastIncomeData)
	var checkType = 0
	for k, v := range infoList {
		//if v.PaidType != checkType {
		if checkType != v.PaidType && checkType != 0 {
			if checkType == 1 || checkType == 10 {
			} else {
				incomeData.Income[checkType] = info
			}
			info = new(back.FastIncomeData)
		}
		checkType = v.PaidType
		info.PaidType = checkType
		OnlineIncomeOne := new(back.OnlineIncomeData)
		OnlineIncomeOne.Id = v.SetId
		OnlineIncomeOne.Sort = v.Sort
		info.OnlineIncome = append(info.OnlineIncome, *OnlineIncomeOne)
		//}
		if len(infoList)-1 == k {
			if checkType == 1 || checkType == 10 {
			} else {
				incomeData.Income[checkType] = info
			}
		}
	}

	//封装 第三方与公司入款 数据
	//incomeData.Income, err = paidTypeBean.GetFastIncomeData(siteInfo.SiteId, siteInfo.SiteIndexId, member.LevelId)

	//m.Datetime = time.Now().String()[:19]
	//m.UrlLink = siteInfo.UrlLink
	//m.SiteName = siteName
	//m.OrderNum = strings.Replace(time.Now().String()[:10], "-", "", 2) + siteId + siteIndexId + strconv.FormatInt(time.Now().Unix(), 10)
	//m.TradeDate = time.Now().String()[:10]
	return ctx.JSON(200, global.ReplyItem(incomeData))
}

//存款数据
func (*OtherController) GetIncomeData(ctx echo.Context) error {
	incomeData := new(back.GetIncomeData)
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	/*//模拟数据
	memberRedis := new(global.MemberRedisToken)
	memberRedis.Id = 5
	memberRedis.SiteIndex = "a"
	memberRedis.Site = "aaa"
	memberRedis.LevelId = "A007"
	memberRedis.Account = "sitong5"
	memberRedis.Type = "pc"
	memberRedis.ExpirTime = 0
	memberRedis.Status = 1*/
	//判断会员Id是否存在
	incomeData.Account = memberRedis.Account
	member, have, err := memberBean.GetInfoById(memberRedis.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//入款上下限
	incomeData.PaySet.OnlineDepositMin = paySet.OnlineDepositMin
	incomeData.PaySet.OnlineDepositMax = paySet.OnlineDepositMax
	incomeData.PaySet.LineDepositMax = paySet.LineDepositMax
	incomeData.PaySet.LineDepositMin = paySet.LineDepositMin
	//支付类型（线上+公司）
	//线上支付类型+支付设定
	infoList, err := paidTypeBean.GetIncomeData(memberRedis.Site, memberRedis.SiteIndex, memberRedis.LevelId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//公司支付类型+收款银行设定
	bankInfo, err := memberCompanyIncomeBean.GetSetAccountInfo(&input.SiteId{memberRedis.Site, memberRedis.SiteIndex})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//封装 第三方与公司入款 数据
	incomeData.Income, err = paidTypeBean.MergeCompanyAndOnlineData(infoList, bankInfo)

	//入款银行
	siteBank, err := bankCardBean.GetAllIncomeBank(member.SiteId, member.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	incomeData.SiteIncomeBank = siteBank

	//m.Datetime = time.Now().String()[:19]
	//m.UrlLink = siteInfo.UrlLink
	//m.SiteName = siteName
	//m.OrderNum = strings.Replace(time.Now().String()[:10], "-", "", 2) + siteId + siteIndexId + strconv.FormatInt(time.Now().Unix(), 10)
	//m.TradeDate = time.Now().String()[:10]
	return ctx.JSON(200, global.ReplyItem(incomeData))
}

//添加一条公司入款记录
func (c *OtherController) AddCompanyIncome(ctx echo.Context) error {
	deposit := new(input.AddCompanyIncome)
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
		return ctx.JSON(200, global.ReplyError(30208, ctx))
	}
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//是否超限
	if paySet.LineDepositMax < deposit.DepositMoney {
		//超过上限了
		return ctx.JSON(200, global.ReplyError(10158, ctx))
	}
	if paySet.LineDepositMin > deposit.DepositMoney {
		//低于下限了
		return ctx.JSON(200, global.ReplyError(10159, ctx))
	}
	//查看该站点下是否开启线上入款优惠,1首次，2每次
	companyIncome := new(schema.MemberCompanyIncome)
	//查看会员是否是首次存款
	have, err = onlineEntryRecordBean.IsFirstIncome(member.Id, member.SiteId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//是否首次存款
	if have {
		companyIncome.IsFirstDeposit = 1
	} else {
		companyIncome.IsFirstDeposit = 2
	}
	//判断优惠是首次还是每次
	if paySet.OnlineIsDepositDiscount == 1 && companyIncome.IsFirstDeposit != 1 { //首次有优惠 且 会员不是首次充值了，则无优惠
	} else { //其他情况，则有优惠
		if deposit.DepositMoney >= paySet.OnlineDiscountStandard { //大于优惠标准
			companyIncome.DepositDiscount = deposit.DepositMoney * paySet.OnlineDiscountPercent * 0.01 //优惠
			if companyIncome.DepositDiscount > paySet.OnlineDiscountUp {                               //大于优惠上限
				companyIncome.DepositDiscount = paySet.OnlineDiscountUp
			}
			if deposit.DepositMoney >= paySet.OnlineOtherDiscountStandard {
				companyIncome.OtherDiscount = deposit.DepositMoney * paySet.OnlineOtherDiscountPercent * 0.01 //其他优惠
				if companyIncome.OtherDiscount > paySet.OnlineOtherDiscountUp {
					companyIncome.OtherDiscount = paySet.OnlineOtherDiscountUp
				}
			}
		}
	}

	var siteStr string
	for _, char := range []rune(member.SiteId) {
		siteStr += strconv.Itoa(int(char - 30))
	}
	companyIncome.OrderNum = strconv.FormatInt(time.Now().UnixNano(), 10) + siteStr + strconv.FormatInt(member.Id, 10) + ""
	//companyIncome.OrderNum = uuid.NewV4().String()                                                                  //订单号
	companyIncome.DepositCount = deposit.DepositMoney + companyIncome.DepositDiscount + companyIncome.OtherDiscount //总额
	companyIncome.ClientType = 1                                                                                    //pc端
	companyIncome.CreateTime = 1                                                                                    //本平台跳到第三方时的时间
	companyIncome.DepositMoney = deposit.DepositMoney                                                               //存入金额
	companyIncome.MemberId = member.Id
	companyIncome.LevelId = member.LevelId
	companyIncome.Account = member.Account
	companyIncome.SiteId = member.SiteId
	companyIncome.SiteIndexId = member.SiteIndexId
	companyIncome.AgencyId = agency.Id
	companyIncome.AgencyAccount = agency.Account
	companyIncome.Status = 4 //未支付
	companyIncome.DepositUsername = deposit.DepositUsername
	companyIncome.DepositMethod = deposit.DepositMethod
	companyIncome.SetId = deposit.SetId
	companyIncome.BankId = deposit.BankId
	companyIncome.BankName = deposit.BankName
	companyIncome.Remark = deposit.Remark

	count, err := memberCompanyIncomeBean.AddCompanyIncome(companyIncome)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10058, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(back.AddBankIn{OrderNumber: companyIncome.OrderNum})) //返回订单号
}

//模拟一条公司入款记录成功
func (c *OtherController) CompanyIncomeResult(ctx echo.Context) error {
	deposit := new(input.CompanyIncomeResult)
	code := global.ValidRequest(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	err := global.DepositCache.Set(deposit.OrderNum, deposit)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//添加线上入款记录(网银、点卡除外)
func (c *OtherController) AddOnlineIncome(ctx echo.Context) error {
	/*
		通过会员信息，封装入款记录参数
		查询所有符合条件的第三方列表
		随机调用一个第三方接口
			添加线上入款记录
			把传入信息提交到该第三方接口
				领取二维码信息或界面
				会员扫码提交信息
			判断第三方返回信息
				如果返回成功
					开启事务
						更新线上入款记录
						更新余额
						更新现金流水记录
						更新稽核
						更新入款统计表
				如果返回失败
					记录这个失败的第三方，下次不再选取
					重新随机调用一个调用第三方接口
	*/

	deposit := new(input.WapOnlineDeposit)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断是否为快速充值
	var nowAccount string
	if deposit.IsFast == 1 { //是快速充值
		nowAccount = deposit.Account
	} else {
		nowAccount = memberRedis.Account
	}
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoBySite(memberRedis.Site, memberRedis.SiteIndex, nowAccount)
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
		return ctx.JSON(200, global.ReplyError(30208, ctx))
	}
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//是否超限
	if paySet.OnlineDepositMax < deposit.Money { //超过上限了
		return ctx.JSON(200, global.ReplyError(10160, ctx))
	}
	if paySet.OnlineDepositMin > deposit.Money { //低于下限了
		return ctx.JSON(200, global.ReplyError(10161, ctx))
	}
	//查看该站点下是否开启线上入款优惠,1首次，2每次
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	//查看会员是否是首次存款
	have, err = onlineEntryRecordBean.IsFirstIncome(member.Id, member.SiteId)
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
	//是否放弃优惠
	if paySet.OnlineIsDeposit == 1 { //放弃优惠
		onlineEntryRecord.IsDiscount = 2 //无优惠
	} else { //不放弃优惠
		//判断优惠是首次还是每次
		if paySet.OnlineIsDepositDiscount == 1 && onlineEntryRecord.IsFirstDeposit != 1 { //首次有优惠 且 会员不是首次充值了，则无优惠
			onlineEntryRecord.IsDiscount = 2 //无优惠
		} else { //其他情况，则有优惠
			if deposit.Money >= paySet.OnlineDiscountStandard { //大于优惠标准
				onlineEntryRecord.DepositDiscount = deposit.Money * paySet.OnlineDiscountPercent * 0.01 //优惠
				if onlineEntryRecord.DepositDiscount > paySet.OnlineDiscountUp {                        //大于优惠上限
					onlineEntryRecord.DepositDiscount = paySet.OnlineDiscountUp
				}
				if deposit.Money >= paySet.OnlineOtherDiscountStandard {
					onlineEntryRecord.OtherDepositDiscount = deposit.Money * paySet.OnlineOtherDiscountPercent * 0.01 //其他优惠
					if onlineEntryRecord.OtherDepositDiscount > paySet.OnlineOtherDiscountUp {
						onlineEntryRecord.OtherDepositDiscount = paySet.OnlineOtherDiscountUp
					}
				}
			}
			onlineEntryRecord.IsDiscount = 1 //有优惠
		}
	}
	//拼接订单号，填充数据
	var siteStr string
	for _, char := range []rune(member.SiteId) {
		siteStr += strconv.Itoa(int(char - 30))
	}
	onlineEntryRecord.ThirdOrderNumber = strconv.FormatInt(time.Now().UnixNano(), 10) + siteStr + strconv.FormatInt(member.Id, 10) + ""
	onlineEntryRecord.SourceDeposit = 1                    //pc端
	onlineEntryRecord.LocalOperateTime = time.Now().Unix() //本平台跳到第三方时的时间
	onlineEntryRecord.AmountDeposit = deposit.Money        //存入金额
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1                  //未支付 赋值线上入款(默认未成功)
	onlineEntryRecord.PaidType = deposit.PaidType //支付类型id

	/*//根据条件，随机选择一条第三方支付平台
	oneThird, oneThirdNum, err := onlinePaidSetupBean.GetOnePaid(&input.GetOnePaid{member.SiteId, member.SiteIndexId, deposit.PaidType, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if oneThirdNum <= 0 {
		return ctx.JSON(200, global.ReplyError(10157, ctx))
	}*/
	//根据传入的id,选择对应的第三方
	oneThird, oneThirdHas, err := onlinePaidSetupBean.GetOnePaidById(
		&input.GetOnePaidById{
			member.SiteId,
			member.SiteIndexId,
			deposit.PaidType,
			member.LevelId,
			deposit.Id})
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !oneThirdHas {
		return ctx.JSON(200, global.ReplyError(10157, ctx))
	}
	onlineEntryRecord.ThirdId = oneThird.PaidPlatform //第三方支付平台id
	onlineEntryRecord.PaidSetupId = int(oneThird.Id)  //线上支付设定id

	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}

	/*fmt.Println(onlineEntryRecord.ThirdOrderNumber)
	fmt.Println(deposit.PaidType)
	fmt.Println(oneThird.MerchatId)
	fmt.Println(oneThird.PaidPlatform)
	fmt.Println(oneThird.PaidCode)*/
	//var newHtml string
	//拼接支付rediskey,用于模板渲染时获取数据
	payRedisKey := "payRedisKey_" + member.SiteId + member.Account
	//获取对应站点的三方对接验证加密记录
	apiClients, has, err := apiClientsBean.GetOneApiClients("zzz") //member.SiteId todo 真实上线时，要对应上站点
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10162, ctx))
	}
	//请求第三方支付接口
	err = payBean.ThirdPayTest(
		&input.ThirdPayData{
			deposit.Money,
			onlineEntryRecord.ThirdOrderNumber,
			deposit.PaidType,
			oneThird.MerchatId,
			oneThird.Id,
			payRedisKey,
			oneThird.PaidCode,
			apiClients.UserId,
			apiClients.Name,
			apiClients.Secret,
			apiClients.SiteId,
			member.SiteIndexId,
			0,
			"",
			""})
	/*fmt.Println(&input.ThirdPayData{
	deposit.Money,
	onlineEntryRecord.ThirdOrderNumber,
	deposit.PaidType,
	oneThird.MerchatId,
	oneThird.PaidPlatform,
	payRedisKey,
	oneThird.PaidCode,
	apiClients.UserId,
	apiClients.Name,
	apiClients.Secret,
	apiClients.SiteId,
	member.SiteIndexId})*/
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//添加线上入款记录(网银）
func (c *OtherController) AddOnlineBankIncome(ctx echo.Context) error {
	/*
		通过会员信息，封装入款记录参数
		查询所有符合条件的第三方列表
		随机调用一个第三方接口
			添加线上入款记录
			把传入信息提交到该第三方接口
				领取二维码信息或界面
				会员扫码提交信息
			判断第三方返回信息
				如果返回成功
					开启事务
						更新线上入款记录
						更新余额
						更新现金流水记录
						更新稽核
						更新入款统计表
				如果返回失败
					记录这个失败的第三方，下次不再选取
					重新随机调用一个调用第三方接口
	*/
	deposit := new(input.OnlineBankDeposit)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断是否为快速充值
	var nowAccount string
	if deposit.IsFast == 1 { //是快速充值
		nowAccount = deposit.Account
	} else {
		nowAccount = memberRedis.Account
	}
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoBySite(memberRedis.Site, memberRedis.SiteIndex, nowAccount)
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
		return ctx.JSON(200, global.ReplyError(30208, ctx))
	}
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//判断是否超限
	if paySet.OnlineDepositMax < deposit.Money { //超过上限了
		return ctx.JSON(200, global.ReplyError(10160, ctx))
	}
	if paySet.OnlineDepositMin > deposit.Money { //低于下限了
		return ctx.JSON(200, global.ReplyError(10161, ctx))
	}
	//查看该站点下是否开启线上入款优惠,1首次，2每次
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	//查看会员是否是首次存款
	have, err = onlineEntryRecordBean.IsFirstIncome(member.Id, member.SiteId)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//判断是否首次存款
	if have {
		onlineEntryRecord.IsFirstDeposit = 1
	} else {
		onlineEntryRecord.IsFirstDeposit = 2
	}
	//判断是否放弃优惠
	if paySet.OnlineIsDeposit == 1 { //放弃优惠
		onlineEntryRecord.IsDiscount = 2 //无优惠
	} else { //不放弃优惠
		//判断优惠是首次还是每次
		if paySet.OnlineIsDepositDiscount == 1 && onlineEntryRecord.IsFirstDeposit != 1 { //首次有优惠 且 会员不是首次充值了，则无优惠
			onlineEntryRecord.IsDiscount = 2 //无优惠
		} else { //其他情况，则有优惠
			if deposit.Money >= paySet.OnlineDiscountStandard { //大于优惠标准
				onlineEntryRecord.DepositDiscount = deposit.Money * paySet.OnlineDiscountPercent * 0.01 //优惠
				if onlineEntryRecord.DepositDiscount > paySet.OnlineDiscountUp {                        //大于优惠上限
					onlineEntryRecord.DepositDiscount = paySet.OnlineDiscountUp
				}
				if deposit.Money >= paySet.OnlineOtherDiscountStandard {
					onlineEntryRecord.OtherDepositDiscount = deposit.Money * paySet.OnlineOtherDiscountPercent * 0.01 //其他优惠
					if onlineEntryRecord.OtherDepositDiscount > paySet.OnlineOtherDiscountUp {
						onlineEntryRecord.OtherDepositDiscount = paySet.OnlineOtherDiscountUp
					}
				}
			}
			onlineEntryRecord.IsDiscount = 1 //有优惠
		}
	}
	//拼接订单号，数据填充
	var siteStr string
	for _, char := range []rune(member.SiteId) {
		siteStr += strconv.Itoa(int(char - 30))
	}
	onlineEntryRecord.ThirdOrderNumber = strconv.FormatInt(time.Now().UnixNano(), 10) + siteStr + strconv.FormatInt(member.Id, 10) + ""
	onlineEntryRecord.SourceDeposit = 1                    //pc端
	onlineEntryRecord.LocalOperateTime = time.Now().Unix() //本平台跳到第三方时的时间
	onlineEntryRecord.AmountDeposit = deposit.Money        //存入金额
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1                  //未支付 赋值线上入款(默认未成功)
	onlineEntryRecord.PaidType = deposit.PaidType //支付类型id

	//根据传入的id,选择对应的第三方
	oneThird, oneThirdHas, err := onlinePaidSetupBean.GetOnePaidById(
		&input.GetOnePaidById{
			member.SiteId,
			member.SiteIndexId,
			deposit.PaidType,
			member.LevelId,
			deposit.Id})
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !oneThirdHas {
		return ctx.JSON(200, global.ReplyError(10157, ctx))
	}
	onlineEntryRecord.ThirdId = oneThird.PaidPlatform //第三方支付平台id
	onlineEntryRecord.PaidSetupId = int(oneThird.Id)  //线上支付设定id

	//判断是否为网银在线 type=1 bank
	if deposit.PaidType == 1 {
		oneThird.PaidCode = deposit.Bank
	} else { //此接口只用于网银，否则报错
		return ctx.JSON(200, global.ReplyError(60034, ctx))
	}

	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}

	/*fmt.Println(onlineEntryRecord.ThirdOrderNumber)
	fmt.Println(deposit.PaidType)
	fmt.Println(oneThird.MerchatId)
	fmt.Println(oneThird.PaidPlatform)
	fmt.Println(oneThird.PaidCode)*/
	//var newHtml string
	//拼接支付rediskey,用于模板渲染时获取数据
	payRedisKey := "payRedisKey_" + member.SiteId + member.Account
	//获取对应站点的三方对接验证加密记录
	apiClients, has, err := apiClientsBean.GetOneApiClients("zzz") //member.SiteId todo 真实上线时，要对应上站点
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10162, ctx))
	}
	//请求第三方支付接口
	err = payBean.ThirdPayTest(
		&input.ThirdPayData{
			deposit.Money,
			onlineEntryRecord.ThirdOrderNumber,
			deposit.PaidType,
			oneThird.MerchatId,
			oneThird.Id,
			payRedisKey,
			oneThird.PaidCode,
			apiClients.UserId,
			apiClients.Name,
			apiClients.Secret,
			apiClients.SiteId,
			member.SiteIndexId,
			0,
			"",
			""})
	/*fmt.Println(&input.ThirdPayData{
	deposit.Money,
	onlineEntryRecord.ThirdOrderNumber,
	deposit.PaidType,
	oneThird.MerchatId,
	oneThird.PaidPlatform,
	payRedisKey,
	oneThird.PaidCode,
	apiClients.UserId,
	apiClients.Name,
	apiClients.Secret,
	apiClients.SiteId,
	member.SiteIndexId})*/
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//线上存款 -点卡//todo 三方开发组未开发完，无法对接，此接口仍待优化
func (c *OtherController) CardDeposit(ctx echo.Context) error {
	deposit := new(input.WapOnlineDepositCard)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoBySite(memberRedis.Site, memberRedis.SiteIndex, memberRedis.Account)
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
		return ctx.JSON(200, global.ReplyError(30208, ctx))
	}
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//判断银行类型Id(卡类型)是否存在 todo:点卡类型暂不做判断
	//have, err = bankCardBean.DepositById(deposit.Bank)
	//if err != nil {
	//	global.GlobalLogger.Error("err:", err)
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//if !have {
	//	return ctx.JSON(200, global.ReplyError(10221, ctx))
	//}

	//是否超限
	if paySet.OnlineDepositMax < deposit.Money { //超过上限了
		return ctx.JSON(200, global.ReplyError(10160, ctx))
	}
	if paySet.OnlineDepositMin > deposit.Money { //低于下限了
		return ctx.JSON(200, global.ReplyError(10161, ctx))
	}
	//查看该站点下是否开启线上入款优惠,1首次，2每次
	onlineEntryRecord := new(schema.OnlineEntryRecord)
	//查看会员是否是首次存款
	have, err = onlineEntryRecordBean.IsFirstIncome(member.Id, member.SiteId)
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
	//是否放弃优惠
	if paySet.OnlineIsDeposit == 1 { //放弃优惠
		onlineEntryRecord.IsDiscount = 2 //无优惠
	} else { //不放弃优惠
		//判断优惠是首次还是每次
		if paySet.OnlineIsDepositDiscount == 1 && onlineEntryRecord.IsFirstDeposit != 1 { //首次有优惠 且 会员不是首次充值了，则无优惠
			onlineEntryRecord.IsDiscount = 2 //无优惠
		} else { //其他情况，则有优惠
			if deposit.Money >= paySet.OnlineDiscountStandard { //大于优惠标准
				onlineEntryRecord.DepositDiscount = deposit.Money * paySet.OnlineDiscountPercent * 0.01 //优惠
				if onlineEntryRecord.DepositDiscount > paySet.OnlineDiscountUp {                        //大于优惠上限
					onlineEntryRecord.DepositDiscount = paySet.OnlineDiscountUp
				}
				if deposit.Money >= paySet.OnlineOtherDiscountStandard {
					onlineEntryRecord.OtherDepositDiscount = deposit.Money * paySet.OnlineOtherDiscountPercent * 0.01 //其他优惠
					if onlineEntryRecord.OtherDepositDiscount > paySet.OnlineOtherDiscountUp {
						onlineEntryRecord.OtherDepositDiscount = paySet.OnlineOtherDiscountUp
					}
				}
			}
			onlineEntryRecord.IsDiscount = 1 //有优惠
		}
	}

	//拼接订单号，填充数据
	var siteStr string
	for _, char := range []rune(member.SiteId) {
		siteStr += strconv.Itoa(int(char - 30))
	}
	onlineEntryRecord.ThirdOrderNumber = strconv.FormatInt(time.Now().UnixNano(), 10) + siteStr + strconv.FormatInt(member.Id, 10) + ""
	onlineEntryRecord.SourceDeposit = 1             //pc端
	onlineEntryRecord.LocalOperateTime = 1          //本平台跳到第三方时的时间
	onlineEntryRecord.AmountDeposit = deposit.Money //存入金额
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1                  //未支付
	onlineEntryRecord.PaidType = deposit.PaidType //支付类型id

	/*//根据条件，随机选择一条第三方支付平台
	oneThird, oneThirdNum, err := onlinePaidSetupBean.GetOnePaid(&input.GetOnePaid{member.SiteId, member.SiteIndexId, deposit.PaidType, member.LevelId})
	//根据传入的id,选择对应的第三方*/
	oneThird, oneThirdHas, err := onlinePaidSetupBean.GetOnePaidById(
		&input.GetOnePaidById{
			member.SiteId,
			member.SiteIndexId,
			deposit.PaidType,
			member.LevelId,
			deposit.Id})
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !oneThirdHas {
		return ctx.JSON(200, global.ReplyError(10157, ctx))
	}
	onlineEntryRecord.ThirdId = oneThird.PaidPlatform //第三方支付平台id
	onlineEntryRecord.PaidSetupId = int(oneThird.Id)  //线上支付设定id

	//插入线上入款纪录
	count, err := onlineEntryRecordBean.Add(onlineEntryRecord)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10218, ctx))
	}
	//拼接支付rediskey,用于模板渲染时获取数据
	payRedisKey := "payRedisKey_" + member.SiteId + member.Account
	//获取对应站点的三方对接验证加密记录
	apiClients, has, err := apiClientsBean.GetOneApiClients("zzz") //member.SiteId todo 真实上线时，要对应上站点
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10162, ctx))
	}
	//请求第三方支付接口
	err = payBean.ThirdPayTest(
		&input.ThirdPayData{
			deposit.Money,
			onlineEntryRecord.ThirdOrderNumber,
			deposit.PaidType,
			oneThird.MerchatId,
			oneThird.Id,
			payRedisKey,
			oneThird.PaidCode,
			apiClients.UserId,
			apiClients.Name,
			apiClients.Secret,
			apiClients.SiteId,
			member.SiteIndexId,
			deposit.CardMoney,
			deposit.CardNumber,
			deposit.CardPassword})
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//第三方支付返回信息回调处理
func (c *OtherController) AddOnlineIncomeCallback(ctx echo.Context) error {

	deposit := &struct {
		Status  bool   `form:"status"`
		Code    int64  `form:"code"`
		Message string `form:"message"`
		Data    string `form:"data"`
	}{
	//Status:true,
	//Code:200,
	//Message:"成功",
	//Data:`"agentId":"zzz","agentNum":"a","amount":"654.00","businessNum":"2010711131","order":"151627073353850802667676711","status":1,"sign":"60eb77aba0c702b35b978a9ca388e29c"`,
	}
	depositDate := struct {
		AgentId     string
		AgentNum    string
		Amount      string
		BusinessNum string
		Order       string
		Status      int
		Sign        string
		ErrorMsg    string
	}{}
	//return ctx.String(200,"SUCCESS")
	//wori, _ := ioutil.ReadAll(ctx.Request().Body)
	//fmt.Println(string(wori))
	//fmt.Println("1111111三方回调成功")
	//return ctx.String(200, "SUCCESS")
	//deposit := new(input.WapOnlineDepositCallback)
	//ctx.Request().Method = echo.GET
	//fmt.Printf("wori1:%#v\n",ctx.Request().Form)
	//fmt.Printf("wori2:%#v\n",ctx.Request().URL.Query())
	//ri,_:=ioutil.ReadAll(ctx.Request().Body)
	//fmt.Printf("wori3:%#v\n",string(ri))

	//fmt.Println("1111111三方回调成功")

	//fmt.Println("SUCCESS")
	//deposit1 := new(input.WapOnlineDepositCallback)
	//x := ctx.Response()
	//return ctx.String(200,"SUCCESS")
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//fmt.Println("wori")
	//return ctx.String(200,"SUCCESS")

	//fmt.Printf(">>1:%#v\n", deposit.Status)
	//fmt.Printf(">>2:%#v\n", deposit.Code)
	//fmt.Printf(">>3:%#v\n", deposit.Message)
	//fmt.Printf(">>4:%#v\n", strings.Replace(deposit.Data, "\\\"", "", -1))
	err := json.Unmarshal([]byte(deposit.Data), &depositDate)
	if err != nil {
		fmt.Println("cuowu :", err.Error())
	}
	//fmt.Println(depositDate.Status)
	//fmt.Println(depositDate.AgentId)
	//fmt.Println(depositDate.AgentNum)
	//fmt.Println(depositDate.Amount)
	//fmt.Println(depositDate.BusinessNum)
	//fmt.Println(depositDate.Order)
	//fmt.Println(depositDate.Sign)
	//fmt.Println(depositDate.ErrorMsg)
	//处理返回结果
	if !deposit.Status { //请求失败
		return ctx.JSON(500, global.ReplyError(10219, ctx))
	}
	/*
		如果失败
			更新入款记录状态，入款失败
		如果成功
			对比签名
			从data中取出数据
			根据订单号，查询入款记录
				根据入款记录信息查询获取相关会员信息
					更新会员出入款纪录
					更新会员入款纪录状态为 2 已支付
					更新会员余额
					添加现金流水
					更新会员稽核纪录

	*/
	fmt.Println(deposit.Data)
	incomeInfo, flag, err := onlineEntryRecordBean.GetInfoByOrder(depositDate.Order) //订单号待返回
	//return ctx.String(200, "SUCCESS")
	//incomeInfo, flag, err := onlineEntryRecordBean.GetInfoByOrder("111") //订单号待返回
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(500, global.ReplyError(50108, ctx)) //订单号不合法（订单号不存在）
	}
	if incomeInfo.Status != 1 {
		fmt.Println("已操作则直接返回")
		return ctx.String(200, "SUCCESS")
	}
	//验证签名
	apiClients, has, err := apiClientsBean.GetOneApiClients("zzz") //incomeInfo.SiteId todo  真实上线时，要对应上站点
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(10162, ctx))
	}
	str := "&&&agentId=#" + depositDate.AgentId + "#&&&&&&agentNum=#" + depositDate.AgentNum + "#&&&&&&amount=#" + depositDate.Amount + "#&&&&&&businessNum=#" + depositDate.BusinessNum + "#&&&&&&clientSecret=#" + apiClients.Secret + "#&&&&&&order=#" + depositDate.Order + "#&&&&&&status=#" + strconv.Itoa(depositDate.Status) + "#&&&"
	switch { //去除空值后签名
	case depositDate.AgentId == "":
		strings.Replace(str, "&&&agentId=##&&&", "", 1)
	case depositDate.AgentNum == "":
		strings.Replace(str, "&&&agentNum=##&&&", "", 1)
	case depositDate.Amount == "":
		strings.Replace(str, "&&&amount=##&&&", "", 1)
	case depositDate.BusinessNum == "":
		strings.Replace(str, "&&&businessNum=##&&&", "", 1)
	case apiClients.Secret == "":
		strings.Replace(str, "&&&clientSecret=##&&&", "", 1)
	case depositDate.Order == "":
		strings.Replace(str, "&&&order=##&&&", "", 1)
	case strconv.Itoa(depositDate.Status) == "":
		strings.Replace(str, "&&&status=##&&&", "", 1)
	}

	sign := global.Md5(global.Md5(str))
	fmt.Println(sign, "-------------------------------")
	if sign != depositDate.Sign { //验签不通过
		return ctx.JSON(200, global.ReplyError(10163, ctx))
	}
	//获取会员信息
	member, have, err := memberBean.GetInfoBySite(incomeInfo.SiteId, incomeInfo.SiteIndexId, incomeInfo.MemberAccount)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//查询会员总余额,入款前的系统余额+各视讯余额
	sumBalance, _, err := memberBalanceConversionBean.GetVideoBalance(incomeInfo.MemberId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	sumBalance += member.Balance
	//更新会员出入款统计纪录
	memberCashCount := new(schema.MemberCashCount)
	memberCashCount.Member = incomeInfo.MemberId                              //会员id
	memberCash, have, err := memberCashCountBean.GetInfo(incomeInfo.MemberId) //查询会员现金统计记录
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have { //无记录，则增加
		memberCashCount.DepositCount = 1                      //累计存款次数
		memberCashCount.DepositMax = incomeInfo.AmountDeposit //最大存款金额
		memberCashCount.DepositNum = incomeInfo.AmountDeposit //累计存款金额
	} else { //有记录，则更新
		memberCashCount.DepositCount = memberCash.DepositCount + 1 //累计存款次数
		if incomeInfo.AmountDeposit > memberCash.DepositMax {      //最大存款金额
			memberCashCount.DepositMax = incomeInfo.AmountDeposit
		} else {
			memberCashCount.DepositMax = memberCash.DepositMax
		}
		memberCashCount.DepositNum = memberCash.DepositNum + incomeInfo.AmountDeposit //累计存款金额
	}
	//更新会员余额
	member.Balance += incomeInfo.AmountDeposit //余额+存款金额+优惠金额
	//更新入款状态+第三方转账时间
	incomeInfo.Status = 2 //已经支付
	incomeInfo.ThirdPayTime = time.Now().Unix()
	//添加现金流水
	cash := new(schema.MemberCashRecord)
	cash.SiteId = member.SiteId
	cash.SiteIndexId = member.SiteIndexId
	cash.MemberId = member.Id
	cash.UserName = member.Account
	cash.AgencyId = incomeInfo.AgencyId                                            //代理id
	cash.AgencyAccount = incomeInfo.AgencyAccount                                  //代理账号
	cash.SourceType = 2                                                            //类型:线上入款
	cash.Type = 1                                                                  //存入
	cash.TradeNo = incomeInfo.ThirdOrderNumber                                     //订单号
	cash.DisBalance = incomeInfo.DepositDiscount + incomeInfo.OtherDepositDiscount //优惠金额
	cash.Balance = incomeInfo.AmountDeposit                                        //金额
	cash.AfterBalance = cash.DisBalance + cash.Balance                             //操作后余额
	cash.ClientType = 1                                                            //客户端类型 pc
	cash.Remark = "线上入款"                                                           //备注
	//更新会员稽核纪录
	aduit := new(schema.MemberAudit)
	//查询对应层级的支付设定id
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//sales_site_pay_set根据查到的支付设定id查询支付设定详情
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	//设置线上入款的优惠金额
	aduit.DepositMoney = incomeInfo.DepositDiscount + incomeInfo.OtherDepositDiscount
	//查询时间最近的一条稽核记录
	memberAudits, have, err := auditsBean.AuditLastOne(incomeInfo.SiteId, incomeInfo.MemberAccount)
	//fmt.Println(memberAudits)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//没有稽核记录 或者 有稽核记录状态为已处理
	if !have || (have && memberAudits.Status == 2) {
		//直接添加一条新的稽核记录
	} else {
		//更新上一条稽核记录的结束时间为当前插入新稽核记录的开始时间
		aduit.Id = memberAudits.Id
		if sumBalance < paySet.AuditRelaxQuota { //会员总余额是否小于常态稽核放宽额度
			onlineDepositBean.AuditStatus = 2 //更新该会员所有稽核状态为已处理
		}
		//插入一条新的稽核记录开始时间为当前时间
	}
	//是否开启了常态稽核 1:开启
	if paySet.OnlineIsNormalAudit == 1 {
		//常态稽核
		aduit.NormalMoney = incomeInfo.AmountDeposit * paySet.OnlineNormalAuditPercent
	}
	//是否开启了综合性稽核 1:开启
	if paySet.OnlineIsMultipleAudit == 1 {
		//综合性稽核
		aduit.MultipleMoney = (incomeInfo.AmountDeposit + aduit.DepositMoney) * float64(paySet.OnlineMultipleAuditTimes)
	}
	aduit.BeginTime = time.Now().Unix()
	aduit.MemberId = member.Id
	aduit.Account = member.Account
	aduit.DepositMoney = incomeInfo.DepositDiscount
	aduit.SiteId = member.SiteId
	aduit.SiteIndexId = member.SiteIndexId
	aduit.Status = 4 //未处理
	aduit.RelaxMoney = int64(paySet.AuditRelaxQuota)
	//更新会员余额
	member.Balance += aduit.DepositMoney
	//组装更新对象
	onlineDepositBean.Audit = aduit
	onlineDepositBean.Cash = memberCashCount
	onlineDepositBean.Member = &member
	onlineDepositBean.OnlineRecord = &incomeInfo
	onlineDepositBean.MemberRecord = cash
	//打印
	//fmt.Println(onlineDepositBean.OnlineRecord)
	//fmt.Println(onlineDepositBean.Audit)
	//fmt.Println(onlineDepositBean.Cash)
	//fmt.Println(onlineDepositBean.Member)
	//fmt.Println(onlineDepositBean.MemberRecord)
	//执行事务更新
	err = onlineEntryRecordBean.Update(onlineDepositBean)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	fmt.Println("SUCCESS")
	return ctx.String(200, "SUCCESS")
	//return ctx.String(200, "SUCCESS")
	//return ctx.NoContent(204)
}

//获取网银在线的银行
func (c *OtherController) GetOnlineIncomeBank(ctx echo.Context) error {
	onlineIncome := new(input.GetOnlineIncomeBank)
	code := global.ValidRequestMember(onlineIncome, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	apiClients, has, err := apiClientsBean.GetOneApiClients("zzz") //incomeInfo.SiteId todo  真实上线时，要对应上站点
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(10162, ctx))
	}
	data := make(url.Values)
	data["clientUserId"] = []string{strconv.FormatInt(apiClients.UserId, 10)}
	data["clientName"] = []string{apiClients.Name}
	data["clientSecret"] = []string{apiClients.Secret}
	data["payId"] = []string{strconv.FormatInt(onlineIncome.PayId, 10)}
	fmt.Println(data)
	res, err := http.PostForm("http://olmanage.pk1358.com/api/v1/bank/list", data) //todo 路径暂时写死，为公测路径
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	fmt.Println(res)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	f := struct {
		Status bool        `json:"status"`
		Code   int64       `json:"code"`
		Data   interface{} `json:"data"`
	}{}
	fmt.Println(string(b))
	err = json.Unmarshal(b, &f)
	fmt.Println(err)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !f.Status {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	fmt.Println("-------------------------", f.Data)
	return ctx.JSON(200, global.ReplyItem(f.Data))
}
