package schema

import "global"

type MemberMessage struct {
	ID          int64  `xorm:"id"`
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //前台id
	Title       string `xorm:"title"`         //标题
	MemberId    int64  `xorm:"member_id"`     //用户id
	Content     string `xorm:"content"`       //内容
	CreateTime  int64  `xorm:"create_time"`   //创建时间
	Mtype       int64  `xorm:"mtype"`         //消息类型
	State       int8   `xorm:"state"`         //1未读,2,已读
	DeleteTime  int64  `xorm:"delete_time"`   //删除时间
}

func (*MemberMessage) TableName() string {
	return global.TablePrefix + "member_message"
}
