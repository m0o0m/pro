package schema

import "global"

//弹窗广告
type SiteAdv struct {
	Id          int64  `xorm:"id"`            //广告id
	SiteId      string `xorm:"site_id"`       //站点id
	SiteIndexId string `xorm:"site_index_id"` //站点前台id
	Title       string `xorm:"title"`         //广告标题
	Content     string `xorm:"content"`       //广告内容
	Type        int8   `xorm:"type"`          //广告类型：1中间，2左下，3右下
	AddTime     int64  `xorm:"add_time"`      //添加时间
	State       int8   `xorm:"state"`         //状态 1启用  2关闭
	BeforeUrl   string `xorm:"before_url"`    //登录前广告链接
	AfterUrl    string `xorm:"after_url"`     //登陆后链接
	IsLink      int8   `xorm:"is_link"`       //1新开页面  2本页跳转
	SiteText    string `xorm:"site_text"`     //当site_id和site_index_id为空时，就是所有站显示，本字段记录不显示的站点
	DeleteTime  int64  `xorm:"delete_time"`   //删除时间   0未删除
	Sort        int    `xorm:"sort"`          //排序
	Remark      string `xorm:"remark"`        //广告备注
	//StartTime   int64  `xorm:"startTime"`     //广告开始时间
	//EndTime     int64  `xorm:"endTime"`       //广告结束时间
}

func (*SiteAdv) TableName() string {
	return global.TablePrefix + "site_adv"
}
