package input

type MessageList struct {
	SiteId      string
	SiteIndexId string
	Mtype       int `query:"mtype"`
}

type Message struct {
	SiteId      string
	SiteIndexId string
	Id          int64 `json:"id" query:"id" valid:"Min(1)"` //消息id
}

//获取个人信息列表
type WapMemberMessageList struct {
	SiteId      string //站点Id
	SiteIndexId string //站点前台Id
	MemberId    int64  //会员Id
	//Mtype       int    `query:"mtype" valid:"Range(3,7);ErrorCode(10210)"` //
}

//获取会员游戏公告列表
type WapMemberNoticeList struct {
	SiteId string //站点Id
	Mtype  int8   `query:"mtype" valid:"Range(3,7);ErrorCode(10210)"` //公告类型（ 3:体育公告 4:彩票公告 5视讯公告 6:扑鱼公告 7:电子公告)
}

//获取游戏公告详情

type WapMemberNoticeInfo struct {
	Id int64 `query:"id" json:"id" valid:"Min(1);ErrorCode(10215)"` //消息Id
}

//获取个人信息详情
type WapMemberMessageInfo struct {
	Id          int64  `query:"id" valid:"Min(1);ErrorCode(10215)"` //消息Id
	MemberId    int64  //会员Id
	SiteId      string //站点Id
	SiteIndexId string //站点前台Id
}

//删除个人消息
type WapMemberMessageDel struct {
	Id          int64  `json:"id" query:"id"` //消息Id
	MemberId    int64  //会员Id
	SiteId      string //站点Id
	SiteIndexId string //站点前台Id
}

//wap前台消息列表
type WapMemberMesList struct {
	MemberId int64 `query:"memberId"` //会员Id
}

//wap修改消息状态
type WapMesStatus struct {
	Id    int64 `json:"id"`    //消息id
	State int64 `json:"state"` //状态
}

//会员消息分页
type MesTerm struct {
	StartTime int64 `query:"start_time"` //
	EndTime   int64 `query:"end_time"`   //
	PageSize  int   `query:"pageSize"`   //
	Page      int   `query:"page"`       //
}
