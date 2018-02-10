package framework

import (
	"config"
	"framework/app"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"models/thirdParty"
	"router/admin"
	"router/agency"
	"router/front"
	"router/wap"
)

func NewApp(cfg *config.Config) (*app.App, error) {
	a := new(app.App)

	a.Cfg = cfg
	a.WebServer = echo.New()
	a.WebServer.HideBanner = true
	//跨域的自定义response的headers的暴露
	defaultCORSConfig := middleware.DefaultCORSConfig
	defaultCORSConfig.ExposeHeaders = append(defaultCORSConfig.ExposeHeaders, "Code")
	defaultCORSConfig.ExposeHeaders = append(defaultCORSConfig.ExposeHeaders, "Platform")
	defaultCORSConfig.ExposeHeaders = append(defaultCORSConfig.ExposeHeaders, "Domain")
	//中间件
	//可以支持跨域请求
	a.WebServer.Use(middleware.CORSWithConfig(defaultCORSConfig))
	//
	a.WebServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 20 << 10, // 1 KB
	}))
	//路由
	a.WebServer.GET("/favicon.ico", func(ctx echo.Context) error {
		return ctx.NoContent(204)
	})
	//测试版本
	a.WebServer.GET("/version/test", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Test)
	})
	//线上版本
	a.WebServer.GET("/version/product", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Product)
	})
	//用户的
	agency.AccountRouter(a.WebServer)
	//现金
	agency.CashRouter(a.WebServer)
	//注单
	agency.NoteRouter(a.WebServer)
	//数据中心
	agency.ReportRouter(a.WebServer)
	//消息中心
	agency.MessageRouter(a.WebServer)
	//资讯系统
	agency.OtherRouter(a.WebServer)
	//资讯管理
	agency.WebSiteRouter(a.WebServer)

	return a, nil
}

func NewAdmin(cfg *config.Config) (*app.App, error) {
	a := new(app.App)

	a.Cfg = cfg
	a.WebServer = echo.New()
	a.WebServer.HideBanner = true
	//跨域的自定义response的headers的暴露
	defaultCORSConfig := middleware.DefaultCORSConfig
	defaultCORSConfig.ExposeHeaders = append(defaultCORSConfig.ExposeHeaders, "Code")
	//中间件
	//可以支持跨域请求
	a.WebServer.Use(middleware.CORSWithConfig(defaultCORSConfig))
	//
	a.WebServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 20 << 10, // 1 KB
	}))
	//路由
	a.WebServer.GET("/favicon.ico", func(ctx echo.Context) error {
		return ctx.NoContent(204)
	})
	//测试版本
	a.WebServer.GET("/version/test", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Test)
	})
	//线上版本
	a.WebServer.GET("/version/product", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Product)
	})
	//a.WebServer.Logger = global.GlobalLogger
	//总后台
	admin.NewAdminRouter(a.WebServer)
	admin.ReportRouter(a.WebServer)
	admin.NoteRouter(a.WebServer)
	return a, nil
}

func NewFront(cfg *config.Config) (*app.App, error) {
	a := new(app.App)
	a.Cfg = cfg
	thirdParty.VideoCfg = &a.Cfg.Video
	a.WebServer = echo.New()
	a.WebServer.HideBanner = true
	//跨域的自定义response的headers的暴露
	defaultCORSConfig := middleware.DefaultCORSConfig
	defaultCORSConfig.ExposeHeaders = append(defaultCORSConfig.ExposeHeaders, "Code")
	//中间件
	//可以支持跨域请求
	a.WebServer.Use(middleware.CORSWithConfig(defaultCORSConfig))
	//
	a.WebServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 20 << 10, // 1 KB
	}))
	//路由
	a.WebServer.GET("/favicon.ico", func(ctx echo.Context) error {
		return ctx.NoContent(204)
	})
	//测试版本
	a.WebServer.GET("/version/test", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Test)
	})
	//线上版本
	a.WebServer.GET("/version/product", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Product)
	})
	//个人中心
	front.Center(a.WebServer)
	//模板
	front.PageTemplate(a.WebServer)
	//wap模板
	front.WapPageTemplate(a.WebServer)
	//wap会员中心
	front.WapMember(a.WebServer)
	return a, nil
}

func NewWap(cfg *config.Config) (*app.App, error) {
	a := new(app.App)

	a.Cfg = cfg
	a.WebServer = echo.New()
	a.WebServer.HideBanner = true
	//跨域的自定义response的headers的暴露
	defaultCORSConfig := middleware.DefaultCORSConfig
	defaultCORSConfig.ExposeHeaders = append(defaultCORSConfig.ExposeHeaders, "Code")
	defaultCORSConfig.ExposeHeaders = append(defaultCORSConfig.ExposeHeaders, "Platform")
	//中间件
	//可以支持跨域请求
	a.WebServer.Use(middleware.CORSWithConfig(defaultCORSConfig))
	//
	a.WebServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 20 << 10, // 1 KB
	}))
	//路由
	a.WebServer.GET("/favicon.ico", func(ctx echo.Context) error {
		return ctx.NoContent(204)
	})
	//测试版本
	a.WebServer.GET("/version/test", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Test)
	})
	//线上版本
	a.WebServer.GET("/version/product", func(ctx echo.Context) error {
		return ctx.String(200, cfg.Version.Product)
	})
	//总后台
	wap.WapRouter(a.WebServer)
	return a, nil
}
