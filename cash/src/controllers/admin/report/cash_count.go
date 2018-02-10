//[控制器] [平台]出入款汇总
package report

import (
	"bytes"
	"controllers"
	"encoding/json"
	"fmt"
	"framework/uuid"
	"github.com/labstack/echo"
	"github.com/tealeg/xlsx"
	"global"
	"models/back"
	"models/input"
	"sync"
)

const INCOME_KEY = "income"

//出入款统计管理
type CashCountController struct {
	controllers.BaseController
}

//出入款统计查询
func (c *CashCountController) GetCashCount(ctx echo.Context) error {
	reqData := new(input.CashCollect)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var sTime, eTime int64
	if reqData.STime != "" && reqData.ETime != "" {
		sTime, eTime, code = global.FormatDay2Timestamp(reqData.STime, reqData.ETime)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	} else {
		sTime = 0
		eTime = 0
	}
	//0公司入款1线上入款2人工存入3会员出款被扣金额4会员出款5给予优惠6人工提出7给予返水
	cashCollectDetails := make([]*back.CashCollectDetails, 8)
	wg := &sync.WaitGroup{}
	wg.Add(6)
	var cashRecordList []*back.CashCollectDetails
	safeErr := global.NewSafeError()
	//(0,1,2)人工入款,公司入款,线上入款
	go func() {
		defer wg.Done()
		var err error
		cashRecordList, err = cashCountBean.CashCollectCount(reqData.SiteId, &global.Times{StartTime: sTime, EndTime: eTime})
		if err != nil {
			safeErr.Push(err.Error())
		}
	}()
	//(3)会员出款被扣金额
	go func() {
		defer wg.Done()
		var err error
		cashCollectDetails[3], err = makeMoneyBean.GetTakeOutCount(reqData.SiteId, sTime, eTime)
		cashCollectDetails[3].SourceType = 3
		if err != nil {
			safeErr.Push(err.Error())
		}
	}()
	//(4)会员出款
	go func() {
		defer wg.Done()
		var err error
		cashCollectDetails[4], err = makeMoneyBean.GetMakeMoneyCount(reqData.SiteId, sTime, eTime)
		cashCollectDetails[4].SourceType = 4
		if err != nil {
			safeErr.Push(err.Error())
		}
	}()
	//(5)给予优惠  3个入款优惠加注册优惠
	go func() {
		defer wg.Done()
		var err error
		cashCollectDetails[5], err = cashRecordBean.GetDiscountCount(reqData.SiteId, sTime, eTime)
		cashCollectDetails[5].SourceType = 5
		if err != nil {
			safeErr.Push(err.Error())
		}
	}()
	//(6)人工出款
	go func() {
		defer wg.Done()
		var err error
		cashCollectDetails[6], err = manualAccessBean.GetOutCount(reqData.SiteId, sTime, eTime)
		cashCollectDetails[6].SourceType = 6
		if err != nil {
			safeErr.Push(err.Error())
		}
	}()
	//(7)给予返水(优惠返水)
	go func() {
		defer wg.Done()
		var err error
		cashCollectDetails[7], err = discountCountBean.GetDiscountCount(reqData.SiteId, &global.Times{StartTime: sTime, EndTime: eTime})
		cashCollectDetails[7].SourceType = 7
		if err != nil {
			safeErr.Push(err.Error())
		}
	}()
	wg.Wait()
	if safeErr.IsValid() {
		global.GlobalLogger.Error("error:%s", safeErr.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//0公司入款1线上入款2人工存入
	cashCollectDetails[0] = &back.CashCollectDetails{SourceType: 0}
	cashCollectDetails[1] = &back.CashCollectDetails{SourceType: 1}
	cashCollectDetails[2] = &back.CashCollectDetails{SourceType: 2}
	for k, _ := range cashRecordList {
		cashCollectDetails[cashRecordList[k].SourceType] = cashRecordList[k]
	}
	var cashCollect back.CashCollect
	cashCollect.Details = cashCollectDetails
	//总收入
	cashCollect.IncomeSum = global.FloatReserve2(
		cashCollectDetails[0].Money +
			cashCollectDetails[1].Money +
			cashCollectDetails[2].Money +
			cashCollectDetails[3].Money)
	//总支出
	cashCollect.OutlaySum = global.FloatReserve2(
		cashCollectDetails[4].Money +
			cashCollectDetails[5].Money +
			cashCollectDetails[6].Money +
			cashCollectDetails[7].Money)
	//实际盈亏
	cashCollect.ProfitLoss = global.FloatReserve2(cashCollect.IncomeSum - cashCollect.OutlaySum)
	//账目统计
	cashCollect.Total = global.FloatReserve2(cashCollect.IncomeSum + cashCollect.OutlaySum)
	key, err := c.saveRedisData(&cashCollect)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("content", cashCollect, "key", key))
}

//导出统计表格
func (c *CashCountController) PutCashCountData(ctx echo.Context) error {
	reqData := new(input.CashCollectExport)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	cashCollect, err := c.getRedisData(reqData.Key)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if cashCollect.Details == nil || len(cashCollect.Details) != 8 {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	js, _ := json.Marshal(cashCollect)
	fmt.Println(string(js))
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "收入"
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "公司入款"
	cell = row.AddCell()
	cell.Value = "线上支付"
	cell = row.AddCell()
	cell.Value = "人工存入"
	cell = row.AddCell()
	cell.Value = "会员出款被扣金额"
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[2].GetString()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[0].GetString()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[1].GetString()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[3].GetString()
	row = sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell = row.AddCell()
	cell.Value = "收入总计:" + fmt.Sprintf("%10.2f", cashCollect.IncomeSum)
	row = sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell = row.AddCell()
	cell.Value = "支出"
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "会员出款"
	cell = row.AddCell()
	cell.Value = "给予优惠"
	cell = row.AddCell()
	cell.Value = "人工提出"
	cell = row.AddCell()
	cell.Value = "给予反水"
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[4].GetString()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[5].GetString()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[6].GetString()
	cell = row.AddCell()
	cell.Value = cashCollect.Details[7].GetString()
	row = sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell = row.AddCell()
	cell.Value = "支出总计:" + fmt.Sprintf("%10.2f", cashCollect.OutlaySum)
	row = sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell = row.AddCell()
	cell.Value = "实际盈亏:" + fmt.Sprintf("%10.2f", cashCollect.ProfitLoss)
	cell = row.AddCell()
	cell.Value = "账目统计:" + fmt.Sprintf("%10.2f", cashCollect.Total)

	var b bytes.Buffer
	file.Write(&b)
	return ctx.Stream(200, "application/vnd.ms-excel", &b)
}

//将汇总数据存入到redis
func (*CashCountController) saveRedisData(src *back.CashCollect) (key string, err error) {
	key = uuid.NewV4().String()
	js, err := json.Marshal(src)
	if err != nil {
		return
	}
	global.GetRedis().HSet(INCOME_KEY, key, js)
	return
}

//从redis中取出汇总数据
func (c *CashCountController) getRedisData(key string) (dst back.CashCollect, err error) {
	src, err := global.GetRedis().HGet(INCOME_KEY, key).Result()
	if err != nil {
		global.GlobalLogger.Error("%s", err.Error())
		return dst, err
	}
	err = json.Unmarshal([]byte(src), &dst)
	if err != nil {
		global.GlobalLogger.Error("%s", err.Error())
		return dst, err
	}
	return dst, err
}
