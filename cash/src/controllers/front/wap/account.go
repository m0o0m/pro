package wap

import (
	"controllers/front/wap/data_merge"
	"fmt"
	"framework/render"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"net/http"
	"strings"
	"time"
	//"github.com/go-redis/redis"
	"encoding/json"
	"strconv"
)

//会员中心
type AccountController struct {
	WapBaseController
}

//会员中心
func (c *AccountController) Index(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	page := new(data_merge.Account)
	//获取会员层级信息
	//获取该会员所处层级详情
	levelInfo, flag, err := memberLevelBean.GetLevelInfo(member.LevelId)
	if err != nil {
		page.IsSelf = 2 //1开启 2未开启
	}
	//找不到该层级
	if !flag || levelInfo.LevelId == "" {
		page.IsSelf = 2 //1开启 2未开启
	} else {
		page.IsSelf = int(levelInfo.IsSelfRebate)
	}
	return c.Render(page, ctx)
}

//会员详情
func (c *AccountController) Info(ctx echo.Context) error {
	return c.Render(new(data_merge.Info), ctx)
}

//获取个人信息
func (c *AccountController) GetInfo(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	mi := new(input.MemberInfoSelf)
	mi.Id = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	data, has, err := baseInfoBean.MemberSelfInfo(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	if data.Realname == "" {
		mbank, has, err := MemberBankBean.GetMemberBankOne(member.Id)
		if has && err == nil {
			_, err = memberBean.UpdateMemberReallname(member.Site, member.SiteIndex, member.Account, mbank.CardName)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
			}
			data.Realname = mbank.CardName
		}
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改会员信息
func (c *AccountController) PutInfo(ctx echo.Context) error {
	base_info := new(input.EditMeans)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	flag := true
	//检验手机
	if len(base_info.PhoneNum) != 0 {
		flag = global.CheckPhoneNumber(base_info.PhoneNum)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20014, ctx))
		}
	}
	//检验邮箱
	if len(base_info.EmailNum) != 0 {
		flag = global.CheckEmail(base_info.EmailNum)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20015, ctx))
		}
	}

	//检验qq
	if len(base_info.QqNum) != 0 {
		flag = global.Checkqq(base_info.QqNum)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20016, ctx))
		}
	}
	//检验身份证
	if len(base_info.Card) != 0 {
		flag = global.CheckIdentity(base_info.Card)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20013, ctx))
		}
	}

	_, err := baseInfoBean.ChangeMeans(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取会员的详细信息
func (*AccountController) MemberDetail(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.MemberInfoSelf)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.Id = member.Id
	data, has, err := baseInfoBean.OneMemberDetail(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改密码
func (c *AccountController) ModifyPas(ctx echo.Context) error {
	return c.Render(new(data_merge.ModifyPas), ctx)
}

//修改密码
func (c *AccountController) ModifyInfo(ctx echo.Context) error {
	return c.Render(new(data_merge.ModifyInfo), ctx)
}

//银行卡
func (c *AccountController) BankCard(ctx echo.Context) error {
	return c.Render(new(data_merge.BankCard), ctx)
}

//获取个人出款银行列表
func (bic *AccountController) Bank(ctx echo.Context) error {
	//用户参数
	base_info := new(input.MemberBankList)
	//登录用户信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	base_info.SiteId = member.Site
	base_info.SiteIndexId = member.SiteIndex
	data, err := baseInfoBean.MemberBankList(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k, v := range data {
		data[k].Card = global.BankStr(v.Card)
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

func bankStr(card string) (cards string) {
	var sss string
	sss = card[:4] + " **** **** **** " + card[len(card)-4:]
	return sss
}

//银行卡
func (c *AccountController) BankCardAdd(ctx echo.Context) error {
	return c.Render(new(data_merge.BankCardAdd), ctx)
}

//银行卡
func (c *AccountController) BankAddInfo(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)

	data, err := bankCardBean.IndexAllBankOut(member.Site, member.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	return ctx.JSON(200, global.ReplyItem(data))
}

//添加出款银行
func (bic *AccountController) BankAdd(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.MemberBankAdd)

	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.SiteIndexId = member.SiteIndex

	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	base_info.MemberId = member.Id
	base_info.SiteId = member.Site

	//检查银行卡是否合法
	flag := global.CheckCardNumber(base_info.Card)
	if !flag {
		return ctx.JSON(200, global.ReplyError(30060, ctx))
	}
	//检验该银行卡是否已经被添加
	_, has, count, err := baseInfoBean.CheckOutBank(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(500, global.ReplyError(30069, ctx))
	}
	if count > 2 {
		return ctx.JSON(500, global.ReplyError(50117, ctx))
	}
	if len(base_info.Password) == 0 && count == 0 {
		return ctx.JSON(500, global.ReplyError(20005, ctx))
	}
	//添加银行卡
	_, err = baseInfoBean.MemberOutBankAdd(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(base_info.Password) != 0 && count == 0 {
		password, _ := global.MD5ByStr(base_info.Password, global.EncryptSalt)
		_, err = memberBean.ChangePassword(member.Id, password, 2)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	return ctx.NoContent(204)
}

//获取会员对应层级对应配置的支付设定
func (c *AccountController) GetPaySetData(ctx echo.Context) error {
	memberInfo := ctx.Get("member").(*global.MemberRedisToken)
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{memberInfo.Site, memberInfo.SiteIndex, memberInfo.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(paySet))
}

//获取会员对应层级对应配置的支付设定--快捷支付时用
func (c *AccountController) GetPaySetDataByAccount(ctx echo.Context) error {
	member := new(input.MemberLevelPaySetAndFirstDeposit)
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//取会员信息
	memberInfo, has1, err := memberBean.GetInfoBySite(memberRedis.Site, memberRedis.SiteIndex, member.Account)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has1 {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, memberInfo.LevelId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(paySet))
}

//存款数据
func (*AccountController) GetIncomeData(ctx echo.Context) error {
	incomeData := new(back.GetIncomeData)
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
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

//快速充值中心存款数据
func (*AccountController) GetFastIncomeData(ctx echo.Context) error {
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
	siteInfo, flag, err := siteDomainBean.GetSiteByDomain(siteDomin)
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

//添加一条公司入款记录
func (c *AccountController) AddCompanyIncome(ctx echo.Context) error {
	//member := new(input.AddCompanyIncome)
	//code := global.ValidRequestMember(member, ctx)
	//if code != 0 {
	//	return ctx.JSON(200, global.ReplyError(code, ctx))
	//}
	//fmt.Println(member)
	//count, err := memberCompanyIncomeBean.AddCompanyIncome(member)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//if count == 0 {
	//	return ctx.JSON(200, global.ReplyError(10058, ctx))
	//}
	//return ctx.NoContent(204)

	deposit := new(input.AddCompanyIncome)
	code := global.ValidRequestMember(deposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取当前会员的信息siteId ,siteIndexId,member_id
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//判断会员Id是否存在
	member, have, err := memberBean.GetInfoById(memberRedis.Id)
	fmt.Println(member.Account)
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
	if paySet.OnlineDepositMax < deposit.DepositMoney {
		//超过上限了
		return ctx.JSON(200, global.ReplyError(10158, ctx))
	}
	if paySet.OnlineDepositMin > deposit.DepositMoney {
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

	companyIncome.DepositCount = deposit.DepositMoney + companyIncome.DepositDiscount + companyIncome.OtherDiscount //总额
	companyIncome.OrderNum = uuid.NewV4().String()                                                                  //订单号
	companyIncome.ClientType = 2                                                                                    //wap端
	companyIncome.CreateTime = 1                                                                                    //本平台跳到第三方时的时间
	companyIncome.DepositMoney = deposit.DepositMoney                                                               //存入金额
	companyIncome.MemberId = member.Id
	companyIncome.LevelId = member.LevelId
	companyIncome.Account = member.Account
	companyIncome.SiteId = member.SiteId
	companyIncome.SiteIndexId = member.SiteIndexId
	companyIncome.AgencyId = agency.Id
	companyIncome.AgencyAccount = agency.Account
	companyIncome.Status = 1 //未支付
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
	return ctx.NoContent(204)
}

//添加线上入款记录(网银、点卡除外)
func (c *AccountController) AddOnlineIncome(ctx echo.Context) error {
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
	onlineEntryRecord.SourceDeposit = 2                    //wap端
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

//添加线上入款记录(网银)
func (c *AccountController) AddOnlineBankIncome(ctx echo.Context) error {
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
	onlineEntryRecord.SourceDeposit = 2                    //wap端
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

//线上存款 -点卡
func (w *AccountController) CardDeposit(ctx echo.Context) error {
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
	//是否超限
	if paySet.OnlineDepositMax < deposit.Money {
		//超过上限了
		return ctx.JSON(200, global.ReplyError(10160, ctx))
	}
	if paySet.OnlineDepositMin > deposit.Money {
		//低于下限了
		return ctx.JSON(200, global.ReplyError(10161, ctx))
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

	//赋值线上入款(默认未成功)
	onlineEntryRecord.ThirdOrderNumber = uuid.NewV4().String() //strings.Replace(time.Now().String()[:10], "-", "", 2) + member.SiteId + member.SiteIndexId + strconv.FormatInt(time.Now().Unix(), 10) //订单号
	onlineEntryRecord.SourceDeposit = 2                        //wap端
	onlineEntryRecord.LocalOperateTime = 1                     //本平台跳到第三方时的时间
	onlineEntryRecord.AmountDeposit = deposit.Money            //存入金额
	onlineEntryRecord.MemberId = member.Id
	onlineEntryRecord.Level = member.LevelId
	onlineEntryRecord.MemberAccount = member.Account
	onlineEntryRecord.SiteId = member.SiteId
	onlineEntryRecord.SiteIndexId = member.SiteIndexId
	onlineEntryRecord.AgencyId = agency.Id
	onlineEntryRecord.AgencyAccount = agency.Account
	onlineEntryRecord.Status = 1                  //未支付
	onlineEntryRecord.PaidType = deposit.PaidType //支付类型id

	//根据条件，随机选择一条第三方支付平台
	oneThird, oneThirdNum, err := onlinePaidSetupBean.GetOnePaid(&input.GetOnePaid{member.SiteId, member.SiteIndexId, deposit.PaidType, member.LevelId})
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if oneThirdNum <= 0 {
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
	onlineDepositBean.Member = &member
	onlineDepositBean.OnlineRecord = onlineEntryRecord
	onlineDepositBean.MemberRecord = cash
	err = onlineEntryRecordBean.Update(onlineDepositBean)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//抢红包 获取已经分配好的红包
func (*AccountController) GetSnatch(ctx echo.Context) error {
	snatch := new(input.GetSnatch)
	code := global.ValidRequest(snatch, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	fmt.Println(snatch)
	member := ctx.Get("member").(*global.MemberRedisToken)

	data, err := redPacketSetBean.GetOne(snatch.SetId) //查询红包活动
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	if data.Id == 0 { //没有该红包活动
		return ctx.JSON(200, global.ReplyError(71106, ctx))
	}
	if data.StartTime > time.Now().Unix() {
		return ctx.JSON(200, global.ReplyError(91201, ctx))
	}
	var Ip string
	if data.IsIp == 2 {
		Ip = ctx.RealIP()
		count, err := redPacketSetBean.IpRed(snatch.SetId, Ip) //查询红包活动
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return render.PageErr(60000, ctx)
		}
		if count > 0 { //没有该红包活动
			return ctx.JSON(200, global.ReplyError(71108, ctx))
		}
	}
	if data.LevelId != "" {
		strArr := strings.Split(data.LevelId, ",")
		for _, v := range strArr {
			if member.LevelId == v {
				return ctx.JSON(200, global.ReplyError(71105, ctx))
			}
		}
		count, err := redPacketSetBean.UserRed(snatch.SetId, member.Id) //查询红包活动
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return render.PageErr(60000, ctx)
		}
		if count >= data.MaxCount { //没有该红包活动
			return ctx.JSON(200, global.ReplyError(71109, ctx))
		}
	}

	//查询可分配的红包
	b, redinfo, err := redPacketSetBean.GetRebInfo(snatch.SetId) //查询红包活动
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	if !b { //没有查到可分配的红包
		return ctx.JSON(200, global.ReplyError(71107, ctx))
	}

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	//修改红包状态
	redinfo.MemberId = member.Id
	redinfo.Account = member.Account
	redinfo.PType = 1
	num, err := redPacketSetBean.SetRebMakeSure(redinfo, sess) //修改红包状态
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//增加会员金额
	code, num, newCashRecord, err := redPacketSetBean.SetRebMemBalance(member.Id, redinfo.Money, member.Site, sess)

	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if num == 0 {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	if data.IsGenerate != 2 {
		return ctx.JSON(200, global.ReplyError(71104, ctx))
	}

	agencyInfo, _, err := thirdAgencyBean.BaseInfo(&input.ThirdAgencyInfo{member.Site, member.SiteIndex, newCashRecord.AgencyId})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	newCashRecord.AgencyAccount = agencyInfo.Account

	//会员现金流水
	num, err = memberCashBean.AddNewRecordSess(&newCashRecord, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if num == 0 {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	err = sess.Commit()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(redinfo))
}

//优惠申请大厅数据请求
func (*AccountController) ApplyData(ctx echo.Context) error {
	memberRedis := ctx.Get("member").(*global.MemberRedisToken)
	//查询活动列表
	list, err := sitePromotionConfig.GetSitePromotionConfig(&input.SitePromotionConfig{memberRedis.Site, memberRedis.SiteIndex})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	applyPros := make([]back.Pro, 0)
	applyPro := new(back.Pro)
	for k, _ := range list {
		applyPro.ProId = list[k].Id
		applyPro.ProTitle = list[k].ProTitle
		applyPro.ProContent = list[k].ProContent
		applyPros = append(applyPros, *applyPro)
	}
	return ctx.JSON(200, global.ReplyItem(applyPros))
}

//优惠活动申请提交
func (*AccountController) ApplySubmit(ctx echo.Context) error {
	apply := new(input.SelfHelpApllyAdd)
	code := global.ValidRequest(apply, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
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

//第三方支付返回信息回调处理
func (*AccountController) AddOnlineIncomeCallback(ctx echo.Context) error {

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
	aduit.Status = 1 //未处理
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
