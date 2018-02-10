package schema

import "global"

//站点域名配置表
type SiteDomain struct {
	Id          int64  `xorm:"'id' PK autoincr"`                 //主键id
	SiteId      string `xorm:"'site_id' notnull"`                //站点id
	SiteIndexId string `xorm:"'site_index_id' notnull"`          //站点前台id
	Domain      string `xorm:"'domain' notnull"`                 //pc域名
	IsDefault   int64  `xorm:"'is_default' notnull default(0)"`  //是否主域名
	CreateTime  int64  `xorm:"'create_time' notnull default(0)"` //创建时间
	DeleteTime  int64  `xorm:"'delete_time' notnull default(0)"` //软删除时间(为0表示未删除)
	IsUsed      int64  `xorm:"'is_used' notnull"`                //是否已经使用
	Type        int    `xorm:"'type' notnull  default(1)"`       //1前台域名，2后台域名，3代理域名
}

func (*SiteDomain) TableName() string {
	return global.TablePrefix + "site_domain"
}

//type SiteDomain struct {
//	Id              int64  `xorm:"'id' PK autoincr"` //主键id
//	SiteId          string `xorm:"site_id"`          //站点id  固定4位
//	SiteIndexId     string `xorm:"site_index_id"`    //站点前台id  固定4位
//	SslCsr          string `xorm:"ssl_csr"`          //csr文件
//	SslKey          string `xorm:"ssl_key"`          //key文件
//	Domain        string `xorm:"pc_domain"`        //pc域名  50
//	Domain       string `xorm:"wap_domain"`       //wap域名  50
//	IsDefault       int    `xorm:"is_default"`       //手机域名和pc域名是否主域名
//	IsUsed          int    `xorm:"is_used"`          //是否已经使用
//	BackstageDomain string `xorm:"backstage_domain"` //后台域名
//	//FileName        map[string]string `xorm:"file_name"`             //文件名["ssl_key":"","ssl_csr":""]
//	CreateTime int64 `xorm:"'create_time' created"` //添加时间
//	DeleteTime int64 `xorm:"delete_time"`           //删除时间
//}
