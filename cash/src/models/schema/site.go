package schema

import (
	"global"
)

//Site 站点表
type Site struct {
	Id                string  `xorm:"'id' PK"`               //主键id
	IndexId           string  `xorm:"index_id"`              //前台id
	SiteName          string  `xorm:"site_name"`             //站点名称
	AgencyId          int64   `xorm:"agency_id"`             //所属开户人id
	ComboId           int64   `xorm:"combo_id"`              //套餐id
	Status            int8    `xorm:"status"`                //状态(1.正常2.关闭)
	CreateTime        int64   `xorm:"'create_time' created"` //创建时间
	DeleteTime        int64   `xorm:"delete_time"`           //删除时间
	DomainUp          int     `xorm:"domain_up"`             //域名上限
	UpCose            float64 `xorm:"up_cose"`               //超过上线收费金额
	IsDefault         int8    `xorm:"is_default"`            //是否默认站点
	SelfHelpSwitch    int8    `xorm:"self_help_switch"`      //自助优惠开关(1.开启2.关闭)
	H5StateSwitch     int8    `xorm:"h5_state_switch"`       //h5动画状态开关(1.开启2.关闭)
	PopoverBgColor    string  `xorm:"popover_bg_color"`      //站点弹窗广告背景颜色
	PopoverTitleColor string  `xorm:"popover_title_color"`   //站点弹窗广告标题颜色
	PopoverBarColor   string  `xorm:"popover_bar_color"`     //站点弹窗广告标题栏颜色
	OnlineTime        int64   `xorm:"online_time"`           //站点上线时间
	IsDownApp         int8    `xorm:"is_down_app"`           //是否可以下载app
	//Theme             string  `xorm:"theme" `                //主题
}

func (*Site) TableName() string {
	return global.TablePrefix + "site"
}
