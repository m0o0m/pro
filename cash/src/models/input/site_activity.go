package input

//获取会员活动列表（Wap)
type WapActivity struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(30016)"` //站点前台id
}

//获取单个会员活动详情
type WapActivityInfo struct {
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(10214)"`                       //主键Id
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`      //站点id
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(30016)"` //站点前台id

}
