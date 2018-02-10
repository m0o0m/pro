package thirdParty

import (
	"config"
	"encoding/json"
	"errors"
	"fmt"
	"framework/napping"
	"global"
	"models/back"
	"models/input"
	"strings"
)

var VideoCfg *config.Video

type VideoGameBean struct {
	config.Video
}

func NewThirdParty() *VideoGameBean {
	vd := new(VideoGameBean)
	vd.ApiUrl = VideoCfg.ApiUrl
	vd.Md5Key = VideoCfg.Md5Key
	vd.DesKey = VideoCfg.DesKey
	return vd
}

var (
	USER_AGENT_PRE  = "Game_Go_"
	USER_HAS_PREFIX = "Lottery_Api_Go_"
)

//type UserData struct {
//	SiteId   string
//	UserName string
//	Platform string //游戏平台
//	AgentId  int64
//	IndexId  string
//	ShId     int64  //股东id
//	UaId     int64  //总代id
//	Media    string //设备类型  wap pc app
//	GameID   string //子游戏id
//	IP       string //ip
//	Lang     string //语言类型
//	Cur      string //货币类型
//	Limit    string //限额
//	Domain   string //登陆的域名
//	IsSw     bool   //是否是试玩
//}

type TransferData struct {
	//UserData
	input.VideoUserData
	TransferType string  //IN OUT
	Credit       float64 //额度不能为负数
	TradeNo      string  //订单号
}

type PasswordData struct {
	SiteId   string
	UserName string
	Password string
	Platform string //游戏平台
}

func (m *VideoGameBean) ForwardGameTest(ld input.VideoUserData) (string, error) {
	//可以试玩的游戏
	return "", nil
}

//第三方游戏处理
func (m *VideoGameBean) ForwardGame(ld input.VideoUserData) (string, error) {
	params := make([]string, 15)
	params[0] = "siteid=" + ld.SiteId
	params[1] = "username=" + ld.UserName
	params[2] = fmt.Sprintf("agent_id=%d", ld.AgentId)
	params[3] = "index_id=" + ld.IndexId
	params[4] = fmt.Sprintf("sh_id=%d", ld.ShId)
	params[5] = fmt.Sprintf("ua_id=%d", ld.UaId)
	params[6] = "cur=" + ld.Cur
	params[7] = "limit=" + ld.Limit
	params[8] = "lang=" + ld.Lang
	params[9] = "gtype=" + ld.Platform
	params[10] = "ip=" + ld.IP
	params[11] = "media=" + ld.Media
	params[12] = "domain=" + ld.Domain
	params[13] = "sw=0"
	params[14] = "subtype=" + ld.GameID
	qParams, err := global.DesEncrypt([]byte(strings.Join(params, "/\\\\/")), []byte(m.DesKey))
	if err != nil {
		global.GlobalLogger.Error("ForwardGame DesEncrypt error %v", err)
		return "", err
	}
	key := global.Md5(string(qParams) + m.Md5Key)

	p := napping.Params{"params": string(qParams), "key": key}
	s := napping.Session{Timeout: 40, Datatype: "xml"}
	s.SetHeader("User-Agent", USER_AGENT_PRE+ld.SiteId)

	resp, err := s.Get(m.ApiUrl+"/CreateMemberAndForwardGame", &p, nil, nil)
	if err != nil {
		body := ""
		if resp != nil {
			body = resp.RawText()
		}
		global.GlobalLogger.Error("ForwardGame  %s error %v", body, err)
		return "", err
	}
	if resp != nil {
		var gameResult back.GameResult
		err = json.Unmarshal([]byte(resp.RawText()), &gameResult)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			return "", err
		}
		if gameResult.Data.Code > 0 {
			return "", errors.New(resp.RawText())
		}
		return gameResult.Data.LoginUrl, nil
	}
	return "", errors.New("unknow")
}

//第三方游戏处理 ,retryCounts - 重试次数
func (m *VideoGameBean) TransferCredit(ld *TransferData, retryCounts ...int64) (transferCreditResult back.TransferCreditResult, err error) {
	// TODO 测试数据 <<
	//temp := back.TransferCreditResult{
	//	Code:    200,
	//	TradeNo: "xxooxxooxoxoxoxoxooxooxo",
	//}
	//if ld.TransferType == "in" {
	//	temp.Balance = 100 - ld.Credit
	//} else {
	//	temp.Balance = 100 + ld.Credit
	//}
	//return temp, nil
	// TODO 测试数据 >>
	var retryCount int64
	if len(retryCounts) == 1 {
		retryCount = retryCounts[0]
	}
	params := make([]string, 16)
	params[0] = "siteid=" + ld.SiteId
	params[1] = "username=" + ld.UserName
	params[2] = fmt.Sprintf("agent_id=%d", ld.AgentId)
	params[3] = "index_id=" + ld.IndexId
	params[4] = fmt.Sprintf("sh_id=%d", ld.ShId)
	params[5] = fmt.Sprintf("ua_id=%d", ld.UaId)
	params[6] = "cur=" + ld.Cur
	params[7] = "limit=" + ld.Limit
	params[8] = "lang=" + ld.Lang
	params[9] = "gtype=" + ld.Platform
	params[10] = "ip=" + ld.IP
	params[11] = "media=" + ld.Media
	params[12] = "domain=" + ld.Domain
	params[13] = fmt.Sprintf("credit=%f", ld.Credit) //"credit="
	params[14] = "type=" + ld.TransferType
	params[15] = "sw=0"
	fmt.Println("md5key:", m.Md5Key, "deskey:", m.DesKey)
	desParams, err := global.DesEncrypt([]byte(strings.Join(params, "/\\\\/")), []byte(m.DesKey))
	if err != nil {
		global.GlobalLogger.Error("TransferCredit DesEncrypt error %v", err)
		return
	}
	key := global.Md5(string(desParams) + m.Md5Key)

	p := napping.Params{"params": string(desParams), "key": key}
	s := napping.Session{Timeout: 10, Datatype: "xml"}
	fmt.Println(USER_AGENT_PRE + ld.SiteId)
	fmt.Println("我发送的数据:", params)
	fmt.Println(m.ApiUrl + "/TransferCredit")
	s.SetHeader("User-Agent", USER_AGENT_PRE+ld.SiteId)
	resp, err := s.Get(m.ApiUrl+"/TransferCredit", &p, nil, nil)
	if err != nil {
		global.GlobalLogger.Error("TransferCredit  %s", err.Error())
		if retryCount <= 1 {
			return
		} else {
			return m.TransferCredit(ld, retryCount-1) //重试
		}
	}
	if resp == nil { //返回值为空,或者返回值无效或者返回值不够
		//return m.TransferCredit(ld, retryCount-1) //重试
		err = errors.New("not fount retry data")
		return
	}
	fmt.Println("服务器返回:", resp.RawText())
	err = json.Unmarshal([]byte(resp.RawText()), &transferCreditResult)
	return
}

//第三方游戏处理
func (m *VideoGameBean) GetBalance(site_id, username, platform string) (float32, error) {
	params := make([]string, 3)
	params[0] = "site_id=" + site_id
	params[1] = "username=" + username
	params[2] = "gtype=" + platform //|隔开

	q_params, err := global.DesEncrypt([]byte(strings.Join(params, "/\\\\/")), []byte(m.DesKey))
	if err != nil {
		global.GlobalLogger.Error("ForwardGame DesEncrypt error %v", err)
		return 0, err
	}
	key := global.Md5(string(q_params) + m.Md5Key)

	p := napping.Params{"params": string(q_params), "key": key}
	s := napping.Session{Timeout: 10, Datatype: "xml"}
	s.SetHeader("User-Agent", USER_AGENT_PRE+site_id)
	resp, err := s.Get(m.ApiUrl+"/GetBalance", &p, nil, nil)
	if err != nil {
		body := ""
		if resp != nil {
			body = resp.RawText()
		}
		global.GlobalLogger.Error("ForwardGame  %s error %v", body, err)
		return 0, err
	}
	return 99, nil
}

//第三方游戏处理
func (m *VideoGameBean) GetBalanceAll(site_id, username, types string) (map[string]float32, error) {
	data := make(map[string]float32)
	params := make([]string, 3)
	params[0] = "site_id=" + site_id
	params[1] = "username=" + username
	params[2] = "types=" + types //|隔开

	q_params, err := global.DesEncrypt([]byte(strings.Join(params, "/\\\\/")), []byte(m.DesKey))
	if err != nil {
		global.GlobalLogger.Error("ForwardGame DesEncrypt error %v", err)
		return data, err
	}
	key := global.Md5(string(q_params) + m.Md5Key)

	p := napping.Params{"params": string(q_params), "key": key}
	s := napping.Session{Timeout: 10, Datatype: "xml"}
	s.SetHeader("User-Agent", USER_AGENT_PRE+site_id)

	resp, err := s.Get(m.ApiUrl+"/GetAllBalance", &p, nil, nil)
	if err != nil {
		body := ""
		if resp != nil {
			body = resp.RawText()
		}
		global.GlobalLogger.Error("ForwardGame  %s error %v", body, err)
		return data, err
	}
	//
	data["ag"] = 100
	return data, nil
}

//修改用户密码
func (m *VideoGameBean) EditUserPassword(site_id, username, platform, password string) error {
	params := make([]string, 3)
	params[0] = "site_id=" + site_id
	params[1] = "username=" + username
	params[2] = "password=" + password //|隔开
	params[3] = "gtype=" + platform    //|隔开

	q_params, err := global.DesEncrypt([]byte(strings.Join(params, "/\\\\/")), []byte(m.DesKey))
	if err != nil {
		global.GlobalLogger.Error("EditUserPassword DesEncrypt error %v", err)
		return err
	}
	key := global.Md5(string(q_params) + m.Md5Key)

	p := napping.Params{"params": string(q_params), "key": key}
	s := napping.Session{Timeout: 10, Datatype: "xml"}
	s.SetHeader("User-Agent", USER_AGENT_PRE+site_id)

	resp, err := s.Get(m.ApiUrl+"/EditUserPassword", &p, nil, nil)
	if err != nil {
		body := ""
		if resp != nil {
			body = resp.RawText()
		}
		global.GlobalLogger.Error("EditUserPassword  %s error %v", body, err)
		return err
	}
	return nil
}
