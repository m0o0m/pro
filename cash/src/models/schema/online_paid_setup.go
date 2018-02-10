package schema

import (
	"global"
)

//线上支付设定
type OnlinePaidSetup struct {
	Id                int     `xorm:"id" json:"id"`                                 //主键Id
	MerchatId         string  `xorm:"merchat_id" json:"merchat_id"`                 //商户号
	PaidPlatform      int     `xorm:"paid_platform" json:"paid_platform"`           //支付平台
	PaidType          int     `xorm:"paid_type" json:"paid_type"`                   //支付类型
	PrivateKey        string  `xorm:"private_key" json:"private_key"`               //密匙
	PublicKey         string  `xorm:"public_key" json:"public_key"`                 //公匙
	BackAddress       string  `xorm:"back_address" json:"back_address"`             //返回地址
	Status            int     `xorm:"status" json:"status"`                         //状态(1.启用2.停用)
	PaidCode          string  `xorm:"paid_code" json:"paid_code"`                   //支付code
	PaidDomain        string  `xorm:"paid_domain" json:"paid_domain"`               //支付域名
	FitforLevel       string  `xorm:"fitfor_level" json:"fitfor_level"`             //适用层级
	IsJumpApp         int     `xorm:"is_jump_app"`                                  //是否跳转app(0.不跳转1.跳转)
	SiteId            string  `xorm:"site_id" json:"site_id"`                       //站点id
	SiteIndexId       string  `xorm:"site_index_id" json:"site_index_id"`           //前台Id
	PaidLimit         float64 `xorm:"paid_limit" json:"paid_limit"`                 //当日支付限额
	SuitableEquipment int     `xorm:"suitable_equipment" json:"suitable_equipment"` //适用设备(1.pc2.wap3.都可以)
	Remark            string  `xorm:"remark" json:"remark"`                         //备注
	MerUrl            string  `xorm:"mer_url" json:"mer_url"`                       //自填写支付网关
	CreateTime        int64   `xorm:"create_time" json:"create_time"`               //创建时间
	DelTime           int64   `xorm:"delete_time" json:"del_time"`                  //删除时间
	Sort              int64   `xorm:"sort" json:"sort"`                             //排序，数字越大，排越前面
}

func (*OnlinePaidSetup) TableName() string {
	return global.TablePrefix + "online_paid_setup"
}
