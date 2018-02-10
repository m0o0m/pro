package input

//添加套餐
type ComboAdd struct {
	ComboName string `json:"combo_name" valid:"Required;MinSize(1);MaxSize(20);ErrorCode(30114)"` //套餐名称
}

//套餐id
type ComboId struct {
	Id int64 `query:"id"  valid:"Min(1);Required;ErrorCode(30104)"` //套餐id
}

//修改套餐
type ComboEdit struct {
	Id        int64  `json:"id"  valid:"Required;ErrorCode(30104)"`                              //套餐id
	ComboName string `json:"comboName" valid:"Required;MinSize(1);MaxSize(20);ErrorCode(30114)"` //套餐名
}

//套餐列表筛选条件
type ComboList struct {
	Status    int8   `query:"status"`    //套餐状态
	ComboName string `query:"comboName"` //套餐名
}

//套餐id和状态
type ComboStatus struct {
	Id     int64 `json:"id" valid:"Required;ErrorCode(30104)"`       //套餐id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(30050)"` //状态
}
