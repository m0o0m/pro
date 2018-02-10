package schema

import "global"

//代理资料
type AgencyThirdInfo struct {
	AgencyId   int64  `xorm:"'agency_id' PK"` //代理id
	ChName     string `xorm:"ch_name"`        //中文名称  50
	UsName     string `xorm:"us_name"`        //英文名称  50
	Card       string `xorm:"card"`           //身份证号 固定20
	Phone      string `xorm:"phone"`          //手机   固定16
	QQ         string `xorm:"qq"`             //qq 固定12
	Email      string `xorm:"email"`          //邮箱  50
	ProvinceId int64  `xorm:"province_id"`    //省id
	CityId     int64  `xorm:"city_id"`        //市id
	AreaId     int64  `xorm:"area_id"`        //区id
	Remark     string `xorm:"remark"`         //备注  255
	SpreadId   string `xorm:"spread_id"`      //推广id
}

func (*AgencyThirdInfo) TableName() string {
	return global.TablePrefix + "agency_third_info"
}
