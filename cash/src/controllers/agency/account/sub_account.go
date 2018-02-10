package account

import (
	"controllers"
	"framework/logger"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/function"
	"models/input"
	"strings"
)

//子账号管理
type SubAccountController struct {
	controllers.BaseController
}

//子账号列表
func (sac *SubAccountController) Index(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	sac.GetParam(listparam, ctx)
	subAccount := new(input.SubAccountList)
	code := global.ValidRequest(subAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//得到站点id和角色id
	subAccount.SiteId = user.SiteId
	subAccount.SiteIndexId = user.SiteIndexId
	subAccount.RoleId = user.RoleId
	subAccount.ParentId = user.Id
	//只有开户人和代理能查看子帐号
	if (user.Level != 1 && user.Level != 4) || user.IsSub == 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	list, count, err := subAccountBeen.GetList(subAccount, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(30036, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//获取子账号权限
func (*SubAccountController) Permission(ctx echo.Context) error {
	subAccount := new(input.SubAccountPermission)
	code := global.ValidRequest(subAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	subAccount.ParentId = user.Id
	data, err := subAccountBeen.GetPermission(subAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var subAgencyPermisionBack back.SubAgencyPermissionBack
	if len(data) == 0 {
		//根据id查询子账号
		info, has, err := subAccountBeen.GetInfos(subAccount.Id, subAccount.SiteId, subAccount.SiteIndexId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30159, ctx))
		} else {
			subAgencyPermisionBack.Account = info.Account
		}
	}
	promiss := new(input.PromissRoleId)
	promiss.Type = subAccount.Type
	//获取module
	module, err := permissionBean.GetModules(promiss)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取权限功能
	var p back.Permissions
	//权限返回列表
	var permissions []back.Permissions
	var permission back.Permission
	var mdl []string
	for m := range module {
		mdl = append(mdl, module[m].Module)
	}
	list, err := permissionBean.GetListByModule(mdl)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k := range module {
		p.Module = module[k].Module
		var ps []back.Permission
		for i := range list {
			if list[i].Module == module[k].Module {
				permission.Id = list[i].Id
				permission.CreateTime = list[i].CreateTime
				permission.PermissionName = list[i].PermissionName
				permission.Route = list[i].Route
				permission.Module = list[i].Module
				permission.Method = list[i].Method
				permission.Status = list[i].Status
				permission.Type = list[i].Type
				permission.IsPermission = list[i].IsPermission
				ps = append(ps, permission)
			}
		}
		p.Permission = ps
		permissions = append(permissions, p)
	}
	//角色权限返回列表
	for k := range data {
		subAgencyPermisionBack.Account = data[k].Account
		for i := range permissions {
			for j := range permissions[i].Permission {
				if data[k].PermissionName == permissions[i].Permission[j].PermissionName && data[k].Module == permissions[i].Permission[j].Module {
					permissions[i].Permission[j].IsPermission = 1
				}
			}
		}
	}
	subAgencyPermisionBack.Permissions = permissions
	//查询该子帐号的资料细项
	info, _, err := permissionBean.GetDetailsMemberByChild(subAccount.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var infoList = make(map[string]interface{})
	infoList["subAgencyPermisionBack"] = subAgencyPermisionBack

	infoPower := strings.Split(info.ChildPower, ",")
	infoPowers := []string{"A1", "A2", "B1", "B2", "C1", "C2", "D1", "D2", "E1", "E2", "F1", "F2", "G1", "G2"}
	type powerStatus struct {
		Id     string `json:"id"`
		Status int8   `json:"status"`
	}
	//会员详细资料细项
	var ps []powerStatus
	var pps powerStatus
	for _, m := range infoPowers {
		i := 0
		if len(infoPower) > 0 {
			for _, x := range infoPower {
				if m == x {
					i = i + 1
				}
			}
			if i < 1 {
				pps.Id = m
				pps.Status = 2
				ps = append(ps, pps)
			} else {
				pps.Id = m
				pps.Status = 1
				ps = append(ps, pps)
			}
		} else {
			pps.Id = m
			pps.Status = 2
			ps = append(ps, pps)
		}
	}
	infoList["infoPower"] = ps
	//站点
	siteInfo, err := permissionBean.GetSiteIndexBySite(user.SiteId, user.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	childSite := strings.Split(info.ChildSite, ",")
	var cH back.SiteIndexBySiteBack
	var cHs []back.SiteIndexBySiteBack
	for _, v := range siteInfo {
		if len(childSite) > 0 {
			u := 0
			for _, b := range childSite {
				if b == v.IndexId {
					u = u + 1
				}
			}
			if u < 1 {
				cH.Id = v.Id
				cH.IndexId = v.IndexId
				cH.Status = 2
				cH.SiteName = v.SiteName
				cHs = append(cHs, cH)
			} else {
				cH.Id = v.Id
				cH.IndexId = v.IndexId
				cH.Status = 1
				cH.SiteName = v.SiteName
				cHs = append(cHs, cH)
			}
		} else {
			cH.Id = v.Id
			cH.IndexId = v.IndexId
			cH.Status = 2
			cH.SiteName = v.SiteName
			cHs = append(cHs, cH)
		}
	}
	infoList["childSite"] = cHs
	return ctx.JSON(200, global.ReplyItem(infoList))
}

//配置子账号权限
func (sac *SubAccountController) PermissionEdit(ctx echo.Context) error {
	subAccountPermission := new(input.PermissionEdit)
	code := global.ValidRequest(subAccountPermission, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	subAccountPermission.SiteId = user.SiteId
	//查看子账号id是否存在
	data, has, err := subAccountBeen.SubAccountIsExist(subAccountPermission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30121, ctx))
	}
	//判断子帐号的power是否合法
	if data.RoleId == 4 {
		//代理子帐号没有修改功能
		power := strings.Split(subAccountPermission.ChildPower, ",")
		for _, v := range power {
			if v == "A2" || v == "B2" || v == "C2" || v == "D2" || v == "E2" || v == "F2" || v == "G2" {
				return ctx.JSON(200, global.ReplyError(50160, ctx))
			}
		}
		//判断勾选前台站点是否跟子帐号的前台id相同
		subAccountPermission.ChildSite = data.SiteIndexId
	}
	//查看权限id是否存在
	if len(subAccountPermission.PermissionId) > 0 {
		has, err = function.IsPermissionById(subAccountPermission.PermissionId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30110, ctx))
		}
	}
	count, err := subAccountBeen.UpdatePermission(subAccountPermission)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30122, ctx))
	}
	return ctx.NoContent(204)
}

//启用/禁用
func (sac *SubAccountController) Status(ctx echo.Context) error {
	subAccount := new(input.SubAccountId)
	code := global.ValidRequest(subAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	agency := ctx.Get("user").(*global.RedisStruct)
	//只有开户人和代理能修改子帐号状态
	if (agency.Level != 1 && agency.Level != 4) || agency.IsSub != 2 {
		return ctx.JSON(200, global.ReplyError(50166, ctx))
	}
	count, err := subAccountBeen.UpdateStatus(subAccount.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30037, ctx))
	}
	return ctx.NoContent(204)
}

//获取子账号基本资料
func (sac *SubAccountController) BaseInfo(ctx echo.Context) error {
	subAccount := new(input.SubAccountId)
	code := global.ValidRequest(subAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//只有开户人和代理能查看子帐号
	if (user.Level != 1 && user.Level != 4) || user.IsSub != 2 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	info, has, err := subAccountBeen.GetInfo(subAccount.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !has {
		global.GlobalLogger.Error(logger.ERROR, err)
		return ctx.JSON(200, global.ReplyError(30038, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//修改子账号基本资料
func (sac *SubAccountController) BaseInfoEdit(ctx echo.Context) error {
	subAccount := new(input.SubAccountEdit)
	code := global.ValidRequest(subAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	users := ctx.Get("user").(*global.RedisStruct)
	//只有开户人能修改子帐号
	if (users.Level != 1 && users.Level != 4) || users.IsSub != 2 {
		return ctx.JSON(200, global.ReplyError(50166, ctx))
	}
	//两次密码不一致
	if subAccount.Password != subAccount.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	if subAccount.Password != "" {
		md5Password, err := global.MD5ByStr(subAccount.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		subAccount.Password = md5Password
	}
	if subAccount.OperatePassword != "" {
		md5OperatePassword, err := global.MD5ByStr(subAccount.OperatePassword, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30045, ctx))
		}
		subAccount.OperatePassword = md5OperatePassword
	}
	count, err := subAccountBeen.UpdateInfo(subAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	//修改成功
	return ctx.NoContent(204)
}

//删除子账号
func (sac *SubAccountController) Delete(ctx echo.Context) error {
	subAccount := new(input.SubAccountId)
	code := global.ValidRequest(subAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	agency := ctx.Get("user").(*global.RedisStruct)
	//只有开户人和代理能删除子帐号状态
	if (agency.Level != 1 && agency.Level != 4) || agency.IsSub != 2 {
		return ctx.JSON(200, global.ReplyError(50167, ctx))
	}
	count, err := subAccountBeen.Delete(subAccount.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30043, ctx))
	}
	return ctx.NoContent(204)
}

//添加子账号
func (sac *SubAccountController) Add(ctx echo.Context) error {
	subAccount := new(input.SubAccountAdd)
	code := global.ValidRequest(subAccount, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	agency := ctx.Get("user").(*global.RedisStruct)
	//只有开户人和代理能添加子帐号状态
	if (agency.Level != 1 && agency.Level != 4) || agency.IsSub != 2 {
		return ctx.JSON(200, global.ReplyError(50167, ctx))
	}
	//查看账号是否存在
	has, err := subAccountBeen.GetAccount(subAccount.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(30039, ctx))
	}
	//两次密码不一致
	if subAccount.Password != subAccount.ConfirmPassword {
		return ctx.JSON(200, global.ReplyError(30046, ctx))
	}
	//MD5加密
	md5Password, err := global.MD5ByStr(subAccount.Password, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(30044, ctx))
	}
	md5OperatePassword, err := global.MD5ByStr(subAccount.OperatePassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(30045, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	//得到站点id和角色id
	subAccount.SiteId = user.SiteId
	subAccount.RoleId = user.RoleId
	subAccount.Level = user.Level
	subAccount.ParentId = user.Id
	subAccount.Password = md5Password
	subAccount.OperatePassword = md5OperatePassword
	subAccount.SiteIndexId = user.SiteIndexId
	count, err := subAccountBeen.Add(subAccount)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30040, ctx))
	}
	return ctx.NoContent(204)
}

//查看子账号口令验证信息
func (*SubAccountController) AccessTokenInfo(ctx echo.Context) error {
	user := ctx.Get("user").(*global.RedisStruct)
	_, err, Level := agencyThirdInfoBean.ThirdAgencyInfoLevel(user.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if Level != 1 {
		//判断登陆账号是否有设置口令验证的权限
		return ctx.JSON(500, global.ReplyError(30110, ctx))
	}

	//获取口令验证数据
	info, has, err := subAccountBeen.SubAccessTokenInfo(user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !has || info.PassKey == "" {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	//口令密钥解密
	info.PassKey, err = global.ParseBase64Str(info.PassKey, user.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(70011, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//修改子账号口令验证信息
func (*SubAccountController) AccessToken(ctx echo.Context) error {
	subAccountToken := new(input.SubAccountToken)
	code := global.ValidRequest(subAccountToken, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if subAccountToken.PassKey != "" && len(subAccountToken.PassKey) != 16 {
		return ctx.JSON(200, global.ReplyError(20024, ctx))
	}

	//获取口令验证数据
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level != 1 || user.RoleId != 1 {
		//判断登陆账号是否有设置口令验证的权限
		return ctx.JSON(500, global.ReplyError(30110, ctx))
	}
	//密钥加密
	subAccountToken.PassKey = global.EncryptBase64Str(subAccountToken.PassKey, user.SiteId)
	count, err := subAccountBeen.SubAccessToken(subAccountToken, user.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30043, ctx))
	}
	return ctx.NoContent(204)
}

//随机生成口令密钥
func (*SubAccountController) GenKey(ctx echo.Context) error {
	reqData := new(input.GenKey)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	key, err := agencySignBean.CreateSecret(reqData.Len)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(key))
}
