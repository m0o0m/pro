package page

import (
	"controllers/front/page/data_merge"
	"fmt"
	"framework/render"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"strings"
)

type HomeController struct {
	PageBaseController
}

//n_index.html页面
func (c *HomeController) NIndex(ctx echo.Context) error {
	return c.Render(new(data_merge.NIndex), ctx)
}

//体育页面
func (c *HomeController) Sports(ctx echo.Context) error {
	return c.Render(new(data_merge.Sports), ctx)
}

//视讯页面
func (c *HomeController) LiveTop(ctx echo.Context) error {
	return c.Render(new(data_merge.LiveTop), ctx)
}

//彩票页面
func (c *HomeController) Lottery(ctx echo.Context) error {
	return c.Render(new(data_merge.Lottery), ctx)
}

//下载专区
func (c *HomeController) Download(ctx echo.Context) error {
	return c.Render(new(data_merge.Download), ctx)
}

//egame.html页面
func (c *HomeController) EGame(ctx echo.Context) error {
	reqDta := new(input.VType)
	code := global.ValidRequest(reqDta, ctx)
	if code != 0 {
		return render.PageErr(code, ctx)
	}
	xx := new(data_merge.Game)
	xx.Type = reqDta.VType
	return c.Render(xx, ctx)
}

//youhui.html页面
func (c *HomeController) YouHui(ctx echo.Context) error {
	return c.Render(new(data_merge.Youhui), ctx)
}

//线路检测页面
func (c *HomeController) Detect(ctx echo.Context) error {
	return c.Render(new(data_merge.Detect), ctx)
}

//会员注册页面
func (c *HomeController) Register(ctx echo.Context) error {
	return c.Render(new(data_merge.Register), ctx)
}

//代理注册页面
func (c *HomeController) AgencyRegister(ctx echo.Context) error {
	return c.Render(new(data_merge.AgencyRegister), ctx)
}

//文案页面
func (c *HomeController) IWord(ctx echo.Context) error {
	combo := new(input.IwordList)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return render.PageErr(code, ctx)
	}
	iword := new(data_merge.IWord)
	iword.Id = combo.Id
	return c.Render(iword, ctx)
}

//点击公告弹出页面
func (c *HomeController) NoticeData(ctx echo.Context) error {
	return c.Render(new(data_merge.NoticeData), ctx)
}

//快速充值中心
func (c *HomeController) QuickPay(ctx echo.Context) error {
	return c.Render(new(data_merge.QuickPay), ctx)
}

//优惠活动大厅
func (c *HomeController) Applypro(ctx echo.Context) error {
	return c.Render(new(data_merge.Applypro), ctx)
}

//会员中心首页--我的账户
func (c *HomeController) MemberAccount(ctx echo.Context) error {
	isOtherData := new(input.IsParameter)
	code := global.ValidRequestMember(isOtherData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	page := new(data_merge.MemberIndex)
	showUrl := ctx.Request().RequestURI
	data := strings.Split(showUrl, "?")
	url := data[0]
	if isOtherData != nil {
		page.IsOther = isOtherData.IsOther
	}
	page.MemberPage = 1
	if url == "/member/account" { //我的账户
		page.MemberPage = 1
	} else if url == "/member/bank" { //存款,公司入款
		page.MemberPage = 2
	} else if url == "/member/withdraw" { //取款
		page.MemberPage = 3
	} else if url == "/member/convert" { //额度转换
		page.MemberPage = 4
	} else if url == "/member/record" { //交易记录
		page.MemberPage = 5
	} else if url == "/member/report" { //报表统计
		page.MemberPage = 6
	} else if url == "/member/spread" { //我要推广
		page.MemberPage = 7
	} else if url == "/member/mescenter" { //消息中心
		page.MemberPage = 8
	} else if url == "/member/bank/third" { //线上第三方支付
		page.MemberPage = 9
	} else if url == "/member/bank/complete" { //线上第三方支付
		page.MemberPage = 10
	} else if url == "/member/draw/write" { //出款写入
		page.MemberPage = 11
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	//获取会员层级信息
	//获取该会员所处层级详情
	levelInfo, flag, err := member_level_bean.GetLevelInfo(member.LevelId)
	if err != nil {
		page.IsSelf = 2 //1开启 2未开启
	}
	//找不到该层级
	if !flag || levelInfo.LevelId == "" {
		page.IsSelf = 2 //1开启 2未开启
	} else {
		page.IsSelf = int(levelInfo.IsSelfRebate)
	}
	return c.RenderNowData(page, ctx)
}

//第三方返回支付页面
func (c *HomeController) PayCallback(ctx echo.Context) error {
	callbackData := new(data_merge.PayCallback)
	member := ctx.Get("member").(*global.MemberRedisToken)
	fmt.Println("payRedisKey_" + member.Site + member.Account)
	newHtml, err := global.GetRedis().Get("payRedisKey_" + member.Site + member.Account).Result()
	if err != nil {
		return render.PageErr(60000, ctx)
	}
	callbackData.NewHtml = newHtml
	return c.Render(callbackData, ctx)
}

//新版本维护功能
func (c *HomeController) Maintain(ctx echo.Context) error {
	return c.Render(new(data_merge.Maintain), ctx)
}
