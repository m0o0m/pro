package member

import (
	"encoding/json"
	"fmt"
	"framework/render"
	"global"
	"models/input"

	"github.com/go-redis/redis"

	"controllers"
	"models/back"
	"time"

	"github.com/labstack/echo"
)

type BaseInfoController struct {
	controllers.BaseController
}

var RANKING_LIST = "ranking_list"

//获取个人基本资料
func (bic *BaseInfoController) Info(ctx echo.Context) error {
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	mi := new(input.MemberInfoSelf)
	mi.Id = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	data, has, err := baseInfoBean.MemberSelfInfo(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//今日交易记录
func (bic *BaseInfoController) PayRecord(ctx echo.Context) error {
	mi := new(input.PayRecordToday)
	code := global.ValidRequestMember(mi, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	mi.Id = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	//获取listparam的数据
	listparam := new(global.ListParams)
	bic.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	if mi.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", mi.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if mi.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", mi.EndTime, loc)
		times.EndTime = et.Unix()
	}
	data, count, err := baseInfoBean.DealRecordMemberSelf(mi, listparam, times)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//今日交易记录
func (bic *BaseInfoController) GetPayRecord(ctx echo.Context) error {
	mi := new(input.PayRecordToday)
	code := global.ValidRequestMember(mi, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	mi.Id = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	//获取listparam的数据
	listparam := new(global.ListParams)
	bic.GetParam(listparam, ctx)
	//开始时间和结束时间的处理（处理日期方式：年-月-日）
	times := new(global.Times)
	loc, _ := time.LoadLocation("Local")
	var endTime int64
	if mi.StartTime != "" {
		st, _ := time.ParseInLocation("2006-01-02", mi.StartTime, loc)
		times.StartTime = st.Unix()
	}
	if mi.EndTime != "" {
		et, _ := time.ParseInLocation("2006-01-02", mi.EndTime, loc)
		endTime = et.Unix()
		dd, _ := time.ParseDuration("24h")
		t1 := et.Add(dd)
		times.EndTime = t1.Unix()
	}
	if endTime < times.StartTime {
		return ctx.JSON(200, global.ReplyError(30139, ctx))
	}
	data, err := baseInfoBean.GetDealRecordMemberSelf(mi, times)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//修改个人基本资料(密码)
func (bic *BaseInfoController) EditInfo(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	//获取用户参数
	base_info := new(input.MemberPassword)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	base_info.Id = member.Id
	//比较密码和重复密码是否相同
	if base_info.Password != base_info.RePassword {
		return ctx.JSON(200, global.ReplyError(20008, ctx))
	}
	//从数据库中查询会员的密码
	data, has, err := baseInfoBean.OneMemberInfoForPassword(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	if data.IsEditPassword == 2 {
		return ctx.JSON(200, global.ReplyError(50115, ctx))
	}
	//原密码加密
	if base_info.OriginalPassword != "" {
		md5Password, err := global.MD5ByStr(base_info.OriginalPassword, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		base_info.OriginalPassword = md5Password
	}
	//如果类型是修改登录密码并且原密码不等于登录密码
	if base_info.Type == 1 && base_info.OriginalPassword != data.Password {
		return ctx.JSON(200, global.ReplyError(20009, ctx))

	}
	//如果类型是修改取款密码并且原密码不等于取款密码
	if base_info.Type == 2 && base_info.OriginalPassword != data.DrawPassword {
		return ctx.JSON(200, global.ReplyError(20009, ctx))
	}
	//密码加密
	if base_info.Password != "" {
		md5Password, err := global.MD5ByStr(base_info.Password, global.EncryptSalt)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(30044, ctx))
		}
		base_info.Password = md5Password
	}
	_, err = baseInfoBean.UpMemberPassword(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取个人出款银行列表
func (bic *BaseInfoController) Bank(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.MemberBankList)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//登录用户信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	data, err := baseInfoBean.MemberBankList(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k, v := range data {
		data[k].Card = global.BankStr(v.Card)
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//添加出款银行
func (bic *BaseInfoController) BankAdd(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.MemberBankAdd)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//登录用户信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id

	//检查银行卡是否合法
	flag := global.CheckCardNumber(base_info.Card)
	if !flag {
		return ctx.JSON(200, global.ReplyError(30060, ctx))
	}
	//检验该银行卡是否已经被添加
	_, has, count, err := baseInfoBean.CheckOutBank(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(500, global.ReplyError(30069, ctx))
	}
	if count > 2 {
		return ctx.JSON(500, global.ReplyError(50117, ctx))
	}
	if len(base_info.Password) == 0 && count == 0 {
		return ctx.JSON(500, global.ReplyError(20005, ctx))
	}
	//添加银行卡
	_, err = baseInfoBean.MemberOutBankAdd(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(base_info.Password) != 0 && count == 0 {
		password, _ := global.MD5ByStr(base_info.Password, global.EncryptSalt)
		_, err = memberBean.ChangePassword(member.Id, password, 2)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}

	return ctx.NoContent(204)
}

//修改出款银行
func (bic *BaseInfoController) BankEdit(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.MemberBankChange)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//登录用户信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	//检查银行卡是否合法
	flag := global.CheckCardNumber(base_info.Card)
	if !flag {
		return ctx.JSON(200, global.ReplyError(30060, ctx))
	}
	//检验该银行卡是否已经被添加
	base_add := new(input.MemberBankAdd)
	base_add.BankId = base_info.BankId
	base_add.Card = base_info.Card
	base_add.MemberId = base_info.MemberId
	_, has, _, err := baseInfoBean.CheckOutBank(base_add)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		return ctx.JSON(500, global.ReplyError(30069, ctx))
	}
	_, err = baseInfoBean.MemberOutBankUpdata(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//删除出款银行
func (bic *BaseInfoController) BankDelete(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.MemberBankDelete)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//登录用户信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	_, err := baseInfoBean.MemberOutBankDelete(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取一条出款银行数据
func (bic *BaseInfoController) OneMemberBankInfo(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.OneMemberBankInfo)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//登录用户信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	data, has, err := baseInfoBean.OneMemberBankInfo(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50116, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//银行卡
func (c *BaseInfoController) BankAddInfo(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)

	data, err := bankCardBean.IndexAllBankOut(member.Site, member.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	return ctx.JSON(200, global.ReplyItem(data))
}

//手机号修改
func (bic *BaseInfoController) PhoneNumEdit(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.PhoneNum)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	valid_code, err := SendCodeForPhone(base_info.PhoneNum)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if base_info.Code != valid_code {
		return ctx.JSON(200, global.ReplyError(20021, ctx))
	}
	//检验手机
	flag := global.CheckPhoneNumber(base_info.PhoneNum)
	if !flag {
		return ctx.JSON(200, global.ReplyError(20014, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	_, err = baseInfoBean.ChangePhoneNum(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//邮箱修改
func (bic *BaseInfoController) EmailNumEdit(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.EmailNum)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//检验邮箱
	flag := global.CheckEmail(base_info.EmailNum)
	if !flag {
		return ctx.JSON(200, global.ReplyError(20015, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	_, err := baseInfoBean.ChangeEmailNum(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//出生日期修改
func (bic *BaseInfoController) BirthNumEdit(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.BirthdayNum)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	_, err := baseInfoBean.ChangeBirthdayNum(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//修改资料
func (bic *BaseInfoController) MeansEdit(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.EditMeans)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.MemberId = member.Id
	flag := true
	//检验手机
	if len(base_info.PhoneNum) != 0 {
		flag = global.CheckPhoneNumber(base_info.PhoneNum)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20014, ctx))
		}
	}
	//检验邮箱
	if len(base_info.EmailNum) != 0 {
		flag = global.CheckEmail(base_info.EmailNum)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20015, ctx))
		}
	}

	//检验qq
	if len(base_info.QqNum) != 0 {
		flag = global.Checkqq(base_info.QqNum)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20016, ctx))
		}
	}
	//检验身份证
	if len(base_info.Card) != 0 {
		flag = global.CheckIdentity(base_info.Card)
		if !flag {
			return ctx.JSON(200, global.ReplyError(20013, ctx))
		}
	}
	_, err := baseInfoBean.ChangeMeans(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取会员的详细信息
func (bic *BaseInfoController) MemberDetail(ctx echo.Context) error {
	//获取用户参数
	base_info := new(input.MemberInfoSelf)
	code := global.ValidRequestMember(base_info, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	base_info.Id = member.Id
	data, has, err := baseInfoBean.OneMemberDetail(base_info)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50007, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//手机验证码
func SendCodeForPhone(phone string) (code string, err error) {
	if phone != "" {
		code = "123456"
	}
	return
}

//登陆验证
func (c *BaseInfoController) GetAjaxLoginStatus(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	mi := new(input.MemberInfoSelf)
	mi.Id = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	data, _, err := baseInfoBean.MemberSelfInfo(mi)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}

	sdata := back.AjaxLoginIn{}
	sdata.Id = data.Id
	sdata.Account = data.Account
	sdata.Realname = data.Realname
	sdata.Balance = data.Balance
	platformBalance, err := memberBalanceConversionBean.GetPlatformBalance(member.Id)
	var da back.MemberBalanceTotalBack
	for _, v := range platformBalance.ProductClassifyBalance {
		da.Type = 1
		da.Balance = v.Balance
		da.Name = v.Platform
		sdata.TBalance = append(sdata.TBalance, da)
	}
	sdata.TBalance = append(sdata.TBalance, back.MemberBalanceTotalBack{"账户余额", platformBalance.AccountBalance, 2})
	dad := platformBalance.AccountBalance + platformBalance.GameBalance
	sdata.TBalance = append(sdata.TBalance, back.MemberBalanceTotalBack{"账户总余额", dad, 3})

	//sdata.TBalance, err = baseInfoBean.MemberBalanceAll(mi)

	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}

	times := new(global.Times)
	timeNow := global.GetCurrentTime()
	times.StartTime = timeNow - 7*24*3600
	times.EndTime = timeNow
	sdata.Count, err = memMessageBean.MemMessageCount(member.Site, member.SiteIndex, member.Id, times)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//获取推广详情
func (m *BaseInfoController) GetMemberRebate(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	//查询会员信息
	memberRebateInfo, err := memberBean.GetMemberSpreadById(member.Site, member.SiteIndex, member.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	memberRebateInfo.SpreadUrl = ctx.Request().Host + "/?u=" + fmt.Sprintf("%d", member.Id)
	//查询推广返佣比例
	//查询商品表信息
	products, err := productBean.GetList(member.Site, member.SiteIndex, &input.ProductList{Status: 1})
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//js, _ := json.Marshal(products)
	//fmt.Println("所有商品:", string(js))
	if len(products) == 0 {
		return ctx.JSON(500, global.ReplyError(70013, ctx))
	}
	//查询站点返佣设定(优惠设定)
	rebateSets, err := rebateSetBean.GetAll(member.Site, member.SiteIndex, products)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//横向排列,便于前端展示
	rates := make(map[string]*back.MemberRebateRate, len(rebateSets))
	for i, rebateSet := range rebateSets {
		if i == 0 {
			rates["valid_bet"] = new(back.MemberRebateRate)
			rates["valid_bet"].Title = "有效打码"
			rates["valid_bet"].Values = make([]float64, len(rebateSets))
		}
		rates["valid_bet"].Values[i] = float64(rebateSet.ValidMoney)
		for _, productRate := range *rebateSet.ProductRates {
			if i == 0 {
				rates[productRate.VType] = new(back.MemberRebateRate)
				rates[productRate.VType].Title = productRate.ProductName
				rates[productRate.VType].Values = make([]float64, len(rebateSets))
			}
			rates[productRate.VType].Values[i] = productRate.Rate
		}
	}
	temp := make([]*back.MemberRebateRate, 0)
	if memberRebateRate, ok := rates["valid_bet"]; ok {
		temp = append(temp, memberRebateRate)
	}
	for _, rate := range rates {
		if rate.Title != "有效打码" {
			temp = append(temp, rate)
		}
	}
	memberRebateInfo.RebateSets = temp

	//先从缓存从取排行榜
	memberRebateInfo.RankingList, err = m.getRanking()
	if err != nil {
		if err == redis.Nil {
			//通过推广设定生成假数据
			set, err := memberSpreadBean.GetSpreadSetBySite(member.Site, member.SiteIndex)
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			rankingNum := set.RankingNum    //人数系数 ,我这里设定生成的假数据最高金额不超过最低的2倍
			rankinMoney := set.RankingMoney //金额系数 ,我这里设定生成的假数据最高金额不超过最低的3倍

			rankingList := back.RankingList{}

			for i := 1; i < 11; i++ {
				rankingList.RankingNumList = append(rankingList.RankingNumList, &back.RankingNumList{
					Id:      int64(i),
					Account: global.GetRandomAccount(),
					Num:     int64(int64(rankingNum)/10*int64(11-i)) + int64(rankingNum),
				})
				rankingList.RankingMoneyList = append(rankingList.RankingMoneyList, &back.RankingMoneyList{
					Id:      int64(i),
					Account: global.GetRandomAccount(),
					Money:   global.FloatReserve2(float64(rankinMoney/5*float64(11-i)) + float64(rankinMoney)),
				})
			}
			memberRebateInfo.RankingList = rankingList
			err = m.saveRanking(&rankingList)
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			return ctx.JSON(200, global.ReplyItem(&memberRebateInfo))
		}
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(&memberRebateInfo))
}

func (m *BaseInfoController) saveRanking(saveData *back.RankingList) error {
	data, err := json.Marshal(saveData)
	if err != nil {
		//global.GlobalLogger.Error("err:%s",err.Error())
		return err
	}
	err = global.GetRedis().SetNX(RANKING_LIST, string(data), global.UpdateRankingExp).Err()
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
	}
	return err
}

func (m *BaseInfoController) getRanking() (saveData back.RankingList, err error) {
	str, err := global.GetRedis().Get(RANKING_LIST).Result()
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		return
	}
	err = json.Unmarshal([]byte(str), &saveData)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
	}
	return
}
