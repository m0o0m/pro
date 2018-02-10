package function

import (
	"errors"
	"global"
	"models/input"
	"models/schema"

	"framework/google_auth"
	"models/back"
)

type AgencySignBean struct{}

//查询登录账号和密码是否正确
func (*AgencySignBean) Login(l *input.AgencySignLogin, siteInfo schema.SiteDomain) (ok bool, err error, agency *schema.Agency) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency = new(schema.Agency)
	agency.Account = l.Account
	agency.SiteId = siteInfo.SiteId
	if siteInfo.Type == 2 { //后台域名只允许开户人和开户人子账号登陆
		agency.RoleId = 1
	}
	if siteInfo.Type == 3 { //代理域名 允许股东总代，代理登陆 以及代理子账号登陆
		sess.In("role_id", []int8{2, 3, 4})
		agency.SiteIndexId = siteInfo.SiteIndexId
	}
	ok, err = sess.Get(agency)
	if err != nil {
		global.GlobalLogger.Error("Login error:%s", err.Error())
		return
	}
	return
}

//登录失败更新字段
func (*AgencySignBean) LoginErrUpdate(a *schema.Agency) (row int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	row, err = sess.ID(a.Id).Cols("login_err_count").Update(a)
	return
}

//登录成功更新字段
func (*AgencySignBean) LoginUpdate(a *schema.Agency) (row int64, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	row, err = sess.ID(a.Id).Cols("login_key", "last_login_ip", "last_login_time", "login_count", "login_ip", "login_time", "is_login").Update(a)
	return
}

//退出
func (*AgencySignBean) Logout(id int64) error {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.IsLogin = 2
	agency.LoginKey = ""
	_, err := sess.ID(id).Cols("login_key", "is_login").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error%s", err.Error())
		return err
	}
	return err
}

//验证密码
func (*AgencySignBean) ValidPassword(old string, id int64) (ok bool, err error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.Id = id
	ok, err = sess.ID(agency.Id).Get(agency)
	if old != agency.Password {
		ok = false
	}
	return
}

//更新密码
func (*AgencySignBean) UpdatePassword(id int64, pwd string) (int64, error) {
	sess := global.GetXorm().NewSession()
	defer sess.Close()
	agency := new(schema.Agency)
	agency.Id = id
	agency.Password = pwd
	row, err := sess.ID(agency.Id).Cols("password").Update(agency)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return row, err
	}
	return row, err
}

//口令验证
func (*AgencySignBean) VerifyCode(secret, code string) bool {
	// 1:30sec
	ga := googleAuth.NewGAuth()
	ret, err := ga.VerifyCode(secret, code, 1)
	if err != nil {
		return false
	}
	return ret
}

//密钥生成
func (*AgencySignBean) CreateSecret(lens ...int) (key *back.Key, err error) {
	ga := googleAuth.NewGAuth()
	secret, err := ga.CreateSecret(lens...)
	if err != nil {
		return nil, errors.New("gen key err")
	}
	return &back.Key{Key: secret}, nil
}
