package account

import (
	"controllers"
	"fmt"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

//代理管理
type ThirdAgencyController struct {
	controllers.BaseController
}

//代理列表
func (tac *ThirdAgencyController) Index(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	tac.GetParam(listparam, ctx)
	thirdAgencyState := new(input.ThirdAgency)
	code := global.ValidRequest(thirdAgencyState, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	//用户id
	user_id := int(user.Id)
	//用户等级
	user_level := user.Level
	if user_level > 4 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	} else {
		switch user_level {
		case 1:
			thirdAgencyState.Id = user_id
		case 2:
			thirdAgencyState.FirstId = user_id
		case 3:
			thirdAgencyState.SecondId = int64(user_id)
		case 4:
			thirdAgencyState.Isvague = 0
			thirdAgencyState.ThirdId = int64(user_id)
		}
	}
	data, count, err := agencyCountBean.GetSearchThirdAgency(thirdAgencyState, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	num := new(back.AgentNumberPerson)
	if len(data) > 0 {
		for _, v := range data {
			if v.Status == 1 {
				num.OpenNum = num.OpenNum + 1
			} else {
				num.CloseNum = num.CloseNum + 1
			}
			if v.IsLogin == 1 {
				num.OnlineNum = num.OnlineNum + 1
			}
		}
	}
	num.TotalNum = count
	var list = make(map[string]interface{})
	list["data"] = data
	list["num"] = num
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(data)), count, ctx))
}

//新增代理
func (tac *ThirdAgencyController) Add(ctx echo.Context) error {
	agency := new(input.AgencyAdd)
	code := global.ValidRequest(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看账号是否存在
	has, err := subAccountBeen.GetAccount(agency.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//两次密码不一致
	if agency.Password != agency.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	md5Password, err := global.MD5ByStr(agency.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(30044, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//得到站点id和角色id
	agency.SiteId = user.SiteId
	agency.Password = md5Password
	agency.Level = user.Level
	if agency.Level != 1 {
		//非开户人登录，上级id默认为登录人id
		agency.ParentId = user.Id
	} else {
		if agency.SiteIndexId == "" {
			return ctx.JSON(200, global.ReplyError(10050, ctx))
		}
	}
	count, err := agencyCountBean.ThirdAgencyAdd(agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30054, ctx))
	}
	return ctx.NoContent(204)
}

//启用/禁用
func (tac *ThirdAgencyController) Status(ctx echo.Context) error {
	thirdAgency := new(input.ThirdAgencyInfo)
	code := global.ValidRequest(thirdAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人和股东和总代才能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level > 3 || user.RoleId > 3 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}
	if user.Level == 3 || user.RoleId == 3 {
		if thirdAgency.Id != user.Id {
			return ctx.JSON(200, global.ReplyError(60001, ctx))
		}
	}
	count, err := thirdAgencyBean.UpdateStatus(thirdAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30037, ctx))
	}
	return ctx.NoContent(204)
}

//获取基本资料
func (tac *ThirdAgencyController) BaseInfo(ctx echo.Context) error {
	thirdAgency := new(input.ThirdAgencyInfo)
	code := global.ValidRequest(thirdAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人和股东和总代和代理都能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level > 4 || user.RoleId > 4 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}
	if user.Level == 4 || user.RoleId == 4 {
		if thirdAgency.Id != user.Id {
			return ctx.JSON(200, global.ReplyError(60001, ctx))
		}
	}
	info, has, err := thirdAgencyBean.BaseInfo(thirdAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//修改基本资料
func (tac *ThirdAgencyController) BaseInfoEdit(ctx echo.Context) error {
	thirdAccount := new(input.FirstAgencyEdit)
	code := global.ValidRequest(thirdAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//两次密码不一致
	if thirdAccount.Password != thirdAccount.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	if thirdAccount.Password != "" {
		md5Password, err := global.MD5ByStr(thirdAccount.Password, global.EncryptSalt)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		thirdAccount.Password = md5Password
	}
	count, err := agencyCountBean.FirstAgencyEdit(thirdAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	//修改成功
	return ctx.NoContent(204)
}

//获取详细资料
func (tac *ThirdAgencyController) DetailInfo(ctx echo.Context) error {
	third := new(input.ThirdInformation)
	//获取结构体的json数据
	code := global.ValidRequest(third, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := agencyThirdInfoBean.ThirdAgencyInfo(third)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var agencyInfo = make(map[string]interface{})
	if !has {
		agencyInfo["data"] = nil
		agencyInfo["bank"] = nil
		agencyInfo["doMain"] = nil
		return ctx.JSON(200, global.ReplyItem(agencyInfo))
	}
	//代理银行卡
	bankInfo, err := agencyThirdInfoBean.ThirdAgencyBankInfo(third.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//代理域名
	doMainInfo, err := agencyThirdInfoBean.ThirdAgencyDomainInfo(third.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	agencyInfo["data"] = data
	agencyInfo["bank"] = bankInfo
	agencyInfo["doMain"] = doMainInfo
	return ctx.JSON(200, global.ReplyItem(agencyInfo))
}

//修改详细资料
func (tac *ThirdAgencyController) DetailInfoEdit(ctx echo.Context) error {
	third := new(input.ThirdInformationUpdata)
	//获取结构体的json数据
	code := global.ValidRequest(third, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	ok := global.CheckEmail(third.Email)
	if !ok {
		return ctx.JSON(200, global.ReplyError(20015, ctx))
	}
	ok = global.CheckIdentity(third.Card)
	if !ok {
		return ctx.JSON(200, global.ReplyError(20013, ctx))
	}
	ok = global.Checkqq(third.QQ)
	if !ok {
		return ctx.JSON(200, global.ReplyError(20016, ctx))
	}
	ok = global.CheckPhoneNumber(third.Phone)
	if !ok {
		return ctx.JSON(200, global.ReplyError(20014, ctx))
	}
	//检验银行卡号是否合法
	if len(third.Cards) > 0 {
		for _, v := range third.Cards {
			ok = global.CheckCardNumber(v)
			if !ok {
				return ctx.JSON(200, global.ReplyError(30060, ctx))
			}
		}
	}
	//检验域名是否合法
	if len(third.Domain) > 0 {
		for _, n := range third.Domain {
			ok = global.DomainCheck(n)
			if !ok {
				return ctx.JSON(200, global.ReplyError(30068, ctx))
			}
		}
	}
	_, num, err := agencyThirdInfoBean.ThirdAgencyInfoUpdata(third)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 1 {
		return ctx.JSON(200, global.ReplyError(50165, ctx))
	}
	return ctx.NoContent(204)
}

//获取会员注册优惠设定
func (tac *ThirdAgencyController) MemberRegDiscountSet(ctx echo.Context) error {
	tdSet := new(input.ThirdDiscountSet)
	//获取结构体的json数据
	code := global.ValidRequest(tdSet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency_schema := new(schema.Agency)
	agency_schema.SiteId = tdSet.SiteId
	agency_schema.Id = tdSet.AccountId
	dl, has, err := agencyBean.GetOneAgencyByid(agency_schema)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	//用户id
	user_id := user.Id
	//用户等级
	user_level := user.Level
	if user_level == 1 && dl.Level >= 1 {
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneThirdDiscountSet(tdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	if user_level == 3 && dl.Level >= 3 {
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneThirdDiscountSet(tdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	if user_level == 4 && dl.Level >= 4 {
		tdSet.AccountId = user_id
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneThirdDiscountSet(tdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	if user_level == 2 && dl.Level >= 2 {
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneThirdOtherDiscountSet(tdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	return ctx.JSON(200, global.ReplyError(50082, ctx))
}

//修改会员注册优惠设定
func (tac *ThirdAgencyController) MemberRegDiscountSetEdit(ctx echo.Context) error {
	tdSetUpdata := new(input.ThirdDiscountUpdata)
	code := global.ValidRequest(tdSetUpdata, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	tdSetUpdata.SiteId = user.SiteId
	user_lever := user.Level
	if user_lever == 2 {
		_, err := agencyMemberRegisterDiscountSetBean.UpdataThirdOtherSet(tdSetUpdata)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.NoContent(204)
	}
	count, err := agencyMemberRegisterDiscountSetBean.UpdataThirdSet(tdSetUpdata)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50172, ctx))
	}
	return ctx.NoContent(204)
}

//推广域名列表
func (tac *ThirdAgencyController) SpreadDomain(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	tac.GetParam(listparam, ctx)
	agencyDomain := new(input.AgencyThirdDomainList)
	code := global.ValidRequest(agencyDomain, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, count, err := agencyDomainBeen.GetList(agencyDomain, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(30063, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//添加推广域名
func (tac *ThirdAgencyController) SpreadDomainAdd(ctx echo.Context) error {
	agencyDomain := new(input.AgencyThirdDomain)
	code := global.ValidRequest(agencyDomain, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//检查域名是否合法
	ok := global.DomainCheck(agencyDomain.Domain)
	if !ok {
		return ctx.JSON(200, global.ReplyError(30068, ctx))
	}
	//查看域名是否存在
	has, err := agencyDomainBeen.GetDomain(agencyDomain.Domain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30065, ctx))
	}
	count, err := agencyDomainBeen.Add(agencyDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30066, ctx))
	}
	return ctx.NoContent(204)
}

//修改推广域名
func (tac *ThirdAgencyController) SpreadDomainEdit(ctx echo.Context) error {
	agencyDomain := new(input.AgencyThirdDomainEdit)
	code := global.ValidRequest(agencyDomain, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//检查域名是否合法
	ok := global.DomainCheck(agencyDomain.Domain)
	if !ok {
		return ctx.JSON(200, global.ReplyError(30068, ctx))
	}
	//查看域名是否存在
	has, err := agencyDomainBeen.GetDomains(agencyDomain.Id, agencyDomain.Domain)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30065, ctx))
	}
	count, err := agencyDomainBeen.UpdateInfo(agencyDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//删除推广域名
func (tac *ThirdAgencyController) SpreadDomainDel(ctx echo.Context) error {
	agencyDomain := new(input.AgencyThirdDomainDel)
	code := global.ValidRequest(agencyDomain, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := agencyDomainBeen.Delete(agencyDomain.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30064, ctx))
	}
	return ctx.NoContent(204)
}

//取所有的总代的id和名称
func (tac *ThirdAgencyController) SecondIdNameByFirst(ctx echo.Context) error {
	site := new(input.SecondIdNameBySite)
	code := global.ValidRequest(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	fmt.Println("site:", site)
	user_level := agency.Level
	if user_level == 1 {
		site.SiteId = agency.SiteId
		data, err := agencyBean.GetAllSecondIdName(site)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	if user_level == 2 {
		site.FirstId = agency.Id
		site.SiteId = agency.SiteId
		data, err := agencyBean.GetAllSecondIdName(site)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	return ctx.JSON(200, global.ReplyItem(nil))
}

//取站点下所有代理id和帐号
func (tac *ThirdAgencyController) ThirdIdNameBySite(ctx echo.Context) error {
	site := new(input.SecondIdNameBySite)
	code := global.ValidRequest(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := agencyBean.GetAllThirdIdName(site)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
