package back

//站点维护信息
type SiteModule struct {
	Id          int64  `xorm:"'id' PK autoincr"`
	VType       string `xorm:"'v_type' PK notnull"`      //唯一性标示ag  og bbin  mg lebo ct具体商品维护  indexid前台维护  admin后台维护
	ProductName string `xorm:"product_name"`             //种类名称
	State       int64  `xorm:"'state' default(1)"`       //1启用2停用
	Content     string `xorm:"content"`                  //维护内容
	SiteIds     string `xorm:"'site_id_s' default('0')"` //站点0全部站点
	FType       int64  `xorm:"'f_type' PK notnull"`      //来源id 1全部 2pc 3wap 4app
}

//全站维护
type SiteMaintenance struct {
	Id          int64  `json:"id"`
	ProductName string `json:"productName"` //商品名
	SiteIdS     string `json:"siteIdS"`     //选择站点
	Content     string `json:"content"`     //内容
}

//全站维护
type SiteIdS struct {
	SiteIdS string `json:"siteIdS"` //站点
}

//站点是否选中
type SiteIsSelect struct {
	Id        string `json:"id"`         //主键id
	IndexId   string `json:"indexId"`    //前台id
	SiteName  string `json:"siteName"`   //站点名称
	IsDefault int8   ` json:"isDefault"` //是否默认站点
	IsSelect  int8   `json:"isSelect"`   //是否选中  0没有  1有
}
