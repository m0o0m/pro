//[控制器] [平台]登录管理
package login

import (
	"fmt"
	"strings"
	"time"

	"controllers"
	"encoding/json"
	"framework/validation"
	"global"
	"models/back"
	"models/input"
	"models/schema"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

//登录管理
type LoginDoController struct {
	controllers.BaseController
}

//登录处理
func (c *LoginDoController) GetLoginDo(ctx echo.Context) error {
	login := new(input.AdminLogin)
	if err := ctx.Bind(login); err != nil {
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	}
	//数据校验
	valid := validation.Validation{}
	ok, err := valid.Valid(login)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !ok {
		for _, e := range valid.Errors {
			global.GlobalLogger.Error("error:%s", e.Error())
			return ctx.JSON(200, global.ReplyError(e.Code(), ctx))
		}
	}
	//验证验证码
	code := ctx.Request().Header.Get("code")
	key, err := GetTokenS(code)
	if err == redis.Nil {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if key == "" || strings.ToLower(key) != strings.ToLower(login.Code) {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	//删除验证码
	err = global.GetRedis().Del(code).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//根据登录账号查询账号
	admin, ok, err := adminBean.GetInfoByAcPa(login.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//未查询到账号
	if !ok {
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//加密密码
	login.Password, err = global.MD5ByStr(login.Password, global.EncryptSalt)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//密码错误
	if admin.Password != login.Password {
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//账号被删除
	if admin.DeleteTime != 0 {
		return ctx.JSON(200, global.ReplyError(20001, ctx))
	}
	//账号被禁用
	if admin.Status != 1 {
		return ctx.JSON(200, global.ReplyError(20002, ctx))
	}
	//获取ip
	loginIp := ctx.RealIP()
	if loginIp == "::1" {
		loginIp = "127.0.0.1"
	}
	//获取ip限制
	if admin.LoginIp != "" {
		//以英文字符的逗号分割
		adminLoginIp := strings.Split(admin.LoginIp, ",")
		ok := false
		for k := range adminLoginIp {
			if loginIp == adminLoginIp[k] {
				ok = true
				break
			}
		}
		if !ok {
			return ctx.JSON(200, global.ReplyError(30250, ctx))
		}
	}
	s := fmt.Sprintf("admin %d", time.Now().UnixNano()+admin.Id)
	//生成tokenkey
	result, err := global.MD5ByBytes([]byte(s), []byte(global.EncryptSalt))
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	data := new(global.AdminRedisStruct)
	data.Id = admin.Id
	data.Account = admin.Account
	data.Status = admin.Status
	data.RoleId = admin.RoleId
	data.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
	data.Type = "admin"

	//序列化
	b, err := json.Marshal(data)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//redis存储
	err = keySet(result, b, admin.LoginKey)
	s, err = GetTokenS(result)
	if err != nil || s == "" {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//刷新数据库
	var adminLogin schema.Admin
	adminLogin.Id = admin.Id
	adminLogin.LoginKey = result
	//登录日志
	var newLoginLog schema.AdminLoginLog
	newLoginLog.Account = admin.Account
	newLoginLog.LoginIp = loginIp
	newLoginLog.Device = 1
	newLoginLog.LoginRole = admin.RoleId
	newLoginLog.LoginTime = global.GetCurrentTime()
	newLoginLog.LoginResult = 1
	count, err := adminBean.RefreshLoginKey(adminLogin, newLoginLog)
	if err != nil || count == 0 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//返回数据
	adminBack := new(back.AdminSign)
	adminBack.Id = admin.Id
	adminBack.Account = admin.Account
	adminBack.Status = admin.Status
	adminBack.RoleId = admin.RoleId
	adminBack.Token = result
	return ctx.JSON(200, global.ReplyItem(adminBack))
}

func keySet(result string, b []byte, beforeKey string) (err error) {
	if beforeKey != "" {
		//删除旧的key
		err = global.GetRedis().Del(beforeKey).Err()
		//将旧的删除
		err = global.GetRedis().LPop(result).Err()
	}
	//存储新token
	err = global.GetRedis().Set(result, b, 0).Err()
	//将推进list
	err = global.GetRedis().RPush("admin_login", result).Err()
	return err
}

//获取登录的时候存储的redis值
func GetTokenS(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return key, err
	}
	return key, err
}

//退出
func (c *LoginDoController) GetLoginOut(ctx echo.Context) error {
	admin := ctx.Get("admin").(*global.AdminRedisStruct)
	token := ctx.Get("token").(string)
	//更改redis
	err := global.GetRedis().Del(token).Err()
	if err != nil {
		global.GlobalLogger.Error("Error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//弹出list
	global.GetRedis().LPop(token).Err()
	//修改数据库
	count, err := adminBean.UpAdminLoginStatus(admin.Id)
	if err != nil || count != 1 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//修改密码
func (c *LoginDoController) PutLoginPassword(ctx echo.Context) error {
	newPassword := new(input.UpdatePassword)
	code := global.ValidRequestAdmin(newPassword, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//根据id取出账号信息
	info, flag, err := adminBean.GetInfomationById(newPassword.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60051, ctx))
	}
	withdraw, err := global.MD5ByStr(newPassword.BeforePassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//旧密码错误
	if withdraw != info.Password {
		return ctx.JSON(200, global.ReplyError(20009, ctx))
	}
	//判断新密码是否和旧密码一致
	if newPassword.NewPassword == newPassword.BeforePassword {
		return ctx.JSON(200, global.ReplyError(60214, ctx))
	}
	//新密码两次输入不一致
	if newPassword.NewPassword != newPassword.RepeatPassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	password, err := global.MD5ByStr(newPassword.NewPassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	newAdmin := new(schema.Admin)
	newAdmin.Id = newPassword.Id
	newAdmin.Password = password
	_, err = adminBean.ChangePassword(newAdmin)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}
