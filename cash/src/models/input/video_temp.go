package input

type VideoAdd struct {
	Name   string `json:"name" valid:"Required;MaxSize(50);ErrorCode(90003)"` //模板名称
	Pid    int8   `json:"pid" valid:"Min(0);ErrorCode(90004)"`                //父级id
	Style  int8   `json:"style" valid:"Min(1);ErrorCode(90005)"`              //样式序号
	Status int8   `json:"status" valid:"Range(1,2);ErrorCode(90006)"`         //是否启用
}
type VideoUpdate struct {
	Id     int64  `json:"id" valid:"Min(1);ErrorCode(90007)"`                 //模板编号
	Name   string `json:"name" valid:"Required;MaxSize(50);ErrorCode(90003)"` //模板名称
	Pid    int8   `json:"pid" valid:"Min(0);ErrorCode(90004)"`                //父级id
	Style  int8   `json:"style" valid:"Min(1);ErrorCode(90005)"`              //样式序号
	Status int8   `json:"status" valid:"Range(1,2);ErrorCode(90006)"`         //是否启用
}
