package back

//后台人员操作日志
type AdminLog struct {
	Id             int64  `json:"id" xorm:"'id' PK autoincr"`            //主键id
	OperateAccount string `json:"operateAccount" xorm:"operate_account"` //操作者账号
	OperateTime    int64  `json:"operateTime" xorm:"operate_time"`       //操作时间
	OperateInfo    string `json:"operateInfo" xorm:"operate_info"`       //操作详情
	OperateContent string `json:"operateContent" xorm:"operate_content"` //请求内容体
	OperatePath    string `json:"operatePath" xorm:"operate_path"`       //操作路径
	Ip             string `json:"ip" xorm:"ip"`                          //操作者iP
	Type           int8   `json:"type" xorm:"type"`                      //操作类型
}
