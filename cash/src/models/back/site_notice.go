package back

//公告
type Notice struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`    //公告标题
	Type    int64  `json:"type"`     //广告类型：1中间，2左下，3右下
	AddTime int64  `json:"add_time"` //公告时间
	State   int64  `json:"state"`    //公告状态1-启用 2-关闭
}

//公告详情
type NoticeInfo struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`   //公告标题
	Type    int64  `json:"type"`    //广告类型：1中间，2左下，3右下
	Content string `json:"content"` //广告内容
}

//公告列表
type NoticeList struct {
	ID            int64  `json:"id" xorm:"'id' PK autoincr"`        //id
	NoticeTitle   string `json:"title" xorm:"notice_title"`         //公告标题
	NoticeDate    int64  `json:"creat_time" xorm:"notice_date"`     //公告时间
	NoticeState   int64  `json:"notice_state" xorm:"notice_state"`  //公告状态
	NoticeContent string `json:"content" xorm:"notice_content"`     //公告内容
	NoticeAssign  string `xorm:"notice_assign" json:"noticeAssign"` //公告分配站点
}

//站点公告列表
type SiteNoticeList struct {
	Id            int64  `xorm:"id" json:"id"`
	NoticeTitle   string `xorm:"notice_title" json:"noticeTitle" `    //公告标题
	NoticeContent string `xorm:"notice_content" json:"noticeContent"` //公告内容
	NoticeCate    int8   `xorm:"notice_cate" json:"noticeCate"`       //公告类型
	NoticeDate    int64  `xorm:"notice_date" json:"noticeDate" `      //公告日期
	NoticeAssign  string `xorm:"notice_assign" json:"noticeAssign"`   //公告分配站点
	NoticeState   int8   `xorm:"notice_state" json:"noticeState" `    //公告状态
}

//多条站点公告(普通公告)的内容
type SiteNoticeContent struct {
	NoticeContent string `xorm:"notice_content"` //多条站点公告内容
}

type MemberNoticeInfo struct {
	NoticeTitle   string `xorm:"notice_title" json:"noticeTitle" `    //公告标题
	NoticeContent string `xorm:"notice_content" json:"noticeContent"` //公告内容
	NoticeDate    string `xorm:"notice_date" json:"noticeDate" `      //公告时间
}

type NoticeData struct {
	NoticeContent string `xorm:"notice_content" json:"notice_content"` //公告内容
	NoticeDate    int64  `xorm:"notice_date" json:"notice_date" `      //公告时间
}

//wap公告类型
type NoticeTypes struct {
	NoticeCate int8                `xorm:"notice_cate" json:"notice_cate"` //公告类型
	NoticeType string              `json:"notice_type"`                    //公告类型（文字）
	NoticeList []WapSiteNoticeList `json:"list"`                           //同类型公告消息
}

//wap站点公告列表
type WapSiteNoticeList struct {
	Id            int64  `xorm:"id" json:"id"`
	NoticeTitle   string `xorm:"notice_title" json:"notice_title" `    //公告标题
	NoticeContent string `xorm:"notice_content" json:"notice_content"` //公告内容
	NoticeCate    int8   `xorm:"notice_cate" json:"notice_cate"`       //公告类型
	NoticeDate    int64  `xorm:"notice_date" json:"notice_date" `      //公告日期
	NoticeAssign  string `xorm:"notice_assign" json:"notice_assign"`   //公告分配站点
	NoticeState   int8   `xorm:"notice_state" json:"notice_state" `    //公告状态
	NoticeDateStr string `json:"notice_date_str"`                      //转换后的公告日期
}

//公告类型
type NoticeTypeList struct {
	NoticeCate int8   `xorm:"notice_cate" json:"notice_cate"` //公告类型
	NoticeType string `json:"notice_type"`                    //公告类型（文字）
}
