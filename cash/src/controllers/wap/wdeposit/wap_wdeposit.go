package wdeposit

import (
	"config"
	"encoding/json"
	"github.com/labstack/echo"
	"global"
	"io/ioutil"
	"models/function"
	"models/input"
	"net/http"
	"net/url"
	"strconv"
)

var (
	sitePaySetBean          = new(function.SitePaySetBean)          //站点支付设定
	memberLevelBean         = new(function.MemberLevelBean)         //会员层级
	memberCompanyIncomeBean = new(function.MemberCompanyIncomeBean) //公司入款
	memberBean              = new(function.MemberBean)              //会员
	agencyBean              = new(function.AgencyBean)              //代理
	onlineEntryRecordBean   = new(function.OnlineEntryRecordBean)   //线上入款纪录
	onlineDepositBean       = new(function.OnlineDeposit)           //线上存款所需的表结构
	bankCardBean            = new(function.BankCardBean)
)

var client = &http.Client{}
var TotalConfig *config.Config //配置文件全局变量

type WapDepositController struct{}

//获取某个支付类型下面支持的银行卡
func (*WapDepositController) GetPaidBank(ctx echo.Context) error {
	newType := new(input.GetBank)
	code := global.ValidRequest(newType, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	paid := strconv.Itoa(newType.PaidType)
	postValues := url.Values{}
	postValues.Add("clientUserId", TotalConfig.Third.ClientUserId)
	postValues.Add("clientName", TotalConfig.Third.ClientName)
	postValues.Add("clientSecret", TotalConfig.Third.ClientSecret)
	postValues.Add("payId", paid)
	resp, err := client.PostForm(TotalConfig.Third.GetBank, postValues)
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
	return ctx.JSON(200, global.ReplyItem(newDataMap["data"]))
}

//获取会员对应层级对应配置的支付设定
func (c *WapDepositController) GetPaySetData(ctx echo.Context) error {
	member := new(input.MemberLevelPaySetAndFirstDeposit)
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := memberLevelBean.MemberLevelPatSetOne(&input.MemberLevelPaySet{member.SiteId, member.SiteIndexId, member.LevelId})
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	paySet, have, err := sitePaySetBean.GetPaySetInfoById(data.PaySetId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !have {
		return ctx.JSON(200, global.ReplyError(10156, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(paySet))
}
