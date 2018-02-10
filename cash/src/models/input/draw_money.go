package input

//wap 取款
type WapDrawMoney struct {
	Money        float64 `json:"money" valid:"MinFloat64(1.00);MaxFloat64(123456.00);ErrorCode(30230)"` //取款金额
	DrawPassword string  `json:"drawPassword" valid:"Required;ErrorCode(30230)"`                        //取款密码
	BankId       int64   `json:"bankId" valid:"Required;ErrorCode(30231)"`                              //出款银行id
}

// 取款
type DrawMoney struct {
	Money        string `json:"money"`        //取款金额
	DrawPassword string `json:"drawPassword"` //取款密码
	BankId       string `json:"bankId" `      //出款银行id
}
