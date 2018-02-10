package input

//视讯ip白名单添加
type GameWhiteAdd struct {
	Ip      string `json:"ip" valid:"Required;MaxSize(20);ErrorCode(50129)"` //ip
	Remarks string `json:"remarks" valid:"MaxSize(255);ErrorCode(20019)"`    //备注
}

//视讯ip白名单修改
type GameWhiteEdit struct {
	Id      int    `json:"id" valid:"Required;Min(1);ErrorCode(30041)"`      //序号
	Ip      string `json:"ip" valid:"Required;MaxSize(20);ErrorCode(50129)"` //ip
	Remarks string `json:"remarks" valid:"MaxSize(255);ErrorCode(20019)"`    //备注
}

//视讯ip白名单删除
type GameWhiteDel struct {
	Id int `json:"id" valid:"Required;Min(1);ErrorCode(30041)"` //序号
}

//视讯ip白名单列表筛选
type GameWhiteList struct {
	Ip string `query:"ip" valid:"MaxSize(20);ErrorCode(50129)"` //ip
}
