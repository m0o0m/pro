package function

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"global"
	"models/back"
	"models/schema"
)

type MemberRebateBean struct {
}

//查询返佣优惠列表
func (m *MemberRebateBean) GetRebateList(siteId, siteIndexId string, sTime, eTime int64) (
	[]back.MemberRebateCommission, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var rebateCommissions []back.MemberRebateCommission
	rebateCommissionSchema := new(schema.MemberRebateCommission)
	if sTime != 0 && eTime != 0 {
		sess.Table(rebateCommissionSchema.TableName()).
			Where("start_time >= ?", sTime). //都是用开始时间判断的
			Where("start_time <= ?", eTime)
	}
	if siteIndexId != "" {
		sess.Where("site_index_id = ?", siteIndexId)
	}
	if siteId != "" {
		sess.Where("site_id = ?", siteId)
	}
	err := sess.Table(rebateCommissionSchema.TableName()).Find(&rebateCommissions)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return rebateCommissions, err
	}
	return rebateCommissions, err
}

//查询所有交叉的项
func (m *MemberRebateBean) GetCross(siteId, siteIndexId string, sTime, eTime int64) (
	[]*schema.MemberRebateCommission, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var commissions []*schema.MemberRebateCommission
	rebateCommissionSchema := new(schema.MemberRebateCommission)
	sTimeStr := fmt.Sprintf("%d", sTime)
	eTimeStr := fmt.Sprintf("%d", eTime)
	//sql := "Select * form " + rebateCommissionSchema.TableName() +
	//	" WHERE (start_time >=" + sTimeStr + " and start_time <=" + eTimeStr + ")" +
	//	" OR (end_time>=" + sTimeStr + " and end_time <=" + eTimeStr + ")"
	sql := "Select * from " + rebateCommissionSchema.TableName() +
		" WHERE (start_time >=" + sTimeStr + " and start_time <=" + eTimeStr + ")" +
		" OR (end_time>=" + sTimeStr + " and end_time <=" + eTimeStr + ")"
	if siteId != "" && len(siteId) < 5 {
		sql += " and site_id = '" + siteId + "'"
	}
	if siteIndexId != "" && len(siteIndexId) < 5 {
		sql += " and site_index_id = '" + siteIndexId + "'"
	}
	err := sess.SQL(sql).Find(&commissions)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return commissions, err
	}
	return commissions, err
}

//存入优惠总计
func (m *MemberRebateBean) SaveRebateCommission(commission *back.MemberRebateCommission, sessArgs ...*xorm.Session) error {
	commissionSchema := new(schema.MemberRebateCommission)
	var sess *xorm.Session
	switch len(sessArgs) {
	case 1:
		sess = sessArgs[0]
	case 0:
		sess = global.GetXorm().NewSession()
	default:
		panic("<sessArgs> incorrect parameter passed")
	}
	updateNum, err := sess.Table(commissionSchema.TableName()).
		InsertOne(commission)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	if updateNum != 1 {
		global.GlobalLogger.Error("error:%s", err.Error())
		return errors.New("insert 0 now")
	}
	return err
}

//存入返佣详情
func (*MemberRebateBean) SaveRebateRecord(rebateRecords []*schema.MemberRebateRecord, sess *xorm.Session) error {
	//num, err := sess.Table(rebateRecords[0].TableName()).InsertMulti(&rebateRecords) //很遗憾,为了拿到返回的id值,只能放弃批量插入
	ins := make([]interface{}, len(rebateRecords))
	for k, _ := range ins {
		ins[k] = rebateRecords[k]
	}
	num, err := sess.Table(rebateRecords[0].TableName()).Insert(ins...)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	if num != int64(len(rebateRecords)) {
		return errors.New("insert err")
	}
	//js, _ := json.Marshal(rebateRecords)
	//fmt.Println("存完数据后id变了没有:", string(js))

	return err
}

//存入返佣详情对应商品信息
func (m *MemberRebateBean) SaveRebateRecordProduct(rebateRecordProducts []*schema.MemberRebateRecordProduct, sess *xorm.Session) (err error) {
	num, err := sess.Table(rebateRecordProducts[0].TableName()).InsertMulti(&rebateRecordProducts)
	if err != nil {
		return
	}
	if num != int64(len(rebateRecordProducts)) {
		return errors.New("insert err")
	}
	return
}

//wap 查看是否开启会员推广
func (*MemberRebateBean) WapMemberSpread(siteId, siteIndexId string) (bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mss := new(schema.MemberSpreadSet)
	has, err := sess.Where("site_id=?", siteId).Where("site_index_id=?", siteIndexId).
		Where("is_open=1").Get(mss)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return has, err
	}
	return has, err
}

//wap 推广返佣比例列表
func (*MemberRebateBean) WapMemberRebate(siteId, siteIndexId string) (memberRebateList []back.WapMemberRebateBack, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mrs := new(schema.MemberRebateSet)
	mrp := new(schema.MemberRebateProduct)
	product := new(schema.Product)
	sql := fmt.Sprintf("%s.id = %s.set_id", mrs.TableName(), mrp.TableName())
	sql2 := fmt.Sprintf("%s.product_id=%s.id", mrp.TableName(), product.TableName())
	sess.Select("valid_money,product_name,rate")
	memberRebate := make([]back.WapMemberRebate, 0)
	err = sess.Table(mrs.TableName()).Join("LEFT", mrp.TableName(), sql).Join("LEFT",
		product.TableName(), sql2).Where(mrs.TableName()+".site_id=?", siteId).Where(mrs.TableName()+
		".site_index_id=?", siteIndexId).Find(&memberRebate)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return memberRebateList, err
	}
	var memberRebates back.WapMemberRebateBack
	var productRebate back.ProductRate
	var validMoneys []int64
	for i := range memberRebate {
		validMoney := memberRebate[i].ValidMoney
		validMoneys = append(validMoneys, validMoney)

	}
	//给数组去重
	validMoneys = RemoveDuplicatesAndEmpty(validMoneys)
	for i := range validMoneys {
		memberRebates.ValidMoney = validMoneys[i]
		memberRebateList = append(memberRebateList, memberRebates)
	}
	for k := range memberRebate {
		for i := range memberRebateList {
			if memberRebate[k].ValidMoney == memberRebateList[i].ValidMoney {
				productRebate.ProductName = memberRebate[k].ProductName
				productRebate.Rate = memberRebate[k].Rate
				memberRebates.ValidMoney = memberRebate[i].ValidMoney
				memberRebateList[i].ProductRate = append(memberRebateList[i].ProductRate, productRebate)
			}
		}
	}
	return memberRebateList, err
}

//wap 获取会员推广设定系数
func (*MemberRebateBean) WapMemberRebateRanking(siteId, siteIndexId string) (float64, bool, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	mss := new(schema.MemberSpreadSet)
	rebateRanking := new(back.WapRebateRanking)
	has, err := sess.Table(mss.TableName()).Where("site_id=?", siteId).
		Where("site_index_id=?", siteIndexId).Get(rebateRanking)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return 0, has, err
	}
	rankingNum := rebateRanking.RankingNum
	return rankingNum, has, err
}
