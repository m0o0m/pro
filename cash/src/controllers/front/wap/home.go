package wap

import (
	"controllers/front/wap/data_merge"
	"framework/render"
	"github.com/labstack/echo"
	"global"
)

type HomeController struct {
	WapBaseController
}

//index选项卡
func (c *HomeController) Index(ctx echo.Context) error {
	return c.Render(new(data_merge.Index), ctx)
}

//手机端登陆页面
func (c *HomeController) Login(ctx echo.Context) error {
	return c.Render(new(data_merge.Login), ctx)
}

//活动页面
func (c *HomeController) Discount(ctx echo.Context) error {
	return c.Render(new(data_merge.Discount), ctx)
}

//会员注册页面
func (c *HomeController) Register(ctx echo.Context) error {
	return c.Render(new(data_merge.Register), ctx)
}

//额度转换
func (c *HomeController) Convert(ctx echo.Context) error {
	return c.Render(new(data_merge.Convert), ctx)
}

//快捷支付页面
func (c *HomeController) Fast(ctx echo.Context) error {
	return c.Render(new(data_merge.Fast), ctx)
}

//存款页面
func (c *HomeController) Bank(ctx echo.Context) error {
	return c.Render(new(data_merge.Bank), ctx)
}

//取款页面
func (c *HomeController) Withdraw(ctx echo.Context) error {
	return c.Render(new(data_merge.Withdraw), ctx)
}

//第三方回调页面
func (c *HomeController) PayCallback(ctx echo.Context) error {
	callbackData := new(data_merge.PayCallback)
	member := ctx.Get("member").(*global.MemberRedisToken)
	newHtml, err := global.GetRedis().Get("payRedisKey_" + member.Site + member.Account).Result()
	if err != nil {
		return render.PageErr(60000, ctx)
	}
	callbackData.NewHtml = newHtml
	return c.Render(callbackData, ctx)
}

//线上入款完成页面
func (c *HomeController) Finished(ctx echo.Context) error {
	return c.Render(new(data_merge.Finished), ctx)
}

//快捷支付入款完成页面
func (c *HomeController) Finished2(ctx echo.Context) error {
	return c.Render(new(data_merge.Finished2), ctx)
}

//公司支付完成页面
func (c *HomeController) Carry(ctx echo.Context) error {
	return c.Render(new(data_merge.Carry), ctx)
}

//投注记录
func (c *HomeController) RecordList(ctx echo.Context) error {
	return c.Render(new(data_merge.Record), ctx)
}

//消息中心
func (c *HomeController) MesCenter(ctx echo.Context) error {
	return c.Render(new(data_merge.MesCenter), ctx)
}

//本周报表页面
func (c *HomeController) Statisticsthis(ctx echo.Context) error {
	return c.Render(new(data_merge.StatisticsThis), ctx)
}

//上周报表页面
func (c *HomeController) Statisticslast(ctx echo.Context) error {
	return c.Render(new(data_merge.StatisticsLast), ctx)
}

//电子
func (c *HomeController) EGame(ctx echo.Context) error {
	return c.Render(new(data_merge.EGame), ctx)
}

//自助反水页面
func (c *HomeController) ReturnWater(ctx echo.Context) error {
	return c.Render(new(data_merge.ReturnWater), ctx)
}

//写入出款数据
func (c *HomeController) DrawWrite(ctx echo.Context) error {
	return c.Render(new(data_merge.DrawWrite), ctx)
}

//自助优惠申请大厅
func (c *HomeController) ApplySelf(ctx echo.Context) error {
	return c.Render(new(data_merge.ApplySelf), ctx)
}
