package account

import (
	"controllers"
	"global"
	"models/input"

	"github.com/labstack/echo"
	"models/back"
	"strconv"
	"time"
)

//层级管理
type MemberLevelController struct {
	controllers.BaseController
}

//层级列表
func (mlc *MemberLevelController) Index(ctx echo.Context) error {
	levelIndex := new(input.LevelIndex)
	code := global.ValidRequest(levelIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	mlc.GetParam(listParams, ctx)
	data, count, err := memberLevelBean.List(levelIndex, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//新增层级
func (mlc *MemberLevelController) Add(ctx echo.Context) error {
	memberLevel := new(input.MemberLevel)
	code := global.ValidRequest(memberLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断当前站点，站点前台的层级名称是否出现重复
	_, have, err := memberLevelBean.LevelGet(memberLevel.LevelId, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10052, ctx))
	}
	//添加层级
	count, err := memberLevelBean.Add(memberLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(10012, ctx))
	}
	return ctx.NoContent(204)
}

//获取层级信息
func (mlc *MemberLevelController) Info(ctx echo.Context) error {
	memberLevInfo := new(input.LevelInfoGet)
	code := global.ValidRequest(memberLevInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, have, err := memberLevelBean.LevelGet(memberLevInfo.LevelId, memberLevInfo.SiteId, memberLevInfo.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10017, ctx))
	}
	//将得到的data里面时间戳转换成日期
	info := new(back.MemberLevel)
	info.DepositCount = data.DepositCount
	info.DepositNum = data.DepositNum
	info.Description = data.Description
	info.LevelId = data.LevelId
	info.Remark = data.Remark
	startTime, _ := strconv.ParseInt(data.StartTime, 10, 64)
	info.StartTime = time.Unix(startTime, 0).UTC().Format("2006-01-02 15:04:05")
	endTime, _ := strconv.ParseInt(data.EndTime, 10, 64)
	info.EndTime = time.Unix(endTime, 0).UTC().Format("2006-01-02 15:04:05")
	return ctx.JSON(200, global.ReplyItem(info))
}

//修改层级信息内容
func (mlc *MemberLevelController) InfoEdit(ctx echo.Context) error {
	memberLevel := new(input.MemberLevelUpdate)
	code := global.ValidRequest(memberLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断当前层级是否存在
	_, have, err := memberLevelBean.LevelGet(memberLevel.OldLevelId, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10017, ctx))
	}
	//判断当前层级是否是默认层级
	have, err = memberLevelBean.LevelDefault(memberLevel.OldLevelId, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if have {
		return ctx.JSON(200, global.ReplyError(10053, ctx))
	}
	//判断当前站点，站点前台的层级名称是否出现重复
	_, have, err = memberLevelBean.LevelGet(memberLevel.NewLevelId, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if have {
		if memberLevel.NewLevelId != memberLevel.OldLevelId {
			return ctx.JSON(200, global.ReplyError(10052, ctx))
		}
	}
	count, err := memberLevelBean.Update(memberLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//回归层级
func (mlc *MemberLevelController) Regress(ctx echo.Context) error {
	memberLevel := new(input.LevelInfoGet)
	code := global.ValidRequest(memberLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	_, have, err := memberLevelBean.LevelGet(memberLevel.LevelId, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10017, ctx))
	}

	count, err := memberLevelBean.ComeBackLevel(memberLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10023, ctx))
	}
	return ctx.NoContent(204)
}

//分层
func (mlc *MemberLevelController) Move(ctx echo.Context) error {
	memberLevel := new(input.MoveMemberLevel)
	code := global.ValidRequest(memberLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//移出分组不存在
	_, have, err := memberLevelBean.LevelGet(memberLevel.MoveOut, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10017, ctx))
	}
	//移入分组不合法
	_, have, err = memberLevelBean.LevelGet(memberLevel.MoveIn, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10017, ctx))
	}
	//判断移入分层名称和移出分层名称是否相同
	if memberLevel.MoveIn == memberLevel.MoveOut {
		return ctx.JSON(200, global.ReplyError(10057, ctx))
	}
	//若移入的层级是默认层级
	have, err = memberLevelBean.LevelDefault(memberLevel.MoveIn, memberLevel.SiteId, memberLevel.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if have {
		backLevel := &input.LevelInfoGet{
			LevelId:     memberLevel.MoveOut,
			SiteId:      memberLevel.SiteId,
			SiteIndexId: memberLevel.SiteIndexId}
		_, err = memberLevelBean.ComeBackLevel(backLevel)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.NoContent(204)
	}
	count, err := memberLevelBean.MoveLevel(memberLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(10056, ctx))
	}
	return ctx.NoContent(204)
}

//获取层级支付设定
func (mlc *MemberLevelController) PaySet(ctx echo.Context) error {
	payset := new(input.MemberLevelPaySet)
	code := global.ValidRequest(payset, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	payset.SiteId = user.SiteId
	data, has, err := memberLevelBean.MemberLevelPatSetOne(payset)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改层级支付设定
func (mlc *MemberLevelController) PaySetEdit(ctx echo.Context) error {
	payset := new(input.MemberLevelPaySetUpdata)
	code := global.ValidRequest(payset, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := memberLevelBean.UpdataMemberLevelPaySet(payset)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50177, ctx))
	}
	return ctx.NoContent(204)
}

//开启/关闭自助返水
func (mlc *MemberLevelController) StatusSelfRebate(ctx echo.Context) error {
	LevelSelfRebate := new(input.MemberLevelSelfRebate)
	code := global.ValidRequest(LevelSelfRebate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := memberLevelBean.SelfRebate(LevelSelfRebate)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50175, ctx))
	}
	return ctx.NoContent(204)
}

//会员详情(列表)
func (mlc *MemberLevelController) MemberList(ctx echo.Context) error {
	memberLevel := new(input.LevelMember)
	code := global.ValidRequest(memberLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParams := new(global.ListParams)
	mlc.GetParam(listParams, ctx)
	data, count, err := memberLevelBean.MemberListInfo(memberLevel, listParams)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listParams, data, int64(len(data)), count, ctx))
}

//锁定/解锁层级
func (mlc *MemberLevelController) Locked(ctx echo.Context) error {
	memberLevel := new(input.LockMember)
	code := global.ValidRequest(memberLevel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := memberLevelBean.LockMember(memberLevel)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50176, ctx))
	}
	return ctx.NoContent(204)
}

//获取层级下拉框
func (mlc *MemberLevelController) MemberLevelDrop(ctx echo.Context) error {
	levelIndex := new(input.LevelIndex)
	code := global.ValidRequest(levelIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := memberLevelBean.MemberLevelDrop(levelIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
