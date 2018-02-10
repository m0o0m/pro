//会员个人基本资料
package memberinfo

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

type MemberInfoController struct {
	controllers.BaseController
}

//会员个人基本资料
func (c *MemberInfoController) MemberSelfInfo(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//获取登录会员的个人资料
	info, has, err := mSIBean.GetMemberInfoSelf(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(70014, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//修改/添加邮箱
func (c *MemberInfoController) EmailAddOrChange(ctx echo.Context) error {
	//获取参数
	email := new(input.EmailAddOrChangeIn)
	code := global.ValidRequestMember(email, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	flag := global.CheckEmail(email.Email)
	if !flag {
		return ctx.JSON(200, global.ReplyError(20015, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	email.Id = member.Id
	count, err := mSIBean.EmailChangeOrAdd(email)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//修改/添加生日
func (c *MemberInfoController) BirthAddOrChange(ctx echo.Context) error {
	//获取参数
	birth := new(input.BirthAddOrChangeIn)
	code := global.ValidRequestMember(birth, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	birth.Id = member.Id
	count, err := mSIBean.BirthChangeOrAdd(birth)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//修改/添加手机号
func (c *MemberInfoController) PhoneAddOrChange(ctx echo.Context) error {
	//获取参数
	phone := new(input.PhoneBind)
	code := global.ValidRequestMember(phone, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//从redis中获取验证码 验证码是否为空
	//phone_code, err := global.GetRedis().Get(phone.Phone).Result()
	//if err != nil || phone_code == "" {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(500, global.ReplyError(50147, ctx))
	//}
	//判断用户输入验证是否正确
	//if phone.PhoneCode != phone_code {
	//	return ctx.JSON(200, global.ReplyError(20021, ctx))
	//}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	phone.Id = member.Id
	count, err := mSIBean.PhoneChangeOrAdd(phone)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//手机验证码（TODO  测试版  只生成4位固定的数字）
func (c *MemberInfoController) PhoneCode(ctx echo.Context) error {
	//获取参数
	ph := new(input.PhoneCode)
	code := global.ValidRequestMember(ph, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	err := global.GetRedis().Set(ph.Phone, "1234", time.Minute*5).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//这里可以加上get去获取刚刚set进去的数据看能不能获取到
	phones, err := global.GetRedis().Get(ph.Phone).Result()
	if err != nil || phones == "" {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, phones)
}

//会员中心主页
func (c *MemberInfoController) MemberHomePage(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	info, has, err := mSIBean.MemberHomePage(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(70014, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//会员修改密码
func (c *MemberInfoController) MemberPasswordChange(ctx echo.Context) error {
	//获取参数
	mc := new(input.PasswordMemberChange)
	code := global.ValidRequestMember(mc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//两次密码不一致
	if mc.Password != mc.RePassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	//新密码不能和原密码一致
	if mc.Password == mc.OriginalPassword {
		return ctx.JSON(200, global.ReplyError(50152, ctx))
	}
	if mc.Type == 1 {
		//密码加密
		if mc.Password != "" {
			md5Password, err := global.MD5ByStr(mc.Password, global.EncryptSalt)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			mc.Password = md5Password
		}
		//原始密码加密
		if mc.OriginalPassword != "" {
			md5OrPassword, err := global.MD5ByStr(mc.OriginalPassword, global.EncryptSalt)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			mc.OriginalPassword = md5OrPassword
		}
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	//根据登陆人的id获取会员的密码
	info, has, err := mSIBean.MemberOneInfo(member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//如果不存在返回  会员不存在或者状态不正常
	if !has {
		return ctx.JSON(200, global.ReplyError(60051, ctx))
	}
	//会员密码不允许修改
	if info.IsEditPassword != 1 {
		return ctx.JSON(200, global.ReplyError(50149, ctx))
	}
	//根据获取的参数中类型判断原密码是否正确
	if mc.Type == 1 {
		//登录密码
		if mc.OriginalPassword != info.Password {
			return ctx.JSON(200, global.ReplyError(20009, ctx))
		}
	} else if mc.Type == 2 {
		//取款密码
		if mc.OriginalPassword != info.DrawPassword {
			return ctx.JSON(200, global.ReplyError(20009, ctx))
		}
	}
	//修改密码
	count, err := mSIBean.MemberSelfPassword(mc, member.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}
