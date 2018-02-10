package input

//子账号基本资料设定(添加)
type SubAccountAdd struct {
	SiteId          string //站点id
	SiteIndexId     string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`                          //站点前台id
	Account         string `json:"account" valid:"Required;MinSize(5);MaxSize(12);ErrorCode(30009)"`         //账号
	Password        string `json:"password" valid:"Required;MinSize(6);MaxSize(12);ErrorCode(30010)"`        //密码
	ConfirmPassword string `json:"confirmPassword" valid:"Required;ErrorCode(30011)"`                        //确认密码
	Username        string `json:"username" valid:"Required;ErrorCode(30014)"`                               //名称
	OperatePassword string `json:"operatePassword" valid:"Required;MinSize(6);MaxSize(12);ErrorCode(30015)"` //操作密码
	Status          int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(30050)"`                      //状态
	RoleId          int64  `json:"roleId"`                                                                   //角色id
	Level           int8   `json:"level"`                                                                    //层级
	ParentId        int64  `json:"parentId"`                                                                 //上级id
}

//子账号列表
type SubAccountList struct {
	SiteId      string //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	IsLogin     int8   `query:"isLogin" valid:"Range(0,2);ErrorCode(30050)"`     //在/离线状态
	Key         string `query:"key"`                                             //下拉框搜索条件
	Value       string `query:"value"`                                           //搜索值
	Status      int8   `query:"status"`                                          //在线状态（1：启用   2:禁用）
	RoleId      int64  `json:"roleId"`                                           //角色id
	Level       int8   `json:"level"`                                            //层级
	ParentId    int64  `json:"parentId"`                                         //上级id
}

//id
type SubAccountId struct {
	SiteId      string //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `json:"id" query:"id" form:"id" valid:"Required;Min(1);ErrorCode(30041)"`
}

//设置子账号权限
type PermissionEdit struct {
	SiteId       string  //站点id
	SiteIndexId  string  `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id           int64   `json:"id"  valid:"Required;ErrorCode(30041)"`           //子账号id
	PermissionId []int64 `json:"permissionId"`                                    //权限id
	ParentId     int64   `json:"parentId"`                                        //上级id
	ChildPower   string  `json:"childPower" valid:"MaxSize(50);ErrorCode(50160)"` //子帐号查看权限
	ChildSite    string  `json:"childSite" valid:"MaxSize(255);ErrorCode(50161)"` //子帐号操作站点
}

//子账号基本资料设定(修改)
type SubAccountEdit struct {
	SiteId          string //站点id
	SiteIndexId     string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"`      //站点前台id
	Id              int64  `json:"id"  valid:"Required;Min(1);ErrorCode(30041)"`         //账号id
	Password        string `json:"password" valid:"MaxSize(12);ErrorCode(30010)"`        //密码
	ConfirmPassword string `json:"confirmPassword" valid:"ErrorCode(30011)"`             //确认密码
	Username        string `json:"username" valid:"Required;ErrorCode(30014)"`           //名称
	OperatePassword string `json:"operatePassword" valid:"MaxSize(12);ErrorCode(30015)"` //操作密码
}

//子账号权限
type SubAccountPermission struct {
	SiteId      string `query:"siteId"`                                          //站点id
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(30041)"`
	ParentId    int64  `query:"parentId"`                                 //上级id
	Type        int8   `query:"type" valid:"Range(1,2);ErrorCode(50113)"` //1.代理权限  2.后台管理权限
}

//子账号口令验证信息
type SubAccountToken struct {
	SiteId  string //站点id
	Status  int    `json:"status" valid:"Range(1,2);ErrorCode(30050)"`  //是否启用（1：启用   2:禁用）
	PassKey string `json:"passKey" valid:"MinSize(8);ErrorCode(30009)"` //密钥
	//UpdateTime int64  `json:"update_time"`                                                       //更新时间
}

//服务器端生成密钥
type GenKey struct {
	SiteId      string
	SiteIndexId string
	Len         int `query:"len" valid:"Min(16);Max(16);ErrorCode(20024)"` //密钥长度,8的倍数
}
