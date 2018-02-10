package cash

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
)

//退佣查询
type RetirementController struct {
	controllers.BaseController
}

//退佣查询列表(get列表)
func (pc *RetirementController) GetList(ctx echo.Context) error {
	combo := new(input.PeriodsGet)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取listparam的数据
	data1 := pc.getSiteList(combo.SiteId, combo.SiteIndexId)
	list, err := retirementBean.RetirementList(combo)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10120, ctx))
	}
	return ctx.JSON(200, global.ReplyCollections("list", list, "data", data1))
}

//退佣查询
func (pc *RetirementController) CheckList(ctx echo.Context) error {
	combo := new(input.CheckList)
	code := global.ValidRequest(combo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//fmt.Println(combo)
	data, list, err := retirementBean.CheckList(combo)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10119, ctx))
	}
	var data3 = new(back.RetirementCheckList)
	var data1 []float64
	var data2 []float64
	for _, v := range data {
		data3.Id = v.Id
		data3.SiteId = v.SiteId
		data3.PeriodsId = v.PeriodsId
		data3.AgencyAccount = v.AgencyAccount
		data3.EffectiveMember = v.EffectiveMember
		data3.BeforeProfit = v.BeforeProfit
		data3.BeforeProfit = v.NowProfit
		data3.BeforeBetting = v.BeforeBetting
		data3.NowBetting = v.NowBetting
		data3.BeforeCost = v.BeforeCost
		data3.BeforeProfit = v.BeforeProfit
		data3.NowCost = v.NowCost
		data3.Rebate = v.Rebate
		data3.RebateWater = v.RebateWater
		data3.Remark = v.Remark
		data1 = append(data1, v.RebateRatio)
		data2 = append(data2, v.WaterRatio)
	}
	data3.RebateRatio = data2
	data3.WaterRatio = data1

	return ctx.JSON(200, global.ReplyCollections("list", list, "data", data3))
}

//获取站点列表
func (pc *RetirementController) getSiteList(SiteId string, SiteIndexId string) []back.SiteList {
	list := new(input.GetSiteList)
	list.SiteId = SiteId
	list.SiteIndexId = SiteIndexId
	data, err := periodsBean.GetSiteList(list.SiteId, list.SiteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	return data
}
