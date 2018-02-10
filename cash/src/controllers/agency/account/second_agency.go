package account

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

//总代理管理
type SecondAgencyController struct {
	controllers.BaseController
}

//总代理列表
func (sac *SecondAgencyController) Index(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	sac.GetParam(listparam, ctx)
	secondAgencyState := new(input.SecondAgency)
	code := global.ValidRequest(secondAgencyState, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	//用户id
	user_id := int(user.Id)
	//用户等级
	user_level := user.Level
	if user_level > 3 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	} else {
		switch user_level {
		case 1:
			secondAgencyState.Id = user_id
		case 2:
			secondAgencyState.FirstId = user_id
		case 3:
			secondAgencyState.Isvague = 0
			secondAgencyState.SecondId = int64(user_id)
		}
	}
	data, count, err := agencyCountBean.GetSearchSecondAgency(secondAgencyState, listparam)
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

//新增总代理
func (sac *SecondAgencyController) Add(ctx echo.Context) error {
	agency := new(input.AgencyAdd)
	code := global.ValidRequest(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查看账号是否存在
	has, err := subAccountBeen.GetAccount(agency.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
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
	count, err := agencyCountBean.SecondAgencyAdd(agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30053, ctx))
	}
	return ctx.NoContent(204)
}

//启用/禁用
func (sac *SecondAgencyController) Status(ctx echo.Context) error {
	secondAgency := new(input.SecondAgencyInfo)
	code := global.ValidRequest(secondAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人和股东才能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level > 2 || user.RoleId > 2 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}
	if user.Level == 2 || user.RoleId == 2 {
		if secondAgency.Id != user.Id {
			return ctx.JSON(200, global.ReplyError(60001, ctx))
		}
	}
	count, err := secondAgencyBean.UpdateStatus(secondAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30037, ctx))
	}
	return ctx.NoContent(204)
}

//获取基本资料
func (sac *SecondAgencyController) BaseInfo(ctx echo.Context) error {
	secondAgency := new(input.SecondAgencyInfo)
	code := global.ValidRequest(secondAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人和股东和总代才能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level > 3 || user.RoleId > 3 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}
	if user.Level == 3 || user.RoleId == 3 {
		if secondAgency.Id != user.Id {
			return ctx.JSON(200, global.ReplyError(60001, ctx))
		}
	}
	info, has, err := secondAgencyBean.BaseInfo(secondAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//修改基本资料
func (sac *SecondAgencyController) BaseInfoEdit(ctx echo.Context) error {
	secondAccount := new(input.FirstAgencyEdit)
	code := global.ValidRequest(secondAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//两次密码不一致
	if secondAccount.Password != secondAccount.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	if secondAccount.Password != "" {
		md5Password, err := global.MD5ByStr(secondAccount.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		secondAccount.Password = md5Password
	}
	count, err := agencyCountBean.FirstAgencyEdit(secondAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	//修改成功
	return ctx.NoContent(204)
}

//获取会员注册优惠设定
func (sac *SecondAgencyController) MemberRegDiscountSet(ctx echo.Context) error {
	sdSet := new(input.SecondDiscountSet)
	code := global.ValidRequest(sdSet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency_schema := new(schema.Agency)
	agency_schema.Id = sdSet.AcountId
	agency_schema.SiteId = sdSet.SiteId
	zd, has, err := agencyBean.GetOneAgencyByid(agency_schema)
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
	if user_level == 1 && zd.Level >= 1 {
		sdSet.UserId = user_id
		sdSet.SiteId = user.SiteId
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneSecondDiscountSet(sdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	if user_level == 2 && zd.Level >= 2 {
		sdSet.UserId = user_id
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneSecondDiscountSet(sdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	if user_level == 3 && zd.Level >= 3 {
		sdSet.AcountId = user_id
		sdSet.UserId = user_id
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneSecondDiscountSet(sdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	return ctx.JSON(200, global.ReplyItem(nil))
}

//修改会员注册优惠设定
func (sac *SecondAgencyController) MemberRegDiscountSetEdit(ctx echo.Context) error {
	sdSetUpdata := new(input.SecondDiscountUpdata)
	code := global.ValidRequest(sdSetUpdata, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	sdSetUpdata.SiteId = user.SiteId
	count, err := agencyMemberRegisterDiscountSetBean.UpdataSecondSet(sdSetUpdata)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50172, ctx))
	}
	return ctx.NoContent(204)
}

//取股东id和名称下拉框
func (sac *SecondAgencyController) FirstIdByStieId(ctx echo.Context) error {
	site := new(input.FirstIdNameBySite)
	code := global.ValidRequest(site, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency := ctx.Get("user").(*global.RedisStruct)
	user_level := agency.Level
	if user_level == 1 {
		site.SiteId = agency.SiteId
		data, err := agencyBean.GetAllFirstIdName(site)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	return ctx.JSON(200, global.ReplyItem(nil))
}
