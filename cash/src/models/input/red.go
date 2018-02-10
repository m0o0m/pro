package input

//抢红包
type GetSnatch struct {
	SetId int64 `query:"rid" valid:"Required;ErrorCode(71021)"`
}
