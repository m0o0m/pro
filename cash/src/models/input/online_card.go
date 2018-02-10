package input

//UpStatus 修改状态请求struct
type UpStatus struct {
	Id     int64 `json:"id" valid:"Min(1);ErrorCode(10236)"`         //三方id
	Status int8  `json:"status" valid:"Range(1,2);ErrorCode(10238)"` //状态
}

//删除
type DelThird struct {
	Id int64 `query:"id" json:"id" valid:"Min(1);ErrorCode(10236)"` //三方id
}

//修改三方网银信息
type UpdateThird struct {
	Id        int64  `json:"id" valid:"Min(1);ErrorCode(10236)"`                 //id
	ThirdName string `json:"thirdName" valid:"RangeSize(1,20);ErrorCode(10237)"` //三方名称
	Status    int8   `json:"status" valid:"Range(1,2);ErrorCode(10238)"`         //状态
	IsOpenIn  int8   `json:"isOpenIn" valid:"Range(1,2);ErrorCode(10239)"`       //是否开启入款
	IsOpenOut int8   `json:"isOpenOut" valid:"Range(1,2);ErrorCode(10240)"`      //是否开启出款
	IsIpLimit int8   `json:"isIpLimit" valid:"Range(1,2);ErrorCode(10241)"`      //是否开启ip限制
	//ModelName string `json:"modelName" valid:"RangeSize(1,20);ErrorCode(10242)"`  //mod名称
	Code    string `json:"code" valid:"RangeSize(1,255);ErrorCode(10243)"` //code码
	BankUrl string `json:"bankUrl" valid:"Url;ErrorCode(10244)"`           //网银支付网关
	PayUrl  string `json:"aliPayUrl" valid:"Url;ErrorCode(10245)"`         //支付网关
}

//添加网银
type AddThird struct {
	ThirdName string `json:"thirdName" valid:"RangeSize(1,20);ErrorCode(10237)"` //三方名称
	Status    int8   `json:"status" valid:"Range(1,2);ErrorCode(10238)"`         //状态
	IsOpenIn  int8   `json:"isOpenIn" valid:"Range(1,2);ErrorCode(10239)"`       //是否开启入款
	IsOpenOut int8   `json:"isOpenOut" valid:"Range(1,2);ErrorCode(10240)"`      //是否开启出款
	IsIpLimit int8   `json:"isIpLimit" valid:"Range(1,2);ErrorCode(10241)"`      //是否开启ip限制
	ModelName string `json:"modelName" valid:"RangeSize(1,20);ErrorCode(10242)"` //mod名称
	Code      string `json:"code" valid:"RangeSize(1,255);ErrorCode(10243)"`     //code码
	BankUrl   string `json:"bankUrl" valid:"Url;ErrorCode(10244)"`               //网银支付网关
	PayUrl    string `json:"aliPayUrl" valid:"Url;ErrorCode(10245)"`             //支付网关
}
