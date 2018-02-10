package cash

import (
	"controllers"
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"io/ioutil"
	"models/back"
	"models/input"
	"net/url"
	"strconv"
)

//入款银行剔除
//出款银行剔除
//第三方银行剔除
type BankCardController struct {
	controllers.BaseController
}

//入款银行(get列表)
func (bc *BankCardController) BankIncome(ctx echo.Context) error {
	//获取用户参数
	bank_card := new(input.InComeList)
	code := global.ValidRequest(bank_card, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	bank_card.SiteId = user.SiteId
	bank_card.IsIncome = 1
	listparam := new(global.ListParams)
	//获取listparam的数据
	bc.GetParam(listparam, ctx)
	data, count, err := bankCardBean.GetAllBank(bank_card, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//启用/剔除入款银行(put)
func (*BankCardController) BankIncomeStatus(ctx echo.Context) error {
	//获取用户参数
	bank_card := new(input.OpenAndRejectBank)
	code := global.ValidRequest(bank_card, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//剔除
	if bank_card.Status == 1 {
		count, err := bankCardBean.RejectBank(bank_card)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50178, ctx))
		}
	} else if bank_card.Status == 2 {
		//开启
		count, err := bankCardBean.OpenBank(bank_card)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50179, ctx))
		}
	}
	return ctx.NoContent(204)
}

//出款银行(get列表)
func (bc *BankCardController) BankOut(ctx echo.Context) error {
	bank_card := new(input.IsOutList)
	code := global.ValidRequest(bank_card, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	bank_card.SiteId = user.SiteId
	bank_card.Isout = 1
	listparam := new(global.ListParams)
	//获取listparam的数据
	bc.GetParam(listparam, ctx)
	data, count, err := bankCardBean.GetAllBankOut(bank_card, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))

}

//启用/剔除出款银行(put)
func (bc *BankCardController) BankOutStatus(ctx echo.Context) error {
	//获取用户参数
	bank_card := new(input.OpenAndRejectBank)
	code := global.ValidRequest(bank_card, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//剔除
	if bank_card.Status == 1 {
		count, err := bankCardBean.RejectOutBank(bank_card)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50178, ctx))
		}
	} else if bank_card.Status == 2 {
		//开启
		count, err := bankCardBean.OpenOutBank(bank_card)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50179, ctx))
		}
	}
	return ctx.NoContent(204)
}

//第三方银行(get列表)
func (bc *BankCardController) BankThird(ctx echo.Context) error {
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//所有三方银行
	newType := new(input.IsThirdList)
	code := global.ValidRequest(newType, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	paid := strconv.Itoa(newType.PaidType)
	uri := "http://olmanage.pk1358.com/api/v1/bank/list"
	postValues := url.Values{}
	postValues.Add("clientUserId", "1")
	postValues.Add("clientName", "pkClient")
	postValues.Add("clientSecret", "h1qN7RYH9xpvugZhaFu5Inmdk6bJyIopJrsbCAmj")
	postValues.Add("payId", paid)
	resp, err := client.PostForm(uri, postValues)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var newDataMap map[string]interface{}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//解析进map
		err = json.Unmarshal(result, &newDataMap)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//失败
	if newDataMap["code"].(float64) != float64(200) || newDataMap["status"] == false {
		return ctx.JSON(200, global.ReplyError(int64(newDataMap["code"].(float64)), ctx))
	}
	//没有数据
	if newDataMap["data"] == "" {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	newType.SiteId = user.SiteId
	third, err := bankCardBean.GetAllBankThird(newType)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	byteData, err := json.Marshal(newDataMap["data"])
	var sss []back.BankListBack
	err = json.Unmarshal(byteData, &sss)
	for k, v := range sss {
		if len(third) > 0 {
			for _, vv := range third {
				if v.Id == vv.BankId {
					sss[k].Status = 2
				}
				sss[k].Status = 1
			}
		} else {
			sss[k].Status = 1
		}
	}
	return ctx.JSON(200, global.ReplyItem(sss))
}

//启用/剔除第三方银行(put)
func (*BankCardController) BankThirdStatus(ctx echo.Context) error {
	//获取用户参数
	bank_card := new(input.OpenAndRejectBank)
	code := global.ValidRequest(bank_card, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//剔除
	if bank_card.Status == 1 {
		count, err := bankCardBean.RejectThirdBank(bank_card)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50178, ctx))
		}
	} else if bank_card.Status == 2 {
		//开启
		count, err := bankCardBean.OpenThirdBank(bank_card)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50179, ctx))
		}
	}
	return ctx.NoContent(204)
}

//站点出款银行下拉框
func (*BankCardController) BankAgencyOutByDrop(ctx echo.Context) error {
	//获取用户参数
	bankCard := new(input.AgencyBankOutByDrop)
	code := global.ValidRequest(bankCard, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	bankCard.SiteId = user.SiteId
	//根据登录人id获取index_id
	info, has, err := bankCardBean.SiteIndexIdByAgencyId(bankCard.Id, user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	bankCard.SiteIndexId = info.SiteIndexId
	//获取所有出款银行
	data, err := bankCardBean.BankCardListDrop()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取所有被剔除的银行
	dataDel, err := bankCardBean.BankCardListDropDel(bankCard)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var list []back.BankCardListDrop
	var li back.BankCardListDrop
	if len(data) > 0 {
		for _, v := range data {
			i := 0
			if len(dataDel) > 0 {
				for _, n := range dataDel {
					if v.Id == n.BankId {
						i = i + 1
					}
				}
				if i < 1 {
					li.Id = v.Id
					li.Title = v.Title
					list = append(list, li)
				}
			} else {
				list = data
			}

		}
	} else {
		list = data
	}
	return ctx.JSON(200, global.ReplyItem(list))
}
