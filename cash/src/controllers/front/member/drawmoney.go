package member

import (
	"encoding/json"
	"fmt"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strconv"
	"time"
)

//取款管理
type DrawBean struct {
}

//取款(没有给稽核日志表添加数据)
func (*DrawBean) Withdrawal(ctx echo.Context) error {
	wapDrawMoney := new(input.DrawMoney)
	code := global.ValidRequest(wapDrawMoney, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//获取登录会员信息
	member := ctx.Get("member").(*global.MemberRedisToken)
	draw_money, _ := strconv.ParseFloat(wapDrawMoney.Money, 64)
	bankId, _ := strconv.ParseInt(wapDrawMoney.BankId, 10, 64)
	//获取会员余额和取款密码
	balance, password, has, err := drawMoney.GetMemberBalanceAndPassword(member.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(20232, ctx))
	}
	//取款金额是否大于会员余额
	if draw_money > balance {
		return ctx.JSON(200, global.ReplyError(20233, ctx))
	}
	//密码加密
	drawPassword, err := global.MD5ByStr(wapDrawMoney.DrawPassword, global.EncryptSalt)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//取款密码是否正确
	if password != drawPassword {
		return ctx.JSON(200, global.ReplyError(20234, ctx))
	}
	//会员出款银行id是否存在
	has, err = drawMoney.IsExistMemberBankById(member.Id, bankId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(50116, ctx))
	}
	//查看出款管理表是否有待审核的数据
	has, err = drawMoney.IsExistPendingReview(member.Id)
	if err != nil {
		global.GlobalLogger.Error("err:", err)
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if has {
		//存在待审核数据就不允许再次出款
		return ctx.JSON(200, global.ReplyError(30236, ctx))
	}
	//出款展示数据 最终成型的稽核数据结构体
	sdata := new(back.MemberAuditOutBack)
	//客户端类型
	system := ctx.Request().Header.Get("platform")
	switch system {
	case "pc":
		sdata.ClientType = 1
	case "wap":
		sdata.ClientType = 2
	case "android":
		sdata.ClientType = 3
	case "ios":
		sdata.ClientType = 4
	}
	//会员操作后余额
	//memberBalance := balance - draw_money
	//查询稽核表
	data, err := drawMoney.GetDrawData(member)
	fmt.Println(data)
	//切割时间
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	var etime = time.Now().Unix() //当前时间
	if data != nil {
		var ftime = data[0].BeginTime       //第一笔稽核开始时间
		today0, today1 := global.GetToday() //获取当天0点和23点59分59秒的时间戳
		var reportValid []back.AuditBet
		var sbettime, ebettime int64
		mp := make(map[int]*back.AuditBetRecord)
		if ftime > today0 && ftime < today1 {
			//稽核计算日期在当天，则只需要查询注单表即可
			sbettime = today0 //注单表起始时间
			ebettime = etime  //注单表起始时间
			//查询注单表
			betValid, err := drawMoney.GetBetValid(member.Account, member.Site, sbettime, ebettime)
			fmt.Println(betValid)
			for _, v := range betValid {
				mp[v.GameType] = &v
			}
			if err != nil {
				global.GlobalLogger.Error("err:s%", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
		} else {
			//切割时间
			sreporttime := global.GetTimeStart(data[0].BeginTime) //统计表起始时间
			ereporttime := today0                                 //统计表结束时间
			sbettime = today0                                     //注单表起始时间
			ebettime = etime                                      //注单表起始时间
			fmt.Println(sreporttime, ereporttime, sbettime, ebettime)
			//查询统计表
			reportValid, err = drawMoney.GetReportValid(member.Account, member.Site, sreporttime, ereporttime)
			if err != nil {
				global.GlobalLogger.Error("err:s%", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			//查询注单表
			betValid, err := drawMoney.GetBetValid(member.Account, member.Site, sbettime, ebettime)
			if err != nil {
				global.GlobalLogger.Error("err:s%", err.Error())
				return ctx.JSON(500, global.ReplyError(60000, ctx))
			}
			for k, v := range betValid {
				mp[v.GameType] = &betValid[k]
			}
			for k, v := range reportValid {
				if m, ok := mp[v.GameType]; ok {
					m.BetValid += reportValid[k].BetValid
				} else {
					temp := back.AuditBetRecord{
						SiteId:      reportValid[k].SiteId,
						SiteIndexId: reportValid[k].SiteIndexId,
						Username:    reportValid[k].Account,
						BetValid:    reportValid[k].BetValid,
						GameType:    reportValid[k].GameType,
					}
					mp[v.GameType] = &temp
				}
			}
			mergeValid, _ := json.Marshal(&mp)
			global.GlobalLogger.Error("%s", string(mergeValid))
		}
		sdata.BeginTime = ftime                     //稽核开始时间
		sdata.EndTime = etime                       //稽核结束时间
		sdata.Money = data[0].Money                 //存款金额
		sdata.NormalMoney = data[0].NormalMoney     //常态稽核金额
		sdata.MultipleMoney = data[0].MultipleMoney //综合稽核金额
		sdata.AdminMoney = data[0].AdminMoney       //扣除的行政费用
		sdata.DepositMoney = data[0].DepositMoney   //优惠金额
		sdata.RelaxMoney = data[0].RelaxMoney       //放宽额度
		//初始化打码数据
		sdata.Vdbet = 0
		sdata.Dzbet = 0
		sdata.Spbet = 0
		sdata.Fcbet = 0
		if mp != nil { //打码数据不为空，整合有效打码
			for _, v := range mp {
				if v.GameType == 1 {
					sdata.Vdbet = v.BetValid
				}
				if v.GameType == 2 || v.GameType == 3 {
					sdata.Dzbet += v.BetValid
				}
				if v.GameType == 4 {
					sdata.Fcbet = v.BetValid
				}
				if v.GameType == 5 {
					sdata.Spbet = v.BetValid
				}
			}
		}
		sdata.Allbet = sdata.Vdbet + sdata.Dzbet + sdata.Fcbet + sdata.Spbet //所有打码合计
		if sdata.Allbet+float64(sdata.RelaxMoney) >= sdata.NormalMoney {     //判断常态稽核是否通过
			sdata.IsNormal = 1 //通过
			sdata.AdminMoney = 0.00
		} else {
			sdata.IsNormal = 2 //未通过       扣除行政费
		}
		if sdata.MultipleMoney > 0 { //开启了综合稽核
			if sdata.Allbet >= sdata.MultipleMoney { //综合稽核是否通过
				sdata.IsMultiple = 1 //通过
				sdata.DepositMoney = 0.00
			} else {
				sdata.IsMultiple = 2 //未通过，扣除优惠
			}
		} else {
			sdata.DepositMoney = 0.00
		}
		//手续费
		//获取站点手续费设置
		sitepay, has, err := drawMoney.GetSiteSet(member.Site, member.SiteIndex)
		if err != nil {
			global.GlobalLogger.Error("err:s%", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if !has {
			return ctx.JSON(200, global.ReplyError(60000, ctx))
		}
		outcharge := 0.00
		if sitepay.IsFree == 2 { // 不免手续费
			//统计当天已经出款次数
			var numOut int64
			numOut, err = drawMoney.NumOutMoney(member.Account, member.Site, today0, today1)
			if numOut > int64(sitepay.FreeNum) {
				outcharge = sitepay.OutCharge
			}
		}
		//需要扣除的费用 優惠稽核 + 常態性稽核 + 手续费
		outfee := sdata.AdminMoney + sdata.DepositMoney + outcharge
		sdata.OutMoney = global.FloatReserve2(draw_money)
		sdata.OutCharge = sdata.OutMoney - global.FloatReserve2(outfee)
		if draw_money > outfee {
			sdata.OutStatus = 1
		} else {
			sdata.OutStatus = 2
		}
	} else {
		sdata.BeginTime = etime
		sdata.EndTime = etime
		sdata.Money = 0              //存款金额
		sdata.AdminMoney = 0         //扣除的行政费用
		sdata.DepositMoney = 0       // 优惠金额
		sdata.MultipleMoney = 0      //综合稽核
		sdata.NormalMoney = 0        //常态稽核
		sdata.RelaxMoney = 0         //放宽额度
		sdata.Vdbet = 0              // 视讯打码
		sdata.Dzbet = 0              //电子打码
		sdata.Spbet = 0              // 体育打码
		sdata.Fcbet = 0              // 彩票打码
		sdata.Allbet = 0             //打码总计
		sdata.IsNormal = 1           //通过   //通过常态稽核 1为通过，2为未通过，未通过扣取行政费用
		sdata.IsMultiple = 1         //通过  //是否通过综合稽核 1为通过，2为未通过，未通过扣取优惠金额
		sdata.Charge = 0             //出款手续费
		sdata.OutMoney = draw_money  //提交出款金额
		sdata.OutCharge = draw_money //实际可出金额
		sdata.OutStatus = 1          //可出款状态 1未可出款 2未不可出款（提出金额减去费用小于0）
	}
	key := uuid.NewV4().String()
	sdata.OrderNum = key
	redisKey := member.Site + "_" + member.SiteIndex + "_" + strconv.FormatInt(member.Id, 10) + "outCharge"
	b, err := json.Marshal(sdata)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	err = global.GetRedis().Set(redisKey, b, 0).Err()
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//redis 存储
func wapMemberRedisSet(result string, b []byte, beforeKey string) (err error) {
	if beforeKey != "" {
		//删除旧的key
		err = global.GetRedis().Del(beforeKey).Err()
		//将旧的删除
		err = global.GetRedis().LPop(result).Err()
	}
	//存储新token
	err = global.GetRedis().Set(result, b, 0).Err()
	//将推进list
	err = global.GetRedis().RPush("member_login", result).Err()
	return err
}

//获取个人银行信息和出款银行列表
func (m *DrawBean) GetMemberBankList(ctx echo.Context) error {
	member := ctx.Get("member").(*global.MemberRedisToken)
	mi := new(input.MemberBankList)
	mi.MemberId = member.Id
	mi.SiteId = member.Site
	mi.SiteIndexId = member.SiteIndex
	data, err := baseInfoBean.MemberBankList(mi)
	if err != nil {
		global.GlobalLogger.Error("err:s%", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//获取单条出款银行卡号
func (m *DrawBean) GetOneBank(ctx echo.Context) error {

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
	return ctx.JSON(200, global.ReplyItem(data))
}
