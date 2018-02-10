package agency

import (
	"controllers/agency/account"

	"router"

	"github.com/labstack/echo"
)

const apiPath = "/api"

func AccountRouter(c *echo.Echo) {
	sign := new(account.AgencySignController)
	c.GET(apiPath+"/captcha", router.VerCode)
	c.POST(apiPath+"/login", sign.Login)
	//加入路由权限的判断
	//e := c.Group("", router.GetRedisToken, router.AgencyPerCheck)
	e := c.Group(apiPath, router.GetRedisToken, router.AgencyAccessLog) //开启代理操作日志
	e.GET("/logout", sign.Logout)
	e.POST("/password", sign.SetPwd)
	e.GET("/menu/role", sign.Menu) //根据角色id获取菜单

	//股东管理
	first := new(account.FirstAgencyController)
	e.GET("/agent/first", first.Index)                             //股东列表
	e.GET("/agent/first/discount", first.MemberRegDiscountSet)     //查询会员优惠
	e.PUT("/agent/first/discount", first.MemberRegDiscountSetEdit) //修改会员注册优惠
	e.POST("/firstAgency", first.Add)                              //增加股东
	e.PUT("/firstAgency/status", first.Status)                     //启用/禁用股东
	e.PUT("/firstAgency", first.BaseInfoEdit)                      //修改详情
	e.GET("/firstAgency/info", first.BaseInfo)                     //查看股东详情

	//总代管理
	second := new(account.SecondAgencyController)
	e.GET("/agent/second", second.Index)                             //总代列表
	e.GET("/agent/second/discount", second.MemberRegDiscountSet)     //查询会员优惠
	e.PUT("/agent/second/discount", second.MemberRegDiscountSetEdit) //修改会员注册优惠
	e.POST("/secondAgency", second.Add)                              //增加总代
	e.PUT("/secondAgency/status", second.Status)                     //启用/禁用总代
	e.PUT("/secondAgency", second.BaseInfoEdit)                      //修改详情
	e.GET("/secondAgency/info", second.BaseInfo)                     //查看总代详情

	//代理管理
	third := new(account.ThirdAgencyController)
	e.GET("/agent/third", third.Index)                             //代理列表
	e.GET("/agent/third/discount", third.MemberRegDiscountSet)     //查询会员优惠
	e.PUT("/agent/third/discount", third.MemberRegDiscountSetEdit) //修改会员注册优惠
	e.POST("/thirdAgency", third.Add)                              //增加代理
	e.PUT("/thirdAgency/status", third.Status)                     //启用/禁用代理
	e.PUT("/thirdAgency", third.BaseInfoEdit)                      //修改详情
	e.GET("/thirdAgency/info", third.BaseInfo)                     //查看代理详情
	e.GET("/agent/third/info", third.DetailInfo)                   //查看代理详细资料
	e.PUT("/agent/third/info", third.DetailInfoEdit)               //修改代理详细资料
	//公用的下拉框
	e.GET("/third/dropf", first.SiteIdByAgencyId)          //开户人下所有站点（下拉框）
	e.GET("/agent/first/drop", second.FirstIdByStieId)     //股东下拉框查询
	e.GET("/agent/second/drop", third.SecondIdNameByFirst) //总代理下拉框  取所有的总代的id和名称
	e.GET("/agent/third/drop", third.ThirdIdNameBySite)    //代理下拉框  取站点下所有代理id和帐号

	e.GET("/agent/domain", third.SpreadDomain)       //查看推广域名
	e.POST("/agent/domain", third.SpreadDomainAdd)   //添加推广域名
	e.PUT("/agent/domain", third.SpreadDomainEdit)   //修改推广域名
	e.DELETE("/agent/domain", third.SpreadDomainDel) //删除推广域名
	//会员账号管理
	member := new(account.MemberController)
	e.GET("/member", member.Index)                      //会员列表
	e.GET("/member/sort/drop", member.MemberSortDrop)   //会员管理排序下拉框
	e.PUT("/member/status", member.Status)              //修改会员状态
	e.GET("/member/infos", member.BaseInfo)             //获取会员基本资料
	e.PUT("/member/base", member.BaseInfoEdit)          //修改会员基本资料
	e.GET("/member/detail", member.DetailInfo)          //获取会员详细资料
	e.GET("/member/bank", member.Bank)                  //获取会员出款银行集合
	e.PUT("/member/detail", member.DetailInfoEdit)      //修改会员详细资料
	e.DELETE("/member/bank", member.MemberBankDel)      //删除会员银行
	e.GET("/member/bank/info", member.BankInfo)         //获取会员银行详情
	e.PUT("/member/bank", member.MemberBankEdit)        //修改会员银行
	e.PUT("/member/offline", member.Offline)            //会员踢线
	e.PUT("/member/batch/status", member.BatchStatus)   //批量禁用启用会员
	e.PUT("/member/batch/offline", member.BatchOffline) //批量踢线会员
	e.GET("/member/bankDrop", member.MemberBankDrop)    //会员银行下拉框
	//代理申请管理
	agencyReg := new(account.AgencyRegController)
	e.GET("/agent/register", agencyReg.Index)                //查看代理申请列表
	e.POST("/agent/register", agencyReg.Add)                 //审核代理申请
	e.DELETE("/agent/register", agencyReg.AgentRegState)     //删除代理申请
	e.GET("/agent/register/setting", agencyReg.Set)          //获取代理注册申请设定
	e.POST("/agent/register/setting", agencyReg.SetEdit)     //修改代理注册申请设定
	e.GET("/agent/register/one", agencyReg.OneAgencyRegById) //查询一条代理申请
	//子账号管理
	subAccount := new(account.SubAccountController)
	e.POST("/agent/sub/permission", subAccount.PermissionEdit) //设置子账号权限
	e.GET("/agent/sub/permission", subAccount.Permission)      //子账号权限列表

	e.GET("/agent/sub/list", subAccount.Index)    //查看子账号列表
	e.GET("/agent/sub/info", subAccount.BaseInfo) //获取子账号详情
	e.PUT("/agent/sub", subAccount.BaseInfoEdit)  //修改子账号基本资料
	e.PUT("/agent/sub/status", subAccount.Status) //启用/禁用子账号
	e.DELETE("/agent/sub", subAccount.Delete)     //删除子账号
	e.POST("/agent/sub", subAccount.Add)          //添加子账号

	e.POST("/agent/subtoken/status", subAccount.AccessToken)  //修改口令验证设置
	e.GET("/agent/subtoken/info", subAccount.AccessTokenInfo) //查看口令验证设置信息
	e.GET("/agent/subtoken", subAccount.GenKey)               //服务器端生成key

	//站点会员注册设定
	memberRegsetting := new(account.MemberRegController)
	e.POST("/member/register/setting", memberRegsetting.RegisterSetEdit) //配置站点会员注册信息
	e.GET("/member/register/setting", memberRegsetting.RegisterSet)      //获取站点会员注册信息
	//会员层级
	memberLevel := new(account.MemberLevelController)
	e.POST("/member/level", memberLevel.Add)                        //站点添加会员层级
	e.GET("/member/level/info", memberLevel.Info)                   //获取会员层级信息
	e.GET("/member/level/list", memberLevel.Index)                  //站点会员层级列表
	e.PUT("/member/level", memberLevel.InfoEdit)                    //修改站点会员层级信息
	e.PUT("/member/level/regress", memberLevel.Regress)             //回归层级
	e.GET("/member/level/drop", memberLevel.MemberLevelDrop)        //会员层级下拉框
	e.PUT("/member/level/selfrebate", memberLevel.StatusSelfRebate) //开启自助返水
	e.GET("/member/level/memberlist", memberLevel.MemberList)       //会员详情列表
	e.PUT("/member/level/lock", memberLevel.Locked)                 //锁定会员层级
	e.PUT("/member/level/move", memberLevel.Move)                   //移动会员分层
	//体系查询
	search := new(account.SearchController)
	e.GET("/agent/search", search.Index) //体系查询

	e.GET("/member/level/payset", memberLevel.PaySet)     //获取层级支付设定
	e.PUT("/member/level/payset", memberLevel.PaySetEdit) //修改层级支付设定
}
