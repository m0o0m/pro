package schema

import "global"

//附件表
type SiteThumb struct {
	Id          int64  `xorm:"'id' PK autoincr" json:"id"`         // 主键id
	SiteId      string `xorm:"site_id" json:"site_id"`             // 站点ID
	SiteIndexId string `xorm:"site_index_id" json:"site_index_id"` // 前台ID
	FilePath    string `xorm:"file_path" json:"file_path"`         // 文件路径
	FileName    string `xorm:"file_name" json:"file_name"`         // 文件名称
	FileType    string `xorm:"file_type" json:"file_type"`         // 文件类型
	FileMd5     string `xorm:"file_md5" json:"file_md5"`
	State       int8   `xorm:"state" json:"state"`             // '状态1-可用 2-禁用'
	AddTime     int64  `xorm:"add_time" json:"add_time"`       // 添加时间
	DeleteTime  int64  `xorm:"delete_time" json:"delete_time"` // 删除时间

}

func (*SiteThumb) TableName() string {
	return global.TablePrefix + "site_thumb"
}
