package back

//站点广告列表
type WebPopList struct {
	Id          int64  `json:"id"`          //广告id
	SiteId      string `json:"siteId"`      //站点id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
	Title       string `json:"title"`       //广告标题
	AddTime     int64  `json:"addTime"`     //添加时间
	State       int8   `json:"state"`       //状态 1启用  2关闭
	Type        int8   `json:"type"`        //广告类型：1中间，2左下，3右下
	BeforeUrl   string `json:"beforeUrl"`   //登录前广告链接
	AfterUrl    string `json:"afterUrl"`    //登录后连接
	IsLink      int64  `json:"isLink"`      //1 新开页面  2本页跳转
	Content     string `json:"content"`     //广告内容 备注
}
