//[控制器] [平台]优惠查询
package customer

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strconv"
	"time"
)

//优惠管理
type DiscountSearchController struct {
	controllers.BaseController
}

//获取优惠列表
func (c *DiscountSearchController) GetDiscountCountList(ctx echo.Context) error {
	//获取用户参数
	dc := new(input.DiscountSearchList)
	code := global.ValidRequestAdmin(dc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := discountSearchBean.DiscountSearchLIst(dc, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//获取优惠总计
func (c *DiscountSearchController) GetDiscountCount(ctx echo.Context) error {
	list := new(input.DiscountAllList)
	code := global.ValidRequestAdmin(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var startTime string
	var endTime string
	nextmonth := list.Month + 1
	if list.Month < 10 {
		startTime = strconv.Itoa(list.Year) + "-0" + strconv.Itoa(list.Month) + "-01 00:00:00"
		if list.Month == 9 {
			endTime = strconv.Itoa(list.Year) + "-" + strconv.Itoa(nextmonth) + "-01 00:00:00"
		} else {
			endTime = strconv.Itoa(list.Year) + "-0" + strconv.Itoa(nextmonth) + "-01 00:00:00"
		}
	} else {
		startTime = strconv.Itoa(list.Year) + "-" + strconv.Itoa(list.Month) + "-01 00:00:00"
		endTime = strconv.Itoa(list.Year) + "-" + strconv.Itoa(nextmonth) + "-01 00:00:00"
	}

	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	start_time, _ := time.ParseInLocation(timeLayout, startTime, loc)
	end_time, _ := time.ParseInLocation(timeLayout, endTime, loc)
	data, err := discountSearchBean.GetDiscountCount(list, start_time.Unix(), end_time.Unix())
	if err != nil {
		return ctx.JSON(200, global.ReplyError(90603, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取优惠明细
func (c *DiscountSearchController) GetDiscountList(ctx echo.Context) error {
	list := new(input.DiscountInfo)
	code := global.ValidRequestAdmin(list, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := discountSearchBean.GetDiscountList(list)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(90605, ctx))
	}
	info := make(map[int64]back.DiscountInfo)
	for _, v := range data {
		info_one := back.DiscountInfo{}
		product_list := back.DiscountProduct{}
		for _, value := range data {
			if v.Id == value.Id {
				info_one.Id = v.Id
				info_one.SiteIndexId = v.SiteIndexId
				info_one.SiteId = v.SiteId
				info_one.Account = v.Account
				info_one.AllMoney = v.AllMoney
				info_one.Betall = v.Betall
				info_one.LevelId = v.LevelId
				info_one.RebateWater = v.RebateWater
				info_one.MemberId = v.MemberId
				info_one.SelfMoney = v.SelfMoney
				info_one.Status = v.Status
				info_one.StartTime = v.StartTime
				info_one.EndTime = v.EndTime
				product_list.Money = v.Money
				product_list.ProductBet = v.ProductBet
				product_list.ProductId = v.ProductId
				product_list.Rate = v.Rate
				info_one.List = append(info_one.List, product_list)
			}
		}
		info[v.Id] = info_one
	}
	return ctx.JSON(200, global.ReplyItem(info))
}
