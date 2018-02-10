package account

import (
	"controllers"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strings"
	"time"
)

//会员管理
type MemberController struct {
	controllers.BaseController
}

//会员列表
func (mc *MemberController) Index(ctx echo.Context) error {
	member := new(input.MemberIndex)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	agency := ctx.Get("user").(*global.RedisStruct)
	member.SiteId = agency.SiteId
	//如果当前登录账号是代理,那么无论前端传过来的代理账号是多少,会员列表的代理账号筛选条件只能是当前登录账号
	if agency.Level == 4 {
		member.AgencyId = agency.Id
	}
	//查看排序的传值（会员视讯余额）是否存在
	if member.SortBy != "" && member.SortBy != "1" && member.SortBy != "2" && member.SortBy != "3" && member.SortBy != "4" {
		msd, err := memberBean.MemberSortDrop(member.SiteId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		var p int8
		for k := range msd {
			if msd[k].Platform == member.SortBy {
				p += 1
			}
		}
		if p < 1 {
			return ctx.JSON(200, global.ReplyError(30247, ctx))
		}
	}
	listParams := new(global.ListParams)
	mc.GetParam(listParams, ctx)
	data, count, err := memberBean.List(member, listParams)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	sTime, eTime := global.GetToday()
	//今日注册人数
	info, err := memberBean.MemberNumberBySite(agency.SiteId, member.SiteIndexId, sTime, eTime)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	info.TotalNum = count
	var list = make(map[string]interface{})
	list["data"] = data
	list["num"] = info
	return ctx.JSON(200, global.ReplyPagination(listParams, list, int64(len(list)), count, ctx))
}

//启用/禁用
func (mc *MemberController) Status(ctx echo.Context) error {
	memberStatus := new(input.MemberStatus)
	code := global.ValidRequest(memberStatus, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断当前会员是否存在
	member, have, err := memberBean.GetInfoById(memberStatus.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10228, ctx))
	}
	//根据会员各个设备的login_key去做删除redis缓存操作.
	//pc端
	if member.PcLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.PcLoginKey).Err()
		if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//wap端
	if member.WapLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.WapLoginKey).Err()
		if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//ios端
	if member.IosLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.IosLoginKey).Err()
		if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//android端
	if member.AndroidLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.AndroidLoginKey).Err()
		if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	ok, err := memberBean.Status(memberStatus)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(10019, ctx))
	}
	return ctx.NoContent(204)
}

//踢线
func (mc *MemberController) Offline(ctx echo.Context) error {
	//获取下线会员的Id主键
	offlineMember := new(input.MemberStatus)
	code := global.ValidRequest(offlineMember, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断当前会员是否存在
	member, have, err := memberBean.GetInfoById(offlineMember.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10228, ctx))
	}
	//根据会员各个设备的login_key去做删除redis缓存操作.
	//pc端
	if member.PcLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.PcLoginKey).Err()
		if err == redis.Nil {
		} else if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//wap端
	if member.WapLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.WapLoginKey).Err()
		if err == redis.Nil {
		} else if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//ios端
	if member.IosLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.IosLoginKey).Err()
		if err == redis.Nil {
		} else if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//android端
	if member.AndroidLoginKey != "" {
		//删除redis
		err := global.GetRedis().Del(member.AndroidLoginKey).Err()
		if err == redis.Nil {
		} else if err != nil {
			global.GlobalLogger.Error("Error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//踢出会员
	count, err := memberBean.OffLine(offlineMember)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(50193, ctx))
	}
	return ctx.NoContent(204)
}

//获取基本资料
func (mc *MemberController) BaseInfo(ctx echo.Context) error {
	member := new(input.MemberStatus)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	ok, err, back := memberBean.Info(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(60009, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(back))
}

//修改基本资料
func (mc *MemberController) BaseInfoEdit(ctx echo.Context) error {
	member := new(input.MemberBaseInfo)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	//只能开户人和代理操作
	user := ctx.Get("user").(*global.RedisStruct)
	if (user.Level != 1 && user.Level != 4) || (user.RoleId != 1 && user.RoleId != 4) {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	if member.NewPassword != member.ReplyPassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	if member.NewPassword != "" {
		pwd, err := global.MD5ByStr(member.NewPassword, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		member.NewPassword = pwd
	}

	ok, err := memberBean.UpdateInfo(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(60086, ctx))
	}

	return ctx.NoContent(204)
}

//获取详细资料
func (mc *MemberController) DetailInfo(ctx echo.Context) error {
	member := new(input.MemberStatus)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	member.SiteId = user.SiteId
	ok, err, back := memberBean.Detail(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(30246, ctx))
	}
	//如果登陆人是子帐号
	if user.IsSub == 1 && back != nil {
		//查询该子帐号的资料细项
		info, _, err := permissionBean.GetDetailsMemberByChild(user.Id)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if user.Level == 4 {
			//判断该子帐号是否有操作会员所在站点的权限
			if len(info.ChildSite) > 0 {
				childSite := strings.Split(info.ChildSite, ",")
				var j int
				for _, k := range childSite {
					if user.SiteIndexId == k {
						j = j + 1
					}
				}
				if j < 1 {
					return ctx.JSON(200, global.ReplyError(50163, ctx))
				}
			} else {
				return ctx.JSON(200, global.ReplyError(50163, ctx))
			}
		}
		if len(info.ChildPower) > 0 {
			childPower := strings.Split(info.ChildPower, ",")
			var a, b, c, d, e, f, g int
			for _, v := range childPower {
				switch v {
				case "A1":
					a = a + 1
				case "B1":
					b = b + 1
				case "C1":
					c = c + 1
				case "D1":
					d = d + 1
				case "E1":
					e = e + 1
				case "F1":
					f = f + 1
				case "G1":
					g = g + 1
				}
			}
			if a < 1 {
				back.Realname = "*"
			}
			if b < 1 && len(back.MemberBank) > 0 {
				back.MemberBank = nil
			}
			if c < 1 {
				back.DrawPassword = "*"
			}
			if d < 1 {
				back.Mobile = "*"
			}
			if e < 1 {
				back.Email = "*"
			}
			if f < 1 {
				back.QQ = "*"
			}
			if g < 1 {
				back.Card = "*"
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(back))
}

//获取会员出款银行列表
func (mc *MemberController) Bank(ctx echo.Context) error {
	member := new(input.MemberStatus)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, count, err := memberBean.Bank(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(data, count))
}

//获取会员出款银行详情
func (mc *MemberController) BankInfo(ctx echo.Context) error {
	member := new(input.MemberBankInfo)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := memberBean.BankInfo(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改详细资料
func (mc *MemberController) DetailInfoEdit(ctx echo.Context) error {
	member := new(input.MemberDetail)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if member.Mobile != "" && member.Mobile != "*" {
		//检验手机号
		has := global.CheckPhoneNumber(member.Mobile)
		if !has {
			return ctx.JSON(200, global.ReplyError(20014, ctx))
		}
	}
	if member.Email != "" && member.Email != "*" {
		//检验邮箱
		has := global.CheckEmail(member.Email)
		if !has {
			return ctx.JSON(200, global.ReplyError(20015, ctx))
		}
	}
	if member.Card != "" && member.Card != "*" {
		//检验省份证
		has := global.CheckIdentity(member.Card)
		if !has {
			return ctx.JSON(200, global.ReplyError(20013, ctx))
		}
	}
	//只能开户人和代理操作/代理子帐号没有修改权限
	user := ctx.Get("user").(*global.RedisStruct)
	if (user.Level != 1 && user.Level != 4) || (user.RoleId != 1 && user.RoleId != 4) || (user.RoleId == 4 && user.IsSub == 1) {
		return ctx.JSON(200, global.ReplyError(50162, ctx))
	}
	member.SiteId = user.SiteId
	if user.RoleId == 1 && user.IsSub == 1 {
		//查询该子帐号的资料细项
		info, _, err := permissionBean.GetDetailsMemberByChild(user.Id)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//判断该子帐号是否有操作会员所在站点的权限
		if len(info.ChildSite) > 0 {
			childSite := strings.Split(info.ChildSite, ",")
			var j int
			for _, k := range childSite {
				if member.SiteIndexId == k {
					j = j + 1
				}
			}
			if j < 1 {
				return ctx.JSON(200, global.ReplyError(50162, ctx))
			}
		}
		if len(info.ChildPower) > 0 {
			childPower := strings.Split(info.ChildPower, ",")
			var a, b, c, d, e, f, g int
			for _, v := range childPower {
				switch v {
				case "A2":
					a = a + 1
				case "B2":
					b = b + 1
				case "C2":
					c = c + 1
				case "D2":
					d = d + 1
				case "E2":
					e = e + 1
				case "F2":
					f = f + 1
				case "G2":
					g = g + 1
				}
			}
			if a < 1 {
				member.Realname = ""
			}
			if b < 1 {
				member.Ids = ""
			}
			if c < 1 {
				member.DrawPassword = ""
			}
			if d < 1 {
				member.Mobile = ""
			}
			if e < 1 {
				member.Email = ""
			}
			if f < 1 {
				member.QQ = ""
			}
			if g < 1 {
				member.Card = ""
			}
		}
	}
	if member.Birthday != "" {
		loc, _ := time.LoadLocation("Local")
		_, err := time.ParseInLocation("2006-01-02", member.Birthday, loc)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(20012, ctx))
		}
	}
	member.SiteIndexId = user.SiteIndexId
	ok, err := memberBean.UpdateDetail(member, user.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(60086, ctx))
	}
	return ctx.NoContent(204)
}

//修改会员银行卡
func (mc *MemberController) MemberBankEdit(ctx echo.Context) error {
	member := new(input.MemberBankEdit)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	//只能开户人和代理操作
	user := ctx.Get("user").(*global.RedisStruct)
	if (user.Level != 1 && user.Level != 4) || (user.RoleId != 1 && user.RoleId != 4) {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	//判断卡号是否存在
	ok, err := memberBean.GetCard(member.Card, member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if ok {
		return ctx.JSON(200, global.ReplyError(30069, ctx))
	}
	count, err := memberBean.UpdateMemberBank(member)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//删除会员银行卡
func (mc *MemberController) MemberBankDel(ctx echo.Context) error {
	member := new(input.MemberStatus)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	//只能开户人和代理操作
	user := ctx.Get("user").(*global.RedisStruct)
	if (user.Level != 1 && user.Level != 4) || (user.RoleId != 1 && user.RoleId != 4) {
		return ctx.JSON(200, global.ReplyError(60001, ctx))
	}

	count, err := memberBean.DelMemberBank(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(30055, ctx))
	}
	return ctx.NoContent(204)
}

//会员排序下拉
func (mc *MemberController) MemberSortDrop(ctx echo.Context) error {
	member := new(input.MemberSortDrop)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := memberBean.MemberSortDrop(member.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//批量启用禁用会员
func (mc *MemberController) BatchStatus(ctx echo.Context) error {
	members := new(input.BatchMember)
	code := global.ValidRequest(members, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断批量会员数
	if len(members.Ids) < 1 {
		return ctx.JSON(200, global.ReplyError(10230, ctx))
	}
	ids := make([]int64, 0)
	//判断当前会员是否存在
	for _, id := range members.Ids {
		member, have, err := memberBean.GetInfoById(id)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !have {
			return ctx.JSON(200, global.ReplyError(10228, ctx))
		}
		if members.Status == 2 && member.Status == 2 {
			//return ctx.JSON(200, global.ReplyError(10233, ctx))
			continue
		}
		if members.Status == 1 && member.Status == 1 {
			//return ctx.JSON(200, global.ReplyError(10234, ctx))
			continue
		}
		//根据会员各个设备的login_key去做删除redis缓存操作.
		//pc端
		if member.PcLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.PcLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		//wap端
		if member.WapLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.WapLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		//ios端
		if member.IosLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.IosLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		//android端
		if member.AndroidLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.AndroidLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		ids = append(ids, id)
	}
	members.Ids = ids
	rows, err := memberBean.BatchStatus(members)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if rows != int64(len(members.Ids)) {
		return ctx.JSON(200, global.ReplyError(10231, ctx))
	}
	return ctx.NoContent(204)
}

//批量踢线会员
func (mc *MemberController) BatchOffline(ctx echo.Context) error {
	//获取批量会员的Id主键
	offlineMembers := new(input.OfflineMember)

	code := global.ValidRequest(offlineMembers, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断批量会员数
	if len(offlineMembers.Ids) < 1 {
		return ctx.JSON(200, global.ReplyError(10230, ctx))
	}
	//判断当前会员是否存在
	ids := make([]int64, 0)
	for _, id := range offlineMembers.Ids {
		member, have, err := memberBean.GetInfoById(id)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !have {
			return ctx.JSON(200, global.ReplyError(10228, ctx))
		}
		if member.IosLoginKey == "" && member.WapLoginKey == "" && member.AndroidLoginKey == "" && member.PcLoginKey == "" {
			//return ctx.JSON(200, global.ReplyError(10235, ctx))
			continue
		}
		//根据会员各个设备的login_key去做删除redis缓存操作.
		//pc端
		if member.PcLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.PcLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		//wap端
		if member.WapLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.WapLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		//ios端
		if member.IosLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.IosLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		//android端
		if member.AndroidLoginKey != "" {
			//删除redis
			err := global.GetRedis().Del(member.AndroidLoginKey).Err()
			if err != nil {
				global.GlobalLogger.Error("Error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		ids = append(ids, id)
	}
	offlineMembers.Ids = ids
	rows, err := memberBean.BatchOffline(offlineMembers)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if rows != int64(len(offlineMembers.Ids)) {
		return ctx.JSON(200, global.ReplyError(10231, ctx))
	}
	return ctx.NoContent(204)
}

//会员银行下拉框
func (mc *MemberController) MemberBankDrop(ctx echo.Context) error {
	mD := new(input.MemberBankDropIn)
	code := global.ValidRequest(mD, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	//根据登录人id获取index_id
	info, has, err := bankCardBean.SiteIndexIdByMemberId(mD.Id, user.SiteId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	//获取所有出款银行
	data, err := bankCardBean.BankCardListDrop()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	bI := new(input.AgencyBankOutByDrop)
	bI.SiteId = user.SiteId
	bI.SiteIndexId = info.SiteIndexId
	//获取所有被剔除的银行
	dataDel, err := bankCardBean.BankCardListDropDel(bI)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var list []back.BankCardListDrop
	var li back.BankCardListDrop
	if len(data) > 0 {
		for _, v := range data {
			i := 0
			if len(dataDel) > 0 {
				for _, n := range dataDel {
					if v.Id == n.BankId {
						i = i + 1
					}
				}
				if i < 1 {
					li.Id = v.Id
					li.Title = v.Title
					list = append(list, li)
				}
			} else {
				list = data
			}

		}
	} else {
		list = data
	}
	return ctx.JSON(200, global.ReplyItem(list))
}
