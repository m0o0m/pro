package schema

import "global"

type MemberAuditLog struct {
	Id          int64  `json:"id"`
	SiteId      string `json:"site_id"`               //站点
	SiteIndexId string `json:"site_index_id"`         //前台
	MemberId    int64  `json:"member_id"`             //会员id
	Account     string `json:"account"`               //会员账号
	UpdateDate  int64  `json:"'update_date' created"` //更新时间
	Content     string `json:"content"`               //日志详细内容
	Type        int64  `json:"type"`                  //类型 1表示清除
}

func (*MemberAuditLog) TableName() string {
	return global.TablePrefix + "member_audit_log"
}
