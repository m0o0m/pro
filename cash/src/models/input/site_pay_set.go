package input

//站点支付设定
type GetPaySetInfoById struct {
	Id int64 `query:"id" json:"id"` //id 通过站点id+站点前台id请求查询到的id
}
