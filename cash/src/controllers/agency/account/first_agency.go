package account

import (
	"controllers"

	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"strings"
)

//股东管理
type FirstAgencyController struct {
	controllers.BaseController
}

//股东列表
func (fac *FirstAgencyController) Index(ctx echo.Context) error {
	//获取用户参数
	firstAgencyState := new(input.FirstAgency)
	code := global.ValidRequest(firstAgencyState, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	//用户id
	user_id := int(user.Id)
	//用户等级
	user_level := user.Level

	//开户人和股东才能获取
	if user_level > 2 || user.RoleId > 2 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	var (
		err   error                  //错误
		count int64                  //返回数据的数量
		data  []back.FirstAgencyBack //返回的数据集合
	)
	listparam := new(global.ListParams)
	//获取listparam的数据
	fac.GetParam(listparam, ctx)
	if user_level == 1 {
		data, count, err = agencyCountBean.GetSearchFirstAgency(firstAgencyState, listparam, user)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	if user_level == 2 {
		//股东
		firstAgencyState.Isvague = 0
		firstAgencyState.FirstId = user_id
		data, count, err = agencyCountBean.GetSearchFirstAgency(firstAgencyState, listparam, user)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
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

//新增股东
func (fac *FirstAgencyController) Add(ctx echo.Context) error {
	firstAgency := new(input.FirstAgencyAdd)
	code := global.ValidRequest(firstAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人身份才能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level != 1 || user.RoleId != 1 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	//查看账号是否存在
	has, err := subAccountBeen.GetAccount(firstAgency.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//两次密码不一致
	if firstAgency.Password != firstAgency.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	md5Password, err := global.MD5ByStr(firstAgency.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//得到站点id和角色id
	firstAgency.SiteId = user.SiteId
	firstAgency.ParentId = user.Id
	firstAgency.Password = md5Password
	count, err := agencyCountBean.FirstAgencyAdd(firstAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30051, ctx))
	}
	return ctx.NoContent(204)
}

//启用/禁用
func (fac *FirstAgencyController) Status(ctx echo.Context) error {
	firstAgenncy := new(input.FirstAgencyInfo)
	code := global.ValidRequest(firstAgenncy, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人身份才能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level != 1 || user.RoleId != 1 {
		return ctx.JSON(200, global.ReplyError(30239, ctx))
	}
	//查看id是否为股东
	has, err := subAccountBeen.GetFirstAgent(firstAgenncy.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30238, ctx))
	}
	count, err := subAccountBeen.UpdateStatus(firstAgenncy.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30037, ctx))
	}
	return ctx.NoContent(204)
}

//获取基本资料
func (fac *FirstAgencyController) BaseInfo(ctx echo.Context) error {
	first_agency := new(input.FirstAgencyInfo)
	code := global.ValidRequest(first_agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人和股东才能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level > 2 || user.RoleId > 2 {
		return ctx.JSON(200, global.ReplyError(30239, ctx))
	}
	if user.Level == 2 || user.RoleId == 2 {
		if first_agency.Id != user.Id {
			return ctx.JSON(200, global.ReplyError(30239, ctx))
		}
	}
	info, has, err := firstAgencyBean.BaseInfo(first_agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30241, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//修改基本资料
func (fac *FirstAgencyController) BaseInfoEdit(ctx echo.Context) error {
	firstAgency := new(input.FirstAgencyEdit)
	code := global.ValidRequest(firstAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开户人和股东才能操作
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level > 2 || user.RoleId > 2 {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	//如果是股东,只能修改自己的资料
	if user.Level == 2 || user.RoleId == 2 {
		if firstAgency.Id != user.Id {
			return ctx.JSON(200, global.ReplyError(60001, ctx))
		}
	}
	//两次密码不一致
	if firstAgency.Password != firstAgency.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	if firstAgency.Password != "" {
		md5Password, err := global.MD5ByStr(firstAgency.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		firstAgency.Password = md5Password
	}
	count, err := agencyCountBean.FirstAgencyEdit(firstAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30005, ctx))
	}
	//修改成功
	return ctx.NoContent(204)
}

//获取会员注册优惠设定
func (fac *FirstAgencyController) MemberRegDiscountSet(ctx echo.Context) error {
	fdSet := new(input.FirstDiscountSet)
	code := global.ValidRequest(fdSet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	agency_schema := new(schema.Agency)
	agency_schema.Id = fdSet.AcountId
	zf, has, err := agencyBean.GetOneAgencyByid(agency_schema)
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
	if user_level == 1 && zf.Level >= 1 {
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneDiscountSet(fdSet)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
	if user_level == 2 && zf.Level >= 2 {
		fdSet.AcountId = user_id
		data, ok, err := agencyMemberRegisterDiscountSetBean.GetOneDiscountSet(fdSet)
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
func (fac *FirstAgencyController) MemberRegDiscountSetEdit(ctx echo.Context) error {
	fdSetUpdata := new(input.FirstDiscountUpdata)
	code := global.ValidRequest(fdSetUpdata, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := agencyMemberRegisterDiscountSetBean.UpdataSet(fdSetUpdata)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50172, ctx))
	}
	return ctx.NoContent(204)
}

//获取开户人下所有的前台站点id
func (fac *FirstAgencyController) SiteIdByAgencyId(ctx echo.Context) error {
	user := ctx.Get("user").(*global.RedisStruct)
	userLever := user.Level
	newSite := new(schema.Site)
	newSite.Id = user.SiteId
	if userLever == 1 {
		data, err := siteOperateBean.GetSiteIndexId(newSite)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		var list []back.IndexBackStruct
		var li back.IndexBackStruct
		if user.IsSub == 1 {
			//查询该子帐号的资料细项
			info, has, err := permissionBean.GetDetailsMemberByChild(user.Id)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			//不存在或者是空值，都默认为所有
			if !has {
				list = data
			}
			if len(info.ChildSite) > 0 {
				siteArr := strings.Split(info.ChildSite, ",")
				for _, v := range siteArr {
					for _, n := range data {
						if v == n.SiteIndexId {
							li.IsDefault = n.IsDefault
							li.SiteIndexId = n.SiteIndexId
							li.SiteId = n.SiteId
							li.SiteName = n.SiteName
							list = append(list, li)
						}
					}
				}
			} else {
				list = data
			}
		} else {
			list = data
		}
		return ctx.JSON(200, global.ReplyItem(list))
	}
	return ctx.JSON(200, global.ReplyItem(nil))
}
