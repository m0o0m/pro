package function

//会员优惠返水记录
type RetreatWaterRecordBean struct {
}

/*
SELECT
  sales_member_retreat_water_record.id,
  sales_member_retreat_water_record.site_id,
  sales_member.third_agency_id,
  member_id,
  sales_member_retreat_water_record.account,
  count(*),
  sum(rebate_water),
  from_unixtime(start_time, '%Y %D %M %h:%i:%s %x')
FROM sales_member_retreat_water_record
JOIN sales_member ON member_id = sales_member.id
GROUP BY member_id;
*/
//优惠统计,查询
// deprecated 应该直接查询discount_report表,方法在discount_count文件中
//func (m *RetreatWaterRecordBean) List(retreatWaterCount *input.DiscountCountList, times *global.Times, params *global.ListParams) (retreatWaterRecordTotal back.DiscountCountTotal, err error) {
//	sess := global.GetXorm().NewSession()
//	defer sess.Close()
//	retreatWaterRecordSchema := new(schema.MemberRetreatWaterRecord)
//	t2 := new(schema.Member).TableName()
//	sess.Table(retreatWaterRecordSchema.TableName()).Alias("t1")
//	if retreatWaterCount.SiteId != "" {
//		sess.Where("t1.site_id = ?", retreatWaterCount.SiteId)
//	}
//	sess.Join("INNER", t2, "t1.member_id = "+t2+".id")
//	if retreatWaterCount.AgencyAccount != "" {
//		//如果代理账号不为空,就得联代理表
//		t3 := new(schema.Agency).TableName()
//		sess.Join("INNER", t3, t2+".third_agency_id = "+t3+".id")
//		sess.Where(t3+".account = ?", retreatWaterCount.AgencyAccount)
//	}
//	times.Make("t1.start_time", sess) //
//	conds := sess.Conds()
//	params.Make(sess) //分页
//	sess.Select("t1.id,t1.site_id," + t2 + ".third_agency_id as agency_id,t1.member_id,t1.account,COUNT(*) as num,SUM(t1.rebate_water) as money,t1.start_time")
//	var retreatWaterRecordList []back.DiscountCountList
//	err = sess.GroupBy("t1.member_id").Find(&retreatWaterRecordList)
//	if err != nil {
//		return
//	}
//	b, err := sess.Table(retreatWaterRecordSchema.TableName()).
//		Alias("t1").
//		Where(conds).
//		Select("COUNT(*) as num,SUM(rebate_water) as money").
//		Get(&retreatWaterRecordTotal)
//	if err != nil {
//		return
//	}
//	if !b {
//		err = errors.New("find 0 row")
//		return
//	}
//	retreatWaterRecordTotal.Content = &retreatWaterRecordList
//	return
//}
