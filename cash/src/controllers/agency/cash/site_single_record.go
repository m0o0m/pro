package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
	"time"
)

//掉单管理
type SiteSingleRecord struct {
	controllers.BaseController
}

//提交掉单申请
func (*SiteSingleRecord) Add(ctx echo.Context) error {
	siteSingleRecord := new(input.SiteSingleRecordAdd)
	code := global.ValidRequest(siteSingleRecord, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	if !(user.Level == 4 || user.Level == 1) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//根据操作者id获取操作者名称
	agency, err := manualAccessBean.GetUserNameById(user.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	siteSingleRecord.AdminUser = agency.Username
	//查看会员账号是否是该操作人所在站点
	has, err := siteSingleRecordBean.IsExistMemberAccount(siteSingleRecord)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	//判断转入项目是否存在
	if siteSingleRecord.Vtype != 0 {
		has, err := memberBalanceConversionBean.IsExistFtype(siteSingleRecord.Vtype)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30147, ctx))
		}
	}
	//判断转出项目是否存在
	if siteSingleRecord.Ctype != 0 {
		has, err := memberBalanceConversionBean.IsExistFtype(siteSingleRecord.Ctype)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30154, ctx))
		}
	}
	//转入项目或者转出项目必须有一个是系统余额
	if siteSingleRecord.Vtype != 0 && siteSingleRecord.Ctype != 0 {
		return ctx.JSON(200, global.ReplyError(30157, ctx))
	}
	count, err := siteSingleRecordBean.Add(siteSingleRecord)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30164, ctx))
	}
	return ctx.NoContent(204)
}

//审核掉单申请
func (*SiteSingleRecord) Check(ctx echo.Context) error {
	siteSingleRecord := new(input.SiteSingleRecordEdit)
	code := global.ValidRequest(siteSingleRecord, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	if user.Level != 1 {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//根据操作者id获取操作者账号
	agency, err := manualAccessBean.GetUserNameById(user.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	siteSingleRecord.UpdateUsername = agency.Username
	count, err := siteSingleRecordBean.Check(siteSingleRecord)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30165, ctx))
	}
	return ctx.NoContent(204)
}

//掉单申请列表
func (ss *SiteSingleRecord) Index(ctx echo.Context) error {
	siteSingleRecord := new(input.SiteSingleRecordList)
	code := global.ValidRequest(siteSingleRecord, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if siteSingleRecord.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", siteSingleRecord.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if siteSingleRecord.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", siteSingleRecord.EndTime, loc)
		times.EndTime = et.Unix()
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	ss.GetParam(listparam, ctx)
	list, err := siteSingleRecordBean.GetList(siteSingleRecord, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}
