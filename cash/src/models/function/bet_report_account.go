package function

import (
	"database/sql"
	"errors"
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
	"time"

	"github.com/go-xorm/xorm"
)

type BetReportAccountBean struct {
}

//优惠统计
func (bra *BetReportAccountBean) CountBetReportAccount(this *input.CountBetReportAccount) (
	back.CountAllBetReportAccountTotalMap, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	hadRetreat := new(back.CountBetReportAccountTotalMap)
	var data back.CountAllBetReportAccountTotalMap
	//unRetreat := new(back.CountBetReportAccountTotal)
	unRetreat := new(back.CountBetReportAccountTotalMap)
	data.StartTime = this.StartTime
	data.EndTime = this.EndTime
	/*优惠统计流程：
	先查询退水记录表，看是否存在时间重合或交叉
	如果 有交叉或重合
		如果完全重合
			对比记录，添加记录
				记录来自查询
				记录来自sales_member_retreat_water_record
					查询退水记录表+退水记录明细表+退水记录商品表 -> 返回对应时间段的数据
	    如果有交叉
			返回err:输入的时间区间不合法
	否则
		添加记录
			记录来自查询
				商品表+站点商品剔除表 -> 本站点最终退水设定的商品
				优惠退水设定表+优惠退水商品表+商品表 -> 本站点最终退水设定详细内容
				查询会员每日打码统计表+商品表 -> 会员在各个商品下的综合打码量
				自助返水表+自助商品表 -> 查询本站点该区间内的自助返水金额，计算出总退水金额
				综合打码量与设定的打码对比取出比率相乘得出 退水结果，再与 优惠上限 对比得出最终结果。
				封装返回结果
	*/
	//--------------------------先查询退水记录表，看是否存在时间重合或交叉----------------------
	var has bool
	_, has, err := bra.recordAcrossBetReportAccount(this)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, err
	}
	if has { //有交叉或重复
		var record back.ListRetreatWater
		var has2 bool
		record, has2, err := bra.recordBetReportAccount(this)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
		if has2 { //重复
			recordsNew := make([]back.RetreatWaterRecord, 0)
			//accounts := make([]string,0)
			betValedsNew := make([]back.BetValidBetReportAccountList, 0)
			var records []back.RetreatWaterRecord
			records, _, err = bra.recordsBetReportAccount(record.Id)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, err
			}
			fmt.Println("记录：", records)
			var betValeds []back.BetValidBetReportAccountList
			betValeds, err = bra.betValidBetReportAccount(this)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, err
			}
			var checkAccount = 0
			for _, b := range betValeds {
				checkAccount = 0
				for _, r := range records {
					if b.Account == r.Account { //交叉的已退水的会员
						recordsNew = append(recordsNew, r)
						//accounts=append(accounts, r.Account)
						checkAccount = 1
						break
					}
				}
				if checkAccount == 0 {
					betValedsNew = append(betValedsNew, b) //未退水的会员
				}
			}
			//fmt.Println("交叉的已退水的会员", recordsNew)
			//fmt.Println("未退水的会员", betValedsNew)
			//--------------------------已退水会员记录列表------------------------
			*hadRetreat = bra.hadRetreatWaterBetReportAccount(recordsNew)

			//--------------------------组装封装[]back.RetreatWaterSetList-------
			var waters []back.RetreatWaterSetList
			waters, err = bra.waterBetReportAccount(this)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, err
			}
			fmt.Println("退水设定列表", waters)
			//--------------------------会员自助返水查询--------------------------
			var am map[string]float64
			am, err = bra.seachWaterSelf(this) //map[account]Money
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return data, err
			}
			//fmt.Println("自助返水记录列表", am)
			//--------------------------综合打码量与设定--------------------------
			*unRetreat = bra.compareBetReportAccount(waters, betValedsNew, am)
			data.BetTotalAll = hadRetreat.BetTotal + unRetreat.BetTotal
			data.PeopleNumAll = hadRetreat.PeopleNum + unRetreat.PeopleNum
			data.MoneyAll = hadRetreat.Money + unRetreat.Money
			data.UnRetreat = *unRetreat
			data.HadRetreat = *hadRetreat
			fmt.Println("已退水", hadRetreat)
			fmt.Println("未退水", unRetreat)
			fmt.Println("最终数据", data)
		} else { //交叉
			err = errors.New("日期区间交叉了！")
			return data, err
		}
	} else { //没记录
		//fmt.Println("没记录1111111111111")
		//--------------------------组装封装[]back.RetreatWaterSetList-------
		var waters []back.RetreatWaterSetList
		waters, err = bra.waterBetReportAccount(this)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
		//fmt.Println("退水设定列表", waters)
		//--------------------------综合打码量提取----------------------------
		var betValeds []back.BetValidBetReportAccountList
		betValeds, err = bra.betValidBetReportAccount(this)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
		//fmt.Println("综合打码量", betValeds)
		//--------------------------会员自助返水查询--------------------------
		var am map[string]float64
		am, err = bra.seachWaterSelf(this) //map[account]Money
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return data, err
		}
		//--------------------------综合打码量与设定--------------------------
		*unRetreat = bra.compareBetReportAccount(waters, betValeds, am)
		data.MoneyAll = unRetreat.Money
		data.PeopleNumAll = unRetreat.PeopleNum
		data.BetTotalAll = unRetreat.BetTotal
		data.UnRetreat = *unRetreat
		data.HadRetreat = *hadRetreat //没有，为空
	}

	return data, err

}

//存总表sales_member_retreat_water
//todo 未写稽核 稽核日志 存现金记录
func (bra *BetReportAccountBean) StoreBetReportAccount(this *input.StoreBetReportAccount, redisKey back.CountAllBetReportAccountTotalMap, user *global.RedisStruct) error {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	sess.Begin()
	/*
		接收参数
		取出redis里的统计数据
		通过会员账号选取redis统计里的数据，并封装存入数据库

		事务开始
			存总表：
				insert 总表
				返回总表主键id
			存明细表+存明细商品表：
				从redis中取出所有传入的member_ids未退水的会员
				根据主键id查询已冲销记录
				判断是否有 未退水与已冲销 交叉
					循环组装明细记录
						没有交叉
							insert 未交叉明细 明细集合对象组装1
								循环组装商品1(无id)
						有交叉
							update 交叉的明细 明细集合对象组装2
								循环组装商品1(有id)
					循环结束
				insert 明细集合对象组装1
				查询 明细集合对象组装1 返回id
				循环组装商品1 组装id  insert 循环组装商品1

				删除 明细集合对象组装2
				删除 循环组装商品1(有id)
				insert 明细集合对象组装2
				insert 循环组装商品1(有id)
		事务结束
	*/
	//audits := make([]*schema.MemberAudit, 0)           //准备存入的稽核
	//cashRecords := make([]*schema.MemberCashRecord, 0) //准备存入的现金记录
	record := new(schema.MemberRetreatWater) //准备存入的总表

	record.SiteId = this.SiteId
	record.SiteIndexId = this.SiteIndexId
	record.AdminUser = user.Account
	record.LevelId = string(user.Level)
	record.StartTime = redisKey.StartTime
	record.EndTime = redisKey.EndTime
	record.CreateTime = time.Now().Unix()
	record.Event = this.Event
	record.NoPeopleNum = 0                          //冲销人数
	record.PeopleNum = redisKey.UnRetreat.PeopleNum //未退水的人数
	record.Money = redisKey.UnRetreat.Money         //未退水的金额
	record.Bet = this.Bet
	var result sql.Result
	result, err := sess.Exec("INSERT INTO sales_member_retreat_water (site_id, site_index_id, admin_user,level_id,start_time,end_time,create_time,event,no_people_num,people_num,money,bet) VALUES(?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE admin_user=values(admin_user),level_id=values(level_id),create_time=values(create_time),event=values(event),no_people_num=no_people_num+values(no_people_num),people_num=people_num+values(people_num),money=money+values(money),bet=values(bet)", record.SiteId, record.SiteIndexId, record.AdminUser, record.LevelId, record.StartTime, record.EndTime, record.CreateTime, record.Event, record.NoPeopleNum, record.PeopleNum, record.Money, record.Bet) //在原记录上累加冲销人数，退水的人数，未退水的金额----insert 总表
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		sess.Rollback()
		return err
	}
	periodsId, err := result.LastInsertId() //----返回总表主键id
	fmt.Println("最后的id", periodsId)
	//---------------------------------添加明细记录----------------------------
	//添加明细记录
	//查询是否有未退水与已冲销交叉
	var cxrecords []back.RetreatWaterRecord2
	fmt.Println("redis的会员", redisKey.UnRetreat.Subtotal)
	//取出所有传入的要退水的会员
	unRetreat := make(map[string]*back.CountBetReportAccountSubtotalMap)
	for k, s := range redisKey.UnRetreat.Subtotal { //----从redis中取出所有传入的member_ids未退水的会员
		for _, a := range this.MemberIds {
			if s.MemberId == a {
				unRetreat[k] = s
			}
		}
	}
	records := make([]schema.MemberRetreatWaterRecord, 0)                       //未退水且无冲销记录
	chrecords := make([]schema.MemberRetreatWaterRecord, 0)                     //重合冲销记录
	recordPros := make([]schema.MemberRetreatWaterRecordProduct, 0)             //商品集合
	recordProsMap := make(map[string]*[]schema.MemberRetreatWaterRecordProduct) //无重合商品map
	chrecordPros := make([]schema.MemberRetreatWaterRecordProduct, 0)           //重合的冲销的商品
	cxrecordIds := make([]int64, 0)                                             //重合的记录id
	recordMemberIds := make([]int64, 0)                                         //无重合的会员id
	ids := make([]back.SearchIdBetReportAccount, 0)                             //添加明细后查询明细id
	cxaccount := make(map[string]int64, 0)                                      //存储冲销与未返水重合的会员
	cxrecords, err = bra.cqrecordsBetReportAccount(periodsId)                   //----根据主键id查询已冲销记录
	for _, c := range cxrecords {                                               //	 ----判断是否有 未退水与已冲销 交叉
		_, ok := unRetreat[c.Account]
		if ok { //交叉的已冲销的会员 更新
			cxaccount[c.Account] = c.Id
		}
	}
	fmt.Println("传入的会员", unRetreat)
	fmt.Println("冲销的会员", cxrecords)
	fmt.Println("交叉的会员", cxaccount)
	//------------------------数据组装-----------------------
	//如果有重合的冲销会员，添加是取出冲销会员，更新冲销会员的状态
	for _, c := range unRetreat { //	----循环组装明细记录
		recordone := new(schema.MemberRetreatWaterRecord)
		recordone.SiteId = this.SiteId
		recordone.SiteIndexId = this.SiteIndexId
		recordone.Account = c.Account
		recordone.StartTime = redisKey.StartTime
		recordone.EndTime = redisKey.EndTime
		recordone.PeriodsId = periodsId
		recordone.MemberId = c.MemberId
		recordone.LevelId = c.LevelId
		recordone.Betall = c.Betall
		recordone.AllMoney = c.AllMoney
		recordone.SelfMoney = c.SelfMoney
		recordone.RebateWater = c.RebateWater
		recordone.Status = 1
		recordone.CreateTime = time.Now().Unix()
		recordPros = nil
		var i int64 = 0
		for _, p := range c.Products {
			recordProone := new(schema.MemberRetreatWaterRecordProduct)
			recordProone.ProductId = p.ProductId
			recordProone.ProductBet = p.ProductBet
			recordProone.Rate = p.Rate
			recordProone.Money = p.Money
			if v, ok := cxaccount[c.Account]; ok { //有冲销记录的要封装id
				if i == 0 {
					cxrecordIds = append(cxrecordIds, v) //交叉的id存入数组
				}
				//cxrecordIds[c.Account]=v
				recordProone.RecordId = cxaccount[c.Account]
				recordone.Id = cxaccount[c.Account]
				chrecordPros = append(chrecordPros, *recordProone) //把所有冲销的商品封装到一个商品slice里 冲销商品
			}
			recordPros = append(recordPros, *recordProone)
			i = i + 1
		}
		if _, ok := cxaccount[c.Account]; !ok { //如果不是重合的冲销会员
			recordProsMap[recordone.Account] = &recordPros        //会员商品集合存入map 未退水商品  ----循环组装商品1(无id)
			records = append(records, *recordone)                 //未退水记录                    ----明细集合对象组装1
			recordMemberIds = append(recordMemberIds, c.MemberId) //未退水无交叉会员id集合，用于insert后查询record_id的条件
		} else { //如果是重合的冲销会员
			//chrecordProsMap[recordone.Account] = recordPros //重合冲销会员商品集合存入map
			chrecords = append(chrecords, *recordone) //冲销记录									 ----明细集合对象组装2
		}
	}
	//--------------------------无交叉的-----------------------
	fmt.Println("未退水无交叉会员id集合", recordMemberIds)
	fmt.Println("未退水的会员数据组装", records)
	fmt.Println("重合的的会员数据组装", chrecords)
	_, err = sess.Table(new(schema.MemberRetreatWaterRecord).TableName()).Insert(records) // ----insert 明细集合对象组装1
	if err != nil {                                                                       //回滚
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	//查询记录ID，组装未交叉的商品集合
	ids, err = bra.searchIdBetReportAccount(recordMemberIds, this, redisKey, sess) //----查询 明细集合对象组装1 返回id
	if err != nil {                                                                //回滚
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	//组装未退水且无交叉的商品集合,并添加到数据库
	fmt.Println("ids", ids)
	fmt.Println("商品map", recordProsMap)
	err = bra.mergeProductBetReportAccount(ids, recordProsMap, sess) //----循环组装商品1 组装id  insert 循环组装商品1
	if err != nil {                                                  //回滚
		sess.Rollback()
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	fmt.Println("未退水与冲销重合的状态更新", chrecords)
	fmt.Println("交叉的id组", cxrecordIds)
	//--------------------------有退水与冲销交叉的-----------------------
	if len(chrecords) > 0 { //如果有冲销的会员要退水，则update(通过delete+insert实现),状态更新为1，商品无需更新
		delRecord := new(schema.MemberRetreatWaterRecord)
		_, err = sess.Table(delRecord.TableName()).In("id", cxrecordIds).Delete(delRecord) //----删除 明细集合对象组装2
		if err != nil {                                                                    //回滚
			sess.Rollback()
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		delRecordPro := new(schema.MemberRetreatWaterRecordProduct)
		_, err = sess.Table(delRecordPro.TableName()).In("record_id", cxrecordIds).Delete(delRecordPro) //----删除 循环组装商品1(有id)
		if err != nil {                                                                                 //回滚
			sess.Rollback()
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		_, err = sess.Table(delRecord.TableName()).Insert(chrecords) //  ----insert 明细集合对象组装2
		if err != nil {                                              //回滚
			sess.Rollback()
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		_, err = sess.Table(delRecordPro.TableName()).Insert(chrecordPros) //  ----insert 循环组装商品1(有id)
		if err != nil {                                                    //回滚
			sess.Rollback()
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
	}
	sess.Commit()
	sess.Close()
	return err
}

//组装未退水且无交叉的商品集合,并添加到数据库
func (*BetReportAccountBean) mergeProductBetReportAccount(ids []back.SearchIdBetReportAccount, recordProsMap map[string]*[]schema.MemberRetreatWaterRecordProduct, sess *xorm.Session) (err error) {
	recordProsMerge := make([]schema.MemberRetreatWaterRecordProduct, 0)
	recordPro := new(schema.MemberRetreatWaterRecordProduct)
	for _, id := range ids {
		v, ok := recordProsMap[id.Account]
		if ok {
			for _, v2 := range *v {
				recordPro.RecordId = id.Id
				recordPro.ProductId = v2.ProductId
				recordPro.ProductBet = v2.ProductBet
				recordPro.Rate = v2.Rate
				recordPro.Money = v2.Money
				recordProsMerge = append(recordProsMerge, *recordPro) //把所有商品都添加到一个集合里
			}
		}
	}
	//sess := global.GetXorm().NewSession()
	//defer sess.Close()
	_, err = sess.Table(recordPro.TableName()).Insert(recordProsMerge)
	return
}

//查询明细记录id
func (*BetReportAccountBean) searchIdBetReportAccount(memberIds []int64, this *input.StoreBetReportAccount, redisKey back.CountAllBetReportAccountTotalMap, sess *xorm.Session) (data []back.SearchIdBetReportAccount, err error) {
	//sess := global.GetXorm().NewSession()
	//defer sess.Close()
	table := new(schema.MemberRetreatWaterRecord)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId) //1
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId) //2
	}
	sess.Where("start_time=?", redisKey.StartTime)
	sess.Where("end_time=?", redisKey.EndTime)
	sess.In("member_id", memberIds)
	err = sess.Table(table.TableName()).
		Select("id,account,member_id").Find(&data)
	return
}

//已退水会员数据组装
func (*BetReportAccountBean) hadRetreatWaterBetReportAccount(records []back.RetreatWaterRecord) (data back.CountBetReportAccountTotalMap) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	//dataRecord := new(back.CountBetReportAccountSubtotalMap)
	//dataRecordPro := new(back.CountBetReportAccountProduct)
	//dataRecords:=make([]back.CountBetReportAccountSubtotal,0)
	//dataRecordPros:=make([]back.CountBetReportAccountProduct,0)
	data.PeopleNum = int64(len(records))
	countSubMap := make(map[string]*back.CountBetReportAccountSubtotalMap)
	for _, r := range records {
		data.BetTotal = data.BetTotal + r.Betall
		data.Money = data.Money + r.RebateWater
		//dataRecord = nil
		dataRecord := new(back.CountBetReportAccountSubtotalMap)
		dataRecord.Account = r.Account
		dataRecord.MemberId = r.MemberId
		dataRecord.LevelId = r.LevelId
		dataRecord.Betall = r.Betall
		dataRecord.AllMoney = r.AllMoney
		dataRecord.SelfMoney = r.SelfMoney
		dataRecord.RebateWater = r.RebateWater
		countProductsMap := make(map[string]*back.CountBetReportAccountProduct)
		for _, p := range r.Params {
			dataRecordPro := new(back.CountBetReportAccountProduct)
			dataRecordPro.ProductId = p.ProductId
			dataRecordPro.ProductName = p.ProductName
			dataRecordPro.Rate = p.Rate
			dataRecordPro.ProductBet = p.ProductBet
			dataRecordPro.Money = p.Money
			//dataRecord.Products[dataRecordPro.ProductName] = dataRecordPro
			countProductsMap[dataRecordPro.ProductName] = dataRecordPro
			//dataRecord.Products = append(dataRecord.Products, *dataRecordPro)
		}
		dataRecord.Products = countProductsMap
		countSubMap[dataRecord.Account] = dataRecord
		//data.Subtotal = append(data.Subtotal, *dataRecord)
	}
	data.Subtotal = countSubMap

	return
}

//查询退水记录表是否有交叉
func (*BetReportAccountBean) recordAcrossBetReportAccount(this *input.CountBetReportAccount) (waters back.ListRetreatWater, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWater)
	//waterRecord:=new(schema.MemberRetreatWaterRecord)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId) //1
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId) //2
	}
	sess.Where("((start_time>=? and start_time<=?) or (end_time>=? and end_time<=?))", this.StartTime, this.EndTime, this.StartTime, this.EndTime)
	has, err = sess.Table(water.TableName()).
		Select("id,admin_user,start_time,end_time,create_time,event,no_people_num,people_num,money,bet").Get(&waters)
	if err != nil {
		return
	}
	return
}

//查询退水记录表是否有记录
func (*BetReportAccountBean) recordBetReportAccount(this *input.CountBetReportAccount) (waters back.ListRetreatWater, has bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWater)
	//waterRecord:=new(schema.MemberRetreatWaterRecord)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId) //1
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId) //2
	}
	sess.Where("start_time=?", this.StartTime)
	sess.Where("end_time=?", this.EndTime)
	//sess.Table(water.TableName()).Select("id,people_num,money").Find()

	has, err = sess.Table(water.TableName()).
		Select("id,admin_user,start_time,end_time,create_time,event,no_people_num,people_num,money,bet").Get(&waters)
	if err != nil {
		return
	}
	return
}

//退水明细记录表 已退水的会员
func (*BetReportAccountBean) recordsBetReportAccount(id int64) (data []back.RetreatWaterRecord, count int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWaterRecord)
	sess.Where("sales_member_retreat_water_record.periods_id=?", id)
	sess.Where("sales_member_retreat_water_record.status=1")
	conds := sess.Conds()
	waterPro := new(schema.MemberRetreatWaterRecordProduct)
	product := new(schema.Product)
	sql2 := fmt.Sprintf("%s.id = %s.record_id", water.TableName(), waterPro.TableName())
	sql3 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), waterPro.TableName())
	data1 := make([]back.RetreatWaterRecordAndProduct, 0)
	err = sess.Table(water.TableName()).
		Select("sales_member_retreat_water_record.id,sales_member_retreat_water_record.account,member_id,level_id,betall,all_money,self_money,rebate_water,sales_member_retreat_water_record.status,sales_member_retreat_water_record.create_time,product_id,product_name,product_bet,rate,money").
		Join("LEFT", waterPro.TableName(), sql2).
		Join("LEFT", product.TableName(), sql3).OrderBy("member_id").
		Find(&data1)
	if err != nil {
		return
	}
	var checkid int64 = 0
	waterPros := make([]back.RetreatWaterRecordProduct, 0)
	waterPro2 := new(back.RetreatWaterRecordProduct)
	water2 := new(back.RetreatWaterRecord)
	for _, d := range data1 {
		//有效打码，上限
		if checkid != d.MemberId {
			if checkid != 0 {
				water2.Params = waterPros    //商品列表 组装到总组装Params参数
				data = append(data, *water2) //总组装
			}
			waterPros = nil
			checkid = d.MemberId

			water2.Account = d.Account
			water2.MemberId = d.MemberId
			water2.LevelId = d.LevelId
			water2.Betall = d.Betall
			water2.AllMoney = d.AllMoney
			water2.SelfMoney = d.SelfMoney
			water2.RebateWater = d.RebateWater
			water2.Status = d.Status
			water2.CreateTime = d.CreateTime
			//组装商品列表
			waterPro2.ProductId = d.ProductId
			waterPro2.ProductName = d.ProductName
			waterPro2.ProductBet = d.ProductBet
			waterPro2.Rate = d.Rate
			waterPro2.Money = d.Money
			waterPros = append(waterPros, *waterPro2)
		} else { //id相同时，只组装商品列表
			waterPro2.ProductId = d.ProductId
			waterPro2.ProductName = d.ProductName
			waterPro2.ProductBet = d.ProductBet
			waterPro2.Rate = d.Rate
			waterPro2.Money = d.Money
			waterPros = append(waterPros, *waterPro2)
		}
	}
	water2.Params = waterPros
	if checkid != 0 {
		data = append(data, *water2)
	}
	count, err = sess.Table(water.TableName()).Where(conds).Count()
	return
}

//退水明细记录表 已冲销的会员
func (*BetReportAccountBean) cqrecordsBetReportAccount(id int64) (data []back.RetreatWaterRecord2, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWaterRecord)
	sess.Where("sales_member_retreat_water_record.periods_id=?", id)
	sess.Where("sales_member_retreat_water_record.status=2")
	//conds := sess.Conds()
	waterPro := new(schema.MemberRetreatWaterRecordProduct)
	product := new(schema.Product)
	sql2 := fmt.Sprintf("%s.id = %s.record_id", water.TableName(), waterPro.TableName())
	sql3 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), waterPro.TableName())
	data1 := make([]back.RetreatWaterRecordAndProduct, 0)
	err = sess.Table(water.TableName()).
		Select("sales_member_retreat_water_record.id,sales_member_retreat_water_record.account,member_id,level_id,betall,all_money,self_money,rebate_water,sales_member_retreat_water_record.status,sales_member_retreat_water_record.create_time,product_id,product_name,product_bet,rate,money").
		Join("LEFT", waterPro.TableName(), sql2).
		Join("LEFT", product.TableName(), sql3).OrderBy("member_id").
		Find(&data1)
	if err != nil {
		return
	}
	var checkid int64 = 0
	waterPros := make([]back.RetreatWaterRecordProduct, 0)
	waterPro2 := new(back.RetreatWaterRecordProduct)
	water2 := new(back.RetreatWaterRecord2)
	for _, d := range data1 {
		//有效打码，上限
		if checkid != d.MemberId {
			if checkid != 0 {
				water2.Params = waterPros    //商品列表 组装到总组装Params参数
				data = append(data, *water2) //总组装
			}
			waterPros = nil
			checkid = d.MemberId

			water2.Id = d.Id
			water2.Account = d.Account
			water2.MemberId = d.MemberId
			water2.LevelId = d.LevelId
			water2.Betall = d.Betall
			water2.AllMoney = d.AllMoney
			water2.SelfMoney = d.SelfMoney
			water2.RebateWater = d.RebateWater
			water2.Status = d.Status
			water2.CreateTime = d.CreateTime
			//组装商品列表
			waterPro2.ProductId = d.ProductId
			waterPro2.ProductName = d.ProductName
			waterPro2.ProductBet = d.ProductBet
			waterPro2.Rate = d.Rate
			waterPro2.Money = d.Money
			waterPros = append(waterPros, *waterPro2)
		} else { //id相同时，只组装商品列表
			waterPro2.ProductId = d.ProductId
			waterPro2.ProductName = d.ProductName
			waterPro2.ProductBet = d.ProductBet
			waterPro2.Rate = d.Rate
			waterPro2.Money = d.Money
			waterPros = append(waterPros, *waterPro2)
		}
	}
	water2.Params = waterPros
	if checkid != 0 {
		data = append(data, *water2)
	}
	//count, err = sess.Table(water.TableName()).Where(conds).Count()
	return
}

//商品剔除,退水设置;剔除后对应的比例，上限，有效总投注;遍历结果，组装返回值，返回[]back.RetreatWaterSetList
func (*BetReportAccountBean) waterBetReportAccount(this *input.CountBetReportAccount) (waters []back.RetreatWaterSetList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	product := new(schema.Product)
	productDel := new(schema.SiteProductDel)
	productRetWater := new(schema.MemberRetreatWaterProduct)
	waterSet := new(schema.MemberRetreatWaterSet)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId) //1
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId) //2
	}
	conds := sess.Conds() //封装条件1，2

	//--------------------------商品剔除,退水设置----------------------
	productIds := make([]int64, 0)
	sql3 := fmt.Sprintf("%s.product_id = %s.id", productDel.TableName(), product.TableName())
	sess.Table("sales_site_product_del").Select("sales_product.id").
		Join("LEFT", product.TableName(), sql3).
		GroupBy("sales_product.id").
		Find(&productIds) //查询要剔除的商品id 条件1，2
	fmt.Println(productIds)
	if err != nil {
		return
	}
	//--------------------无需会员层级条件,组合4个条件进行查询-----------
	sess.NotIn("sales_product.id", productIds)                    //加入剔除商品条件  3
	sess.Where("sales_member_retreat_water_set.delete_time=?", 0) //4
	sess.Where(conds)                                             //1,2
	//查询本站点的所有商品剔除后对应的比例，上限，有效总投注
	data1 := make([]back.RetreatWaterSetListAndProduct, 0)
	sql := fmt.Sprintf("%s.id = %s.set_id", waterSet.TableName(), productRetWater.TableName())
	sql2 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), productRetWater.TableName())
	err = sess.Table(waterSet.TableName()).
		Select("sales_member_retreat_water_set.id,valid_money,discount_up,product_id,product_name,rate").
		Join("LEFT", productRetWater.TableName(), sql).
		Join("LEFT", product.TableName(), sql2).OrderBy("valid_money desc,sales_member_retreat_water_set.id,product_id").
		Find(&data1)
	if err != nil {
		return
	}

	var checkid int64 = 0
	//waters := make([]back.RetreatWaterSetList, 0)
	waterPros := make([]back.RetreatWaterProductList, 0)
	water2 := new(back.RetreatWaterSetList)
	waterPro2 := new(back.RetreatWaterProductList)
	for _, d := range data1 {
		if checkid != d.Id { //id不同时，组装有效打码，上限，商品列表，总组装
			if checkid != 0 {
				water2.Params = waterPros        //商品列表 组装到总组装Params参数
				waters = append(waters, *water2) //总组装
			}
			waterPros = nil
			checkid = d.Id

			//有效打码，上限
			water2.ValidMoney = d.ValidMoney
			water2.DiscountUp = d.DiscountUp
			//组装商品列表
			waterPro2.ProductId = d.ProductId
			waterPro2.ProductName = d.ProductName
			waterPro2.Rate = d.Rate
			waterPros = append(waterPros, *waterPro2)
		} else { //id相同时，只组装商品列表
			waterPro2.ProductId = d.ProductId
			waterPro2.ProductName = d.ProductName
			waterPro2.Rate = d.Rate
			waterPros = append(waterPros, *waterPro2)
		}

	}
	water2.Params = waterPros
	if checkid != 0 {
		waters = append(waters, *water2)
	}
	//查询本站点商品剔除后对应的比例，上限，有效总投注
	//fmt.Println(waters)
	return
}

//综合打码量提取
func (*BetReportAccountBean) betValidBetReportAccount(this *input.CountBetReportAccount) (betValid []back.BetValidBetReportAccountList, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()                       //封装条件1，2
	betRepAc := new(schema.BetReportAccount) //综合打码量统计表约束
	if this.System == 1 {                    //通过会员层级查会员
		if this.SiteId != "" {
			sess.Where("site_id=?", this.SiteId) //1
		}
		if this.SiteIndexId != "" {
			sess.Where("site_index_id=?", this.SiteIndexId) //2
		}
		sess.In("level_id", this.LevelId)
		accounts := make([]string, 0)
		err = sess.Table("sales_member").Select("account").Find(&accounts) //使用条件1，2
		if err != nil {
			return
		}
		sess.Where("sales_bet_report_account.site_id=?", this.SiteId)            //1
		sess.Where("sales_bet_report_account.site_index_id=?", this.SiteIndexId) //2
		sess.In("sales_bet_report_account.account", accounts)                    //3
	} else if this.System == 2 { //直接查会员
		sess.Where("sales_bet_report_account.site_id=?", this.SiteId)            //1
		sess.Where("sales_bet_report_account.site_index_id=?", this.SiteIndexId) //2
		sess.In("sales_bet_report_account.account", this.Account)                //3
	}
	sess.Where("sales_bet_report_account.create_time>=?", this.StartTime)
	sess.Where("sales_bet_report_account.create_time<=?", this.EndTime)
	sql4 := fmt.Sprintf("%s.v_type = %s.v_type", betRepAc.TableName(), "sales_product")
	sql5 := fmt.Sprintf("%s.account = %s.account", betRepAc.TableName(), "sales_member")

	err = sess.Table(betRepAc.TableName()).Select("sales_bet_report_account.account,member_id,sales_member.level_id,bet_valid,sales_product.id as product_id,sales_product.product_name").
		Join("LEFT", "sales_product", sql4).
		Join("LEFT", "sales_member", sql5).OrderBy("account").
		Find(&betValid)
	if err != nil {
		return
	}
	return
}

//会员自助返水查询
func (*BetReportAccountBean) seachWaterSelf(this *input.CountBetReportAccount) (selfWater map[string]float64, err error) {
	//selfWater:=make(map[string]float64)//返回map
	waterSelf := make([]back.ListRetreatWaterSelf, 0)
	water := new(schema.MemberRetreatWaterSelf)
	sess := global.GetXorm().Table(water.TableName())
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("create_time>=?", this.StartTime)
	sess.Where("create_time<=?", this.EndTime)
	sess.In("account", this.Account)
	//sess.Where("order_num=?", this.OrderNum)
	err = sess.Table(water.TableName()).
		Select("id,order_num,account,betting,money,create_time").Find(&waterSelf)
	if err != nil {
		return
	}
	m := make(map[string]float64)
	for _, b := range waterSelf {
		m[b.Account] = b.Money //返回map
	}
	selfWater = m
	fmt.Println(selfWater)
	return
}

//综合打码量与设定的打码对比,自助返水小计,组装数据
func (*BetReportAccountBean) compareBetReportAccount(rws []back.RetreatWaterSetList, bvbral []back.BetValidBetReportAccountList, am map[string]float64) (data back.CountBetReportAccountTotalMap) {
	//productMap := make(map[int64]back.RetreatWaterProductList)
	//waterRetTotal:=make([]back.CountBetReportAccountTotal,0)
	//waterRetSubtotal := make([]back.CountBetReportAccountSubtotal, 0)
	//countProducts := make([]back.CountBetReportAccountProduct, 0)

	//countProduct := new(back.CountBetReportAccountProduct)
	//countSub := new(back.CountBetReportAccountSubtotalMap)
	countSubMap := make(map[string]*back.CountBetReportAccountSubtotalMap)

	for _, b := range bvbral { //循环打码量
		data.BetTotal = data.BetTotal + b.BetValid //累加总有效投注  三
		for i, r := range rws {                    //循环优惠设定 按validbet排序   1条打码量 对应 1条设定 中的对应 1个商品
			countProductsMap := make(map[string]*back.CountBetReportAccountProduct)
			if b.BetValid >= float64(r.ValidMoney) { // 判断打码量是否大于设定打码量
				v, ok := countSubMap[b.Account] //判断是否添加此条打码的会员
				if ok {                         //如果已有会员信息 修改会员Map
					for _, p := range r.Params { //找打码量里与设定里对应的商品比例进行计算
						if p.ProductId == b.ProductId { //退水额计算
							v2, ok2 := v.Products[b.ProductName]
							if ok2 {
								v2.ProductBet = b.BetValid                         //有效投注额  8-3
								if p.Rate*b.BetValid/100 > float64(r.DiscountUp) { //判断退水金额是否大于优惠上限
									v2.Money = float64(r.DiscountUp) //8-4
								} else {
									v2.Money = p.Rate * b.BetValid / 100 //8-4
								}
								fmt.Println("已加会员返水", v2.Money)
								v.Betall = v.Betall + b.BetValid         //4
								v.RebateWater = v.RebateWater + v2.Money //7
								v.AllMoney = v.SelfMoney + v.RebateWater //5

								data.Money = data.Money + v2.Money //累加返水总计  二
								break
							}

						}
					}
				} else { //如果没有会员信息 第一次添加 加入会员Map
					countSub := new(back.CountBetReportAccountSubtotalMap)
					countSub.Account = b.Account   //1
					countSub.MemberId = b.MemberId //2
					countSub.LevelId = b.LevelId   //3
					countSub.Betall = b.BetValid   //4
					//找打码量里与设定里对应的商品比例进行计算
					for _, p := range r.Params { //循环所有商品并赋值，存储
						countProduct := new(back.CountBetReportAccountProduct)
						countProduct.ProductId = p.ProductId     //商品id  8-1
						countProduct.ProductName = p.ProductName //商品名  8-2
						countProduct.Rate = p.Rate               //比例  8-5
						countProduct.ProductBet = 0.00           //有效投注额  8-3
						countProduct.Money = 0.00                //退水金额  8-4
						if p.ProductId == b.ProductId {          //退水额计算
							countProduct.ProductBet = b.BetValid //有效投注额  8-3
							//fmt.Println("有效投注额",b.BetValid)
							if p.Rate*b.BetValid/100 > float64(r.DiscountUp) { //判断退水金额是否大于优惠上限
								countProduct.Money = float64(r.DiscountUp) //8-4
							} else {
								countProduct.Money = p.Rate * b.BetValid / 100 //8-4
							}
						}
						countProductsMap[countProduct.ProductName] = countProduct //统计的商品集合map
						fmt.Println("统计的商品集合", countProduct)
						//fmt.Println("统计的商品集合map",countProductsMap[countProduct.ProductName])
						countSub.RebateWater = countProduct.Money //返水小计 7
						fmt.Println("返水额", countProduct.Money)
						fmt.Println("返水小计", countSub.RebateWater)
						data.Money = data.Money + countProduct.Money //累加返水总计  二
					}
					//会员自助返水小计
					selfWaterMoney, ok := am[countSub.Account]
					if !ok {
						countSub.SelfMoney = 0.00 //6
					} else {
						countSub.SelfMoney = selfWaterMoney //6
					}
					countSub.AllMoney = countSub.SelfMoney + countSub.RebateWater //返水总额小计 5
					countSub.Products = countProductsMap                          //8
					countSubMap[countSub.Account] = countSub                      //加入会员Map
					fmt.Println("会员Map", countSubMap[countSub.Account])
				}
				break //找到对应最大打码量的设定，跳出
			} else { //如果打码量小于所有设定打码量，则退水为0
				if i == len(rws)-1 { //没有找到优惠设定对应设定，则退水为0
					v, ok := countSubMap[b.Account] //判断是否添加此条打码的会员
					if ok {                         //如果已有会员信息 修改会员Map
						for _, p := range r.Params { //找打码量里与设定里对应的商品比例进行计算
							if p.ProductId == b.ProductId { //退水额计算
								v2, ok2 := v.Products[b.ProductName]
								if ok2 {
									v2.ProductBet = b.BetValid       //有效投注额  8-3
									v2.Money = 0.00                  //退水直接设置为零  8-4
									v.Betall = v.Betall + b.BetValid //4
									break
								}
							}
						}
					} else { //如果没有会员信息 第一次添加 加入会员Map
						countSub := new(back.CountBetReportAccountSubtotalMap)
						countSub.Account = b.Account   //1
						countSub.MemberId = b.MemberId //2
						countSub.LevelId = b.LevelId   //3
						countSub.Betall = b.BetValid   //4
						//找打码量里与设定里对应的商品比例进行计算
						for _, p := range r.Params { //循环所有商品并赋值，存储
							countProduct := new(back.CountBetReportAccountProduct)
							countProduct.ProductId = p.ProductId     //商品id  8-1
							countProduct.ProductName = p.ProductName //商品名  8-2
							countProduct.Rate = p.Rate               //比例  8-5
							countProduct.ProductBet = 0.00           //有效投注额  8-3
							countProduct.Money = 0.00                //退水金额  8-4

							if p.ProductId == b.ProductId { //退水额计算
								countProduct.ProductBet = b.BetValid //有效投注额  8-3
								countProduct.Money = 0.00            //退水直接设置为零  8-4
							}
							countProductsMap[countProduct.ProductName] = countProduct //统计的商品集合map
							countSub.RebateWater = countProduct.Money                 //返水小计 7
						}
						//会员自助返水小计
						selfWaterMoney, ok := am[countSub.Account]
						if !ok {
							countSub.SelfMoney = 0.00 //6
						} else {
							countSub.SelfMoney = selfWaterMoney //6
						}
						countSub.AllMoney = countSub.SelfMoney + countSub.RebateWater //返水总额小计 5
						countSub.Products = countProductsMap                          //8
						countSubMap[countSub.Account] = countSub                      //加入会员Map
					}
				}
			}
		}

	}
	data.PeopleNum = int64(len(countSubMap)) //累计返水总人数 一
	data.Subtotal = countSubMap              //四

	return
}
