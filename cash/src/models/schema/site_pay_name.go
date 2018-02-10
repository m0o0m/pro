package schema

import "global"

//后台充值第三方配置表
type SitePayName struct {
	ID          int64  `xorm:"id PK autoincr"` //
	PayName     string `xorm:"pay_name" `      // 第三方名字,银行卡名字
	PayType     int64  `xorm:"pay_type" `      //第三方支付类型
	State       int64  `xorm:"state" `         //状态1停用2启用
	PayId       string `xorm:"pay_id" `        //商户ID,如果type类型是银行卡的话,就是银行卡号
	PayKey      string `xorm:"pay_key" `       //密钥
	FUrl        string `xorm:"f_url" `         //支付域名
	Vircarddoin string `xorm:"vircarddoin" `   //国付宝终端号
	TerminalId  string `xorm:"terminal_id" `   //宝付终端标号
	Type        int64  `xorm:"type" `          //1第三方,2银行卡
	MyName      string `xorm:"my_name" `       //收款人姓名
	Address     string `xorm:"address" `       //银行卡开户行地址
	Lid         int64  `xorm:"lid" `           //默认未分层,站点套餐id
}

func (m *SitePayName) TableName() string {
	return global.TablePrefix + "site_pay_name"
}
