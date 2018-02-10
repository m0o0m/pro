package conversion

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//额度转换
type WapBalanceConversionController struct {
	controllers.BaseController
}

//额度转换-获取各平台余额&&一键刷新
func (*WapBalanceConversionController) WapGetPlatformBalance(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	//查询商品表信息
	data, err := memberBalanceConversion.GetPlatformBalance(member.Id)
	//data, err := memberBalanceConversion.GetPlatformBalance(member.Id, products)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//额度转换-单个平台余额刷新
func (*WapBalanceConversionController) PlatformBalanceRefresh(ctx echo.Context) error {
	platformBalanceRefresh := new(input.PlatformBalanceRefresh)
	code := global.ValidRequestMember(platformBalanceRefresh, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	data, err := memberBalanceConversion.PlatformBalanceRefresh(member.Id, platformBalanceRefresh.PlatformId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//额度转换-余额转换
func (bpc *WapBalanceConversionController) WapBalanceConversion(ctx echo.Context) error {
	wapMemberBalanceConversion := new(input.WapMemberBalanceConversion)
	code := global.ValidRequestMember(wapMemberBalanceConversion, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if wapMemberBalanceConversion.Money < 10 {
		return ctx.JSON(200, global.ReplyError(30183, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	wapMemberBalanceConversion.MemberId = member.Id
	wapMemberBalanceConversion.SiteId = member.Site
	wapMemberBalanceConversion.SiteIndexId = member.SiteIndex
	//判断转入项目是否存在
	if wapMemberBalanceConversion.FromType != 0 {
		has, err := memberBalanceConversion.IsExistFtype(wapMemberBalanceConversion.FromType)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30147, ctx))
		}
	}
	//判断转出项目是否存在
	if wapMemberBalanceConversion.ForType != 0 {
		has, err := memberBalanceConversion.IsExistFtype(wapMemberBalanceConversion.ForType)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30154, ctx))
		}
	}
	//转入项目或者转出项目必须有一个是系统余额
	if wapMemberBalanceConversion.FromType != 0 && wapMemberBalanceConversion.ForType != 0 {
		return ctx.JSON(200, global.ReplyError(30157, ctx))
	}
	//获取站点下套餐id
	site, err := memberBalanceConversion.GetSiteCombo(wapMemberBalanceConversion.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取开户人视讯余额
	videoBalance, err := memberBalanceConversion.GetAgency(wapMemberBalanceConversion.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if wapMemberBalanceConversion.ForType == 0 { //转出项目为系统余额
		balance, err := memberBalanceConversion.WapGetBalance(wapMemberBalanceConversion.MemberId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//系统余额是否大于转出金额
		if balance < wapMemberBalanceConversion.Money {
			return ctx.JSON(200, global.ReplyError(30146, ctx))
		}
		//根据转入项目和套餐id获取手续费占成比
		proportion, err := memberBalanceConversion.GetProductProportion(wapMemberBalanceConversion.FromType, site.ComboId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(70029, ctx))
		}
		if len(proportion) == 0 {
			//手续费必须要有
			return ctx.JSON(500, global.ReplyError(70029, ctx))
		}
		//将占比排序（从大到小）
		for i := 0; i < len(proportion); i++ {
			for j := i + 1; j < len(proportion); j++ {
				if proportion[i] < proportion[j] {
					proportion[i], proportion[j] = proportion[j], proportion[i]
				}
			}
		}
		//手续费
		wapMemberBalanceConversion.Fee = proportion[0] * 0.01 * wapMemberBalanceConversion.Money
		//判断视讯余额是否大于扣除的手续费
		if videoBalance.VideoBalance < wapMemberBalanceConversion.Fee {
			return ctx.JSON(200, global.ReplyError(30158, ctx))
		}
	} else { //其他金额转换到系统金额
		/*//根据会员账号获取会员id
		memberId, err := member_balance_conversion.GetMemberId(memberBalanceConversion.Account)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}*/
		//获取转出项目的余额
		info, _, err := memberBalanceConversion.GetMoneyByVideo(wapMemberBalanceConversion.MemberId, wapMemberBalanceConversion.ForType)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//根据转入项目和套餐id获取手续费占成比
		proportion, err := memberBalanceConversion.GetProductProportion(wapMemberBalanceConversion.ForType, site.ComboId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(70029, ctx))
		}
		if len(proportion) == 0 {
			//手续费必须要有
			return ctx.JSON(500, global.ReplyError(70029, ctx))
		}
		//将占比排序（从大到小）
		for i := 0; i < len(proportion); i++ {
			for j := i + 1; j < len(proportion); j++ {
				if proportion[i] < proportion[j] {
					proportion[i], proportion[j] = proportion[j], proportion[i]
				}
			}
		}
		//手续费
		wapMemberBalanceConversion.Fee = proportion[0] * 0.01 * wapMemberBalanceConversion.Money
		//判断转出项目的余额是否大于转出金额
		if info.Balance < wapMemberBalanceConversion.Money {
			return ctx.JSON(200, global.ReplyError(30146, ctx))
		}
	}
	count, err := memberBalanceConversion.WapBalanceConversionDo(wapMemberBalanceConversion)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30148, ctx))
	}
	return ctx.NoContent(204)
}

//会员中心--会员余额刷新
func (*WapBalanceConversionController) WapBalance(ctx echo.Context) error {
	//获取会员登录信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	data, err := memberBalanceConversion.MemberBalance(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
