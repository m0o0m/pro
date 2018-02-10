package input

type NoticeList struct {
	SiteId        string
	SiteIndexId   string
	NoticeCate    []int64 `query:"notice_cate"`    //公告分类
	NoticeContent string  `query:"notice_content"` //公告内容
	Page          int     `query:"page"`           //页码
	PageSize      int     `query:"pageSize"`       //单页条数
}

type Notice struct {
	SiteId      string
	SiteIndexId string
	Id          int64 `json:"id" query:"id" valid:"Required;Min(1)"` //公告id
	PageSize    int   `query:"pageSize"`                             //
	Page        int   `query:"page"`                                 //
}
type UpdateNotice struct {
	SiteId      string
	SiteIndexId string
	Id          int64  `json:"id"`      //公告id,如果为0,就是更新所有
	Title       string `json:"title"`   //公告标题
	Content     string `json:"content"` //公告内容
}

type DelNotice struct {
	SiteId      string
	SiteIndexId string
	Id          int64 `json:"id" valid:"Required;Min(1)"` //公告id
}

//站点公告列表
type SiteNoticeList struct {
	StartTime     string `query:"startTime" json:"startTime"`
	EndTime       string `query:"endTime" json:"endTime"`
	NoticeContent string `query:"noticeContent" json:"noticeContent"` //公告内容(模糊搜索)
}

//添加站点公告
type SiteNoticeAdd struct {
	NoticeCate    int    `json:"noticeCate" valid:"Range(1,7);ErrorCode(70025)"`         //公告分类
	NoticeAssign  string `json:"noticeAssign" valid:"RangeSize(1,200);ErrorCode(10209)"` //公告推送站点 1默认推送所有站点
	NoticeTitle   string `json:"noticeTitle" valid:"RangeSize(0,100);ErrorCode(100208)"` //标题
	NoticeContent string `json:"noticeContent"`                                          //内容
	NoticeState   int8   `json:"noticeState" valid:"Range(1,2);ErrorCode(30050)"`        //状态1开启  2关闭
}

//修改站点公告状态
type SiteNoticeState struct {
	Id     int64 `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`
	Status int8  `json:"status" valid:"Required;Range(1,2);ErrorCode(60117)"` //公告状态  1开启 2关闭
}

//修改某条站点公告信息
type SiteNoticeUpdate struct {
	Id            int64  `json:"id" valid:"Min(1);ErrorCode(10207)"`
	NoticeCate    int    `json:"noticeCate" valid:"Range(1,7);ErrorCode(70025)"`         //公告分类
	NoticeAssign  string `json:"noticeAssign" valid:"RangeSize(1,200);ErrorCode(10209)"` //公告推送站点 1默认推送所有站点
	NoticeTitle   string `json:"noticeTitle" valid:"RangeSize(0,100);ErrorCode(10208)"`  //标题
	NoticeContent string `json:"noticeContent"`                                          //内容
}

//批量删除公告信息
type SiteNoticeDel struct {
	Ids []int64 `json:"ids"`
}
