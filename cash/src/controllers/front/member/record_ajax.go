package member

import (
	"controllers"
	"encoding/json"
	"framework/render"
	"github.com/labstack/echo"
	"global"
	"html/template"
	"models/back"
	"models/function"
	"models/input"
	"strconv"
	"strings"
	"time"
)

type RecordAjax struct {
	controllers.BaseController
}

//获取注册设定
func (c *RecordAjax) GetRegisterSet(ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	regSet := new(function.RegisterStatusBean)
	data, err := regSet.GetRegSet(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取会员消息记录
func (c *RecordAjax) GetMesList(ctx echo.Context) error {
	mesTerm := new(input.MesTerm)
	code := global.ValidRequestMember(mesTerm, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberInfo := ctx.Get("member").(*global.MemberRedisToken)
	times := new(global.Times)
	timeNow := global.GetCurrentTime()
	times.StartTime = timeNow - 7*24*3600
	times.EndTime = timeNow
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := MemberSelfInfoBean.GetMesList(memberInfo.Site, memberInfo.SiteIndex, memberInfo.Id, times, listparam)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//获取注单记录
func (c *RecordAjax) GetRecordInfo(ctx echo.Context) error {
	member := new(input.RecordInfoList)
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if member.GameName == "" {
		return ctx.JSON(200, global.ReplyError(60302, ctx))
	}
	memberinfo := ctx.Get("member").(*global.MemberRedisToken)
	//获取listparam的数据
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := BetRecordInfoBean.GetMemberRecord(memberinfo.Site, memberinfo.SiteIndex, memberinfo.Account, member, listparam)

	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//修改消息状态
func (c *RecordAjax) PutMesStatus(ctx echo.Context) error {
	member := new(input.WapMesStatus)
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberinfo := ctx.Get("member").(*global.MemberRedisToken)
	_, has, err := MemberSelfInfoBean.MemberOneMesInfo(member.Id)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(10211, ctx))
	}
	data, err := MemberSelfInfoBean.PutMesStatus(memberinfo.Site, memberinfo.SiteIndex, member)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(500, global.ReplyError(91110, ctx))
	}
	return ctx.NoContent(204)
}

//站点优惠活动
type SiteActivity struct {
	Id      int64         `xorm:"'id' PK autoincr"` //id
	TopId   int64         `xorm:"top_id"`           //上级栏目
	Title   string        `xorm:"title"`            //标题
	Content template.HTML `xorm:"content"`          //内容
	Img     string        `xorm:"img"`              //标题图片路径
}

func (c *RecordAjax) GetActivity(ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	//优惠活动查询
	youhuiList, err := siteIWordBean.IndexActivityList(siteId, siteIndexId, 2)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	yhTitleData := []back.YouhuiDataInfo{}
	yhData := []SiteActivity{}
	//获取分类活动
	for _, v := range youhuiList {
		if v.TopId == 0 && len(v.TypeName) != 0 {
			titleInfo := back.YouhuiDataInfo{}
			titleInfo.Id = v.Id
			titleInfo.Title = v.Title
			yhTitleData = append(yhTitleData, titleInfo)
		} else {
			info := SiteActivity{}
			info.Id = v.Id
			info.Title = v.Title
			info.TopId = v.TopId
			info.Img = v.Img
			info.Content = template.HTML(v.Content)
			yhData = append(yhData, info)
		}
	}
	sdata := make(map[string]interface{})
	sdata["yhTitleData"] = yhTitleData
	sdata["yhData"] = yhData
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//获取交易记录
func (c *RecordAjax) GetCashList(ctx echo.Context) error {
	member := new(input.WapMemberCashRecord)
	code := global.ValidRequestMember(member, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberinfo := ctx.Get("member").(*global.MemberRedisToken)
	times := new(global.Times)
	times.StartTime = member.StartTime
	times.EndTime = member.EndTime
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	if member.SourceType == 13 {
		data, count, err := MemberCashRecordBean.GetWapCompanyRecord(memberinfo.Site, memberinfo.SiteIndex, memberinfo.Account, memberinfo.Id, member, times, listparam)
		if err != nil {
			global.GlobalLogger.Error("err:s%", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
	} else {
		data, count, err := MemberCashRecordBean.GetWapCashRecord(memberinfo.Site, memberinfo.SiteIndex, memberinfo.Account, member, times, listparam)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
	}
}

//获取电子导航
func (c *RecordAjax) GetGameTitle(ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	//电子导航查询
	gameModel, err := noteGameBean.IndexGameTitle(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	gameType, err := noteGameBean.GameDataType()
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	gameTitleList := strings.Split(gameModel, ",")
	GameTitle := []string{}
	for _, v := range gameTitleList {
		for _, value := range gameType {
			rs := []rune(v)
			rl := len(v)
			r := string(rs[:rl-3])
			if rl > 0 && (r+"h5") == value.Type {
				GameTitle = append(GameTitle, strings.ToUpper(r))
			}
		}
	}
	return ctx.JSON(200, global.ReplyItem(GameTitle))
}

//删除消息
func (c *RecordAjax) DeleteMes(ctx echo.Context) error {
	mes_id := new(input.WapMemberMessageDel)
	code := global.ValidRequestMember(mes_id, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := memMessageBean.MemberDelMessage(mes_id.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(10212, ctx))
	}
	return ctx.NoContent(204)
}

//获取电子游戏列表
func (c *RecordAjax) GetGameData(ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)
	reqDta := new(input.VType)
	code := global.ValidRequest(reqDta, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}

	result := back.EgameDataInfo{}
	//电子信息查询
	gameList, err := noteGameBean.IndexGameList(siteId, siteIndexId, reqDta.VType)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	result.Count = len(gameList)
	result.Wh = 1
	result.Data = gameList
	result.Type = reqDta.VType

	return ctx.JSON(200, global.ReplyItem(result))
}

//获取公告分类列表
func (c *RecordAjax) GetNoticeList(ctx echo.Context) error {
	noticeId := new(input.Notice)
	code := global.ValidRequestMember(noticeId, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberinfo := ctx.Get("member").(*global.MemberRedisToken)
	times := new(global.Times)
	timeNow := global.GetCurrentTime()
	times.StartTime = timeNow - 7*24*3600
	times.EndTime = timeNow
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	listdata, count, err := NoticeBean.GetNoticeList(memberinfo.Site, noticeId.Id, times, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	return ctx.JSON(200, global.ReplyPagination(listparam, listdata, int64(len(listdata)), count, ctx))
}

//红包
func (c *RecordAjax) RedPacketLog(ctx echo.Context) error {
	siteId, _ := ctx.Get("site_id").(string)
	siteIndexId, _ := ctx.Get("site_index_id").(string)

	//红包返回数据
	sdata := []back.RedPacket{}

	//获取红包数据
	data, err := redPacketSetBean.SiteFind(siteId, siteIndexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return render.PageErr(60000, ctx)
	}
	for _, v := range data {
		info := back.RedPacket{}
		info.Id = v.Id
		info.Title = v.Title
		info.MaxCount = v.MaxCount
		info.StartTime = v.StartTime
		info.InStartTime = v.InStartTime
		info.InEndTime = v.InEndTime
		info.InSum = v.InSum
		info.AuditStartTime = v.AuditStartTime
		info.AuditEndTime = v.AuditEndTime
		info.BetSum = v.BetSum
		info.LevelId = v.LevelId
		info.TotalMoney = v.TotalMoney
		info.MinMoney = v.MinMoney
		info.RedNum = v.RedNum
		info.CreateTime = v.CreateTime
		info.IsIp = v.IsIp
		info.StyleId = v.StyleId
		info.IsShow = v.IsShow
		info.AppointMoney = v.AppointMoney
		info.RedType = v.RedType
		info.Status = v.Status
		info.IsGenerate = v.IsGenerate
		timeNum := time.Now().Unix()
		info.Opencount = timeNum - v.StartTime
		info.Closecount = v.EndTime - timeNum
		info.Pic = 1
		sdata = append(sdata, info)
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//出款数据写入
func (c *RecordAjax) DrawWriteData(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	redisKey := member.Site + "_" + member.SiteIndex + "_" + strconv.FormatInt(member.Id, 10) + "outCharge"
	s, err := c.getWapMemberRedis(redisKey) //从redis取数据
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	ShowOutData := new(back.AddOutData)
	err = json.Unmarshal([]byte(s), ShowOutData)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//是否允许出款
	if ShowOutData.OutStatus == 2 {
		return ctx.JSON(200, global.ReplyError(30237, ctx))
	}
	//获取会员代理id
	memberInfo, has, err := ManualAccessBean.GetWapInfo(member.Account, member.Site, member.SiteIndex)
	if !has {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//根据代理id获取代理账号
	agencyinfo, err := ManualAccessBean.GetAgencyAccount(memberInfo.ThirdAgencyId)
	//开启事务
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	data, err := memberBean.DelMoneyById(member.Site, member.Id, ShowOutData.OutMoney, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		sess.Rollback()
		return ctx.JSON(200, global.ReplyError(91011, ctx))
	}
	//现金表数据
	cashData := new(back.WapCashRecord)
	cashData.TradeNo = ShowOutData.OrderNum
	cashData.ClientType = ShowOutData.ClientType
	cashData.Balance = ShowOutData.OutMoney
	cashData.DisBalance = ShowOutData.DepositMoney
	cashData.AgencyId = memberInfo.ThirdAgencyId
	cashData.AgencyAccount = memberInfo.Account
	cashData.MemberId = member.Id
	cashData.SourceType = 4
	cashData.SiteIndexId = member.SiteIndex
	cashData.SiteId = member.Site
	cashData.UserName = member.Account
	cashData.AfterBalance = data
	cashResult, err := drawMoney.WriteCashRecord(cashData, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(91010, ctx))
	}
	if cashResult == 0 {
		sess.Rollback()
		return ctx.JSON(200, global.ReplyError(91012, ctx))
	}
	if memberInfo.FirstDepositTime > 0 {
		ShowOutData.IsFirst = 2
	}
	//出款表数据
	saveMakeData := new(back.SaveMakeMoney)
	saveMakeData.OutStatus = 5
	saveMakeData.OrderNum = ShowOutData.OrderNum
	saveMakeData.Charge = ShowOutData.Charge
	saveMakeData.IsFirst = ShowOutData.IsFirst
	saveMakeData.SiteId = member.Site
	saveMakeData.SiteIndexId = member.SiteIndex
	saveMakeData.MemberId = member.Id
	saveMakeData.UserName = member.Account
	saveMakeData.LevelId = member.LevelId
	saveMakeData.AgencyId = memberInfo.ThirdAgencyId
	saveMakeData.AgencyAccount = agencyinfo.Account
	saveMakeData.OutwardNum = ShowOutData.OutMoney
	saveMakeData.OutwardMoney = ShowOutData.OutCharge
	saveMakeData.Balance = data
	saveMakeData.FavourableMoney = ShowOutData.DepositMoney
	saveMakeData.ExpeneseMoney = ShowOutData.AdminMoney
	saveMakeData.DoAgencyAccount = "0"
	saveMakeData.DoAgencyId = 0
	saveMakeData.OutTime = 0
	saveMakeData.Remark = "出款"
	saveMakeData.DoUrl = ctx.Request().Host
	saveMakeData.OrderNum = ShowOutData.OrderNum //订单号
	if ShowOutData.DepositMoney > 0 {
		saveMakeData.FavourableOut = 1
	} else {
		saveMakeData.FavourableOut = 0
	}
	saveMakeData.CreateTime = ShowOutData.EndTime
	saveMakeData.ClientType = ShowOutData.ClientType
	cashRecord, err := drawMoney.WriteChargeRecord(saveMakeData, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if cashRecord == 0 {
		sess.Rollback()
		return ctx.JSON(200, global.ReplyError(91010, ctx))
	}
	sess.Commit()
	global.GetRedis().Del(redisKey)
	return ctx.JSON(200, global.ReplyItem(back.AddBankOut{OrderNumber: ShowOutData.OrderNum})) //返回订单号
}

//出款结果
func (c *RecordAjax) DrawWriteResult(ctx echo.Context) error {
	reqData := new(input.CompanyIncomeResult)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	err := global.WithdrawalCache.Set(reqData.OrderNum, reqData)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//出款页面加载数据
func (c *RecordAjax) GetDrawData(ctx echo.Context) error {
	memberinfo := ctx.Get("member").(*global.MemberRedisToken)
	sitepay, has, err := drawMoney.GetSiteSet(memberinfo.Site, memberinfo.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50084, ctx))
	}
	data, err := MemberBankBean.MemberBankList(memberinfo.Id)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for k, v := range data {
		v.Card = global.BankStr(v.Card)
		data[k].Card = v.Card
	}
	poundage, has, err := drawMoney.GetPoundageSet(memberinfo.Site, memberinfo.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(10201, ctx))
	}
	//realname
	has, err, realnameinfo := memberBean.AccountInfo(memberinfo.Account, memberinfo.Site, memberinfo.SiteIndex)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(20011, ctx))
	}
	memberOutInfo := new(back.MemberBankSiteSet)
	memberOutInfo.Poundage = poundage
	memberOutInfo.BankList = data
	memberOutInfo.SiteSet = sitepay
	memberOutInfo.RealName = realnameinfo

	return ctx.JSON(200, global.ReplyItem(memberOutInfo))
}

//获取登录的时候存储的redis值
func (c *RecordAjax) getWapMemberRedis(token string) (string, error) {
	key, err := global.GetRedis().Get(token).Result()
	return key, err
}

//获取5大模块下游戏分类
func (c *RecordAjax) GetGameList(ctx echo.Context) error {
	gametype := new(input.ProductTypeId)
	code := global.ValidRequestMember(gametype, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	memberinfo := ctx.Get("member").(*global.MemberRedisToken)
	data, err := productBean.GetGameList(memberinfo.Site, memberinfo.SiteIndex, gametype.TypeId)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取单条出款银行卡号
func (c *RecordAjax) GetOneBank(ctx echo.Context) error {

	bankinfo := new(input.OneMemberBankInfo)
	code := global.ValidRequest(bankinfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	bankinfo.MemberId = member.Id
	data, has, err := MemberBankBean.GetOneBank(bankinfo)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	data.Card = global.BankStr(data.Card)
	return ctx.JSON(200, global.ReplyItem(data))
}

//检测支付密码是否正确
func (c *RecordAjax) CheckDrawPass(ctx echo.Context) error {
	bankinfo := new(input.CheckDrawPass)
	code := global.ValidRequest(bankinfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	bankinfo.MemberId = member.Id
	data, has, err := MemberSelfInfoBean.CheckDrawPass(bankinfo)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(500, global.ReplyError(20018, ctx))
	}
	//加密密码
	password, err := global.MD5ByStr(bankinfo.DrawPassword, global.EncryptSalt)
	//fmt.Println("用户输入的加密密码：",password)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data.DrawPassword != password {
		return ctx.JSON(200, global.ReplyError(30234, ctx))
	}
	return ctx.NoContent(204)
}

//
//刷新用户余额
func (c *RecordAjax) AjaxGetBalance(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	balance, err := memberBalanceConversionBean.WapGetBalance(member.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(balance))
}

//获取优惠申请标题列表
func (c *RecordAjax) AjaxGetApplyTitleList(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	siteInfo := new(input.SitePromotionConfig)
	siteInfo.SiteId = member.Site
	siteInfo.SiteIndexId = member.SiteIndex
	data, err := SitePromotionConfigBean.GetSiteProConfig(siteInfo)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取彩票下列表
func (c *RecordAjax) AjaxFcList(ctx echo.Context) error {
	gameTypeInfo := new(input.GameList)
	code := global.ValidRequest(gameTypeInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var FcGameList []back.FcGameList
	switch gameTypeInfo.GameType {
	case "pk_fc":
		data, err := BetRecordInfoBean.GetPkGameList()
		if err != nil {
			global.GlobalLogger.Error("err:s%", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		var FcGame back.FcGameList
		for _, v := range data {
			FcGame.GameName = v.Name
			FcGame.GameType = v.Type
			FcGameList = append(FcGameList, FcGame)
		}
	case "cs_fc":
		data, err := BetRecordInfoBean.GetCsGameList()
		if err != nil {
			global.GlobalLogger.Error("err:s%", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		var FcGame back.FcGameList
		for _, v := range data {
			FcGame.GameName = v.CsName
			FcGame.GameType = v.CsType
			FcGameList = append(FcGameList, FcGame)
		}
	case "eg_fc":
		data, err := BetRecordInfoBean.GetEgGameList()
		if err != nil {
			global.GlobalLogger.Error("err:s%", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		var FcGame back.FcGameList
		for _, v := range data {
			FcGame.GameName = v.EgName
			FcGame.GameType = v.EgType
			FcGameList = append(FcGameList, FcGame)
		}
	}
	return ctx.JSON(200, global.ReplyItem(FcGameList))
}

// 会员优惠申请查询 MemberAutoApplypro
func (c *RecordAjax) AjaxGetApplyList(ctx echo.Context) error {
	getInfo := new(input.WapMemberCashRecord)
	code := global.ValidRequestMember(getInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	member := ctx.Get("member").(*global.MemberRedisToken)
	times := new(global.Times)
	times.StartTime = getInfo.StartTime
	times.EndTime = getInfo.EndTime
	listparam := new(global.ListParams)
	c.GetParam(listparam, ctx)
	data, count, err := selfHelpApplyforBean.GetMemberApplyList(member.Id, member.Site, member.SiteIndex, getInfo.OrderNum, times, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, data, int64(len(data)), count, ctx))
}

//查询会员出款是否有未出款
func (c *RecordAjax) AjaxGetDrawData(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	data, has, err := drawMoney.IsDrawData(member.Site, member.SiteIndex, member.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		data = nil
	}
	return ctx.JSON(200, global.ReplyItem(data))
}
