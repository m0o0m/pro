package router

import (
	"bytes"
	"encoding/json"
	"framework/logger"
	"framework/uuid"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/lifei6671/gocaptcha"
	"github.com/mssola/user_agent"
	"global"
	"io/ioutil"
	"math/rand"
	"models/function"
	"models/schema"
	"strings"
	"time"
	"unsafe"

	"encoding/base64"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PromHandler(ctx echo.Context) error {
	promhttp.Handler().ServeHTTP(ctx.Response(), ctx.Request())
	return nil

}

//校验
func MemberCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//确定平台
		var system = "pc"
		ua := user_agent.New(c.Request().UserAgent())
		if ua.Mobile() {
			system = "wap"
		}

		//获取token
		ck, err := c.Cookie("loginBack")
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return echo.ErrUnauthorized
		}
		token := ck.Value
		//redis token check
		//取出redis里存储的data，更改刷新过期时间
		result, err := global.GetRedis().Get(token).Result()
		if err == redis.Nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return echo.ErrUnauthorized
		} else if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if result == "" {
			return echo.ErrUnauthorized
		}
		memberBean := new(function.MemberBean)
		flag, err := memberBean.GetLoginKey(token, system)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if !flag {
			global.GlobalLogger.Error("error:flag false")
			return echo.ErrUnauthorized
		}
		//解析
		results := new(global.MemberRedisToken)
		err = json.Unmarshal([]byte(result), &results)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		//刷新时间
		if system == "pc" {
			results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Pc).Unix()
		} else if system == "wap" {
			results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Wap).Unix()
		} else if system == "ios" {
			results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Ios).Unix()
		} else {
			results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Android).Unix()
		}
		b, err := json.Marshal(results)
		if err != nil {
			return err
		}
		err = global.GetRedis().Set(token, b, 0).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return echo.ErrUnauthorized
		}
		c.Set("member", results) //set进去，后面调用
		c.Set("token", token)
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//中间校验
func GetRedisToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取token
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return c.JSON(401, global.ReplyError(20025, c))
		}
		token := strings.Split(authorization, " ")
		if len(token) < 2 || token[1] == "" {
			return c.JSON(401, global.ReplyError(20025, c))
		}

		//取出redis里存储的data，更改刷新过期时间
		result, err := global.GetRedis().Get(token[1]).Result()
		if err == redis.Nil {
			return c.JSON(401, global.ReplyError(20025, c))
		} else if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if result == "" {
			return c.JSON(401, global.ReplyError(20025, c))
		}
		//数据库查询key是否存在
		agency_bean := new(function.AgencyBean)
		info, ok, err := agency_bean.GetInfoByLoginKey(strings.TrimSpace(token[1]))
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if !ok {
			return c.JSON(401, global.ReplyError(20025, c))
		}
		//如果账号被删除或者被禁用
		if info.DeleteTime != 0 || info.Status != 1 {
			return c.JSON(401, global.ReplyError(20025, c))
		}

		//查询ip限制
		ipSetBean := new(function.IpSetBean)
		var accountType int
		if info.RoleId == 1 {
			accountType = 1
		} else {
			accountType = 2
		}
		ip := c.RealIP()
		if ip == "::1" {
			ip = "127.0.0.1"
		}
		ok, err = ipSetBean.IpCheck(accountType, info.SiteId, ip)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if !ok {
			return c.JSON(401, global.ReplyError(30250, c))
		}

		s := new(global.RedisStruct)
		err = json.Unmarshal([]byte(result), &s)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}

		s.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
		b, err := json.Marshal(s)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}

		err = global.GetRedis().Set(token[1], b, 0).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		c.Set("user", s) //set进去，后面调用
		c.Set("token", token[1])
		//获取之后根据获取到的序列化的数据更新过期时间,重新set
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//admin
func AdminRedisCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取token判断
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return c.JSON(401, global.ReplyError(20025, c))
		}
		token := strings.Split(authorization, " ")
		if len(token) < 2 || token[1] == "" {
			return c.JSON(401, global.ReplyError(20025, c))
		}
		//从redis获取，获取不到就是过期
		result, err := global.GetRedis().Get(token[1]).Result()
		if err == redis.Nil {
			return c.JSON(401, global.ReplyError(20025, c))
		} else if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return c.JSON(401, global.ReplyError(20025, c))
		}
		if result == "" {
			return c.JSON(401, global.ReplyError(20025, c))
		}
		//从数据库获取获取不到就是过期
		admin_bean := new(function.AdminBean)
		_, flag, err := admin_bean.GetInfoByToken(token[1])
		if err != nil {
			return err
		}
		if !flag {
			return c.JSON(401, global.ReplyError(20025, c))
		}

		//解析
		results := new(global.AdminRedisStruct)
		err = json.Unmarshal([]byte(result), &results)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		//刷新时间
		results.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
		b, err := json.Marshal(results)
		if err != nil {
			return err
		}
		//获取之后根据获取到的序列化的数据更新过期时间,重新set
		err = global.GetRedis().Set(token[1], b, 0).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return c.JSON(401, global.ReplyError(20025, c))
		}
		c.Set("admin", results) //set进去，后面调用
		c.Set("token", token[1])
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//agency路由权限验证
func AgencyPerCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		agency := c.Get("user").(*global.RedisStruct)
		//判断当前账号是否为子账号
		url := c.Request().URL.Path
		method := c.Request().Method
		roleBean := new(function.RoleBean)
		if agency.IsSub == 1 {
			//获取该子账号的权限map
			permissionId, err := roleBean.GetPermissionIdByAgencyId(agency.Id)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return err
			}
			data, err := roleBean.GetRoutesByPermissionId(permissionId)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return err
			}
			if v, ok := data[url]; ok {
				if v == method {
					return next(c)
				}
			}
			return c.JSON(401, global.ReplyError(30239, c))
		}
		//获取角色的权限map
		data, err := roleBean.GetRoutesByRole(agency.RoleId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if v, ok := data[url]; ok {
			if v == method {
				return next(c)
			}
		}
		//返回http状态码405
		return c.JSON(401, global.ReplyError(30239, c))

	}
}

//admin路由权限验证
func AdminPerCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		admin := c.Get("user").(*global.AdminRedisStruct)
		url := c.Request().URL.Path
		method := c.Request().Method
		roleBean := new(function.RoleBean)
		//获取角色的权限map
		data, err := roleBean.GetRoutesByRole(admin.RoleId)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if v, ok := data[url]; ok {
			if v == method {
				return next(c)
			}
		}
		//返回http状态码405
		return c.JSON(401, global.ReplyError(30239, c))
	}
}

func Langs(ctx echo.Context) error {

	return ctx.JSONBlob(200, []byte(`[
   {
      "key": "cn",
      "alt": "China",
      "title": "简体中文"
   },
   {
      "key": "us",
      "alt": "United States",
      "title": "English (US)"
   },
   {
      "key": "zh",
      "alt": "China",
      "title": "繁体中文"
   }
]`))
}

func Activitys(ctx echo.Context) error {
	return ctx.JSONBlob(200, []byte(`{
   "activities": {
      "last_update": "12/12/2013 9:43AM",
      "types": [
         {
            "title": "Msgs",
            "name": "msgs",
            "length": 14
         },
         {
            "title": "Notify",
            "name": "notify",
            "length": 3
         }
      ]
   }
}`))
}

func Activity(ctx echo.Context) error {
	return ctx.JSONBlob(200, []byte(`{
   "title": "Notify",
   "length": 3,
   "data": [
      {
         "icon": "fa-user",
         "message": "2 new users just signed up! martin.luther and kevin.reliey",
         "time": "1 min ago..."
      },
      {
         "icon": "fa-facebook",
         "message": "Facebook recived +33 unique likes",
         "time": "4 hrs ago..."
      },
      {
         "icon": "fa-check",
         "message": "2 projects were completed on time! Submitted for your approval - Click here",
         "time": "1 day ago..."
      }
   ]
}`))
}

//验证码路由
func VerCode(ctx echo.Context) error {
	p, err := global.GetExecPath()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//选择随机的字体文件方式
	err = gocaptcha.ReadFonts(p+"/etc/fonts", ".ttf")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//初始化一个验证码对象
	captchaImage, err := gocaptcha.NewCaptchaImage(150, 50, gocaptcha.RandLightColor())
	//画上一条随机直线
	captchaImage.Drawline(0)
	//画随机噪点
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexMedium)
	//画随机文字噪点
	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexMedium)
	//画验证码文字
	str := gocaptcha.RandText(4)
	//str := "1234"
	captchaImage.DrawText(str)
	//画边框
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x00FFFFFF))
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	strRand := uuidStr()

	err = global.GetRedis().Set(strRand, str, 0).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	reader := new(bytes.Buffer)
	captchaImage.SaveImage(reader, gocaptcha.ImageFormatJpeg)
	encodeString := base64.StdEncoding.EncodeToString(reader.Bytes())
	data := &struct {
		Code  string `json:"code"`
		Image string `json:"image"`
	}{
		Code:  strRand,
		Image: "data:image/jpeg;base64," + encodeString,
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

func AppVerCode(ctx echo.Context) error {
	strRand := uuidStr()
	r := rand.New(rand.NewSource(rand.Int63n(time.Now().Unix())))
	tpl := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUvWXYZ0123456789"
	b := make([]byte, 4)
	for i := 0; i < 4; i++ {
		b[i] = tpl[r.Int31n(int32(len(tpl)))]
	}
	code := *(*string)(unsafe.Pointer(&b))
	err := global.GetRedis().Set(strRand, code, time.Second*60*1).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data := &struct {
		Code string `json:"code"`
	}{
		Code: strRand,
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

func GetImage(ctx echo.Context) error {
	code := ctx.QueryParam("code") //客户端类型
	if code == "" {
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	}
	captcha, err := global.GetRedis().Get(code).Result()
	if err == redis.Nil {
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	} else if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if captcha == "" {
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	}

	p, err := global.GetExecPath()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//选择随机的字体文件方式
	err = gocaptcha.ReadFonts(p+"/etc/fonts", ".ttf")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//初始化一个验证码对象
	captchaImage, err := gocaptcha.NewCaptchaImage(150, 50, gocaptcha.RandLightColor())
	//画上一条随机直线
	captchaImage.Drawline(0)
	//画随机噪点
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexMedium)
	//画随机文字噪点
	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexMedium)
	captchaImage.DrawText(captcha)
	//画边框
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x00FFFFFF))
	return captchaImage.SaveImage(ctx.Response(), gocaptcha.ImageFormatJpeg)
}

//根据uuid生成唯一字符串
func uuidStr() string {
	newUuid := uuid.NewV1()
	return newUuid.String()
}

//生成唯一字符串
func randStr() string {
	r := rand.New(rand.NewSource(rand.Int63n(time.Now().Unix())))
	tpl := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUvWXYZ0123456789"
	b := make([]byte, 16)
	for i := 0; i < 16; i++ {
		b[i] = tpl[r.Int31n(int32(len(tpl)))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

//Admin操作日志
func AdminAccessLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		admin := c.Get("admin").(*global.AdminRedisStruct)
		adminLog := new(schema.AdminLog)
		adminLog.OperateAccount = admin.Account     //操作账号
		adminLog.OperateTime = time.Now().Unix()    //操作时间
		adminLog.OperatePath = c.Request().URL.Path //请求路由
		adminLog.Ip = c.RealIP()                    //请求Ip
		if adminLog.Ip == "::1" {
			adminLog.Ip = "127.0.0.1"
		}
		method := c.Request().Method //请求方法

		//根据path和method查询出权限名称
		var roleBean = new(function.RoleBean)
		p, ok, err := roleBean.GetPermissionByPM(adminLog.OperatePath, method, "admin")
		if ok {
			adminLog.OperateInfo = p.Module + "-" + p.PermissionName //模块名和权限名拼接成操作详情
		}

		if method == "GET" {
			adminLog.OperateContent = c.QueryParams().Encode()
		} else {
			ctype := c.Request().Header.Get("Content-Type")
			switch {
			case strings.HasPrefix(ctype, "application/json"):
				data, err := ioutil.ReadAll(c.Request().Body)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return c.JSON(500, global.ReplyError(60000, c))
				}
				adminLog.OperateContent = string(data)
				//重新写进body
				c.Request().Body = ioutil.NopCloser(bytes.NewReader(data))
			case strings.HasPrefix(ctype, "application/x-www-form-urlencoded"),
				strings.HasPrefix(ctype, "multipart/form-data"):
				data, err := c.FormParams()
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return c.JSON(500, global.ReplyError(60000, c))
				}
				adminLog.OperateContent = data.Encode()
			default:
				adminLog.OperateContent = ""
			}
		}
		switch method {
		case "POST": // 增加
			adminLog.Type = 1
		case "DELETE": // 删除
			adminLog.Type = 2
		case "GET": // 查看
			adminLog.Type = 3
		default: // 修改
			adminLog.Type = 4
		}
		sess := global.GetXorm().NewSession()
		defer sess.Close()
		_, err = sess.InsertOne(adminLog)
		if err != nil {
			global.GlobalLogger.Error(logger.ERROR, err.Error())
			return c.JSON(500, global.ReplyError(60000, c))
		}
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//校验
func AppCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//获取token
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return echo.ErrUnauthorized
		}
		token := strings.Split(authorization, " ")
		if len(token) < 2 || token[1] == "" {
			return echo.ErrUnauthorized
		}
		//redis token check
		//取出redis里存储的data，更改刷新过期时间
		result, err := global.GetRedis().Get(token[1]).Result()
		if err == redis.Nil {
			return echo.ErrUnauthorized
		} else if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if result == "" {
			return echo.ErrUnauthorized
		}

		//database token check
		system := c.Request().Header.Get("platform")
		if system == "" || (system != "ios" && system != "android") {
			return echo.ErrUnauthorized
		}
		memberBean := new(function.MemberBean)
		flag, err := memberBean.GetLoginKey(token[1], system)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		if !flag {
			return echo.ErrUnauthorized
		}
		//解析
		results := new(global.MemberRedisToken)
		err = json.Unmarshal([]byte(result), &results)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		//刷新时间
		if system == "ios" {
			results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Ios).Unix()
		} else {
			results.ExpirTime = time.Now().Add(global.DefaultRedisExp.Android).Unix()
		}
		b, err := json.Marshal(results)
		if err != nil {
			return err
		}
		err = global.GetRedis().Set(token[1], b, 0).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return echo.ErrUnauthorized
		}
		c.Set("member", results) //set进去，后面调用
		c.Set("token", token[1])
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}

//Agency操作日志
func AgencyAccessLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		agency := c.Get("user").(*global.RedisStruct)
		agencyLog := new(schema.AgencyLog)
		agencyLog.SiteId = agency.SiteId                //站点id
		agencyLog.SiteIndexId = agency.SiteIndexId      //站点前台id
		agencyLog.OperateAccount = agency.Account       //操作账号
		agencyLog.OperateTime = global.GetCurrentTime() //操作时间
		agencyLog.OperatePath = c.Request().URL.Path    //请求路由
		agencyLog.Ip = c.RealIP()                       //请求Ip
		if agencyLog.Ip == "::1" {
			agencyLog.Ip = "127.0.0.1"
		}
		method := c.Request().Method //请求方法

		//根据path和method查询出权限名称
		var roleBean = new(function.RoleBean)
		p, ok, err := roleBean.GetPermissionByPM(agencyLog.OperatePath, method, "agency")
		if ok {
			agencyLog.OperateInfo = p.Module + "-" + p.PermissionName //模块名和权限名拼接成操作详情
		}

		if method == "GET" {
			agencyLog.OperateContent = c.QueryParams().Encode()
		} else {
			ctype := c.Request().Header.Get("Content-Type")
			switch {
			case strings.HasPrefix(ctype, "application/json"):
				data, err := ioutil.ReadAll(c.Request().Body)
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return c.JSON(500, global.ReplyError(60000, c))
				}
				agencyLog.OperateContent = string(data)
				//重新写进body
				c.Request().Body = ioutil.NopCloser(bytes.NewReader(data))
			case strings.HasPrefix(ctype, "application/x-www-form-urlencoded"),
				strings.HasPrefix(ctype, "multipart/form-data"):
				data, err := c.FormParams()
				if err != nil {
					global.GlobalLogger.Error("error:%s", err.Error())
					return c.JSON(500, global.ReplyError(60000, c))
				}
				agencyLog.OperateContent = data.Encode()
			default:
				agencyLog.OperateContent = ""
			}
		}
		switch method {
		case "POST": // 增加
			agencyLog.Type = 1
		case "DELETE": // 删除
			agencyLog.Type = 2
		case "GET": // 查看
			agencyLog.Type = 3
		default: // 修改
			agencyLog.Type = 4
		}
		sess := global.GetXorm().NewSession()
		defer sess.Close()
		_, err = sess.InsertOne(agencyLog)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return c.JSON(500, global.ReplyError(60000, c))
		}
		global.GlobalLogger.Debug("过中间件,URI:%s", c.Request().RequestURI)
		return next(c)
	}
}
