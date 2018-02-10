package input

//站点公告弹窗管理 修改
type AdvUpdate struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90226);"`                      //id
	SiteId      string `json:"siteId" `                                                  //站点id
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Title       string `json:"title" valid:"MaxSize(200);ErrorCode(10050)"`              //广告标题
	Content     string `json:"content"`                                                  //广告内容
	Type        int8   `json:"type" valid:"Min(1);Max(3);ErrorCode(90227)"`              //广告类型：1中间，2左下，3右下
	State       int8   `json:"state" valid:"Min(1);ErrorCode(90228)"`                    //状态 1启用  2关闭
}

//站点公告列表
type AdvList struct {
	SiteId      string `query:"siteId"`                                                    //站点id
	SiteIndexId string `query:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Type        int64  `query:"type" valid:"Range(0,3);ErrorCode(90227);"`                 //广告类型：1中间，2左下，3右下
}

//广告列表
type AdvListBySite struct {
	SiteId      string `query:"siteId"`                                           //站点id
	SiteIndexId string `query:"siteIndexId"  valid:"MaxSize(4);ErrorCode(10050)"` //前台站点id
	Title       string `query:"title" valid:"MaxSize(20);ErrorCode(30193);"`      //广告标题
}

//广告详情
type AdvListBySiteDetail struct {
	SiteId      string `query:"siteId"`                                           //站点id
	SiteIndexId string `query:"siteIndexId"  valid:"MaxSize(4);ErrorCode(10050)"` //前台站点id
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(30041);"`              //广告id
}

//站点公告详情
type AdvListDetail struct {
	Id          int64  `query:"Id" valid:"Min(1);ErrorCode(30041)"`
	SiteId      string `query:"siteId"`      //站点id
	SiteIndexId string `query:"siteIndexId"` //前台站点id
}

//站点公告新增
type AdvAdd struct {
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	Title       string `json:"title" valid:"MaxSize(200);ErrorCode(10050)"`               //广告标题
	Content     string `json:"content"`                                                   //广告内容
	Type        int8   `json:"type" valid:"Min(1);Max(3);ErrorCode(90227)"`               //广告类型：1中间，2左下，3右下
	AddTime     int64  `json:"addTime" `                                                  //添加时间
	State       int8   `json:"state" valid:"Min(1);ErrorCode(90228)"`                     //状态 1启用  2关闭
}

//弹窗广告删除
type UpdateDeleteTime struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode();"`                            //弹窗广告id
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	DeleteTime  int64  `json:"deleteTime"`                                                //删除时间
}

//修改站点公告弹窗配置
type UpdateConfig struct {
	SiteId      string `json:"siteId" xorm:"id" valid:"MaxSize(4);ErrorCode(60105)"`             //站点id
	SiteIndexId string `json:"siteIndexId" xorm:"index_id"  valid:"MaxSize(4);ErrorCode(10050)"` //前台站点id
	BgColor     string `json:"popoverBgColor" xorm:"popover_bg_color"`                           //站点弹窗广告背景颜色
	TitleColor  string `json:"popoverTitle_color" xorm:"popover_title_color"`                    //站点弹窗广告标题颜色
	BarColor    string `json:"popover_bar_color" xorm:"popover_bar_color"`                       //站点弹窗广告标题栏颜色
}

//修改站点公告弹窗配置
type UpdateConfigDetail struct {
	SiteId      string `json:"siteId"`                                          //站点id
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //前台站点id
}

//广告修改
type PopUpdate struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90226);"`                       //弹窗广告id
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	BeforeUrl   string `json:"beforeUrl" valid:"MaxSize(500);ErrorCode(90229)"`           //登录前广告链接
	AfterUrl    string `json:"afterUrl" valid:"MaxSize(500);ErrorCode(90230)"`            //登录后连接
	State       int8   `json:"state" valid:"Min(1);ErrorCode(90228)"`                     //状态 1启用  2关闭
	Type        int8   `query:"type" valid:"Min(1);Max(3);ErrorCode(90227);"`             //广告类型：1中间，2左下，3右下
	Title       string `json:"title" valid:"MaxSize(200);ErrorCode(10050)"`               //广告标题
	Content     string `json:"content"`                                                   //广告内容
	IsLink      int8   `json:"isLink" valid:"Min(1);ErrorCode(90231)"`                    //1 新开页面  2本页跳转
}

//站点广告新增
type PopAdd struct {
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	BeforeUrl   string `json:"beforeUrl" valid:"MaxSize(500);ErrorCode(90229)"`           //登录前广告链接
	AfterUrl    string `json:"afterUrl" valid:"MaxSize(500);ErrorCode(90230)"`            //登录后连接
	State       int8   `json:"state" valid:"Min(1);ErrorCode(90228)"`                     //状态 1启用  2关闭
	//Type        int8   `query:"type"`               //广告类型：1中间，2左下，3右下
	Title   string `json:"title" valid:"MaxSize(200);ErrorCode(10050)"` //广告标题
	Content string `json:"content"`                                     //广告内容
	//IsLink      int8   `json:"isLink" valid:"Min(1);ErrorCode(90231)"`                     //1 新开页面  2本页跳转
	AddTime int64 `json:"addTime"` //
}
type UpdatePopStatus struct {
	Id          int64  `json:"id" valid:"Min(1);ErrorCode(90226);"`                       //广告id
	SiteId      string `json:"siteId"`                                                    //站点id
	SiteIndexId string `json:"siteIndexId"  valid:"Required;MaxSize(4);ErrorCode(10050)"` //前台站点id
	State       int8   `json:"state" valid:"Range(1,2);ErrorCode(30050)"`                 //状态                                                      //状态
}

//总后台广告列表
type AdminAdvertList struct {
	Title string `query:"title"`                                     //广告标题（名称）
	State int8   `query:"state" valid:"Range(0,2);ErrorCode(30050)"` //1-启用 2-关闭
}

//总后台添加广告
type AdminAdvertAdd struct {
	Sort      int    `json:"sort"`                                       //排序
	Title     string `json:"title"`                                      //广告名称
	BeforeUrl string `json:"beforeUrl"`                                  //登录前广告链接
	AfterUrl  string `json:"afterUrl"`                                   //登录后广告链接
	State     int8   `json:"state" valid:"Range(1,2);ErrorCode(30050)"`  //广告状态（1开启  2关闭）
	Content   string `json:"content"`                                    //广告内容
	Remark    string `json:"remark"`                                     //广告备注
	SiteText  string `json:"siteText"`                                   //剔除站点
	IsLink    int8   `json:"isLink" valid:"Range(1,2);ErrorCode(30259)"` //1新开页面  2本页跳转
}

//总后台修改广告
type AdminAdvertPut struct {
	Id        int64  `json:"id"  valid:"Required;ErrorCode(50013)"`
	Sort      int    `json:"sort"`                                       //排序
	Title     string `json:"title"`                                      //广告名称
	BeforeUrl string `json:"before_url"`                                 //登录前广告链接
	AfterUrl  string `json:"after_url"`                                  //登录后广告链接
	State     int8   `json:"state" valid:"Range(1,2);ErrorCode(30050)"`  //广告状态（1开启  2关闭）
	Content   string `json:"content"`                                    //广告内容（图片+文字）
	Remark    string `json:"remark"`                                     //广告备注
	SiteText  string `json:"site_text"`                                  //剔除站点
	IsLink    int8   `json:"isLink" valid:"Range(1,2);ErrorCode(30259)"` //1新开页面  2本页跳转
}

//总后台修改状态
type AdminAdvertState struct {
	Id    int64 `json:"id" valid:"Required;ErrorCode(50013)"`
	State int8  `json:"state" valid:"Range(1,2);ErrorCode(30050)"` //状态
}

//总后台修改排序
type AdminAdvertSort struct {
	Id   int64 `json:"id" valid:"Required;ErrorCode(50013)"`
	Sort int   `json:"sort" ` //排序
}

//总后台详情
type AdminAdvertInfo struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(50013)"`
}
