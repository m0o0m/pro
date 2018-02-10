package back

//会员消息详情
type MemberMessage struct {
	ID          int64  `xorm:"id" json:"id"`
	SiteId      string `xorm:"site_id" json:"site_id"`             //站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //前台id
	Title       string `xorm:"title" json:"title"`                 //标题
	MemberId    int64  `xorm:"member_id" json:"member_id"`         //用户id
	Content     string `xorm:"content" json:"content"`             //内容
	CreateTime  int64  `xorm:"create_time" json:"create_time"`     //创建时间
	Mtype       int64  `xorm:"mtype" json:"mtype"`                 //消息类型
	State       int8   `xorm:"state" json:"state"`                 //1未读,2,已读
	DeleteTime  int64  `xorm:"delete_time" json:"delete_time"`     //删除时间
}

//会员消息列表
type MemberMessageList struct {
	ID         int64  `json:"id" xorm:"'id' PK autoincr"`     //id
	Title      string `json:"title" xorm:"title"`             //消息标题
	CreateTime int64  `json:"create_time" xorm:"create_time"` //创建时间
	State      int8   `json:"state" xorm:"state"`             //消息状态
}

//会员站点公告列表
type MemberNoticeList struct {
	Id           int64  `xorm:"id" json:"id"`                        //主键Id
	NoticeTitle  string `xorm:"notice_title" json:"noticeTitle"`     //公告标题
	NoticeCotent string `xorm:"notice_content" json:"noticeContent"` //公告内容
	NoticeDate   string `xorm:"notice_date" json:"noticeDate"`       //公告时间
	NoticeState  int8   `xorm:"notice_state" json:"noticeState"`     //公告状态
}

//获取个人信息列表
type WapMemberMessageList struct {
	ID         int64  `xorm:"id" json:"id"`
	Title      string `xorm:"title" json:"title"`            //标题
	Content    string `xorm:"content" json:"content"`        //内容
	CreateTime int64  `xorm:"create_time" json:"createTime"` //创建时间
	State      int8   `xorm:"state" json:"state"`            //1未读,2,已读
}

//获取个人消息
type WapMemberMessageInfo struct {
	Title      string `xorm:"title" json:"title"`            //标题
	Content    string `xorm:"content" json:"content"`        //内容
	CreateTime int64  `xorm:"create_time" json:"createTime"` //创建时间
}
