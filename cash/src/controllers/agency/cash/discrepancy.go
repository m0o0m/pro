package cash

import (
	"time"

	"controllers"
	"global"
	"models/back"
	"models/function"
	"models/input"

	"github.com/labstack/echo"
)

type CashController struct {
	controllers.BaseController
}

//添加一条人工存款(post)
func (*CashController) ManualAccessDo(ctx echo.Context) error {
	manualAccess := new(input.ManualAccessAdd)
	code := global.ValidRequest(manualAccess, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//存款金额的小数点后是不是保留两位
	//codes := global.KeepTwoFloat(manualAccess.Money)
	//if codes != 0 {
	//	return ctx.JSON(200, global.ReplyError(code, ctx))
	//}
	//获取操作人信息
	user := ctx.Get("user").(*global.RedisStruct)
	manualAccess.DoAgencyId = user.Id
	manualAccess.DoAgencyAccount = user.Account
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//会员账号是否存在
	has, err := manualAccessBean.GetMemberInfo(manualAccess.Account, user.SiteId, user.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	}
	count, err := manualAccessBean.Add(manualAccess)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30123, ctx))
	}
	return ctx.NoContent(204)
}

//添加多条人工存款(post)
func (*CashController) ManualAccessBatchDo(ctx echo.Context) error {
	manualAccess := new(input.AddManualAccess)
	code := global.ValidRequest(manualAccess, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//存款金额的小数点后是不是保留两位
	//codes := global.KeepTwoFloat(manualAccess.Money)
	//if codes != 0 {
	//	return ctx.JSON(200, global.ReplyError(code, ctx))
	//}
	//获取操作人信息
	user := ctx.Get("user").(*global.RedisStruct)
	manualAccess.DoAgencyId = user.Id
	manualAccess.DoAgencyAccount = user.Account
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	if manualAccess.Types == 1 { //批量方式:账号
		//会员账号是否存在
		member, err := function.IsExistMemberAccount(manualAccess.Account, manualAccess.SiteId, manualAccess.SiteIndexId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if len(member) != len(manualAccess.Account) {
			return ctx.JSON(200, global.ReplyError(30138, ctx))
		}
		//添加数据
		count, err := manualAccessBean.AddBatchAccount(manualAccess)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(200, global.ReplyError(30142, ctx))
		}
	} else if manualAccess.Types == 2 { //批量方式:层级
		//会员层级是否存在会员
		member, err := function.IsExistMemberLevel(manualAccess.LevelId, manualAccess.SiteId, manualAccess.SiteIndexId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if len(member) != len(manualAccess.LevelId) {
			return ctx.JSON(200, global.ReplyError(50192, ctx))
		}
		//添加数据
		count, err := manualAccessBean.AddBatchLevel(manualAccess)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(200, global.ReplyError(30142, ctx))
		}
	}
	return ctx.NoContent(204)
}

//人工存取款历史记录(get列表)
func (cc *CashController) ManualAccess(ctx echo.Context) error {
	manualAccess := new(input.ManualAccessList)
	code := global.ValidRequest(manualAccess, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if manualAccess.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", manualAccess.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if manualAccess.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", manualAccess.EndTime, loc)
		times.EndTime = et.Unix()
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	cc.GetParam(listparam, ctx)
	list, err := manualAccessBean.GetList(manualAccess, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//人工取款(//TODO:稽核方面还没考虑)
func (*CashController) ManualWithdrawal(ctx echo.Context) error {
	manualAccess := new(input.ManualWithdrawalAdd)
	code := global.ValidRequest(manualAccess, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//取款金额的小数点后是不是保留两位
	codes := global.KeepTwoFloat(manualAccess.Money)
	if codes != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作人信息
	user := ctx.Get("user").(*global.RedisStruct)
	manualAccess.DoAgencyId = user.Id
	manualAccess.DoAgencyAccount = user.Account
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//获取会员信息
	member, has, err := function.GetMemberInfo(manualAccess.Account, manualAccess.SiteId, manualAccess.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30138, ctx))
	} else {
		//判断取款金额是否大于余额
		if manualAccess.Money > member.Balance {
			return ctx.JSON(200, global.ReplyError(30145, ctx))
		}
	}
	count, err := manualAccessBean.Withdrawals(manualAccess)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30143, ctx))
	}
	return ctx.NoContent(204)
}

//添加额度转换(post)
func (*CashController) BalanceConversionDo(ctx echo.Context) error {
	memberBalanceConversion := new(input.MemberBalanceConversionAdd)
	code := global.ValidRequest(memberBalanceConversion, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	memberBalanceConversion.DoUserId = user.Id
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//判断转入项目是否存在
	if memberBalanceConversion.FromType != 0 {
		has, err := memberBalanceConversionBean.IsExistFtype(memberBalanceConversion.FromType)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30147, ctx))
		}
	}
	//判断转出项目是否存在
	if memberBalanceConversion.ForType != 0 {
		has, err := memberBalanceConversionBean.IsExistFtype(memberBalanceConversion.ForType)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(30154, ctx))
		}
	}
	//转入项目或者转出项目必须有一个是系统余额
	if memberBalanceConversion.FromType != 0 && memberBalanceConversion.ForType != 0 {
		return ctx.JSON(200, global.ReplyError(30157, ctx))
	}
	//获取站点下套餐id
	site, err := memberBalanceConversionBean.GetSiteCombo(memberBalanceConversion.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//获取开户人视讯余额
	videoBalance, err := memberBalanceConversionBean.GetAgency(memberBalanceConversion.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if memberBalanceConversion.ForType == 0 { //转出项目为系统余额
		balance, err := memberBalanceConversionBean.GetBalance(memberBalanceConversion)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//系统余额是否大于转出金额
		if balance < memberBalanceConversion.Money {
			return ctx.JSON(200, global.ReplyError(30146, ctx))
		}
		//根据转入项目和套餐id获取手续费占成比
		proportion, err := memberBalanceConversionBean.GetProductProportion(memberBalanceConversion.FromType, site.ComboId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//将占比排序（从大到小）
		for i := 0; i < len(proportion); i++ {
			for j := i + 1; j < len(proportion); j++ {
				if proportion[i] < proportion[j] {
					proportion[i], proportion[j] = proportion[j], proportion[i]
				}
			}
		}
		//手续费
		memberBalanceConversion.Margin = proportion[0] * 0.01 * memberBalanceConversion.Money
		//判断视讯余额是否大于扣除的手续费
		if videoBalance.VideoBalance < memberBalanceConversion.Margin {
			return ctx.JSON(200, global.ReplyError(30158, ctx))
		}
	} else { //其他金额转换到系统金额
		//根据会员账号获取会员id
		memberId, err := memberBalanceConversionBean.GetMemberId(memberBalanceConversion.Account)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//获取转出项目的余额
		info, _, err := memberBalanceConversionBean.GetMoneyByVideo(memberId, memberBalanceConversion.ForType)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//根据转入项目和套餐id获取手续费占成比
		proportion, err := memberBalanceConversionBean.GetProductProportion(memberBalanceConversion.ForType, site.ComboId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//将占比排序（从大到小）
		for i := 0; i < len(proportion); i++ {
			for j := i + 1; j < len(proportion); j++ {
				if proportion[i] < proportion[j] {
					proportion[i], proportion[j] = proportion[j], proportion[i]
				}
			}
		}
		//手续费
		memberBalanceConversion.Margin = proportion[0] * 0.01 * memberBalanceConversion.Money
		//判断转出项目的余额是否大于转出金额
		if info.Balance < memberBalanceConversion.Money {
			return ctx.JSON(200, global.ReplyError(30146, ctx))
		}
	}
	count, err := memberBalanceConversionBean.BalanceConversionDo(memberBalanceConversion)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(30148, ctx))
	}
	return ctx.NoContent(204)
}

//额度转换记录(get列表)
func (cc *CashController) BalanceConversion(ctx echo.Context) error {
	mbc := new(input.MemberBalanceConversionList)
	code := global.ValidRequest(mbc, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if mbc.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", mbc.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if mbc.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", mbc.EndTime, loc)
		times.EndTime = et.Unix()
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	cc.GetParam(listparam, ctx)
	//获取操作者信息
	user := ctx.Get("user").(*global.RedisStruct)
	mbc.SiteIndexId = user.SiteIndexId
	list, count, err := memberBalanceConversionBean.GetList(mbc, listparam, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//出入账目汇总(get)
func (*CashController) ManualAccessCollect(ctx echo.Context) error {
	manualAccess := new(input.ManualAccessLists)
	code := global.ValidRequest(manualAccess, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if manualAccess.DateTime != "" {
		st, _ := time.ParseInLocation("2006-01-02", manualAccess.DateTime, loc)
		times.StartTime = st.Unix()
		dd, _ := time.ParseDuration("24h")
		ets := st.Add(dd)
		times.EndTime = ets.Unix()
	}
	//公司入款的收入金额，交易人数，交易笔数
	companyIntoMoney, err := manualAccessBean.CompanyIntoMoney(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//线上支付的收入金额，交易人数，交易笔数
	onlinePayment, err := manualAccessBean.OnlinePayment(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//人工存入的收入金额，交易人数，交易笔数
	depositManually, err := manualAccessBean.DepositManually(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//会员出款被扣除的收入金额，交易人数，交易笔数
	memberPaymentDeducted, err := manualAccessBean.MemberPaymentDeducted(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//会员出款的支出明细，交易人数，交易笔数
	memberPayment, err := manualAccessBean.MemberPayment(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//给予优惠的支出明细，交易人数，交易笔数
	giveDiscount, err := manualAccessBean.GiveDiscount(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//人工提现的支出明细，交易人数，交易笔数
	manualWithdrawal, err := manualAccessBean.ManualWithdrawal(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//给予返水的支出明细，交易人数，交易笔数
	giveBackWater, err := manualAccessBean.GiveBackWater(manualAccess, times)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var info = make(map[string]interface{})
	info["companyIntoMoney"] = companyIntoMoney           //公司入款
	info["onlinePayment"] = onlinePayment                 //线上支付
	info["depositManually"] = depositManually             //人工存入
	info["memberPaymentDeducted"] = memberPaymentDeducted //会员出款被扣除
	info["memberPayment"] = memberPayment                 //会员出款
	info["giveDiscount"] = giveDiscount                   //给予优惠
	info["manualWithdrawal"] = manualWithdrawal           //人工提现
	info["giveBackWater"] = giveBackWater                 //给予返水
	//汇总
	var sumMary back.Summary
	sumMary.DepositMoney = companyIntoMoney.Money + onlinePayment.Money + depositManually.Money + memberPaymentDeducted.Money
	sumMary.DepositPeople = companyIntoMoney.PeopleNum + onlinePayment.PeopleNum + depositManually.PeopleNum + memberPaymentDeducted.PeopleNum
	sumMary.DepositCount = companyIntoMoney.Count + onlinePayment.Count + depositManually.Count + memberPaymentDeducted.Count
	sumMary.PayoutMoney = memberPayment.Money + giveDiscount.Money + manualWithdrawal.Money + giveBackWater.Money
	sumMary.PayoutPeople = memberPayment.PeopleNum + giveDiscount.PeopleNum + manualWithdrawal.PeopleNum + giveBackWater.PeopleNum
	sumMary.PayoutCount = memberPayment.Count + giveDiscount.Count + manualWithdrawal.Count + giveBackWater.Count
	sumMary.ProfitLoss = sumMary.DepositMoney - sumMary.PayoutMoney
	info["sumMary"] = sumMary
	return ctx.JSON(200, global.ReplyItem(info))
}

//出款管理(get出款列表)
func (ca *CashController) ListCash(ctx echo.Context) error {
	companList := new(input.OutMoneyList)
	code := global.ValidRequest(companList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	ca.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if companList.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", companList.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if companList.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", companList.EndTime, loc)
		times.EndTime = et.Unix()
	}
	infolist, err := makeMoneyBean.OutMoney(companList, times, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(infolist))
}

//出款管理(put预备出款)
func (*CashController) ManageCashReady(ctx echo.Context) error {
	prepare := new(input.PrepareOutMoney)
	code := global.ValidRequest(prepare, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	prepare.AgencyId = user.Id
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//首先判断是不是已经在预备出款状态
	info, flag, err := makeMoneyBean.GetInfoById(prepare.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.JSON(500, global.ReplyError(50184, ctx))
	}
	//已经处于预备出款状态
	if info.OutStatus == 2 && info.DoAgencyId != 0 {
		return ctx.JSON(200, global.ReplyError(60024, ctx))
	}
	//处于其他状态
	if info.OutStatus != 5 {
		return ctx.JSON(200, global.ReplyError(60052, ctx))
	}
	//看操作人是否存在
	agency := new(function.AgencyBean)
	infos, flags, err := agency.GetAgency(prepare.AgencyId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flags {
		return ctx.JSON(200, global.ReplyError(60054, ctx))
	}
	//预备出款
	count, err := makeMoneyBean.PrepareOutMoney(prepare, infos, user)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50184, ctx))
	}
	return ctx.NoContent(204)
}

//出款管理(put确定出款)
func (*CashController) ManageCashDo(ctx echo.Context) error {
	confirm := new(input.ConfirmOutMoney)
	code := global.ValidRequest(confirm, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	confirm.AgencyId = user.Id
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//判断操作人是否能够匹配
	info, flag, err := makeMoneyBean.GetInfoById(confirm.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//没有搜索到该条记录
	if !flag {
		return ctx.JSON(500, global.ReplyError(60045, ctx))
	}
	//处于预备出款状态，判断操作人
	if info.OutStatus == 2 {
		if info.DoAgencyId != confirm.AgencyId {
			return ctx.JSON(200, global.ReplyError(60044, ctx))
		}
	}
	//已经拒绝出款或者取消的出款
	if info.OutStatus == 3 || info.OutStatus == 4 {
		return ctx.JSON(200, global.ReplyError(60047, ctx))
	}

	//看操作人是否存在
	agency := new(function.AgencyBean)
	infos, flags, err := agency.GetAgency(confirm.AgencyId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flags {
		return ctx.JSON(200, global.ReplyError(60054, ctx))
	}

	//确认
	count, err := makeMoneyBean.ConfirmOutMoney(confirm, infos)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50185, ctx))
	}
	return ctx.NoContent(204)
}

//出款管理(put取消出款)
func (*CashController) ManageCashChannel(ctx echo.Context) error {
	cancle := new(input.CancleOutMoney)
	code := global.ValidRequest(cancle, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	cancle.AgencyId = user.Id
	if !(user.RoleId == 1 || user.RoleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	info, flag, err := makeMoneyBean.GetInfoById(cancle.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//没有搜索到该条记录
	if !flag {
		return ctx.JSON(500, global.ReplyError(60045, ctx))
	}
	//如果处于预备出款状态，判断操作人有没有权限
	if info.OutStatus == 2 {
		if info.DoAgencyId != cancle.AgencyId {
			return ctx.JSON(200, global.ReplyError(60044, ctx))
		}
	}
	//如果已经出款就不能在进行任何操作
	if info.OutStatus == 1 {
		return ctx.JSON(200, global.ReplyError(60046, ctx))
	}
	//已经拒绝的出款订单
	if info.OutStatus == 4 {
		return ctx.JSON(200, global.ReplyError(60049, ctx))
	}
	//看操作人是否存在
	agency := new(function.AgencyBean)
	infos, flags, err := agency.GetAgency(cancle.AgencyId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flags {
		return ctx.JSON(200, global.ReplyError(60054, ctx))
	}
	//修改
	count, err := makeMoneyBean.CancleOutMoney(cancle, info.OutwardNum, infos)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50186, ctx))
	}
	return ctx.NoContent(204)
}

//出款管理(put拒绝出款)
func (*CashController) ManageCashRefuse(ctx echo.Context) error {
	refuse := new(input.RefuseOutMoney)
	code := global.ValidRequest(refuse, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	refuse.AgencyId = user.Id
	if roleId != 1 && roleId != 5 {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	info, flag, err := makeMoneyBean.GetInfoById(refuse.Id)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//没有搜索到该条记录
	if !flag {
		return ctx.JSON(500, global.ReplyError(60045, ctx))
	}
	//如果出款状态处于预备出款，就判断操作人是否匹配
	if info.OutStatus == 2 {
		if info.DoAgencyId != refuse.AgencyId {
			return ctx.JSON(200, global.ReplyError(60044, ctx))
		}
	}
	//如果已经出款就不能在进行任何操作
	if info.OutStatus == 1 {
		return ctx.JSON(200, global.ReplyError(60046, ctx))
	}
	//已经取消的出款订单
	if info.OutStatus == 3 {
		return ctx.JSON(200, global.ReplyError(60048, ctx))
	}
	//看操作人是否存在
	agency := new(function.AgencyBean)
	infos, flags, err := agency.GetAgency(refuse.AgencyId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flags {
		return ctx.JSON(200, global.ReplyError(60054, ctx))
	}
	//操作
	count, err := makeMoneyBean.RefuseOutMoney(refuse, infos.Account)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50187, ctx))
	}
	return ctx.NoContent(204)
}

//公司入款列表
func (ca *CashController) CompanyIncome(ctx echo.Context) error {
	companList := new(input.CompanyIncomeList)
	code := global.ValidRequest(companList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	ca.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if companList.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", companList.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if companList.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", companList.EndTime, loc)
		times.EndTime = et.Unix()
	}
	//金额下限大于金额上限
	if companList.LowerLimit > companList.UpperLimit {
		return ctx.JSON(200, global.ReplyError(30186, ctx))
	}
	infolist, err := memberCompanyIncomeBean.GetInfoList(companList, times, listparam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(infolist))
}

//确定一条公司入款
func (*CashController) CompanyIncomeDo(ctx echo.Context) error {
	confirmIncome := new(input.ConfirmCompany)
	code := global.ValidRequest(confirmIncome, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	count, err := memberCompanyIncomeBean.ConfirmCompanyIncome(confirmIncome)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50182, ctx))
	}
	return ctx.NoContent(204)
}

//取消一条公司入款
func (*CashController) CompanyIncomeChannel(ctx echo.Context) error {
	companyIncome := new(input.CancleIncome)
	code := global.ValidRequest(companyIncome, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	if companyIncome.Status != 2 && companyIncome.Status != 3 {
		return ctx.JSON(200, global.ReplyError(60055, ctx))
	}
	count, err := memberCompanyIncomeBean.CancleCompanyIncome(companyIncome, user)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(200, global.ReplyError(50183, ctx))
	}
	return ctx.NoContent(204)
}

//线上入款列表
func (ca *CashController) OnlineIncomeList(ctx echo.Context) error {
	depositList := new(input.OnlineDepositList)
	code := global.ValidRequest(depositList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listParam := new(global.ListParams)
	//获取listparam的数据
	ca.GetParam(listParam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if depositList.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", depositList.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if depositList.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", depositList.EndTime, loc)
		times.EndTime = et.Unix()
	}
	//金额下限大于金额上限
	if depositList.LowerLimitMoney > depositList.UpperLimitMoney {
		return ctx.JSON(200, global.ReplyError(30186, ctx))
	}
	infolist, err := onlineEntryRecordBean.DepositList(depositList, times, listParam)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(infolist))
}

//确定一条线上入款
func (*CashController) OnlineIncomeDo(ctx echo.Context) error {
	inputDeposit := new(input.ConfirmDeposit)
	code := global.ValidRequest(inputDeposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	count, err := onlineEntryRecordBean.ConfirmOnlineDeposit(inputDeposit)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(50180, ctx))
	}
	return ctx.NoContent(204)
}

//取消一条线上入款
func (*CashController) OnlineIncomeChannel(ctx echo.Context) error {
	cancleDeposit := new(input.CancleDeposit)
	code := global.ValidRequest(cancleDeposit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//角色权限认证[开户人和平台管理员可以操作]
	user := ctx.Get("user").(*global.RedisStruct)
	roleId := user.RoleId
	if !(roleId == 1 || roleId == 5) {
		return ctx.JSON(200, global.ReplyError(60044, ctx))
	}
	//取消
	count, err := onlineEntryRecordBean.CancleOneDeposit(cancleDeposit)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		return ctx.JSON(500, global.ReplyError(50181, ctx))
	}
	return ctx.NoContent(204)
}

//获取拒绝，取消出款原因
func (*CashController) OutRemark(ctx echo.Context) error {
	outRemark := new(input.OutRemark)
	code := global.ValidRequest(outRemark, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, err := onlineEntryRecordBean.OutRemark(outRemark)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//根据会员账号获取会员账号，余额，姓名
func (*CashController) MemberInfo(ctx echo.Context) error {
	member := new(input.MemberInfo)
	code := global.ValidRequest(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	info, has, err := memberBean.GetMemberInfo(member)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30246, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//根据会员id和转出项目获取会员余额
func (*CashController) GetMoney(ctx echo.Context) error {
	outType := new(input.OutType)
	code := global.ValidRequest(outType, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//根据会员账号获取会员id
	memberId, err := function.GetMemberIdByAccount(outType.Account)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var info back.MemberProductClassifyBalance
	if outType.ForType == 0 {
		info, _, err = memberBalanceConversionBean.GetMoneyByOutTypes(memberId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	} else {
		info, _, err = memberBalanceConversionBean.GetMoneyByVideo(memberId, outType.ForType)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//获取会员余额统计数据
func (*CashController) GetBlance(ctx echo.Context) error {
	MemberClassifyBalance := new(input.MemberClassifyBalance)
	code := global.ValidRequest(MemberClassifyBalance, ctx)

	if len(MemberClassifyBalance.SiteId) == 0 { //站点id为空的时候使用默认站点信息
		userinfo := ctx.Get("user").(*global.RedisStruct)
		MemberClassifyBalance.SiteId = userinfo.SiteId
	}

	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	info, err := memberBalanceConversionBean.ClassifyBalance(MemberClassifyBalance)

	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(info))
}

//交易平台列表
func (*CashController) Index(ctx echo.Context) error {
	data, err := productBean.TypeList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取所有代理id和账号
func (*CashController) GetAgency(ctx echo.Context) error {
	agency := new(input.SiteId)
	code := global.ValidRequest(agency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := memberCompanyIncomeBean.GetAgency(agency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取所有收款账号id和账号
func (*CashController) GetSetAgency(ctx echo.Context) error {
	setAgency := new(input.SiteId)
	code := global.ValidRequest(setAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := memberCompanyIncomeBean.GetSetAccount(setAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取层级下拉框
func (*CashController) MemberDrop(ctx echo.Context) error {
	levelIndex := new(input.MemberLevels)
	code := global.ValidRequest(levelIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	levelIndex.SiteId = user.SiteId
	data, err := memberLevelBean.Memberdrop(levelIndex)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//入款商户下拉
func (*CashController) ThirdPaidList(ctx echo.Context) error {
	paidList := new(input.MemberLevels)
	code := global.ValidRequest(paidList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := memberLevelBean.ThirdPaidList(paidList)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}
