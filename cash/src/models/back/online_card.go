package back

//OnlineThird 站点管理-三方操作返回列表
type OnlineThird struct {
	Id        int    `xorm:"id" json:"id"`                  //id
	ThirdName string `xorm:"title" json:"thirdName"`        //三方名称
	Status    int8   `xorm:"pay_status" json:"status"`      //状态
	IsOpenIn  int8   `xorm:"deposit_state" json:"isOpenIn"` //是否开启入款
	IsOpenOut int8   `xorm:"out_state" json:"isOpenOut"`    //是否开启出款
	IsIpLimit int8   `xorm:"ip_state" json:"isIpLimit"`     //是否开启ip限制
	ModelName string `xorm:"pay_models" json:"modelName"`   //mod名称
	Code      string `xorm:"pay_code" json:"code"`          //code码
	BankUrl   string `xorm:"bank_url" json:"bankUrl"`       //网银支付网关
	PayUrl    string `xorm:"ali_pay_url" json:"aliPayUrl"`  //支付网关
}
