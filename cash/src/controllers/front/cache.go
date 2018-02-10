package front

import (
	"controllers"
	d1 "controllers/front/page/data_merge"
	d2 "controllers/front/wap/data_merge"
	"framework/render"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//pc需要缓存的页面
var pcPageDatas = []render.PcPageData{
	new(d1.NIndex),         //主页
	new(d1.Applypro),       //优惠活动大厅
	new(d1.QuickPay),       //快速充值中心
	new(d1.NoticeData),     //广告弹出框
	new(d1.VideoRule),      //视讯游戏规则
	new(d1.Wapview),        //手机下注
	new(d1.LoginInfo),      //登录信息条款页
	new(d1.Download),       //下载专区
	new(d1.Isoview),        //iso教程
	new(d1.Register),       //注册
	new(d1.AgencyRegister), //代理注册
	new(d1.IWord),          //文案页面
	new(d1.Detect),         //线路检测
}

//wap需要缓存的页面
var wapPageDatas = []render.WapPageData{
	new(d2.Index),     //首页
	new(d2.Convert),   //额度转换
	new(d2.Bank),      //存款
	new(d2.Finished),  //存款支付完成
	new(d2.Carry),     //公司入款提交完成
	new(d2.Fast),      //快速充值中心
	new(d2.Register),  //注册
	new(d2.MesCenter), //消息中心
	new(d2.Withdraw),  //存款
	new(d2.Record),    //游戏公告
}

//缓存操作相关
type CacheController struct {
	controllers.BaseController
}

//缓存站点皮肤信息
func (*CacheController) CacheTheme(ctx echo.Context) error {
	themes, err := siteBean.GetThemeAll()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for _, v := range themes {
		global.ThemeCache.Store(v.SiteId+"$"+v.SiteIndexId, v.ThemeName)
	}
	return ctx.NoContent(204)
}

//主动缓存单站点所有页面
func (*CacheController) GenPageCacheBySite(ctx echo.Context) error {
	reqDta := new(input.From)
	code := global.ValidRequest(reqDta, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct) //代理的siteId和indexId
	if reqDta.ClientType == 1 {
		for _, pcPageData := range pcPageDatas {
			_, code := render.GenPcCache(user.SiteId, user.SiteIndexId, pcPageData)
			if code != 0 {
				return ctx.JSON(200, global.ReplyError(code, ctx))
			}
		}
	} else {
		for _, wapPageData := range wapPageDatas {
			_, code := render.GenWapCache(user.SiteId, user.SiteIndexId, wapPageData)
			if code != 0 {
				return ctx.JSON(200, global.ReplyError(code, ctx))
			}
		}
	}
	return ctx.NoContent(204)
}

//单站点批量页面缓存 = 删除
func (*CacheController) DelPageCacheBySite(ctx echo.Context) error {
	reqData := new(input.From)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct) //代理的siteId和indexId
	if reqData.ClientType == 1 {
		for _, pcPageData := range pcPageDatas {
			err := render.DelCache(user.SiteId, user.SiteIndexId, pcPageData)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(71035, ctx))
			}
		}
	} else {
		for _, wapPageData := range wapPageDatas {
			err := render.DelCache(user.SiteId, user.SiteIndexId, wapPageData)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(71035, ctx))
			}
		}
	}
	return ctx.NoContent(204)
}
