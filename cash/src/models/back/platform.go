package back

//交易平台列表
type Platform struct {
	Id       int64  `json:"id"`
	Status   int8   `json:"status"`   //状态  1：启用  2：禁用
	Platform string `json:"platform"` //角色名称
}
