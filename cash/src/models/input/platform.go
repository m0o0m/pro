package input

// 获取交易平台信息筛选条件
type PlatformList struct {
	Platform string `query:"platform"`
	Status   int8   `query:"status" valid:"Range(0,2);ErrorCode(50106)"` //状态
}

// 获取一条交易平台信息
type OnePlatformInfo struct {
	Id int64 `query:"id" valid:"Required;Min(0);ErrorCode(50013)"` //id
}

// 修改交易平台状态
type UpdatePlatformStatus struct {
	Id     int64 `json:"id" valid:"Required;Min(0);ErrorCode(30161)"`         //id
	Status int8  `json:"status" valid:"Required;Range(0,2);ErrorCode(50106)"` //状态
}

// 修改交易平台信息
type UpdatePlatform struct {
	Id       int64  `json:"id" valid:"Required;Min(0);ErrorCode(30161)"` //id
	Platform string `json:"platform" valid:"Required;ErrorCode(80001)"`
	Status   int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(50106)"` //状态
}

// 新增交易平台
type AddPlatform struct {
	Platform string `json:"platform" valid:"Required;ErrorCode(80001)"`
	Status   int8   `json:"status" valid:"Required;Range(0,2);ErrorCode(50106)"` //状态
}

// 删除交易平台
type DeletePlatform struct {
	Id int64 `json:"id" valid:"Required;Min(0);ErrorCode(50013)"` //id
}
