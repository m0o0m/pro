package input

type Search struct {
	SiteId      string //用户站点ID
	SiteIndexId string
	Account     string `query:"account" valid:"MaxSize(10);ErrorCode(50010)"` //用户查询帐号
	Level       int8   `query:"level" valid:"Min(1);ErrorCode(50000)"`        //用户查询等级
	Isvague     int8   `query:"isvague" valid:"Range(0,2);ErrorCode(50001)"`  //是否模糊查询
}
