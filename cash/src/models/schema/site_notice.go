package schema

import "global"

type SiteNotice struct {
	ID             int64  `xorm:"'id' PK autoincr"`
	NoticeTitle    string `xorm:"notice_title"`        //公告标题
	NoticeContent  string `xorm:"notice_content"`      //公告内容
	NoticeCate     int    `xorm:"notice_cate"`         //公告分类
	NoticeDate     int64  `xorm:"notice_date created"` //公告时间
	NoticeState    int8   `xorm:"notice_state"`        //公告状态
	NoticePosition int8   `xorm:"notice_position"`     //公告弹窗出现的位置(1.中间弹窗)
	NoticeAssign   string `xorm:"notice_assign"`       //公告推送站点(1.全部2.里面写的哪几个站点就推送到哪几个站点)
	DeleteTime     int64  `xorm:"delete_time"`         //公告删除标志(0:表示未删除,其余表示删除的时间)
}

func (*SiteNotice) TableName() string {
	return global.TablePrefix + "site_notice"
}
