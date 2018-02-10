package input

//添加会员银行卡
type MemberBankAddIn struct {
	BankId      int64  `json:"bankId" valid:"Required;Min(1);ErrorCode(30059)"` //银行类型id
	MemberId    int64  //会员id
	CardNumber  string `json:"cardNumber" valid:"Required;MaxSize(19);ErrorCode(30060)"`  //银行卡号
	CardMan     string `json:"cardMan" valid:"Required;MaxSize(20);ErrorCode(30061)"`     //开户人
	CardAddress string `json:"cardAddress" valid:"Required;MaxSize(20);ErrorCode(30062)"` //开户行
}

//银行卡绑定、解绑
type MemberBankUnBindIn struct {
	Id int64 `json:"id" valid:"Min(1);ErrorCode(50013)"` //银行卡id
}

//银行卡删除
type MemberBankDeleteIn struct {
	Id       int64 `json:"id" valid:"Min(1);ErrorCode(50013)"` //银行卡id
	MemberId int64 //会员id
}

//会员银行卡详情
type MemberBankCardDetailsIn struct {
	Id       int64 `query:"id" valid:"Min(1);ErrorCode(50013)"` //银行卡id
	MemberId int64 //会员id
}
