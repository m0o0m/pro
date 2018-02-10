//优惠查询
package input

//站点优惠列表
type DiscountSearchList struct {
	SiteId      string `query:"site_id" valid:"MaxSize(4);ErrorCode(50058)"`       //站点id
	SiteIndexId string `query:"site_index_id" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//优惠总计
type DiscountAllList struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(50058)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	Year        int    `query:"year" valid:"Min(4);ErrorCode(90601)"`                     //统计年份
	Month       int    `query:"month" valid:"Min(1);ErrorCode(90602)"`                    //统计月份
}

//获取优惠明细
type DiscountInfo struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(50058)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `query:"id"`                                              //数据id
}
