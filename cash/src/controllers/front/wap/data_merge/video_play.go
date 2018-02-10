package data_merge

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"models/thirdParty"
	"net/http"
	"time"
)

type VideoPlay struct{}

//wap视讯电子游戏链接请求
func (c *VideoPlay) PostVideoPlay(ctx echo.Context) error {
	ip := ctx.RealIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	//跳转登陆闭包
	redirectFunc := func(c echo.Context) error {
		// delete cookie
		ck := new(http.Cookie)
		ck.Name = "loginBack"
		ck.Value = ""
		loc, _ := time.LoadLocation("Local")
		ck.Expires, _ = time.ParseInLocation("2006-01-02", "2017-01-01", loc) //将过期时间设置为2017年
		c.SetCookie(ck)
		return c.Redirect(http.StatusMovedPermanently, "/index")
	}
	videoParam := new(input.VideoPlay)
	code := global.ValidRequest(videoParam, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	videoBean := thirdParty.NewThirdParty() //第三方游戏地址处理
	var m input.VideoUserData
	m.SiteId = siteId
	m.IndexId = siteIndexId
	m.IsSw = false
	m.Limit = "10000"
	m.IP = ip
	m.Lang = "zh"
	m.Cur = "RMB"
	m.Media = "wap"
	m.Platform = c.GetGameType(videoParam.VType)
	if videoParam.GameId != "" {
		m.GameID = videoParam.GameId
	}
	//获取token,如果没有,说明是试玩
	ck, err := ctx.Cookie("loginBack")
	if err != nil {
		// TODO 试玩
		global.GlobalLogger.Error("error:%s", err.Error())
		m.IsSw = true
		url, _ := videoBean.ForwardGame(m)
		if len(url) == 0 {
			global.GlobalLogger.Error("error type: %s gameid %s,error:%s", videoParam.VType, videoParam.GameId, err.Error())
			return ctx.HTML(200, c.GetErr(m.Platform))
		}
		return ctx.HTML(200, c.GetUrlStr(m.Platform, url))
	} else {
		// TODO 非试玩
		token := ck.Value
		global.GlobalLogger.Debug("token:%s", token)
		//redis token check
		//取出redis里存储的data，更改刷新过期时间
		result, err := global.GetRedis().Get(token).Result()
		if err == redis.Nil {
			global.GlobalLogger.Error("redis not found error:%s", err.Error())
			return redirectFunc(ctx)
		} else if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc(ctx)
		}
		if result == "" {
			return redirectFunc(ctx)
		}
		flag, err := memberBean.GetLoginKey(token, "wap")
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc(ctx)
		}
		if !flag {
			return redirectFunc(ctx)
		}
		member := new(global.MemberRedisToken)
		err = json.Unmarshal([]byte(result), &member)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc(ctx)
		}
		//********** 以上为登陆信息校验,校验失败就返回登陆页面 ***********
		m.UserName = member.Account
		//根据站点账号取出对应的会员信息
		info, flag, err := memberBean.GetInfoBySite(m.SiteId, m.IndexId, m.UserName)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//账号不存在,已经被删除在获取的时候已经被排除出去，或者是该站点下面不存在该账号
		if !flag {
			return ctx.JSON(200, global.ReplyError(30138, ctx))
		}
		m.ShId = info.FirstAgencyId
		m.UaId = info.SecondAgencyId
		m.AgentId = info.ThirdAgencyId
		m.Domain = ctx.Request().Host

		url, err := videoBean.ForwardGame(m)
		if len(url) == 0 {
			global.GlobalLogger.Error("error type: %s gameid %s,error:%s", videoParam.VType, videoParam.GameId, err.Error())
			return ctx.JSON(200, c.GetErr(m.Platform))
		}
		return ctx.HTML(200, c.GetUrlStr(m.Platform, url))
	}
	return ctx.HTML(200, global.GetError(60235, ctx))
}

func (c *VideoPlay) GetErr(vType string) string {
	return `
<html>
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>` + vType + `</title>
</head>
<body>
    接口不通,<a href="/index"><点击></a>跳转首页
</body>
</html>`
}
func (c *VideoPlay) GetUrlStr(vType, url string) string {
	return `
<html>
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>` + vType + `</title>
</head>
<body>
    游戏加载中...
</body>
<script>
    window.location.href="` + url + `"
</script>
</html>`
}

//整理游戏传递平台参数
func (c *VideoPlay) GetGameType(vType string) (gType string) {
	switch vType {
	case "eg_dz":
		gType = "eg"
	case "ag_dz", "ag_by":
		gType = "ag"
	case "bbin_dz", "bbin_sp":
		gType = "bbin"
	case "pt_dz":
		gType = "pt"
	case "hb_dz":
		gType = "hb"
	case "gd_dz":
		gType = "gddz"
	case "gpi_dz":
		gType = "gpidz"
	case "gg_dz":
		gType = "gg"
	case "pk_fc":
		gType = "pk"
	case "cs_fc":
		gType = "cs"
	case "eg_fc":
		gType = "egtc"
	case "pk_sp":
		gType = "pk"
	case "im_sp":
		gType = "im"
	case "sb_sp":
		gType = "sb"
	default:
		gType = vType
	}
	return
}
