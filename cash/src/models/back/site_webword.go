package back

//查询首页文案返回信息(列表)
type IndexWord struct {
	Id          int64  `xorm:"id" json:"id"`                     //文案id
	SiteId      string `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` //站点前台id
	Title       string `xorm:"title" json:"title"`               //标题
	State       int8   `xorm:"state" json:"state"`               //状态 1-启用 2-关闭
	Sort        int64  `xorm:"sort" json:"sort"`                 //排序
	Itype       int64  `xorm:"itype" json:"iType"`               //类型代码
	TypeName    string `xorm:"type_name" json:"typeName"`        //类型名称
}

//查询文案内容(列表)
type IndexContentWord struct {
	Id          int64  `xorm:"id" json:"id"`                     //文案id
	SiteId      string `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` //站点前台id
	Title       string `xorm:"title" json:"title"`               //标题
	State       int8   `xorm:"state" json:"state"`               //状态 1-启用 2-关闭
	Sort        int64  `xorm:"sort" json:"sort"`                 //排序
	Itype       int64  `xorm:"itype" json:"iType"`               //类型代码
	TypeName    string `xorm:"type_name" json:"typeName"`        //类型名称
	Content     string `xorm:"content" json:"content"`           //内容
}

//查询优惠文案返回信息(列表)
type IndexActivityWord struct {
	Id          int64  `xorm:"id" json:"id"`                     //文案id
	SiteId      string `xorm:"site_id" json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id" json:"siteIndexId"` //站点前台id
	Title       string `xorm:"title" json:"title"`               //标题
	State       int8   `xorm:"state" json:"state"`               //状态 1-启用 2-关闭
	Sort        int64  `xorm:"sort" json:"sort"`                 //排序
	Img         string `xorm:"img" json:"img"`                   //标题图片路径
	TypeName    string `xorm:"type_name" json:"typeName"`        //类型名称
}

//查询首页文案返回信息(单条)
type IndexWordInfo struct {
	Id          int64  `xorm:"id" json:"id"`                       //文案id
	SiteId      string `xorm:"site_id" json:"site_id"`             //站点id
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` //站点前台id
	Title       string `xorm:"title" json:"title"`                 //标题
	Content     string `xorm:"content" json:"content"`             //内容
	//TypeName    string `xorm:"type_name" json:"type_name"`         //类型名称
}

//查询优惠文案返回信息(单条)
type IndexWordActivityInfo struct {
	Id      int64  `xorm:"id" json:"id"`           //文案id
	Title   string `xorm:"title" json:"title"`     //标题
	Content string `xorm:"content" json:"content"` //内容
}

//查询线路检测返回数据
type SiteDetect struct {
	Id          int64  `json:"id"`          //文案id
	SiteIndexId string `json:"siteIndexId"` //站点前台id
	Domain      string `json:"domain"`      //域名
	Content     string `json:"content"`     //域名
	Protocol    int8   `json:"protocol"`    //1 http: 2 https:
}
