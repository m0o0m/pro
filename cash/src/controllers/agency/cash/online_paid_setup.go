package cash

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"controllers"
	"global"
	"models/input"
	"models/schema"

	"config"
	"fmt"
	"github.com/labstack/echo"
	"models/back"
)

var client = &http.Client{}
var ConfigThird config.ThirdInterface

type OnlinePaidSetupController struct {
	controllers.BaseController
}

//func OnlinePaidSetupList() 线上支付设定列表
func (oc *OnlinePaidSetupController) OnlinePaidSetupList(ctx echo.Context) error {
	thirdList := new(input.ThirdList)
	code := global.ValidRequest(thirdList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	oc.GetParam(listparam, ctx)
	infolist, count, err := onlinePaidSetupBean.PaidList(thirdList, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(infolist, count))
}

//func StopThisOnlineSetup() 停用该线上支付设定
func (*OnlinePaidSetupController) StopThisOnlineSetup(ctx echo.Context) error {
	stopPaidSetup := new(input.StopThisPaidSetup)
	code := global.ValidRequest(stopPaidSetup, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//判断是否存在
	_, flag, err := onlinePaidSetupBean.GetOnePaidSet(stopPaidSetup.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60023, ctx))
	}
	//停用
	count, err := onlinePaidSetupBean.StopThirdPaid(stopPaidSetup)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(30037, ctx))
	}
	return ctx.NoContent(204)
}

//func GetThisOnlineDepositRecord() 获取该支付设定的存款记录
func (oc *OnlinePaidSetupController) GetThisOnlineDepositRecord(ctx echo.Context) error {
	infoSetup := new(input.GetInfoSetupDeposit)
	code := global.ValidRequest(infoSetup, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	oc.GetParam(listparam, ctx)
	infolist, count, err := onlineEntryRecordBean.GetRecordByPaidId(infoSetup, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, infolist, int64(len(infolist)), count, ctx))
}

//func DelThisPaidSetup() 删除该支付设定
func (*OnlinePaidSetupController) DelThisPaidSetup(ctx echo.Context) error {
	delPaidSetup := new(input.DelThisPaidSetup)
	code := global.ValidRequest(delPaidSetup, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//判断是否存在
	_, flag, err := onlinePaidSetupBean.GetOnePaidSet(delPaidSetup.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(200, global.ReplyError(60023, ctx))
	}
	//删除
	count, err := onlinePaidSetupBean.DelThirdPaid(delPaidSetup)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//func AddNewOnlinePaidSetup() 新增加线上支付设定[给的五个接口第三个]
func (*OnlinePaidSetupController) AddNewOnlinePaidSetup(ctx echo.Context) error {
	newPaid := new(input.NewThirdPayOnline)
	code := global.ValidRequest(newPaid, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//判断是否存在该设定【支付类型和商户号判断】
	info, flag, err := onlinePaidSetupBean.GetInfoBy(newPaid)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if flag || info.Id != 0 {
		return ctx.JSON(200, global.ReplyError(60219, ctx))
	}

	newPlatForm := strconv.Itoa(newPaid.PaidPlatform)
	newPayType := strconv.Itoa(newPaid.PaidType)
	isApp := strconv.Itoa(newPaid.IsApp)
	postValues := url.Values{}
	postValues.Add("clientUserId", ConfigThird.ClientUserId)
	postValues.Add("clientName", ConfigThird.ClientName)
	postValues.Add("clientSecret", ConfigThird.ClientSecret)
	postValues.Add("agentLine", newPaid.SiteId)         //代理线
	postValues.Add("subAgentLine", newPaid.SiteIndexId) //子代理线
	postValues.Add("notifyUrl", newPaid.BackAddress)    //回调域名
	postValues.Add("merchantId", newPaid.MerchatId)     //商户id[这个更改过后叫商户号,新的商户id是merid]
	postValues.Add("payId", newPlatForm)                //平台
	postValues.Add("payType", newPayType)               //支付类型
	postValues.Add("privateKey", newPaid.PrviateKey)    //私钥
	postValues.Add("publicKey", newPaid.PublicKey)      //公钥
	postValues.Add("levelId", newPaid.LevelId)          //层级
	postValues.Add("code", newPaid.PaidCode)            //支付编码
	postValues.Add("merUrl", newPaid.MerUrl)            //自填写网关
	postValues.Add("isApp", isApp)                      //是否跳转app(1.可以2.不可以)

	resp, err := client.PostForm(ConfigThird.NewSetup, postValues)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	defer resp.Body.Close()
	var newDatas input.AddSetupParse
	if resp.StatusCode == 200 {
		//读取返回
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//将返回结果解析到map
		var newDataMap map[string]interface{}
		err = json.Unmarshal(result, &newDataMap)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}

		//失败
		if newDataMap["status"] == false || newDataMap["code"].(float64) != float64(200) {
			global.GlobalLogger.Error("error:%s,code:%v", "Third-party interface requests failed", newDataMap["code"])
			return ctx.JSON(200, global.ReplyError(60042, ctx))
		}
		//返回数据为空
		if newDataMap["data"] == "" {
			global.GlobalLogger.Error("error:%s,code:%v", "Third-party interface requests failed", newDataMap["code"])
			return ctx.JSON(200, global.ReplyError(60042, ctx))
		}
		//将map中间的数据解析到struct
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
		//成功之后插入本地数据库
		count, err := onlinePaidSetupBean.AddNew(newPaid, &newDatas)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50174, ctx))
		}
	}
	return ctx.NoContent(204)
}

//func GetThirdPayList() 获取第三方平台(给的五个接口第四个)
func (*OnlinePaidSetupController) GetThirdPayList(ctx echo.Context) error {
	postValues := url.Values{}
	postValues.Add("clientUserId", ConfigThird.ClientUserId)
	postValues.Add("clientName", ConfigThird.ClientName)
	postValues.Add("clientSecret", ConfigThird.ClientSecret)
	resp, err := client.PostForm(ConfigThird.GetThirdApi, postValues)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var newData []back.BackSelectThird
	var backThirdData []back.BackSelectThirdJson
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
		//将map中间的数据解析到struct
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
		//返回数据库的解析
		err = json.Unmarshal(byteData, &backThirdData)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//判断数据库中间是否存在数据，存在更新，不存在插入
		flag, err := onlineIncomeThirdBean.IsExistData()
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if flag { //刷新
			//采用删除所有然后再次插入
			count, err := onlineIncomeThirdBean.DelAllData(newDatas)
			if err != nil || count == 0 {
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		} else {
			//插入
			count, err := onlineIncomeThirdBean.BatchInsertData(newDatas)
			if err != nil || count == 0 {
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		return ctx.JSON(200, global.ReplyItem(backThirdData))
	} else {
		//失败,没有数据从本地数据库获取
		newData, err = onlineIncomeThirdBean.GetThirdNameAndId()
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(newData))
	}
}

//func ChangeOnlineSetup() 修改该线上支付设定
func (*OnlinePaidSetupController) ChangeOnlineSetup(ctx echo.Context) error {
	//角色权限认证[开户人和平台管理员可以操作]首先判断本地是存在该设定，才能修改线上和本地数据库
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//数据获取
	newPaid := new(input.ChangeThisPaidSetup)
	code := global.ValidRequest(newPaid, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//是否存在
	info, flag, err := onlinePaidSetupBean.GetOnePaidSet(newPaid.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//本地不存在
	if !flag || info.Id == 0 {
		return ctx.JSON(200, global.ReplyError(60023, ctx))
	}
	//请求线上
	newPlatForm := strconv.Itoa(newPaid.PayId)
	newPayType := strconv.Itoa(newPaid.PayType)
	isApp := strconv.Itoa(newPaid.IsApp)
	id := strconv.Itoa(newPaid.Id)
	postValues := url.Values{}
	postValues.Add("clientUserId", ConfigThird.ClientUserId)
	postValues.Add("clientName", ConfigThird.ClientName)
	postValues.Add("clientSecret", ConfigThird.ClientSecret)
	postValues.Add("notifyUrl", newPaid.BackAddress)    //回调域名
	postValues.Add("merId", id)                         //标识我是要修改的字段(这个值是表的主键id，不传值代表新增加，传值代表修改)
	postValues.Add("payId", newPlatForm)                //平台
	postValues.Add("agentLine", newPaid.Site)           //代理线
	postValues.Add("merchatId", newPaid.MerchatId)      //商户id
	postValues.Add("subAgentLine", newPaid.SiteIndexId) //子代理线
	postValues.Add("payType", newPayType)               //支付类型
	postValues.Add("privateKey", newPaid.PrivateKey)    //私钥
	postValues.Add("publicKey", newPaid.PublicKey)      //公钥
	postValues.Add("levelId", newPaid.FitforLevel)      //层级
	postValues.Add("code", newPaid.PaidCode)            //支付编码
	postValues.Add("merUrl", newPaid.MerUrl)            //自填写网关
	postValues.Add("isApp", isApp)                      //是否跳转app(1.可以2.不可以)

	resp, err := client.PostForm(ConfigThird.NewSetup, postValues)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode, "response code")
	if resp.StatusCode == 200 {
		//读取返回
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//将返回结果解析到map
		var newDataMap map[string]interface{}
		err = json.Unmarshal(result, &newDataMap)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//失败
		if newDataMap["status"] == false || newDataMap["code"].(float64) != float64(200) {
			global.GlobalLogger.Error("error:%s,code:%v", "Third-party interface requests failed", newDataMap["code"])
			return ctx.JSON(200, global.ReplyError(60313, ctx))
		}
		//返回数据为空
		if newDataMap["data"] == "" {
			global.GlobalLogger.Error("error:%s,code:%v", "Third-party interface requests failed", newDataMap["code"])
			return ctx.JSON(200, global.ReplyError(60313, ctx))
		}
		//修改本地
		count, err := onlinePaidSetupBean.ChangeSet(newPaid)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(500, global.ReplyError(50173, ctx))
		}
	} else {
		global.GlobalLogger.Error("error:%s", "The resp.StatusCode is not 200,Requesting third is failed")
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//func GetPaidSetup() 获取线上支付设定
func (*OnlinePaidSetupController) GetPaidSetup(ctx echo.Context) error {
	infoSetup := new(input.GetInfoSetup)
	code := global.ValidRequest(infoSetup, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	////判断本地是否存在或者已经删除
	//info, flag, err := onlinePaidSetupBean.GetOnePaidSet(infoSetup.Id)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//if !flag {
	//	return ctx.JSON(200, global.ReplyError(60023, ctx))
	//}
	////本地存在线上拉取选择性的更新本地，返回给前台
	//paidType := strconv.Itoa(info.PaidType)
	//paidId := strconv.Itoa(info.PaidPlatform)
	//merId := strconv.Itoa(info.Id)
	//status := strconv.Itoa(info.Status)
	//postValues := url.Values{}
	//postValues.Add("clientUserId", ConfigThird.ClientUserId)
	//postValues.Add("clientName", ConfigThird.ClientName)
	//postValues.Add("clientSecret", ConfigThird.ClientSecret)
	//postValues.Add("agentLine", info.SiteId)         //代理线
	//postValues.Add("subAgentLine", info.SiteIndexId) //子代理线
	//postValues.Add("merchantId", info.MerchatId)     //商户号
	//postValues.Add("payId", paidId)                  //平台
	//postValues.Add("merId", merId)                   //merid
	//postValues.Add("payState", status)               //状态
	//postValues.Add("payType", paidType)              //支付类型
	//
	//resp, err := client.PostForm(ConfigThird.GetSetup, postValues)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return ctx.JSON(200, global.ReplyError(60000, ctx))
	//}
	//defer resp.Body.Close()
	//var backInfo schema.OnlinePaidSetup
	//var newDatas input.OnlinePaidSetParse
	//if resp.StatusCode == 200 {
	//	//读取返回
	//	result, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	//将返回结果解析到map
	//	var newDataMap map[string]interface{}
	//	err = json.Unmarshal(result, &newDataMap)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	//解析到刷新数据库的struct
	//	byteData, err := json.Marshal(newDataMap["data"])
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err)
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	err = json.Unmarshal(byteData, &newDatas)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	//失败
	//	if newDataMap["status"] == false || newDataMap["code"].(float64) != float64(200) {
	//		global.GlobalLogger.Error("error:%s,code:%d", "Third-party interface requests failed", newDataMap["code"])
	//		return ctx.JSON(200, global.ReplyError(60042, ctx))
	//	}
	//	//返回数据为空
	//	if newDataMap["data"] == "" {
	//		global.GlobalLogger.Error("error:%s,code:%d", "Third-party interface requests failed", newDataMap["code"])
	//		return ctx.JSON(200, global.ReplyError(60042, ctx))
	//	}
	//	//成功之后将数据选择性的更新本地数据库
	//	count, err := onlinePaidSetupBean.UpdateLocalData(&newDatas)
	//	if err != nil || count != 1 {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//	//然后重新获取返回给前台
	//	backInfo, _, err = onlinePaidSetupBean.GetOnePaidSet(infoSetup.Id)
	//	if err != nil {
	//		global.GlobalLogger.Error("error:%s", err.Error())
	//		return ctx.JSON(500, global.ReplyError(60000, ctx))
	//	}
	//} else {
	//	global.GlobalLogger.Error("error:%s", "The resp.StatusCode is not 200,Resquesting third interface is failed")
	//	return ctx.JSON(500, global.ReplyError(60000, ctx))
	//}
	//如果该接口需要请求第三方然后刷新本地的话，打开上面的代码，将下面的代码注释
	backInfo, flag, err := onlinePaidSetupBean.GetOnePaidSet(infoSetup.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag || backInfo.Id == 0 {
		return ctx.JSON(200, global.ReplyError(60023, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(backInfo))
}

//func GetOnlinePaidType() 获取线上支付类型
func (*OnlinePaidSetupController) GetOnlinePaidType(ctx echo.Context) error {
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	postValues := url.Values{}
	postValues.Add("clientUserId", ConfigThird.ClientUserId)
	postValues.Add("clientName", ConfigThird.ClientName)
	postValues.Add("clientSecret", ConfigThird.ClientSecret)
	resp, err := client.PostForm(ConfigThird.PaidType, postValues)
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
		//失败
		if newDataMap["code"].(float64) != float64(200) || newDataMap["status"] == false {
			return ctx.JSON(200, global.ReplyError(int64(newDataMap["code"].(float64)), ctx))
		}
		//没有数据
		if newDataMap["data"] == "" {
			return ctx.JSON(200, global.ReplyItem(nil))
		}
		//将map中间的数据解析出来
		byteData, err := json.Marshal(newDataMap["data"])
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		var newData []schema.PaidType
		err = json.Unmarshal(byteData, &newData)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//首先判断数据库是否存在数据
		flag, err := paidTypeBean.ExistDataPaid()
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !flag { //不存在数据
			_, err := paidTypeBean.AddNew(newData)
			if err != nil {
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		} else {
			//存在先删除，然后插入
			_, err := paidTypeBean.DelAllData(newData)
			if err != nil {
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		}
		return ctx.JSON(200, global.ReplyCollection(newData, int64(len(newData))))
	} else {
		//如果请求第三方失败,获取数据库的数据
		infolist, err := paidTypeBean.GetTbaleData()
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(infolist))
	}
}

//func GetPaidBank() 获取某个支付方式下面支持的银行卡
func (*OnlinePaidSetupController) GetPaidBank(ctx echo.Context) error {
	newType := new(input.GetBank)
	code := global.ValidRequest(newType, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	paid := strconv.Itoa(newType.PaidType)
	postValues := url.Values{}
	postValues.Add("clientUserId", ConfigThird.ClientUserId)
	postValues.Add("clientName", ConfigThird.ClientName)
	postValues.Add("clientSecret", ConfigThird.ClientSecret)
	postValues.Add("payId", paid)
	resp, err := client.PostForm(ConfigThird.GetBank, postValues)
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
