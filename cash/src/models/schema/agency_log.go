package schema

import "global"

//后台人员操作日志记录
type AgencyLog struct {
	Id             int64  `xorm:"'id' PK autoincr"` //主键id
	SiteId         string `xorm:"site_id"`          //站点id
	SiteIndexId    string `xorm:"site_index_id"`    //站点前台id
	OperateAccount string `xorm:"operate_account"`  //操作者账号
	OperateTime    int64  `xorm:"operate_time"`     //操作时间
	OperateInfo    string `xorm:"operate_info"`     //操作详情
	OperateContent string `xorm:"operate_content"`  //操作内容
	OperatePath    string `xorm:"operate_path"`     //操作路径
	Ip             string `xorm:"ip"`               //操作者iP
	Type           int8   `xorm:"type"`             //操作类型1,增加2，删除，3，查看，4，修改
}

func (*AgencyLog) TableName() string {
	return global.TablePrefix + "agency_log"
}
