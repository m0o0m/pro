package input

//开户人列表
type AccountHolderList struct {
	Status  int8   `query:"status" valid:"Range(0,2);ErrorCode(50085)"`   //状态
	IsLogin int8   `query:"isLogin" valid:"Range(0,2);ErrorCode(50083)"`  //在线状态
	Account string `query:"account" valid:"MaxSize(30);ErrorCode(50010)"` //搜索帐号
	PaiXu   string `query:"paiXu"`
	ShunXu  bool   `query:"shunXu"`
	SiteId  string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`
}

//增加开户人
type AddAccountHolder struct {
	Account         string  `json:"account" valid:"RangeSize(6,12);ErrorCode(50059)"`         //帐号
	Password        string  `json:"password" valid:"RangeSize(6,12);ErrorCode(50060)"`        //密码
	RePassword      string  `json:"rePassword" valid:"RangeSize(6,12);ErrorCode(50061)"`      //重复密码
	Username        string  `json:"username" valid:"RangeSize(2,30);ErrorCode(50063)"`        //开户人名称
	OperatePassword string  `json:"operatePassword" valid:"RangeSize(6,12);ErrorCode(50062)"` //操作密码
	ManageDomain    string  `json:"manageDomain" valid:"Required;Url;ErrorCode(50194)"`       //客户后台域名
	Status          int8    `json:"status" valid:"Range(1,2);ErrorCode(50064)"`               //开户人状态
	Remark          string  `json:"remark" valid:"Nullable(0,120);ErrorCode(50196)"`          //备注
	SiteName        string  `json:"siteName" valid:"RangeSize(2,25);ErrorCode(50065)"`        //站点名称
	Site            string  `json:"site" valid:"RangeSize(1,4);ErrorCode(50058)"`             //站点Id
	SiteIndex       string  `json:"siteIndex" valid:"RangeSize(1,4);ErrorCode(10050)"`        //站点前台ID
	AgencyDomain    string  `json:"agencyDomain" valid:"Required;Url;ErrorCode(50195)"`       //代理后台域名
	DomainUp        int     `json:"domainUp" valid:"Required;Min(1);ErrorCode(50067)"`        //前台域名上限
	UpCose          float64 `json:"upCose" valid:"MinFloat64(1.00);ErrorCode(50068)"`         //超过上线收费金额
	ComboId         int64   `json:"comboId" valid:"Required;Min(1);ErrorCode(50066)"`         //套餐id
}

//帐号  id(get)
type AccountNameId struct {
	Account string `query:"account" valid:"MaxSize(30);ErrorCode(50059)"` //帐号
	Id      int64  `query:"id" valid:"ErrorCode(30041)"`                  //id
	Status  int8   `query:"status" valid:"ErrorCode(50085)"`              //状态
}

//帐号  id(post)
type HolderNameId struct {
	Id     int64 `json:"id" valid:"Required;Min(1);ErrorCode(30041)"` //id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"`  //状态                                     //状态
}

//更新开户人信息
type UpdataAccountHolder struct {
	Id              int64  `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`             //id
	Account         string `json:"account" valid:"Required;MaxSize(30);ErrorCode(50059)"`   //帐号
	Password        string `json:"password" valid:"Nullable(6,12);ErrorCode(50060)"`        //密码
	RePassword      string `json:"rePassword" valid:"ErrorCode(50061)"`                     //重复密码
	OperatePassword string `json:"operatePassword" valid:"Nullable(6,12);ErrorCode(50062)"` //操作密码
	Username        string `json:"username" valid:"Required;MaxSize(50);ErrorCode(50063)"`  //开户人名称
	Status          int8   `json:"status" valid:"Required;Range(1,2);ErrorCode(30050)"`     //状态
}

//查询站点会员注册优惠设定
type SiteMemberReg struct { //站点前台ID
	Site      string `query:"site" valid:"Required;MaxSize(4);ErrorCode(50058)"`       //站点Id
	SiteIndex string `query:"site_index" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台ID
}

//修改站点会员注册优惠设定
type UpdataSiteMemberReg struct {
	Site      string  `json:"site" valid:"Required;MaxSize(4);ErrorCode(50058)"`       //站点Id
	SiteIndex string  `json:"site_index" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台ID
	Offer     float64 `json:"offer" valid:"Required;ErrorCode(50072)"`                 //加入会员赠送优惠金额
	AddMosaic int64   `json:"add_mosaic" valid:"Required;ErrorCode(50073)"`            //优惠打码倍数
	IsIp      int8    `json:"is_ip" valid:"Required;ErrorCode(50074)"`
	IsClear   int8    `json:"is_clear" valid:"Required;ErrorCode(50078)"` //是否限制IP 1:是2:否
}

//删除开户人
type DelAccountHolder struct {
	Id int64 `json:"id" valid:"Required;Min(1);ErrorCode(30041)"` //id
}

//开户人详情
type AccountHolderInfoIn struct {
	Id int64 `query:"id" valid:"Required;Min(1);ErrorCode(30041)"` //id
}

//增加管理员
type AddAccount struct {
	Site      string `json:"site" valid:"Required;MaxSize(4);ErrorCode(50058)"`      //站点Id
	SiteIndex string `json:"siteIndex" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台ID
	SiteName  string `json:"siteName" valid:"Required;MaxSize(50);ErrorCode(50065)"` //开户人后台域名
	Account   string `json:"account" valid:"Required;MaxSize(30);ErrorCode(50059)"`  //帐号
	Username  string `json:"username" valid:"Required;MaxSize(50);ErrorCode(50063)"` //开户人名称
	Remark    string `json:"remark"`                                                 //备注
	Password  string `json:"password"`                                               //密码
}
