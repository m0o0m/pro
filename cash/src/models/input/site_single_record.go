package input

//视讯额度掉单申请(添加)
type SiteSingleRecordAdd struct {
	SiteId      string  `json:"siteId"`                                     //站点ID
	SiteIndexId string  `json:"siteIndexId"`                                //多站点ID
	Username    string  `json:"username" valid:"Required;ErrorCode(30124)"` //会员账号
	AdminUser   string  `json:"adminUser"`                                  //提交人
	Money       float64 `json:"money" valid:"Required;ErrorCode(30150)"`    //交易额度
	Ctype       int64   `json:"ctype"`                                      //转出方
	Vtype       int64   `json:"vtype"`                                      //转入方
	DoTime      string  `json:"doTime" valid:"Required;ErrorCode(30163)"`   //掉单时间
	Remark      string  `json:"remark"`                                     //备注
}

//视讯额度掉单申请(审核)
type SiteSingleRecordEdit struct {
	SiteId         string `json:"site_id"`
	SiteIndexId    string `json:"site_index_id"` //多站点ID
	Id             int64  `json:"id" valid:"Required;ErrorCode(30041)"`
	UpdateUsername string `json:"update_username"`                          //操作人
	Type           int8   `json:"type" valid:"Range(1,3);ErrorCode(30162)"` //1表示掉单审核中，2表示审核通过，3无效申请
}

//视讯额度掉单申请(列表)
type SiteSingleRecordList struct {
	SiteId      string `query:"siteId"`
	SiteIndexId string `query:"siteIndexId"` //站点前台id
	Type        int8   `query:"type"`        //交易别 1表示掉单审核中，2表示审核通过，3无效申请
	Ctype       string `query:"ctype"`       //转出方
	Vtype       string `query:"vtype"`       //转入方
	Username    string `query:"username"`    //会员账号
	StartTime   string `query:"startTime"`   //开始时间
	EndTime     string `query:"endTime"`     //结束时间
}
