package back

import "html/template"

//站点公告弹窗列表
type WebAdvList struct {
	Id          int64  `xorm:"id"  json:"id"`                     //广告id
	SiteId      string `xorm:"site_id"  json:"siteId"`            //站点id
	SiteIndexId string `xorm:"site_index_id"  json:"siteIndexId"` //站点前台id
	Title       string `xorm:"title"  json:"title"`               //广告标题
	Content     string `xorm:"content"  json:"content"`           //广告内容
	Type        int8   `xorm:"type"  json:"type"`                 //广告类型：1中间，2左下，3右下
	AddTime     int64  `xorm:"add_time"  json:"addTime"`          //添加时间
	State       int8   `xorm:"state"  json:"state"`               //状态 1启用  2关闭
	BeforeUrl   string `xorm:"before_url"  json:"beforeUrl" `     //登录前广告链接
	AfterUrl    string `xorm:"after_url"  json:"afterUrl" `       //登陆后链接
	IsLink      int8   `xorm:"is_link"  json:"isLink" `           //1新开页面  2本页跳转
	SiteText    string `xorm:"site_text"  json:"siteText" `       //当site_id和site_index_id为空时，就是所有站显示，本字段记录不显示的站点
	DeleteTime  int64  `xorm:"delete_time"  json:"deleteTime"`    //删除时间   0未删除
}

//站点公告弹窗列表
type WebPopAdv struct {
	Id      int64         `xorm:"id"  json:"id"`           //广告id
	Title   string        `xorm:"title"  json:"title"`     //广告标题
	Content template.HTML `xorm:"content"  json:"content"` //广告内容
	State   int8          `xorm:"state"  json:"state"`     //状态 1启用  2关闭
}

type WebAdvColor struct {
	PopoverBgColor    string `xorm:"popover_bg_color" json:"popover_bg_color"`       //站点弹窗广告背景颜色
	PopoverTitleColor string `xorm:"popover_title_color" json:"popover_title_color"` //站点弹窗广告标题颜色
	PopoverBarColor   string `xorm:"popover_bar_color" json:"popover_bar_color"`     //站点弹窗广告标题栏颜色
	AdWay             int8   `xorm:"ad_way" json:"ad_way"`                           //方式（1，方框，2镂空）
}

//公告弹窗配置详情
type WebAdvConfigDetail struct {
	SiteId            string `xorm:"id" json:"siteId"`
	SiteIndexId       string `xorm:"index_id"  json:"siteIndexId"`
	PopoverBgColor    string `xorm:"popover_bg_color" json:"popoverBgColor"`       //站点弹窗广告背景颜色
	PopoverTitleColor string `xorm:"popover_title_color" json:"popoverTitleColor"` //站点弹窗广告标题颜色
	PopoverBarColor   string `xorm:"popover_bar_color" json:"popoverBarColor"`     //站点弹窗广告标题栏颜色
}

//总后台广告列表
type AdminAdvert struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`     //广告名称（标题）
	Sort      int    `json:"sort"`      //排序
	Content   string `json:"content"`   //广告预览（内容）
	BeforeUrl string `json:"beforeUrl"` //登录前广告链接
	AfterUrl  string `json:"afterUrl"`  //登录后广告链接
	State     int8   `json:"state"`     //广告状态（1开启  2关闭）
	Remark    string `json:"remark"`    //广告备注
	AddTime   int64  `json:"addTime"`   //添加时间
}

//总后台广告详情
type AdminAdvertInfo struct {
	Id        int64  `json:"id"`
	Sort      int    `json:"sort"`       //排序
	Title     string `json:"title"`      //广告名称
	BeforeUrl string `json:"before_url"` //登录前广告链接
	AfterUrl  string `json:"after_url"`  //登录后广告链接
	State     int8   `json:"state"`      //广告状态（1开启  2关闭）
	Content   string `json:"content"`    //广告内容（图片+文字）
	Remark    string `json:"remark"`     //广告备注
	SiteText  string `json:"site_text"`  //剔除站点
	IsLink    int8   `json:"isLink"`     //1新开页面  2本页跳转
}
