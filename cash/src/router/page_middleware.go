package router

import (
	"controllers/front/page/data_merge"
	"encoding/json"
	"framework/render"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/mssola/user_agent"
	"global"
	"models/back"
	"models/function"
	"models/input"
	"net/http"
	"strings"
	"time"
)

//判断是否是手机端还是pc端来决定不写路由跳转到哪
func WhoIndex(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ua := user_agent.New(c.Request().UserAgent())
		if ua.Mobile() {
			c.Redirect(http.StatusMovedPermanently, "/m/index")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/index")
		}
		return next(c)
	}
}

//页面中间件(通过域名获取site_id,site_index_id)
func PageCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//手机pc切换
		ua := user_agent.New(c.Request().UserAgent())
		if ua.Mobile() {
			global.GlobalLogger.Debug("pc goto mobile")
			c.Redirect(http.StatusMovedPermanently, "/m/index")
			return next(c)
		}
		host := c.Request().Host
		siteDomainBean := new(function.SiteDomainBean)
		siteDomain, b, err := siteDomainBean.GetSiteByDomain(strings.Split(host, ":")[0])
		if err != nil {
			global.GlobalLogger.Error("domain error %s %v ", host, err)
			return echo.ErrNotFound
		}
		if !(b) {
			global.GlobalLogger.Warn("domain error %s %v ", host, err)
			return echo.ErrNotFound
		}
		if siteDomain.SiteId == "" {
			global.GlobalLogger.Warn("domain SiteId error %s %v ", host, err)
			return echo.ErrNotFound
		}
		if siteDomain.SiteIndexId == "" {
			global.GlobalLogger.Warn("domain SiteIndexId error %s %v ", host, err)
			return echo.ErrNotFound
		}
		//是否整站维护,如果是,就跳到维护页面
		siteModuleInf, ok := global.SiteModuleCache.Load(global.GenKey("all", "pc", "indexid"))
	OK:
		if ok {
			maintenance := new(data_merge.Maintenance)
			maintenance.WebHome = 1
			siteModule, _ := siteModuleInf.(*back.SiteModule)
			maintenance.Content = siteModule.Content
			//bytes, ok := render.GetCache(siteDomain.SiteId, siteDomain.SiteIndexId, maintenance)
			//if ok {
			//	return c.HTMLBlob(200, bytes)
			//}
			bytes, code := render.GenPcCache(siteDomain.SiteId, siteDomain.SiteIndexId, maintenance)
			if code == 0 {
				return c.HTMLBlob(200, bytes)
			}
			return render.PageErr(code, c)
		}
		siteModuleInf, ok = global.SiteModuleCache.Load(global.GenKey(siteDomain.SiteId, "pc", "indexid"))
		if ok {
			goto OK
		}
		c.Set("site_id", siteDomain.SiteId) //set进去，后面调用
		c.Set("site_index_id", siteDomain.SiteIndexId)
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//校验登陆
func PcMemberCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取token
		ck, err := c.Cookie("loginBack")
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			c.Redirect(http.StatusFound, "/index")
			return next(c)
		}
		token := ck.Value
		global.GlobalLogger.Error("token:%s", token)
		redirectFunc := func() error {
			// delete cookie
			ck.Value = ""
			loc, _ := time.LoadLocation("Local")
			ck.Expires, _ = time.ParseInLocation("2006-01-02", "2017-01-01", loc) //将过期时间设置为2017年
			c.SetCookie(ck)
			c.Redirect(http.StatusMovedPermanently, "/index")
			return next(c)
		}
		global.GlobalLogger.Debug("token:%s", token)
		//redis token check
		//取出redis里存储的data，更改刷新过期时间
		result, err := global.GetRedis().Get(token).Result()
		if err == redis.Nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		} else if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		if result == "" {
			return redirectFunc()
		}

		memberBean := new(function.MemberBean)
		flag, err := memberBean.GetLoginKey(token, "pc")
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		if !flag {
			return redirectFunc()
		}
		//解析
		results := new(global.MemberRedisToken)
		err = json.Unmarshal([]byte(result), &results)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		//刷新时间
		results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Pc).Unix()
		b, err := json.Marshal(results)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		err = global.GetRedis().Set(token, b, 0).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		c.Set("member", results) //set进去，后面调用
		c.Set("token", token)
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//pc单商品是否维护判断的页面中间件
func SingleProductCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//第一个中间件已经获取了siteId和siteIndexId
		siteId, _ := c.Get("site_id").(string)
		siteIndexId, _ := c.Get("site_index_id").(string)
		global.GlobalLogger.Debug("middleware: %s %s", siteId, siteIndexId)
		reqDta := new(input.VType)
		global.ValidRequest(reqDta, c)
		if reqDta.VType != "" {
			//只拦截有参的
			//商品是否维护维护,如果是,就跳到维护页面
			siteModuleInf, ok := global.SiteModuleCache.Load(global.GenKey("all", "pc", reqDta.VType))
		OK:
			if ok {
				maintenance := new(data_merge.Maintenance)
				maintenance.WebHome = 2
				siteModule, _ := siteModuleInf.(*back.SiteModule)
				maintenance.Content = siteModule.Content
				//不读缓存
				//bytes, ok := render.GetCache(siteId, siteIndexId, maintenance)
				//if ok {
				//	return c.HTMLBlob(200, bytes)
				//}
				bytes, code := render.GenPcCache(siteId, siteIndexId, maintenance)
				if code == 0 {
					return c.HTMLBlob(200, bytes)
				}
				return render.PageErr(code, c)
			}
			siteModuleInf, ok = global.SiteModuleCache.Load(global.GenKey(siteId, "pc", reqDta.VType))
			if ok {
				goto OK
			}

		}
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//wap页面中间件(通过域名获取site_id,site_index_id)
func WapPageCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//手机pc切换
		ua := user_agent.New(c.Request().UserAgent())
		if !ua.Mobile() {
			c.Redirect(http.StatusMovedPermanently, "/index")
			return next(c)
		}
		host := c.Request().Host
		siteDomainBean := new(function.SiteDomainBean)
		siteDomain, b, err := siteDomainBean.GetSiteByDomain(strings.Split(host, ":")[0])
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			return echo.ErrNotFound
		}
		if !(b) {
			return echo.ErrNotFound
		}
		if siteDomain.SiteId == "" {
			return echo.ErrNotFound
		}
		if siteDomain.SiteIndexId == "" {
			return echo.ErrNotFound
		}
		c.Set("site_id", siteDomain.SiteId) //set进去，后面调用
		c.Set("site_index_id", siteDomain.SiteIndexId)
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//wap校验登陆
func WapMemberCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		//获取token
		ck, err := c.Cookie("loginBack")
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			c.Redirect(http.StatusFound, "/m/login")
			return next(c)
		}
		token := ck.Value
		redirectFunc := func() error {
			// delete cookie
			ck.Value = ""
			loc, _ := time.LoadLocation("Local")
			ck.Expires, _ = time.ParseInLocation("2006-01-02", "2017-01-01", loc) //将过期时间设置为2017年
			c.SetCookie(ck)
			c.Redirect(http.StatusMovedPermanently, "/m/login")
			return next(c)
		}

		//redis token check
		//取出redis里存储的data，更改刷新过期时间
		result, err := global.GetRedis().Get(token).Result()
		if err == redis.Nil {
			return redirectFunc()
		} else if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		if result == "" {
			return redirectFunc()
		}

		memberBean := new(function.MemberBean)
		flag, err := memberBean.GetLoginKey(token, "wap")
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		if !flag {
			return redirectFunc()
		}
		//解析
		results := new(global.MemberRedisToken)
		err = json.Unmarshal([]byte(result), &results)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		//刷新时间
		results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Pc).Unix()
		b, err := json.Marshal(results)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		err = global.GetRedis().Set(token, b, 0).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return redirectFunc()
		}
		c.Set("member", results) //set进去，后面调用
		c.Set("token", token[1])
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}
