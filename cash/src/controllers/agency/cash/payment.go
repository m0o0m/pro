package cash

import (
	"controllers"
	"global"
	"models/back"
	"models/input"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

//入款银行设定/线上支付设定
type PaymentController struct {
	controllers.BaseController
}

//入款银行设定(get列表)
func (pc *PaymentController) BankIncomeList(ctx echo.Context) error {
	/*
		查询条件
		联表  site_bank_income_set 、sales_bank、site_bank_income_member_level
		统计数据
	*/
	//获取登陆人信息
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_bank := new(input.BankInList)
	pay_bank.SiteId = user.SiteId
	pay_bank.SiteIndexId = user.SiteIndexId
	code := global.ValidRequest(pay_bank, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	pc.GetParam(listparam, ctx)
	data, count, err := paymentBean.FindBankInList(pay_bank, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//添加一条入款银行设定(post)
func (pc *PaymentController) BankIncomeDo(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数

	SiteIndexId := ctx.FormValue("siteIndexId") //前台id
	if SiteIndexId == "" {
		return ctx.JSON(200, global.ReplyError(10050, ctx))
	}
	//会员层级id,用逗号隔开
	level := ctx.FormValue("level")
	if level == "" {
		return ctx.JSON(200, global.ReplyError(10014, ctx))
	}
	//银行id
	BankId, err := strconv.ParseInt(ctx.FormValue("bankId"), 10, 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//支付方式
	pay_id, err := strconv.ParseInt(ctx.FormValue("payId"), 10, 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//银行帐号
	Account := ctx.FormValue("account")
	if Account == "" {
		return ctx.JSON(200, global.ReplyError(50010, ctx))
	}
	//开户银行
	OpenBank := ctx.FormValue("openBank")
	if OpenBank == "" {
		return ctx.JSON(200, global.ReplyError(50102, ctx))
	}
	//收款人
	Payee := ctx.FormValue("payee")
	if Payee == "" {
		return ctx.JSON(200, global.ReplyError(50103, ctx))
	}
	//停用金额
	StopBalance, err := strconv.ParseFloat(ctx.FormValue("stopBalance"), 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//备注
	Remark := ctx.FormValue("remark")
	//状态
	Status, err := strconv.Atoi(ctx.FormValue("status"))
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	pay_bank := new(input.BankInAdd)
	if pay_id == 2 || pay_id == 3 {
		//二维码
		ErWeiMa, err := ctx.FormFile("erWeiMa")
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		ErWeiMaFile, err := global.ReadByte(ErWeiMa)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		pay_bank.QrCode = string(ErWeiMaFile)
	}
	pay_bank.Payee = Payee
	pay_bank.Account = Account
	pay_bank.StopBalance = StopBalance
	pay_bank.Status = Status
	pay_bank.Remark = Remark
	pay_bank.SiteIndexId = SiteIndexId
	pay_bank.SiteId = user.SiteId
	pay_bank.Level = strings.Split(level, ",")
	pay_bank.BankId = BankId
	pay_bank.OpenBank = OpenBank
	pay_bank.PayTypeId = pay_id

	//判断表中是否已经存在
	_, has, err := paymentBean.GetOnePay(pay_bank)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(200, global.ReplyError(50107, ctx))
	}
	//添加入款银行设定
	count, err := paymentBean.Add(pay_bank)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50174, ctx))
	}
	return ctx.NoContent(204)
}

//获取一条入款银行设定信息(get)
func (pc *PaymentController) BankIncome(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_bank := new(input.OneBankPaySet)
	code := global.ValidRequest(pay_bank, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, has, err := paymentBean.GetOneBankPaySet(pay_bank)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(6000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyItem(nil))
	}
	levelId, err := paymentBean.GetLelvelId(data.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(6000, ctx))
	}
	var levelIds []string
	for k := range levelId {
		levelIds = append(levelIds, levelId[k].LevelId)
	}
	oneBankPaySetBack := new(back.OneBankPaySetBack)
	oneBankPaySetBack.LevelId = levelIds
	oneBankPaySetBack.OneBankPaySet = data
	return ctx.JSON(200, global.ReplyItem(oneBankPaySetBack))
}

//修改一条入款银行设定信息(put)
func (pc *PaymentController) BankIncomeEdit(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	SiteIndexId := ctx.FormValue("siteIndexId") //前台id
	if SiteIndexId == "" {
		return ctx.JSON(200, global.ReplyError(10050, ctx))
	}
	//会员层级id,用逗号隔开
	level := ctx.FormValue("level")
	if level == "" {
		return ctx.JSON(200, global.ReplyError(10014, ctx))
	}
	//银行id
	BankId, err := strconv.ParseInt(ctx.FormValue("bankId"), 10, 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//支付方式
	pay_id, err := strconv.ParseInt(ctx.FormValue("payId"), 10, 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//银行帐号
	Account := ctx.FormValue("account")
	if Account == "" {
		return ctx.JSON(200, global.ReplyError(50010, ctx))
	}
	//开户银行
	OpenBank := ctx.FormValue("openBank")
	if OpenBank == "" {
		return ctx.JSON(200, global.ReplyError(50102, ctx))
	}
	//收款人
	Payee := ctx.FormValue("payee")
	if Payee == "" {
		return ctx.JSON(200, global.ReplyError(50103, ctx))
	}
	//停用金额
	StopBalance, err := strconv.ParseFloat(ctx.FormValue("stopBalance"), 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//备注
	Remark := ctx.FormValue("remark")
	//id
	id, err := strconv.ParseInt(ctx.FormValue("id"), 10, 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	pay_bank := new(input.BankInUpdata)
	if pay_id == 2 || pay_id == 3 {
		//二维码
		ErWeiMa, err := ctx.FormFile("erWeiMa")

		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}

		ErWeiMaFile, err := global.ReadByte(ErWeiMa)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		pay_bank.QrCode = string(ErWeiMaFile)
	}
	pay_bank.Id = id
	pay_bank.Payee = Payee
	pay_bank.Account = Account
	pay_bank.StopBalance = StopBalance
	pay_bank.Remark = Remark
	pay_bank.SiteIndexId = SiteIndexId
	pay_bank.SiteId = user.SiteId
	pay_bank.Level = strings.Split(level, ",")
	pay_bank.BankId = BankId
	pay_bank.OpenBank = OpenBank
	pay_bank.PayTypeId = pay_id
	//修改入款银行设定
	count, err := paymentBean.UpdataBankPaySet(pay_bank)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50173, ctx))
	}
	return ctx.NoContent(204)
}

//修改入款银行状态(put)
func (pc *PaymentController) BankIncomeStatus(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_bank := new(input.UpdataStatus)
	code := global.ValidRequest(pay_bank, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := paymentBean.ChangeStatus(pay_bank)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(30037, ctx))
	}
	return ctx.NoContent(204)
}

//删除一条入款银行
func (pc *PaymentController) BankIncomeDelete(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_bank := new(input.DeletePaySet)
	code := global.ValidRequest(pay_bank, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	count, err := paymentBean.DeteleOnePaySet(pay_bank)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30291, ctx))
	}
	return ctx.NoContent(204)
}

//存款记录
func (pc *PaymentController) DepositRecord(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_bank := new(input.DepositRecord)
	code := global.ValidRequest(pay_bank, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	pay_bank.SiteId = user.SiteId
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if pay_bank.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", pay_bank.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if pay_bank.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", pay_bank.EndTime, loc)
		times.EndTime = et.Unix()
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	pc.GetParam(listparam, ctx)
	data, count, err := paymentBean.CheckingDepositRecords(pay_bank, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//适用层级
func (pc *PaymentController) ApplicationLevel(ctx echo.Context) error {
	//获取登录用户资料
	user := ctx.Get("user").(*global.RedisStruct)
	user_level := user.Level
	if user_level > 1 {
		return ctx.JSON(200, global.ReplyError(50082, ctx))
	}
	//获取参数
	pay_bank := new(input.ApplicationLevel)
	code := global.ValidRequest(pay_bank, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	pay_bank.SiteId = user.SiteId
	data, err := paymentBean.TopClass(pay_bank)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
