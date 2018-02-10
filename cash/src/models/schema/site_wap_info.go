package schema

import "global"

//平台账号
type SiteWapInfo struct {
	Id            int64  `xorm:"'id' PK autoincr"` //主键id
	SiteId        string `xorm:"site_id"`          //站点id
	SiteIndexId   string `xorm:"site_index_id"`    //站点前台id
	AppUrl        string `xorm:"app_url"`          //app下载地址
	WapColor      string `xorm:"wap_color"`        //wap头部颜色
	WapBottom     string `xorm:"wap_bottom"`       //wap端底部文案
	AutoLinkName  string `xorm:"auto_link_name"`   //自定义链接名称
	AutoLinkUrl   string `xorm:"auto_link_url"`    //自定义链接url
	WapQuick      int8   `xorm:"wap_quick"`        //1默认不开启，2为开启
	IsDownload    int8   `xorm:"is_download"`      //是否允许下载
	WebsiteAppUrl string `xorm:"website_app_url"`  //官网app下载地址
}

func (*SiteWapInfo) TableName() string {
	return global.TablePrefix + "site_wap_info"
}
