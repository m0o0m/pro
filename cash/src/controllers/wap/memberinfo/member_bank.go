package memberinfo

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//会员银行卡
type MemberBankController struct {
	controllers.BaseController
}

//会员银行卡
func (c *MemberBankController) MemberBankList(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//会员银行卡获取
	data, err := mBBean.MemberBankListById(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//添加银行卡
func (c *MemberBankController) MemberBankAdd(ctx echo.Context) error {
	//获取参数
	mb := new(input.MemberBankAddIn)
	code := global.ValidRequestMember(mb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//银行卡校验
	flag := global.CheckCardNumber(mb.CardNumber)
	if !flag {
		return ctx.JSON(200, global.ReplyError(30060, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	mb.MemberId = member.Id
	//先根据登陆人id取出所有的银行卡信息
	data, err := mBBean.MemberBankListById(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//判断银行卡数是否足够
	if len(data) >= 3 {
		return ctx.JSON(200, global.ReplyError(50117, ctx))
	}
	//再判断该会员是否已添加过该银行卡号
	for _, v := range data {
		if v.Card == mb.CardNumber {
			return ctx.JSON(200, global.ReplyError(30069, ctx))
		}
	}
	//添加银行卡
	count, err := mBBean.MemberAddBankCard(mb)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//银行卡解绑
func (c *MemberBankController) MemberBankUnBind(ctx echo.Context) error {
	//获取参数
	mb := new(input.MemberBankUnBindIn)
	code := global.ValidRequestMember(mb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//查询该银行卡id是否存在
	_, has, err := mBBean.BeMemberBankById(mb.Id, member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50150, ctx))
	}
	//解绑
	count, err := mBBean.MemberBankUnBundling(mb.Id, member.Id, 1)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30292, ctx))
	}
	return ctx.NoContent(204)
}

//银行卡绑定
func (c *MemberBankController) MemberBankBind(ctx echo.Context) error {
	//获取参数
	mb := new(input.MemberBankUnBindIn)
	code := global.ValidRequestMember(mb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//查询该银行卡id是否存在,以及是否已经存在绑定的银行卡
	data, err := mBBean.MemberBankListById(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//遍历会员的所有银行卡
	var i = 0
	for _, v := range data {
		//已有默认银行卡
		if v.IsDefaultBank == 1 {
			return ctx.JSON(200, global.ReplyError(50151, ctx))
		}
		//判断前端上传id是否在所获取的id中
		if mb.Id == v.Id {
			i = i + 1
		}
	}
	//i不等于1说明id不存在
	if i != 1 {
		return ctx.JSON(200, global.ReplyError(50150, ctx))
	}
	//绑定
	count, err := mBBean.MemberBankUnBundling(mb.Id, member.Id, 2)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30293, ctx))
	}
	return ctx.NoContent(204)
}

//银行卡删除
func (c *MemberBankController) MemberBankDelete(ctx echo.Context) error {
	//获取参数
	mb := new(input.MemberBankDeleteIn)
	code := global.ValidRequestMember(mb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//查询该银行卡id是否存在
	_, has, err := mBBean.BeMemberBankById(mb.Id, member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50150, ctx))
	}
	//删除
	mb.MemberId = member.Id
	count, err := mBBean.MemberBankDelete(mb)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//银行卡下拉框
func (c *MemberBankController) MemberBankDrop(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	data, err := mBBean.BankDropBySiteAndSiteIndexId(member.Site, member.SiteIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//银行卡详情
func (c *MemberBankController) MemberBankCardOneInfo(ctx echo.Context) error {
	//获取参数
	mb := new(input.MemberBankCardDetailsIn)
	code := global.ValidRequestMember(mb, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	mb.MemberId = member.Id
	//查询
	info, has, err := mBBean.MemberBankCardDetails(mb)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50150, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}
