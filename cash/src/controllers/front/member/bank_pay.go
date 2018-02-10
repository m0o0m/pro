package member

import (
	"controllers"
	"errors"
	"fmt"
	"framework/uuid"
	"github.com/go-xorm/xorm"
	"github.com/golyu/sql-build"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"models/thirdParty"
	"strings"
	"sync"
)

type BankPayController struct {
	controllers.BaseController
}

// 额度转换
func (m *BankPayController) BalanceConversion(ctx echo.Context) error {
	ip := ctx.RealIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	// DESCRIPTION:注意:转入转出,针对的主体是视讯平台,不是本平台
	balanceConvert := new(input.BalanceConvert)
	code := global.ValidRequestMember(balanceConvert, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//类型:只能为转入in或者转出out
	if balanceConvert.TType != "out" && balanceConvert.TType != "in" {
		return ctx.JSON(200, global.ReplyError(60231, ctx))
	}
	//客户端只能是pc,wap,app
	if balanceConvert.Media != "pc" && balanceConvert.Media != "wap" && balanceConvert.Media != "app" {
		return ctx.JSON(200, global.ReplyError(60230, ctx))
	}

	if balanceConvert.Money < 10 {
		return ctx.JSON(200, global.ReplyError(30183, ctx))
	}
	currentTime := global.GetCurrentTime()
	// DESCRIPTION:获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)

	num, err := platformBean.IsExist(balanceConvert.PlatformId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num == 0 {
		//视讯平台不存在
		return ctx.JSON(500, global.ReplyError(60232, ctx))
	}
	// DESCRIPTION:初始化数据库sess
	sess := global.GetXorm().NewSession()
	sess.Begin()
	defer sess.Close()

	// DESCRIPTION:获取会员信息
	memberInfo, err := memberBean.GetMemberBySiteAccount(member.Site, member.Account, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	// DESCRIPTION:获取站点代理(业主)信息
	agency, err := memberBalanceConversionBean.GetAgency(member.Site, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	// DESCRIPTION:会员对应视讯额度表
	memberVideoTemp := new(schema.MemberProductClassifyBalance)
	memberVideoTemp.MemberId = member.Id
	memberVideoTemp.PlatformId = balanceConvert.PlatformId
	// DESCRIPTION:会员额度转换
	memberBalanceConvertTemp := new(schema.MemberBalanceConversion)
	memberBalanceConvertTemp.Money = float64(balanceConvert.Money)
	memberBalanceConvertTemp.DoUserId = member.Id
	memberBalanceConvertTemp.Account = member.Account
	memberBalanceConvertTemp.DoUserType = 2
	memberBalanceConvertTemp.SiteId = member.Site
	memberBalanceConvertTemp.SiteIndexId = member.SiteIndex
	memberBalanceConvertTemp.MemberId = member.Id
	memberBalanceConvertTemp.AgencyId = memberInfo.ThirdAgencyId
	memberBalanceConvertTemp.TradeNo = uuid.NewV4().String() //订单号
	memberBalanceConvertTemp.UpdateTime = currentTime

	// DESCRIPTION:会员现金记录表赋值
	cashRecord := new(schema.MemberCashRecord)
	cashRecord.SiteId = member.Site
	cashRecord.SiteIndexId = member.SiteIndex
	cashRecord.MemberId = member.Id
	cashRecord.UserName = member.Account
	cashRecord.AgencyId = memberInfo.ThirdAgencyId
	cashRecord.SourceType = 8
	cashRecord.Balance = float64(balanceConvert.Money) //操作金额
	cashRecord.Remark = "会员自助额度转换"

	// DESCRIPTION:站点现金记录表赋值
	siteCashRecord := new(schema.SiteCashRecord)
	siteCashRecord.SiteId = member.Site
	siteCashRecord.SiteIndexId = member.SiteIndex
	siteCashRecord.Remark = "会员自助额度转换"
	siteCashRecord.Money = float64(balanceConvert.Money)
	siteCashRecord.AdminName = member.Account
	siteCashRecord.CashType = 1                       //1代表额度转换
	siteCashRecord.VdType = balanceConvert.PlatformId //视讯类型
	siteCashRecord.CreateTime = currentTime           //当前时间

	// DESCRIPTION:进行转换前,需要传给对方的数据
	transferData := new(thirdParty.TransferData)
	transferData.Credit = float64(balanceConvert.Money)              //金额
	transferData.TradeNo = memberBalanceConvertTemp.TradeNo          //订单号
	transferData.SiteId = member.Site                                //站点
	transferData.IndexId = member.SiteIndex                          //前台
	transferData.UserName = member.Account                           //账号
	transferData.AgentId = memberInfo.ThirdAgencyId                  //代理
	transferData.UaId = memberInfo.SecondAgencyId                    //总代
	transferData.ShId = memberInfo.FirstAgencyId                     //股东
	transferData.Media = balanceConvert.Media                        //设备类型,pc wap app
	transferData.GameID = ""                                         //游戏子id
	transferData.IP = ip                                             //ip
	transferData.Lang = "ch"                                         //语言
	transferData.Cur = "CNY"                                         //货币类型
	transferData.Limit = "100000"                                    //限额
	transferData.Domain = ctx.Request().Host                         //登陆时域名
	transferData.IsSw = false                                        //试玩
	transferData.TransferType = balanceConvert.TType                 //out or in
	transferData.Platform = strings.ToLower(balanceConvert.Platform) //平台

	// DESCRIPTION:根据站点层级占成比,计算需要扣除或者加上的的站点视讯保证金
	//  获取站点对应层级的商品对应的占成比
	proportion, err := levelBean.GetProportionBySite(member.Site, balanceConvert.PlatformId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60223, ctx))
	}
	// 保证金
	marginMoney := global.FloatReserve2(proportion * 0.01 * float64(balanceConvert.Money))
	if marginMoney <= 0 {
		global.GlobalLogger.Error("not found margin(保证金)")
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var UpdateMoneyById func(string, int64, float64, ...*xorm.Session) (float64, error)
	var UpdateVideoMoneyById func(string, int64, float64, ...*xorm.Session) (float64, error)
	var count int64
	if balanceConvert.TType == "in" {
		// DESCRIPTION:转入 系统余额 --> 视讯余额

		// DESCRIPTION:会员余额是否足够转出金额
		if memberInfo.Balance < float64(balanceConvert.Money) {
			sess.Rollback()
			return ctx.JSON(200, global.ReplyError(30145, ctx))
		}
		// DESCRIPTION:判断视讯保证金是否小于扣除的保证金
		if agency.VideoBalance < marginMoney {
			global.GlobalLogger.Error("error:agency video balance:%f,fee:%f", agency.VideoBalance, marginMoney)
			sess.Rollback()
			return ctx.JSON(200, global.ReplyError(30158, ctx))
		}
		// DESCRIPTION:给现金记录表赋值
		cashRecord.Type = 2 //取出
		// DESCRIPTION:给站点现金记录表赋值
		siteCashRecord.DoType = 1 //取出
		// DESCRIPTION:会员额度转换表赋值
		memberBalanceConvertTemp.ForType = balanceConvert.PlatformId
		memberBalanceConvertTemp.FromType = 0

		UpdateMoneyById = memberBean.DelMoneyById
		UpdateVideoMoneyById = agencyBean.DelVideoMoneyById
	} else {
		// DESCRIPTION:转出 视讯余额 --> 系统余额
		// DESCRIPTION:视讯额度是否足够
		info, has, err := memberBalanceConversionBean.GetMoneyByVideo(member.Id, balanceConvert.PlatformId, sess)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			global.GlobalLogger.Error("error:没有转换过任何视讯,所以不存在余额")
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if info.Balance < float64(balanceConvert.Money) {
			sess.Rollback()
			return ctx.JSON(200, global.ReplyError(30146, ctx))
		}
		// DESCRIPTION:给现金记录表赋值
		cashRecord.Type = 1 //存入
		// DESCRIPTION:给站点现金记录表赋值
		siteCashRecord.DoType = 2 //取出
		// DESCRIPTION:会员额度转换表赋值
		memberBalanceConvertTemp.ForType = 0
		memberBalanceConvertTemp.FromType = balanceConvert.PlatformId

		UpdateMoneyById = memberBean.AddMoneyById
		UpdateVideoMoneyById = agencyBean.AddVideoMoneyById
	}

	// DESCRIPTION:额度转换
	videoBean := thirdParty.NewThirdParty()
	result, resultErr := videoBean.TransferCredit(transferData, 3) //重试3次

	if resultErr != nil { // DESCRIPTION:发生未知错误
		global.GlobalLogger.Error("resultErr:%s", resultErr.Error())
		memberBalanceConvertTemp.Status = 2 //会员额度转换 2失败
		siteCashRecord.State = 2            //站点现金记录 2掉单
		// DESCRIPTION:会员额度转换表插入
		count, err = sess.InsertOne(memberBalanceConvertTemp)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		// DESCRIPTION:未知错误,但是是转出(本系统加款),,就只在额度转换表中插入一个失败的记录
		if balanceConvert.TType == "out" {
			err = sess.Commit()
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
				sess.Rollback()
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			//掉单
			return ctx.JSON(500, global.ReplyError(60234, ctx))
		}
		// DESCRIPTION:****** 如果是转出 没有明确接收到服务器返回的错误code,均需要扣钱 ******
		// DESCRIPTION:会员表扣款
		newMoney, err := memberBean.DelMoneyById(memberInfo.SiteId, memberInfo.Id, float64(balanceConvert.Money), sess)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			sess.Rollback()
			return err
		}
		// DESCRIPTION:代理表扣除保证金
		newVideoBalance, err := agencyBean.DelVideoMoneyById(memberInfo.SiteId, memberInfo.ThirdAgencyId, marginMoney, sess)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			sess.Rollback()
			return err
		}
		// DESCRIPTION:给现金记录表赋值并插入
		cashRecord.AfterBalance = newMoney // 操作后余额
		count, err = sess.InsertOne(cashRecord)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			sess.Rollback()
			return err
		}
		if count == 0 {
			sess.Rollback()
			return errors.New("insert cashRecord 0 row")
		}
		// DESCRIPTION:给站点现金记录表赋值并插入
		siteCashRecord.Balance = newVideoBalance //站点视讯 余额
		count, err = sess.InsertOne(siteCashRecord)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			sess.Rollback()
			return err
		}
		if count == 0 {
			sess.Rollback()
			return errors.New("insert siteCashRecord 0 row")
		}
		err = sess.Commit()
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//掉单
		memberBalance, err := m.getBalance(member.Id, member.Account, memberInfo.Realname, newMoney)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}

		return ctx.JSON(500, global.ReplyErrItem(60224, memberBalance, ctx))
	}
	// DESCRIPTION:明确失败
	if result.Data.Code != 0 {
		sess.Rollback()
		return ctx.JSON(200, global.ReplyError(result.Data.Code, ctx))
	}
	// DESCRIPTION:明确成功
	siteCashRecord.State = 1            //站点现金记录 正常
	memberBalanceConvertTemp.Status = 1 //会员额度转换 1成功,
	// DESCRIPTION:会员视讯余额表插入或更新
	memberVideoTemp.Balance = result.Data.Balance
	sql, err := sqlBuild.Insert(memberVideoTemp.TableName()).
		Value(memberVideoTemp).
		OrUpdate().
		String()
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	execResult, err := sess.Exec(sql)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count, err := execResult.RowsAffected(); err != nil || count == 0 {
		if err != nil {
			sess.Rollback()
			if err != nil {
				global.GlobalLogger.Error("err:%s", err.Error())
			}
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//count更新0行不为错误,因为有可能就是存了这么多
		global.GlobalLogger.Info("err:sales_member_product_classify_balance insertOrUpdate 0 row")
	}
	// DESCRIPTION:会员额度转换表插入
	count, err = sess.InsertOne(memberBalanceConvertTemp)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if count == 0 {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	// DESCRIPTION:会员表扣款或加款
	newMoney, err := UpdateMoneyById(memberInfo.SiteId, memberInfo.Id, float64(balanceConvert.Money), sess)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	// DESCRIPTION:代理表扣除或加上保证金
	newVideoBalance, err := UpdateVideoMoneyById(memberInfo.SiteId, memberInfo.ThirdAgencyId, marginMoney, sess)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	// DESCRIPTION:给现金记录表赋值并插入
	cashRecord.AfterBalance = newMoney // 操作后余额
	count, err = sess.InsertOne(cashRecord)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	if count == 0 {
		sess.Rollback()
		return errors.New("insert cashRecord 0 row")
	}
	// DESCRIPTION:给站点现金记录表赋值并插入
	siteCashRecord.Balance = newVideoBalance //站点视讯 余额
	count, err = sess.InsertOne(siteCashRecord)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	if count == 0 {
		sess.Rollback()
		return errors.New("insert siteCashRecord 0 row")
	}
	err = sess.Commit()
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	memberBalance, err := m.getBalance(member.Id, member.Account, memberInfo.Realname, newMoney)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(memberBalance))
}

//查询返回
func (m *BankPayController) getBalance(memberId int64, account, name string, money float64) (interface{}, error) {
	// DESCRIPTION:查询余额
	sumBalance, balances, err := memberBalanceConversionBean.GetVideoBalance(memberId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return nil, err
	}
	memberBalance := new(back.MemberPlatformBalance)
	memberBalance.AccountBalance = money
	memberBalance.Account = account
	memberBalance.Realname = name
	memberBalance.GameBalance = sumBalance
	memberBalance.ProductClassifyBalance = balances
	return memberBalance, err
}

//个人余额
func (*BankPayController) Balance(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	data, err := memberBalanceConversionBean.GetPlatformBalance(member.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//视讯金额一键回归
func (c *BankPayController) Flyback(ctx echo.Context) error {
	ip := ctx.RealIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	media := "pc"
	currentTime := global.GetCurrentTime()
	// DESCRIPTION:获取操作者信息
	member := ctx.Get("member").(*global.MemberRedisToken)

	// DESCRIPTION:获取会员所有视讯额度
	_, memberProductClassifyBalancesTemp, err := memberBalanceConversionBean.GetVideoBalance(member.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var memberProductClassifyBalances []back.ProductClassifyBalance
	for _, v := range memberProductClassifyBalancesTemp {
		if v.Balance > 0 {
			memberProductClassifyBalances = append(memberProductClassifyBalances, v)
		}
	}
	if len(memberProductClassifyBalances) == 0 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60225, ctx))
	}

	// DESCRIPTION:获取会员信息
	memberInfo, err := memberBean.GetMemberBySiteAccount(member.Site, member.Account)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	// DESCRIPTION:会员额度转换表
	memberBalanceConvertTemp := new(schema.MemberBalanceConversion)
	memberBalanceConvertTemp.FromType = 0
	memberBalanceConvertTemp.DoUserId = memberInfo.Id
	memberBalanceConvertTemp.Account = memberInfo.Account
	memberBalanceConvertTemp.DoUserType = 2
	memberBalanceConvertTemp.SiteId = memberInfo.SiteId
	memberBalanceConvertTemp.SiteIndexId = memberInfo.SiteIndexId
	memberBalanceConvertTemp.MemberId = memberInfo.Id
	memberBalanceConvertTemp.AgencyId = memberInfo.ThirdAgencyId
	memberBalanceConvertTemp.UpdateTime = currentTime

	// DESCRIPTION:会员现金记录表
	cashRecord := new(schema.MemberCashRecord)
	cashRecord.SiteId = memberInfo.SiteId
	cashRecord.SiteIndexId = memberInfo.SiteIndexId
	cashRecord.MemberId = memberInfo.Id
	cashRecord.UserName = memberInfo.Account
	cashRecord.AgencyId = memberInfo.ThirdAgencyId
	cashRecord.SourceType = 8
	cashRecord.Remark = "视讯额度一键回归"
	cashRecord.Type = 1 //存入

	// DESCRIPTION:站点现金记录表赋值
	siteCashRecord := new(schema.SiteCashRecord)
	siteCashRecord.SiteId = memberInfo.SiteId
	siteCashRecord.SiteIndexId = memberInfo.SiteIndexId
	siteCashRecord.Remark = "视讯额度一键回归"
	siteCashRecord.AdminName = memberInfo.Account
	siteCashRecord.CashType = 1             //1代表额度转换
	siteCashRecord.DoType = 1               //存入
	siteCashRecord.CreateTime = currentTime //当前时间

	// DESCRIPTION:进行转换前,需要传给对方的数据
	transferData := new(thirdParty.TransferData)
	transferData.SiteId = member.Site               //站点
	transferData.IndexId = member.SiteIndex         //前台
	transferData.UserName = member.Account          //账号
	transferData.AgentId = memberInfo.ThirdAgencyId //代理
	transferData.UaId = memberInfo.SecondAgencyId   //总代
	transferData.ShId = memberInfo.FirstAgencyId    //股东
	transferData.Media = media                      //设备类型,pc wap app
	transferData.GameID = ""                        //游戏子id
	transferData.IP = ip                            //ip
	transferData.Lang = "ch"                        //语言
	transferData.Cur = "CNY"                        //货币类型
	transferData.Limit = "100000"                   //限额
	transferData.Domain = ctx.Request().Host        //登陆时域名
	transferData.IsSw = false                       //试玩
	transferData.TransferType = "out"               //out or in

	errs := make([]error, len(memberProductClassifyBalances))
	wg := new(sync.WaitGroup)
	var addSumMoney float64 //总共加了多少钱
	moneyMutex := new(sync.Mutex)
	for i, _ := range memberProductClassifyBalances {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errs[i] = c.flybackOne(
				&memberInfo,
				&memberProductClassifyBalances[i],
				*memberBalanceConvertTemp,
				*cashRecord,
				*siteCashRecord,
				*transferData,
			)
			if errs[i] == nil {
				moneyMutex.Lock()
				addSumMoney += memberProductClassifyBalances[i].Balance
				moneyMutex.Unlock()
			}
		}(i)
	}
	wg.Wait()
	// DESCRIPTION:获取各视讯平台余额
	sumBalance, balances, err := memberBalanceConversionBean.GetVideoBalance(member.Id)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
	}
	memberBalance := new(back.BalanceFlyback)
	memberBalance.AccountBalance = global.FloatReserve2(memberInfo.Balance + addSumMoney)
	memberBalance.Account = memberInfo.Account
	memberBalance.Realname = memberInfo.Realname
	memberBalance.GameBalance = sumBalance
	memberBalance.ProductClassifyBalance = balances
	for _, err := range errs {
		if err == nil {
			memberBalance.SuccessNum++
		} else {
			memberBalance.FailureNum++
		}
	}
	return ctx.JSON(200, global.ReplyItem(memberBalance))
}

//回归单个项目 return 单个平台余额信息 当前时间 设备类型 ip地址 当前站点网址 保证金
func (*BankPayController) flybackOne(
	member *schema.Member, //会员信息
	videoBalance *back.ProductClassifyBalance, //平台余额信息
	memberBalanceConvertTemp schema.MemberBalanceConversion, //会员额度转换
	cashRecord schema.MemberCashRecord, //会员现金记录
	siteCashRecord schema.SiteCashRecord, //站点现金记录
	transferData thirdParty.TransferData, //发起请求转换的数据
) (err error) {
	if videoBalance.Balance <= 0 {
		return
	}
	// DESCRIPTION:进行转换前,需要传给对方的数据赋值
	transferData.Credit = videoBalance.Balance                     //金额
	transferData.TradeNo = memberBalanceConvertTemp.TradeNo        //订单号
	transferData.Platform = strings.ToLower(videoBalance.Platform) //平台
	// DESCRIPTION:额度转换
	videoBean := thirdParty.NewThirdParty()
	result, resultErr := videoBean.TransferCredit(&transferData, 3) //重试3次
	if resultErr != nil {                                           // TODO 发生未知错误
		global.GlobalLogger.Error("resultErr:%s", resultErr.Error())
		memberBalanceConvertTemp.Status = 2 //会员额度转换 2失败
		siteCashRecord.State = 2            //站点现金记录 2掉单
		// DESCRIPTION:未知错误,但是是转入,就只在额度转换表中插入一个失败的记录
		sess := global.GetXorm().NewSession()
		defer sess.Close()

		var count int64
		// DESCRIPTION:会员额度转换表插入
		count, err = sess.InsertOne(&memberBalanceConvertTemp)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			return err
		}
		if count == 0 {
			err = errors.New("insert memberBalanceConvertTemp 0 row")
			return err
		}
		err = sess.Commit()
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			return err
		}
		return resultErr
	}
	if result.Data.Code != 0 { // DESCRIPTION:明确失败
		return errors.New("convert failure,error code:" + fmt.Sprintf("%d", result.Data.Code))
	}

	// DESCRIPTION:初始化数据库sess
	sess := global.GetXorm().NewSession()
	sess.Begin()
	defer sess.Close()

	// DESCRIPTION:会员额度转换表赋值
	memberBalanceConvertTemp.TradeNo = uuid.NewV4().String() //订单号
	memberBalanceConvertTemp.ForType = videoBalance.PlatformId
	memberBalanceConvertTemp.Money = videoBalance.Balance

	// DESCRIPTION:根据站点层级占成比,计算需要扣除或者加上的的站点视讯保证金
	//  获取站点对应层级的商品对应的占成比
	proportion, err := levelBean.GetProportionBySite(member.SiteId, videoBalance.PlatformId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return err
	}
	// 保证金
	margin := global.FloatReserve2(proportion * 0.01 * videoBalance.Balance)
	if margin <= 0 {
		global.GlobalLogger.Error("not found margin(保证金)")
		sess.Rollback()
		return errors.New("not found margin(保证金)")
	}

	// DESCRIPTION:****** 如果是转出 没有明确接收到服务器返回的错误code,均需要扣钱 ******
	// DESCRIPTION:明确成功
	siteCashRecord.State = 1            //站点现金记录 正常
	memberBalanceConvertTemp.Status = 1 //会员额度转换 1成功,
	// DESCRIPTION: 给会员视讯余额更新为0值
	err = memberBalanceConversionBean.ResetVideoBalance(member.Id, videoBalance.PlatformId, sess)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	// DESCRIPTION:会员表加款-更新
	newMoney, err := memberBean.AddMoneyById(member.SiteId, member.Id, videoBalance.Balance, sess)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	// DESCRIPTION:会员现金记录表赋值
	cashRecord.Balance = videoBalance.Balance //操作金额
	cashRecord.AfterBalance = newMoney        //操作后余额
	// DESCRIPTION:代理表加上保证金-更新
	newVideoBalance, err := agencyBean.AddVideoMoneyById(member.SiteId, member.ThirdAgencyId, margin, sess)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	// DESCRIPTION:给站点现金记录表赋值
	siteCashRecord.Money = videoBalance.Balance     //操作额度
	siteCashRecord.Balance = newVideoBalance        //站点视讯 余额
	siteCashRecord.VdType = videoBalance.PlatformId //视讯类型
	var count int64
	// DESCRIPTION:会员额度转换表插入
	count, err = sess.InsertOne(&memberBalanceConvertTemp)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	if count == 0 {
		sess.Rollback()
		return errors.New("insert memberBalanceConvertTemp 0 row")
	}

	// DESCRIPTION:会员现金记录表插入
	count, err = sess.InsertOne(&cashRecord)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	if count == 0 {
		sess.Rollback()
		return errors.New("insert cashRecord 0 row")
	}
	// DESCRIPTION:站点现金记录表插入
	count, err = sess.InsertOne(&siteCashRecord)
	if err != nil {
		global.GlobalLogger.Error("err:%s", err.Error())
		sess.Rollback()
		return err
	}
	if count == 0 {
		sess.Rollback()
		return errors.New("insert siteCashRecord 0 row")
	}
	return sess.Commit()
}
