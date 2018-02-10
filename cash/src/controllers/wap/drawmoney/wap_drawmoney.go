package drawmoney

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//取款管理
type WapDrawMoneyController struct {
	controllers.BaseController
}

//获取会员账号，余额，出款银行卡信息
func (*WapDrawMoneyController) GetMemberInfo(ctx echo.Context) error {
	//获取登录会员信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	data, err := drawMoney.GetMemberInfo(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//取款(没有给稽核日志表添加数据)
func (*WapDrawMoneyController) Withdrawal(ctx echo.Context) error {
	wapDrawMoney := new(input.WapDrawMoney)
	code := global.ValidRequestMember(wapDrawMoney, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登录会员信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//查看出款管理表是否有待审核的数据
	has, err := drawMoney.IsExistPendingReview(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		//存在待审核数据就不允许再次出款
		return ctx.JSON(200, global.ReplyError(30236, ctx))
	}
	//获取会员余额和取款密码
	balance, password, has, err := drawMoney.GetMemberBalanceAndPassword(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30232, ctx))
	}
	//取款金额是否大于会员余额
	if wapDrawMoney.Money > balance {
		return ctx.JSON(200, global.ReplyError(30233, ctx))
	}
	//密码加密
	drawPassword, err := global.MD5ByStr(wapDrawMoney.DrawPassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//取款密码是否正确
	if password != drawPassword {
		return ctx.JSON(200, global.ReplyError(30234, ctx))
	}
	//会员出款银行id是否存在
	has, err = drawMoney.IsExistMemberBank(member.Id, wapDrawMoney.BankId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50116, ctx))
	}
	//会员操作后余额
	memberBalance := balance - wapDrawMoney.Money
	//取款
	count, err := drawMoney.Withdrawal(member, memberBalance, wapDrawMoney.Money)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30235, ctx))
	}
	return ctx.NoContent(204)
}

//取款进度
func (*WapDrawMoneyController) DrawalProgress(ctx echo.Context) error {
	//获取登录会员信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	data, err := drawMoney.DrawalProgress(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//如果实际出款金额小于或等于0，则提示无法出款
	if data.OutwardMoney <= 0 {
		return ctx.JSON(200, global.ReplyError(30237, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
