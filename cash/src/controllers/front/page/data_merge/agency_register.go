package data_merge

import (
	"fmt"
	"models/back"
)

type AgencyRegister struct {
	HeaderFooter
	Key
	RegisterSet *back.SiteAgencyRegisterSet //
	AgencyForm  []back.AgencyForm           //
	RegStatus   int8                        // 1.启用 2.不启用
}

//得到代理注册页面数据
func (m *AgencyRegister) GetData(siteId string, siteIndexId string) (interface{}, error) {
	err := m.initHeaderFooter(siteId, siteIndexId)
	if err != nil {
		return m, err
	}
	data, has, err := siteAgencyRegSet.SiteIdExist(siteId, siteIndexId)
	if !has {
		m.RegStatus = 2
	} else {
		m.RegStatus = data.RegisterProxy
	}
	m.RegisterSet = data
	var Agencyform []back.AgencyForm

	if data.RegisterProxy == 1 {
		var formdata back.AgencyForm
		formdata.Name = "user_name"
		formdata.Notice = "请输入真实姓名"
		formdata.Title = "真实姓名"
		Agencyform = append(Agencyform, formdata)
	}
	if data.ChineseNickname == 1 {
		var formdata back.AgencyForm
		formdata.Name = "chinese_nick_name"
		formdata.Notice = "请输入中文昵稱"
		formdata.Title = "中文昵称"
		Agencyform = append(Agencyform, formdata)
	}
	if data.EnglishNickname == 1 {
		var formdata back.AgencyForm
		formdata.Name = "english_nick_name"
		formdata.Notice = "请输入英文昵稱"
		formdata.Title = "英文昵称"
		Agencyform = append(Agencyform, formdata)
	}
	if data.NeedCard == 1 {
		var formdata back.AgencyForm
		formdata.Name = "card"
		formdata.Notice = "请输入身份号"
		formdata.Title = "身份证"
		Agencyform = append(Agencyform, formdata)
	}
	if data.NeedEmail == 1 {
		var formdata back.AgencyForm
		formdata.Name = "email"
		formdata.Notice = "请输入邮箱"
		formdata.Title = "邮箱"
		Agencyform = append(Agencyform, formdata)
	}
	if data.NeedPhone == 1 {
		var formdata back.AgencyForm
		formdata.Name = "mobile"
		formdata.Notice = "请输入手机号码"
		formdata.Title = "手机号码"
		Agencyform = append(Agencyform, formdata)
	}
	if data.NeedQq == 1 {
		var formdata back.AgencyForm
		formdata.Name = "qq"
		formdata.Notice = "请输入qq"
		formdata.Title = "qq"
		Agencyform = append(Agencyform, formdata)
	}
	if data.PromoteWebsite == 1 {
		var formdata back.AgencyForm
		formdata.Name = "promote_website"
		formdata.Notice = "请输入推广网址"
		formdata.Title = "推广网址"
		Agencyform = append(Agencyform, formdata)
	}
	if data.OtherMethod == 1 {
		var formdata back.AgencyForm
		formdata.Name = "other_method"
		formdata.Notice = "请输入其他方法"
		formdata.Title = "其他方法"
		Agencyform = append(Agencyform, formdata)
	}

	//fmt.Printf("%+v\n",AgencyForm)
	m.AgencyForm = Agencyform
	fmt.Printf("%+v\n", m)
	return m, err
}

//得到页面
func (m *AgencyRegister) GetPage() []string {
	return []string{ZHUCE_DAILI, HEADER, HEADER_, FOOTER, FOOTER_, POP_ADV}
}

//得到<视讯电子彩票体育的页面>
func (m *AgencyRegister) GetSubPage() map[string]string {
	return nil
}
