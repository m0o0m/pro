package cash

import (
	"controllers"
	"encoding/json"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"time"
)

const (
	RETREAT_KEY = "retreat"
)

//优惠统计
type BetReportAccountController struct {
	controllers.BaseController
}

//优惠统计
func (brac *BetReportAccountController) CountBetReportAccount(ctx echo.Context) error {
	countBet := new(input.CountBetReportAccount)
	code := global.ValidRequest(countBet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var err error
	countBet.StartTime, countBet.EndTime, code = global.FormatDay2Timestamp(countBet.StartTimeStr, countBet.EndTimeStr)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if countBet.EndTime-countBet.StartTime != 0 { //不能跨天
		return ctx.JSON(500, global.ReplyError(10097, ctx))
	}
	if time.Now().Unix()-countBet.EndTime < 86400 { //不能为当天
		return ctx.JSON(500, global.ReplyError(10098, ctx))
	}
	countInfo, err := betReportAccountBean.CountBetReportAccount(countBet)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//存入redis
	key, err := brac.saveRedisData(&countInfo)
	if err != nil {
		global.GlobalLogger.Error("redis error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("data", countInfo, "key", key))
	//return ctx.JSON(200, global.ReplyItem(countInfo))
}

//优惠统计-存入
func (brac *BetReportAccountController) StoreBetReportAccount(ctx echo.Context) error {
	storeBet := new(input.StoreBetReportAccount)
	code := global.ValidRequest(storeBet, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(storeBet.MemberIds) == 0 {
		return ctx.JSON(200, global.ReplyError(10096, ctx))
	}
	retreatCommission, err := brac.getRedisData(storeBet.Key)
	user := ctx.Get("user").(*global.RedisStruct)
	err = betReportAccountBean.StoreBetReportAccount(storeBet, retreatCommission, user)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//将统计出的数据存入到redis
func (brac *BetReportAccountController) saveRedisData(src *back.CountAllBetReportAccountTotalMap) (key string, err error) {
	key = uuid.NewV4().String()
	js, err := json.Marshal(src)
	if err != nil {
		return
	}
	global.GetRedis().HSet(RETREAT_KEY, key, js)
	return
}

//从redis中取出统计信息
func (brac *BetReportAccountController) getRedisData(key string) (dst back.CountAllBetReportAccountTotalMap, err error) {
	src, err := global.GetRedis().HGet(RETREAT_KEY, key).Result()
	if err != nil {
		global.GlobalLogger.Error("%s", err.Error())
		return
	}
	err = json.Unmarshal([]byte(src), &dst)
	if err != nil {
		global.GlobalLogger.Error("%s", err.Error())
	}
	return
}
