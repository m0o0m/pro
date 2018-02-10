package input

//添加红包设置
type RedBagData struct {
	SiteId      string  `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(10050)"`      //操作站点id
	SiteIndexId string  `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台id
	Account     string  `query:"account" valid:"MinSize(4);MaxSize(12);ErrorCode(30124)"`  //会员账号
	Type        int8    `query:"type" valid:"Range(1,2);ErrorCode(10141)"`                 //类型 1存款 2打码量
	Value       float64 `query:"value" valid:"Required;ErrorCode(10141)"`                  //存款或打码量的值
}
