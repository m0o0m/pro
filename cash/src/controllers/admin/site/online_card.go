//[控制器] [平台]第三方列表管理   只做列表 其他独立
package site

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"controllers"
	"global"
	"models/input"
	"models/schema"

	"github.com/labstack/echo"
)

var client = &http.Client{}
var TotalConfig *config.Config

//银行列表管理
type OnlineCardController struct {
	controllers.BaseController
}

//第三方列表查询
func (c *OnlineCardController) GetOnlineCardList(ctx echo.Context) error {
	listparam := new(global.ListParams)
	//获取listparam的数据
	c.GetParam(listparam, ctx)
	list, count, err := onlineIncomeThirdBean.GetThirdNeedData(listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//第三方添加
func (c *OnlineCardController) PostOnlineCardAdd(ctx echo.Context) error {
	onlineCard := new(input.AddThird)
	code := global.ValidRequestAdmin(onlineCard, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := onlineIncomeThirdBean.Add(onlineCard)
	if err != nil || count != 1 {
		return ctx.JSON(200, global.ReplyError(10248, ctx))
	}
	return ctx.NoContent(204)
}

//第三方状态修改【只修改了本地，第三方的这边是没有权限修改的】
func (c *OnlineCardController) PutOnlineCardStatusUpdate(ctx echo.Context) error {
	changeStatus := new(input.UpStatus)
	code := global.ValidRequestAdmin(changeStatus, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	_, err := onlineIncomeThirdBean.ChangeThirdStatus(changeStatus)
	if err != nil {
		global.GlobalLogger.Error("error:s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//第三方删除
func (c *OnlineCardController) PutOnlineCardDel(ctx echo.Context) error {
	onlineCard := new(input.DelThird)
	code := global.ValidRequestAdmin(onlineCard, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断Id是否存在
	_, have, err := onlineIncomeThirdBean.GetInfoById(onlineCard.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10246, ctx))
	}
	count, err := onlineIncomeThirdBean.Del(onlineCard)
	if err != nil || count != 1 {
		return ctx.JSON(200, global.ReplyError(10247, ctx))
	}
	return ctx.NoContent(204)
}

//同步第三方列表同步到缓存redis
func (c *OnlineCardController) PostOnlineCardRedis(ctx echo.Context) error {
	postValues := url.Values{}
	postValues.Add("clientUserId", TotalConfig.Third.ClientUserId)
	postValues.Add("clientName", TotalConfig.Third.ClientName)
	postValues.Add("clientSecret", TotalConfig.Third.ClientSecret)
	resp, err := client.PostForm(TotalConfig.Third.GetThirdApi, postValues)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var newDatas []schema.OnlineIncomeThird
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
		if newDataMap["code"].(float64) != float64(200) || newDataMap["status"] == false {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}

		//map解析
		byteData, err := json.Marshal(newDataMap["data"])
		if err != nil {
			global.GlobalLogger.Error("error:%s", err)
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//插入数据库的解析
		err = json.Unmarshal(byteData, &newDatas)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}

		//判断数据库中间是否存在数据，存在更新，不存在插入
		flag, err := onlineIncomeThirdBean.IsExistData()
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if flag {
			//同步到redis
			err = synchronizToRedis(result)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			//采用删除所有然后再次插入
			count, err := onlineIncomeThirdBean.DelAllData(newDatas)
			if err != nil || count == 0 {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		} else {
			//同步到redis
			err = synchronizToRedis(result)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			//插入
			count, err := onlineIncomeThirdBean.BatchInsertData(newDatas)
			if err != nil || count == 0 {
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
	} else {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//将从第三方获取的数据同步缓存到redis
func synchronizToRedis(thirdData []byte) error {
	var err error
	//缓存到redis
	err = global.GetRedis().Set(TotalConfig.RedisThirdSet, string(thirdData), 0).Err()
	return err
}

//三方网银信息修改
func (c *OnlineCardController) UpdateOnlineCard(ctx echo.Context) error {
	onlineCard := new(input.UpdateThird)
	code := global.ValidRequestAdmin(onlineCard, ctx)
	fmt.Printf("%+v\n", onlineCard)
	fmt.Println("code:", code)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//判断Id是否存在
	_, have, err := onlineIncomeThirdBean.GetInfoById(onlineCard.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10246, ctx))
	}
	count, err := onlineIncomeThirdBean.UpdateInfo(onlineCard)
	if err != nil || count != 1 {
		return ctx.JSON(200, global.ReplyError(10247, ctx))
	}
	return ctx.NoContent(204)
}
