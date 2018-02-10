package input

//H5动画设置查询
type SiteH5Set struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	Status      int8   `query:"status" valid:"Range(0,2);ErrorCode(30050)"`      //状态 1正常2禁用
}

//H5动画设置修改
type PutSiteH5Set struct {
	Site   []site `json:"site" valid:"Required;MaxSize(4);ErrorCode(60105)"` //站点id和站点前台id集合
	Status int8   `json:"status" valid:"Required;ErrorCode(30050)"`          //状态 1开启  2关闭 上传的多少就将数据库改成多少
}
type site struct {
	SiteId      string `json:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `json:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//公告弹窗列表
type NoticePopupList struct {
	SiteId      string `query:"siteId" valid:"MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `query:"siteIndexId" valid:"MaxSize(4);ErrorCode(10050)"` //站点前台id
	PageSize    int    `query:"pageSize"`                                        //每页显示
	Page        int    `query:"page"`                                            //页码
}

//查看广告弹窗配置
type GetNoticePopupSet struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//修改广告弹窗配置
type PutNoticePopupSet struct {
	SiteId            string `json:"siteId" valid:"Required;ErrorCode(60105)"`
	SiteIndexId       string `json:"siteIndexId" valid:"Required;ErrorCode(10050)"`       //站点前台id
	PopoverBgColor    string `json:"popoverBgColor" valid:"Required;ErrorCode(30189)"`    //站点弹窗广告背景颜色
	PopoverTitleColor string `json:"popoverTitleColor" valid:"Required;ErrorCode(30190)"` //站点弹窗广告标题颜色
	PopoverBarColor   string `json:"popoverBarColor" valid:"Required;ErrorCode(30191)"`   //站点弹窗广告标题栏颜色
}

//添加广告
type NoticeAdd struct {
	SiteId      string `json:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `json:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
	Title       string `json:"title" valid:"Required;ErrorCode(30193)"`                  //标题
	Type        int8   `json:"type" valid:"Range(1,3);ErrorCode(30194)"`                 //1中间，2左下，3右下
	Content     string `json:"content" valid:"Required;ErrorCode(30195)"`                //内容
}

//公告弹框设置列表
type GetNotice struct {
	SiteId      string `query:"siteId" valid:"Required;MaxSize(4);ErrorCode(60105)"`
	SiteIndexId string `query:"siteIndexId" valid:"Required;MaxSize(4);ErrorCode(10050)"` //站点前台id
}

//公告弹框设置详情
type NoticeInfo struct {
	Id int64 `query:"id" valid:"Required;ErrorCode(30041)"` //id
}

//编辑公告弹框设置
type NoticeEdit struct {
	Id      int64  `json:"id" valid:"Required;ErrorCode(30041)"`
	Title   string `json:"title" valid:"Required;ErrorCode(30193)"`   //标题
	Type    int8   `json:"type" valid:"Range(1,3);ErrorCode(30194)"`  //1中间，2左下，3右下
	Content string `json:"content" valid:"Required;ErrorCode(30195)"` //内容
}

//删除公告弹框设置
type NoticePopupDel struct {
	Id int64 `json:"id" valid:"Required;ErrorCode(30041)"`
}
