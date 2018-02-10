package schema

import "global"

//多站点文案类型
type IwordCate struct {
	Id   int64  `xorm:"'id' PK autoincr"` //类型id
	Name string `xorm:"name"`             //文案类型名称
}

func (*IwordCate) TableName() string {
	return global.TablePrefix + "site_iword_cate"
}
