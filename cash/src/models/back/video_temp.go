package back

type VideoStyle struct {
	Id     int64  `xorm:"id" json:"id"`         //视讯模板id
	Name   string `xorm:"name" json:"name"`     //类型名字
	Pid    int8   `xorm:"pid" json:"pid"`       //父级
	Aid    int8   `xorm:"aid" json:"aid"`       //中间关联字段
	Style  int8   `xorm:"style" json:"style"`   //视讯样式
	Remark string `xorm:"remark" json:"remark"` //备注预留字段
	Status int8   `xorm:"status" json:"status"` //视讯模板开关
	//List   []VideoStyle `json:"list"`                 //子级菜单
}
