package back

//获取维护列表
type MainList struct {
	Id          int64  `json:"id"`           //
	VType       string `json:"v_type"`       //唯一性标示
	ProductName string `json:"product_name"` //种类名称
	State       int64  `json:"state"`        //暂时不用 状态
	Content     string `json:"content"`      //维护内容
	Sids        string `json:"sids"`         //站点 0全部站点
	Rid         int64  `json:"rid"`          //来源id
}

//获取站点列表
type DomainList struct {
	Id          int64        `json:"'id' PK autoincr"` //主键id
	SiteId      string       `json:"site_id"`          //站点id  固定4位
	SiteIndexId string       `json:"site_index_id"`    //站点前台id  固定4位
	List        []DomainList `json:"list"`             //站点下列表
	IsPmOne     int64        `json:"is_pm_one"`        //客户后台代理后台维护状态
	IsWapOne    int64        `json:"is_wap_one"`       //wap维护状态
	IsPcOne     int64        `json:"is_pc_one"`        //pc维护状态
	IsFcOne     int64        `json:"is_fc_one"`        //彩票维护状态
	IsSpOne     int64        `json:"is_sp_one"`        //体育维护状态
	IsVdOne     int64        `json:"is_vd_one"`        //视讯维护状态
	IsDzOne     int64        `json:"is_dz_one"`        //电子维护状态
}

//获取后台域名列表
type DomainInfoList struct {
	Id          int64  `json:"'id' PK autoincr"` //主键id
	SiteId      string `json:"site_id"`          //站点id  固定4位
	SiteIndexId string `json:"site_index_id"`    //站点前台id  固定4位
	QQ          string `json:"qq"`               //客服qq
	Wechat      string `json:"wechat"`           //客服微信
	Phone       string `json:"phone"`            //客服电话
	Email       string `json:"email"`            //客服邮箱
}

//查询条件
type ConditionList struct {
	SiteId      string `json:"site_id"`       //
	SiteIndexId string `json:"site_index_id"` //
	DeleteTime  int64  `json:"delete_time"`   //
}

//存入redis中的维护数据
type SaveMainTain struct {
	SiteId      string     `json:"site_id"`       //
	SiteIndexId string     `json:"site_index_id"` //
	Content     string     `json:"content"`       //维护内容
	Module      []InfoList `json:"module"`        //具体维护游戏类别
	QQ          string     `json:"qq"`            //客服qq
	Wechat      string     `json:"wechat"`        //客服微信
	Phone       string     `json:"phone"`         //客服电话
	Email       string     `json:"email"`         //客服邮箱
}

//获取电子 视讯下级菜单
type InfoList struct {
	ProductName string `json:"product_name"` //商品名
	VType       string `json:"v_type"`       //游戏类型
	IsState     int64  `json:"is_state"`     //维护状态 0 未维护 1 已维护
}
