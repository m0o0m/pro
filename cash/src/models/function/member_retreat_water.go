package function

import (
	"fmt"
	"global"
	"models/back"
	"models/input"
	"models/schema"
)

type MemberRetreatWaterBean struct {
}

//优惠查询列表
func (*MemberRetreatWaterBean) SearchMemberRetreatWater(this *input.ListRetreatWater, listparam *global.ListParams) (data []back.ListRetreatWater, count int64, err error) {
	water := new(schema.MemberRetreatWater)
	sess := global.GetXorm().Table(water.TableName())
	defer sess.Close()
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("start_time>=?", this.StartTime)
	sess.Where("end_time<=?", this.EndTime)
	conds := sess.Conds()
	//获得分页记录
	listparam.Make(sess)
	err = sess.Table(water.TableName()).
		Select("id,admin_user,start_time,end_time,create_time,event,no_people_num,people_num,money,bet").Find(&data)
	if err != nil {
		return
	}
	count, err = sess.Table(water.TableName()).Where(conds).Count()
	return
}

//优惠查询详情
func (*MemberRetreatWaterBean) DetailMemberRetreatWater(this *input.DetailRetreatWater) ([]back.RetreatWaterRecord, int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	var data []back.RetreatWaterRecord
	water := new(schema.MemberRetreatWaterRecord)
	if this.SiteId != "" {
		sess.Where("site_id=?", this.SiteId)
	}
	if this.SiteIndexId != "" {
		sess.Where("site_index_id=?", this.SiteIndexId)
	}
	sess.Where("sales_member_retreat_water_record.periods_id=?", this.Id)
	conds := sess.Conds()
	waterPro := new(schema.MemberRetreatWaterRecordProduct)
	product := new(schema.Product)
	sql2 := fmt.Sprintf("%s.id = %s.record_id", water.TableName(), waterPro.TableName())
	sql3 := fmt.Sprintf("%s.id = %s.product_id", product.TableName(), waterPro.TableName())
	data1 := make([]back.RetreatWaterRecordAndProduct, 0)
	err := sess.Table(water.TableName()).
		Select("sales_member_retreat_water_record.id,sales_member_retreat_water_record.account,member_id,level_id,betall,all_money,self_money,rebate_water,sales_member_retreat_water_record.status,sales_member_retreat_water_record.create_time,product_id,product_name,product_bet,rate,money").
		Join("LEFT", waterPro.TableName(), sql2).
		Join("LEFT", product.TableName(), sql3).OrderBy("member_id").
		Find(&data1)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, 0, err
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
		} else {
			//id相同时，只组装商品列表
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
	count, err := sess.Table(water.TableName()).Where(conds).Count()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return data, count, err
	}
	return data, count, err
}

//优惠查询冲销
func (*MemberRetreatWaterBean) EditMemberRetreatWater(this *input.EditRetreatWater) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	water := new(schema.MemberRetreatWaterRecord)
	water.Status = 2
	count, err := sess.Table(water.TableName()).In("id", this.Id).
		Cols("status").Update(water)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return count, err
	}
	return count, err
}
