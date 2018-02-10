package rebate

import (
	"controllers"
	"encoding/json"
	"fmt"
	"framework/uuid"
	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

const (
	REBATE_KEY = "rebate"
)

type RebateController struct {
	controllers.BaseController
}

//会员返佣统计
func (c *RebateController) Count(ctx echo.Context) error {
	reqData := new(input.MemberRebateCount)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	sTime, eTime, code := global.FormatDay2Timestamp(reqData.STime, reqData.ETime)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询时间区间是否存在(交叉也属于存在)
	commissions, err := rebateBean.GetCross(reqData.SiteId, reqData.SiteIndexId, sTime, eTime)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//只要有2个和以上的交叉项,就认为是绝对有交叉的,如果只有一个,有可能是重叠的
	if len(commissions) >= 2 {
		//交叉
		return ctx.JSON(500, global.ReplyError(70027, ctx))
	} else if len(commissions) == 1 {
		if commissions[0].StartTime != eTime || commissions[0].EndTime != eTime {
			//依旧交叉
			return ctx.JSON(500, global.ReplyError(70027, ctx))
		} else {
			//重叠
		}
	}

	//查询商品表信息
	products, err := productBean.GetList(reqData.SiteId, reqData.SiteIndexId, &input.ProductList{Status: 1})
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//js, _ := json.Marshal(products)
	//fmt.Println("所有商品:", string(js))
	if len(products) == 0 {
		return ctx.JSON(500, global.ReplyError(70013, ctx))
	}
	//查询站点返佣设定(优惠设定)
	rebateSets, err := rebateSetBean.GetAll(reqData.SiteId, reqData.SiteIndexId, products)
	js, _ := json.Marshal(rebateSets)
	fmt.Println("最初的设置:", string(js))
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//查询有效会员(有推广)的详细信息
	validMembers, err := memberBean.GetValidMemberList(reqData.SiteId, reqData.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(validMembers) == 0 {
		return ctx.JSON(500, global.ReplyError(70014, ctx))
	}
	//fmt.Println("有效:",len(validMembers))
	//for _, page := range validMembers {
	//	fmt.Printf("有效会员:%#v\n",page)
	//}

	preRebateRecordTotal := new(back.PreRebateRecordTotal) //最下一排的总计
	preRebateRecordTotalProductMap := make(map[string]*back.PreRebateRecordProduct)
	preRebateRecordTotal.AllProduct = &preRebateRecordTotalProductMap
	// --  根据有效会员,创建对应预返佣记录
	preRebateRecordMap := make(map[int64]*back.PreRebateRecord)
	for _, validMember := range validMembers {
		preRebateRecord := new(back.PreRebateRecord)

		//用来方便计算并存入redis的
		preRebateRecord.Id = validMember.Id
		preRebateRecord.Agency = validMember.Agency
		preRebateRecord.Account = validMember.Account

		//加上所有商品的结构
		preRebateRecordProductMap := make(map[string]*back.PreRebateRecordProduct)
		var preRebateRecordProductSlice []*back.PreRebateRecordProduct
		for _, product := range products {
			//单个会员总计结构初始化
			preRebateRecordProductMap[product.VType] = &back.PreRebateRecordProduct{ProductName: product.ProductName, VType: product.VType}
			//顺便将统计总计的结构初始化
			preRebateRecordTotalProductMap[product.VType] = &back.PreRebateRecordProduct{ProductName: product.ProductName, VType: product.VType}

			preRebateRecordProductSlice = append(preRebateRecordProductSlice, preRebateRecordProductMap[product.VType])
		}
		preRebateRecord.AllProduct = &preRebateRecordProductMap
		preRebateRecordMap[validMember.Id] = preRebateRecord
	}
	//fmt.Printf("有效会员 %#v\n", preRebateRecordMap)
	//for k, page := range preRebateRecordMap {
	//	fmt.Printf("有效会员 %d %#v  ::%#v\n", k, page, page.AllProduct)
	//}

	//统计被推广的会员的打码信息(预返佣信息,这里只有打码,返佣在后面计算)
	preRebateRecordProducts, err := betReportBean.CountValidBet(reqData.SiteId, reqData.SiteIndexId, sTime, eTime)
	//fmt.Println("长度", len(preRebateRecordProducts))
	//for _, page := range preRebateRecordProducts {
	//	fmt.Printf("详细打码信息: %#v\n", page)
	//}
	//打码赋值给对应会员预返佣记录
	// --- 将打码数据插入到每个对应会员
	for _, preRebateRecordProduct := range preRebateRecordProducts {
		preRebateRecord, ok := preRebateRecordMap[preRebateRecordProduct.SpreadId]
		if !ok {
			//这个会员名下有会员产生了打码,但是该站点没有这个会员,不参与计算,实际环境中,不会有这种情况
			global.GlobalLogger.Error("该站点不存在该会员 %d", preRebateRecordProduct.SpreadId)
			continue
		}
		preRebateRecordProductOk, ok := (*preRebateRecord.AllProduct)[preRebateRecordProduct.VType] //单个会员单个商品打码
		if !ok {
			//这个打码在剔除表中
			global.GlobalLogger.Error("已剔除 %d", preRebateRecordProduct.VType)
			continue
		}

		//打码
		preRebateRecordProductOk.ProductBet = preRebateRecordProduct.ProductBet
		preRebateRecordProductOk.ProductId = preRebateRecordProduct.ProductId

		preRebateRecord.AllBet += preRebateRecordProductOk.ProductBet //单个会员总打码

		preRebateRecordTotal.AllBet += preRebateRecordProductOk.ProductBet //多个会员总计打码
	}
	// --- 根据每个会员打码数据计算每个会员应该返佣
	for k, preRebateRecord := range preRebateRecordMap {
		//计算应该采用哪个返佣设置
		var rebateSet *back.MemberRebateSet
		for i, v := range rebateSets {
			if preRebateRecord.AllBet > float64(v.ValidMoney) {
				rebateSet = rebateSets[i]
				break
			}
		}
		//打码量没有达到,不予返佣
		if rebateSet == nil {
			continue
		}
		err := c.computeRebate(preRebateRecordMap[k], rebateSet, preRebateRecordTotal)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//将总计表统计好
	var rebateCommission back.MemberRebateCommission //总计打码
	rebateCommission.SiteIndexId = reqData.SiteIndexId
	rebateCommission.SiteId = reqData.SiteId
	rebateCommission.StartTime = sTime
	rebateCommission.EndTime = eTime
	rebateCommission.NoPeopleNum = 0
	rebateCommission.PeopleNum = preRebateRecordTotal.People
	rebateCommission.TotalBet = preRebateRecordTotal.AllBet
	rebateCommission.Money = preRebateRecordTotal.Rebate
	rebateCommission.RebateRecordMap = &preRebateRecordMap

	//返佣的会员数
	preRebateRecordTotal.People = int64(len(preRebateRecordMap))

	//存入redis
	key, err := c.saveRedisData(&rebateCommission)
	if err != nil {
		global.GlobalLogger.Error("redis error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//变换格式,前端展示
	preRebateRecordShows := make([]*back.PreRebateRecordShow, 0)
	for k, _ := range preRebateRecordMap {
		preRebateRecordShow := new(back.PreRebateRecordShow)
		preRebateRecordShow.Id = preRebateRecordMap[k].Id
		preRebateRecordShow.Agency = preRebateRecordMap[k].Agency
		preRebateRecordShow.Account = preRebateRecordMap[k].Account
		preRebateRecordShow.AllBet = preRebateRecordMap[k].AllBet
		preRebateRecordShow.Rebate = preRebateRecordMap[k].Rebate
		for v, _ := range *preRebateRecordMap[k].AllProduct {
			preRebateRecordShow.AllProduct = append(preRebateRecordShow.AllProduct, (*preRebateRecordMap[k].AllProduct)[v])
		}
		preRebateRecordShows = append(preRebateRecordShows, preRebateRecordShow)
	}
	preRebateRecordTotalShow := new(back.PreRebateRecordTotalShow)
	preRebateRecordTotalShow.People = preRebateRecordTotal.People
	preRebateRecordTotalShow.AllBet = preRebateRecordTotal.AllBet
	preRebateRecordTotalShow.Rebate = preRebateRecordTotal.Rebate
	for k, _ := range *preRebateRecordTotal.AllProduct {
		preRebateRecordTotalShow.AllProduct = append(preRebateRecordTotalShow.AllProduct, (*preRebateRecordTotal.AllProduct)[k])
	}

	return ctx.JSON(200, global.ReplyCollections("content", preRebateRecordShows, "total", preRebateRecordTotalShow, "key", key))
}

//将统计出的数据存入到redis
func (*RebateController) saveRedisData(src *back.MemberRebateCommission) (key string, err error) {
	key = uuid.NewV4().String()
	js, err := json.Marshal(src)
	if err != nil {
		return
	}
	global.GetRedis().HSet(REBATE_KEY, key, js)
	return
}

//从redis中取出统计信息
func (c *RebateController) getRedisData(key string) (dst back.MemberRebateCommission, err error) {
	src, err := global.GetRedis().HGet(REBATE_KEY, key).Result()
	if err != nil {
		global.GlobalLogger.Error("%s", err.Error())
		return
	}
	err = json.Unmarshal([]byte(src), &dst)
	if err != nil {
		global.GlobalLogger.Error("%s", err.Error())
	}
	return
}

//根据有效打码,返佣设置,计算出应该返佣
func (*RebateController) computeRebate(record *back.PreRebateRecord, sets *back.MemberRebateSet, total *back.PreRebateRecordTotal) error {
	//js, _ := json.Marshal(sets.ProductRates)
	//fmt.Printf("rate  %s \n", string(js))

	for _, set := range *sets.ProductRates {
		singleProduct, ok := (*record.AllProduct)[set.VType]
		totalProduct, ok2 := (*total.AllProduct)[set.VType]
		if !ok {
			return global.GlobalLogger.Error("error:member <%s> not found product <%s>", record.Account, set.VType)
		}
		if !ok2 {
			return global.GlobalLogger.Error("error:count product is not exist<%s>", record.Account, set.VType)
		}
		singleProduct.Money = global.FloatReserve2(singleProduct.ProductBet * set.Rate / 100)              //单个会员的单个商品返点 (保留2位小数,4舍五入)
		record.Rebate = global.FloatReserve2(record.Rebate + singleProduct.Money)                          //单个会员返点总计
		totalProduct.Money = global.FloatReserve2(totalProduct.Money + singleProduct.Money)                //多个会员单个商品返点
		totalProduct.ProductBet = global.FloatReserve2(totalProduct.ProductBet + singleProduct.ProductBet) //多个会员单个商品打码
	}
	if record.Rebate > float64(sets.DiscountUp) {
		record.Rebate = float64(sets.DiscountUp)
	}
	total.Rebate = global.FloatReserve2(total.Rebate + record.Rebate) //多个会员返点总计
	return nil
}

//会员返佣统计-存入
func (c *RebateController) Commit(ctx echo.Context) error {
	reqData := new(input.MemberRebateCommit)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if len(reqData.MemberIds) == 0 {
		return ctx.JSON(200, global.ReplyError(70019, ctx))
	}
	rebateCommission, err := c.getRedisData(reqData.Key)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(70019, ctx))
	}
	rebateCommission.Event = reqData.Event //事件
	rebateCommission.Bet = reqData.BetRate //综合打码倍率

	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()

	var audits []*schema.MemberAudit           //准备存入的稽核
	var cashRecords []*schema.MemberCashRecord //现金记录

	//存入优惠总计
	err = rebateBean.SaveRebateCommission(&rebateCommission, sess)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var rebateRecords []*schema.MemberRebateRecord //需要存入的返佣记录
	//js,_:=json.Marshal(rebateCommission)
	//fmt.Println("这个:",reqData.MemberIds,string(js))

	for _, memberId := range reqData.MemberIds {
		preRebateRecord, ok := (*rebateCommission.RebateRecordMap)[memberId]
		if !ok {
			global.GlobalLogger.Error("There is no need to deposit the rebate data")
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		//返佣为0的不用存
		if preRebateRecord.Rebate == 0 {
			continue
		}
		//将优惠记录准备好
		rebateRecord := new(schema.MemberRebateRecord)
		rebateRecord.SiteId = rebateCommission.SiteId
		rebateRecord.SiteIndexId = rebateCommission.SiteIndexId
		rebateRecord.PeriodsId = rebateCommission.Id  //期数id(总计id)
		rebateRecord.MemberId = memberId              // 所属会员
		rebateRecord.Betting = preRebateRecord.AllBet //总投注
		rebateRecord.Rebate = preRebateRecord.Rebate  //总返佣
		rebateRecord.Status = 1
		for _, preRebateProduct := range *preRebateRecord.AllProduct {
			rebateRecordProduct := new(schema.MemberRebateRecordProduct)
			rebateRecordProduct.ProductId = preRebateProduct.ProductId   //商品id
			rebateRecordProduct.ProductBet = preRebateProduct.ProductBet //打码
			rebateRecordProduct.Money = preRebateProduct.Money           //返佣
			rebateRecord.RebateRecordProducts = append(rebateRecord.RebateRecordProducts, rebateRecordProduct)
		}
		rebateRecords = append(rebateRecords, rebateRecord)
	}
	//存入返佣记录
	err = rebateBean.SaveRebateRecord(rebateRecords, sess)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	var rebateRecordProducts []*schema.MemberRebateRecordProduct // 需要存入的返佣的详情
	for i := range rebateRecords {
		for k := range rebateRecords[i].RebateRecordProducts {
			rebateRecords[i].RebateRecordProducts[k].RecordId = rebateRecords[i].Id
			if rebateRecords[i].RebateRecordProducts[k].ProductId > 0 {
				rebateRecordProducts = append(rebateRecordProducts, rebateRecords[i].RebateRecordProducts[k])
			}
		}
	}
	//存入返佣记录对应的商品
	err = rebateBean.SaveRebateRecordProduct(rebateRecordProducts, sess)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//多个会员加款
	// --- 查询会员信息,因为是多个会员,为了避免反复操作后面要用到,所以先查出来
	var memberIds []int64 //会员id
	for _, rebateRecord := range rebateRecords {
		memberIds = append(memberIds, rebateRecord.MemberId)
	}
	auditMembers, err := memberBean.GetAuditMemberByIds(memberIds, sess) //待加款的会员
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(auditMembers) != len(rebateRecords) {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(70024, ctx))
	}
	auditMemberMap := make(map[int64]*back.RebateAuditMember)
	for k, _ := range auditMembers {
		auditMemberMap[auditMembers[k].Id] = &auditMembers[k]
	}
	// --- 在程序中加款后,在数据库中更新
	for _, rebateRecord := range rebateRecords {
		auditMember, ok := auditMemberMap[rebateRecord.MemberId]
		if !ok {
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(70024, ctx))
		}
		auditMember.Amount = rebateRecord.Rebate                                              //操作金额 (返佣多少钱)
		auditMember.Balance = global.FloatReserve2(auditMember.Balance + rebateRecord.Rebate) //剩余多少钱

		//设置好需要插入的现金记录
		cashRecord := new(schema.MemberCashRecord)
		cashRecord.MemberId = rebateRecord.MemberId
		cashRecord.SiteId = rebateRecord.SiteId
		cashRecord.SiteIndexId = rebateRecord.SiteIndexId
		cashRecord.UserName = auditMember.Account            //会员账号
		cashRecord.AgencyId = auditMember.AgencyId           //代理id
		cashRecord.AgencyAccount = auditMember.AgencyAccount //代理账号
		cashRecord.SourceType = 11                           //会员返佣
		cashRecord.Type = 1                                  //存入
		cashRecord.TradeNo = ""                              //下单单号
		cashRecord.Balance = auditMember.Amount              //操作金额
		cashRecord.AfterBalance = auditMember.Balance        //操作后余额
		cashRecord.Remark = "member rebate"                  //备注
		cashRecord.ClientType = reqData.ClientType           //客户端类型
		cashRecords = append(cashRecords, cashRecord)

		//稽核记录
		audit := new(schema.MemberAudit)
		audit.SiteId = rebateRecord.SiteId
		audit.SiteIndexId = rebateRecord.SiteIndexId
		audit.MemberId = rebateRecord.MemberId
		audit.Account = auditMember.Account
		audit.BeginTime = rebateCommission.StartTime
		audit.EndTime = rebateCommission.EndTime
		audit.NormalMoney = 0                                                                      //常态稽核,这里属于综合稽核
		audit.MultipleMoney = global.FloatReserve2(rebateRecord.Rebate * float64(reqData.BetRate)) //综合稽核
		audit.Money = rebateRecord.Rebate
		audit.Status = 1 // 未操作
		audits = append(audits, audit)
	}
	//
	//js, _ := json.Marshal(&auditMembers)
	//fmt.Println("update money ", string(js))

	//更新多个会员的金额,这里先轮询,之后改异步
	for k, _ := range auditMembers {
		err = memberBean.UpdateMoney(&auditMembers[k])
		if err != nil {
			sess.Rollback()
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	//添加多条现金记录
	_, err = cashRecordBean.AddCashRecordMulti(cashRecords, sess)
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//存入多条稽核
	num, err := auditBean.InsertAuditMulti(audits, sess)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != int64(len(audits)) {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	num, err = auditBean.InsertAuditMulti(audits, sess)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != int64(len(audits)) {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	err = sess.Commit()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//会员返佣-冲销
func (c *RebateController) Writeoff(ctx echo.Context) error {
	reqData := new(input.RebateWriteoff)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//查询返佣详情
	rebateRecords, err := rebateRecordBean.GetRebateRecordById(reqData.RecordIds)

	if len(rebateRecords) != len(reqData.RecordIds) {
		return ctx.JSON(500, global.ReplyError(70021, ctx))
	}
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	var cashRecords []*schema.MemberCashRecord //现金记录
	var auditLogs []*schema.MemberAuditLog     //稽核日志
	//冲销掉返佣记录
	num, err := rebateRecordBean.WriteoffRebateRecordById(reqData.RecordIds, sess)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if num != int64(len(reqData.RecordIds)) {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(70022, ctx))
	}
	//删除对应商品
	_, err = rebateRecordBean.DelRebateRecordProductByRecordIds(reqData.RecordIds, sess)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//多个会员扣款
	// --- 查询会员信息,因为是多个会员,为了避免反复操作后面要用到,所以先查出来
	var memberIds []int64 //会员id
	for _, rebateRecord := range rebateRecords {
		memberIds = append(memberIds, rebateRecord.MemberId)
	}
	auditMembers, err := memberBean.GetAuditMemberByIds(memberIds, sess) //待扣款的会员
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if len(auditMembers) != len(rebateRecords) {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(70024, ctx))
	}
	auditMemberMap := make(map[int64]*back.RebateAuditMember)
	for k, _ := range auditMembers {
		auditMemberMap[auditMembers[k].Id] = &auditMembers[k]
	}
	// --- 在程序中扣款后,在数据库中更新
	for _, rebateRecord := range rebateRecords {
		auditMember, ok := auditMemberMap[rebateRecord.MemberId]
		if !ok {
			sess.Rollback()
			return ctx.JSON(500, global.ReplyError(70024, ctx))
		}
		auditMember.Amount = rebateRecord.Rebate                                              //操作金额 (冲销多少钱)
		auditMember.Balance = global.FloatReserve2(auditMember.Balance - rebateRecord.Rebate) //剩余多少钱

		//设置好需要插入的现金记录
		cashRecord := new(schema.MemberCashRecord)
		cashRecord.MemberId = rebateRecord.MemberId
		cashRecord.SiteId = rebateRecord.SiteId
		cashRecord.SiteIndexId = rebateRecord.SiteIndexId
		cashRecord.UserName = auditMember.Account            //会员账号
		cashRecord.AgencyId = auditMember.AgencyId           //代理id
		cashRecord.AgencyAccount = auditMember.AgencyAccount //代理账号
		cashRecord.SourceType = 11                           //会员返佣
		cashRecord.Type = 2                                  //取出
		cashRecord.TradeNo = ""                              //下单单号
		cashRecord.Balance = auditMember.Amount              //操作金额
		cashRecord.AfterBalance = auditMember.Balance        //操作后余额
		cashRecord.Remark = "Member commission rebates"      //备注
		cashRecord.ClientType = reqData.ClientType           //客户端类型
		cashRecords = append(cashRecords, cashRecord)

		//设置好需要插入的稽核日志
		auditLog := new(schema.MemberAuditLog)
		auditLog.SiteId = rebateRecord.SiteId
		auditLog.SiteIndexId = rebateRecord.SiteIndexId
		auditLog.MemberId = auditMember.Id
		auditLog.Account = auditMember.Account
		auditLog.Content = auditMember.Account + "Member commission rebates, cancel audit" //备注
		auditLog.Type = 1
		auditLogs = append(auditLogs, auditLog)
	}
	//更新多个会员的金额,这里先轮询,之后改异步
	for k, _ := range auditMembers {
		err = memberBean.UpdateMoney(&auditMembers[k])
		if err != nil {
			sess.Rollback()
			global.GlobalLogger.Error("error:%s", err.Error())
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}

	//添加多条现金记录
	_, err = cashRecordBean.AddCashRecordMulti(cashRecords, sess)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//根据时间和会员id修改稽核状态
	_, err = auditBean.OverAudit(rebateRecords[0].CreateTime, memberIds)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//插入稽核日志
	_, err = auditLogBean.InsertMulti(auditLogs, sess)
	if err != nil {
		sess.Rollback()
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	err = sess.Commit()
	if err != nil {
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//会员返佣查询-列表
func (c *RebateController) List(ctx echo.Context) error {
	reqData := new(input.RebateList)
	code := global.ValidRequest(reqData, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var sTime int64
	var eTime int64
	if reqData.Year != "" && reqData.Month != "" {
		sTime, eTime, code = global.FormatMonth2Timestamp(reqData.Year + "-" + reqData.Month)
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
	}
	rebateCommissions, err := rebateBean.GetRebateList(reqData.SiteId, reqData.SiteIndexId, sTime, eTime)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyCollection(rebateCommissions, int64(len(rebateCommissions))))
}

//会员返佣查询-明细
func (c *RebateController) Details(ctx echo.Context) error {
	reqData := new(input.RebateDetails)
	code := global.ValidRequest(reqData, ctx)

	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	var result = make(map[string]interface{}) //结果根节点
	var total = new(back.SumRebateDetail)
	total.SumPeople = reqData.SumPeople
	total.NoPeople = reqData.NoPeople
	var sumProducts = make(map[string]*back.RebateRecordProduct) //按商品分类的总打码和总返佣
	total.AllProduct = &sumProducts

	//根据期数id查询详情
	rebateRecords, err := rebateRecordBean.GetRebateRecord(reqData.PeriodsId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//查询出商品信息
	products, err := productBean.GetList(reqData.SiteId, reqData.SiteIndexId, &input.ProductList{Status: 1})
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	for i, _ := range rebateRecords {
		total.AllBet += rebateRecords[i].AllBet
		total.Rebate += rebateRecords[i].Rebate
		allInfo := make(map[string]*back.RebateRecordProduct)
		rebateRecords[i].AllProduct = &allInfo
		for _, product := range products {
			if i == 0 {
				sumProducts[product.ProductName] = new(back.RebateRecordProduct)
			}
			(*rebateRecords[i].AllProduct)[product.ProductName] = new(back.RebateRecordProduct)
		}
	}

	//根据详情id查询商品金额
	var recordIds []int
	for _, v := range rebateRecords {
		recordIds = append(recordIds, v.Id)
	}
	var rebateRecordProducts []back.RebateRecordProduct
	if len(recordIds) > 0 {
		rebateRecordProducts, err = rebateRecordProductBean.GetRecordProductListByRecordIds(recordIds)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	indexes := make(map[int]int) //快速获取下标
	for k, rebateRecord := range rebateRecords {
		indexes[rebateRecord.Id] = k
	}

	for i, rebateRecordProduct := range rebateRecordProducts {
		detail := rebateRecords[indexes[rebateRecordProduct.RecordId]]
		(*detail.AllProduct)[rebateRecordProduct.ProductName] = &rebateRecordProducts[i]

		temp2, ok := sumProducts[rebateRecordProduct.ProductName]
		if ok {
			temp2.ProductBet += rebateRecordProduct.ProductBet
			temp2.Money += rebateRecordProduct.Money
		} else {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
	}
	result["data"] = rebateRecords
	result["total"] = total
	return ctx.JSON(200, global.ReplyItem(&result))
}
