package site

import (
	"strconv"
	"strings"
	"time"

	"controllers"
	"framework/validation"
	"global"
	"models/input"
	"models/schema"

	"github.com/labstack/echo"
)

//SiteController 站点域名配置的Controller
type SetupController struct {
	controllers.BaseController
}

//添加域名配置
func (*SetupController) AddDomainSetup(ctx echo.Context) error {
	var flag bool
	var err error
	domain := ctx.FormValue("domain")          //域名
	siteId := ctx.FormValue("site")            //站点id
	siteInDexId := ctx.FormValue("site_index") //前台id
	//不能为空的校验
	if siteInDexId == "" {
		return ctx.JSON(200, global.ReplyError(60102, ctx))
	}
	if siteId == "" {
		return ctx.JSON(200, global.ReplyError(60101, ctx))
	}

	if domain == "" {
		return ctx.JSON(200, global.ReplyError(60100, ctx))
	}
	//域名格式校验
	flag = global.DomainCheck(domain)
	if !flag {
		return ctx.JSON(200, global.ReplyError(60013, ctx))
	}
	//域名是否已经使用
	flag, err = domainIsAlreadyCheck(domain)
	if flag {
		return ctx.JSON(200, global.ReplyError(60015, ctx))
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}

	siteDomain := new(schema.SiteDomain)
	siteDomain.Domain = domain
	siteDomain.SiteId = siteId
	siteDomain.IsUsed = 2

	siteDomain.SiteIndexId = siteInDexId
	siteDomain.CreateTime = time.Now().UTC().Unix()
	siteDomain.DeleteTime = 0

	//ssl key and csr
	sslKeyFile, err := ctx.FormFile("ssl_key_file")
	if err != nil {
		//文件获取失败
		return ctx.JSON(200, global.ReplyError(60022, ctx))
	}
	//文件格式校验
	keyWord := strings.Split(sslKeyFile.Filename, ".")
	if keyWord[1] != "key" {
		return ctx.JSON(200, global.ReplyError(60018, ctx))
	}
	//siteDomain.FileName = map[string]string{
	//	"ssl_key": sslKeyFile.Filename,
	//	"ssl_csr": sslKeyCsr.Filename,
	//}
	//判断是否主域名
	flag, err = siteDomainBean.IsMainDomain(siteId, siteInDexId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		siteDomain.IsDefault = 1
	} else {
		siteDomain.IsDefault = 2
	}

	_, err = siteDomainBean.Add(siteDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//编辑域名配置
func (*SetupController) EditDomainSetup(ctx echo.Context) error {
	var flag bool
	var err error

	domain := ctx.FormValue("domain") //域名
	siteInDexId := ctx.FormValue("site_index")
	siteId := ctx.FormValue("site")
	domainId := ctx.FormValue("id")
	//用来判断修改的时候是否修改配置文件
	//isChangeConfig, _ := strconv.ParseInt(ctx.FormValue("is_change"), 10, 64)
	//校验不能为空
	if domainId == "" {
		return ctx.JSON(200, global.ReplyError(60109, ctx))
	}
	if siteInDexId == "" {
		return ctx.JSON(200, global.ReplyError(60102, ctx))
	}
	if siteId == "" {
		return ctx.JSON(200, global.ReplyError(60101, ctx))
	}
	if domain == "" {
		return ctx.JSON(200, global.ReplyError(60012, ctx))
	}
	id, err := strconv.ParseInt(domainId, 10, 64)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}

	//域名格式校验
	flag = global.DomainCheck(domain)
	if !flag {
		return ctx.JSON(200, global.ReplyError(60014, ctx))
	}

	domainSet := new(schema.SiteDomain)
	domainSet.SiteId = siteId
	domainSet.SiteIndexId = siteInDexId
	domainSet.Domain = domain

	//是否已经使用
	flag, err = domainIsAlreadyCheck(domain)
	if flag {
		//this wapdomain is already used
		return ctx.JSON(200, global.ReplyError(60016, ctx))
	}
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	domainSet.Domain = domain
	domainSet.Id = id
	//if isChangeConfig == 1 {
	//	//ssl key and csr
	//	sslKeyFile, err := ctx.FormFile("ssl_key_file")
	//	if err != nil {
	//		//文件获取失败
	//		return ctx.JSON(500, global.ReplyError(60022, ctx))
	//	}
	//	//文件格式校验
	//	keyWord := strings.Split(sslKeyFile.Filename, ".")
	//	if keyWord[1] != "key" {
	//		return ctx.JSON(200, global.ReplyError(60018, ctx))
	//	}
	//
	//	srcKeyFile, err := global.ReadByte(sslKeyFile)
	//	domainSet.SslKey = string(srcKeyFile)
	//	sslKeyCsr, err := ctx.FormFile("ssl_key_csr")
	//	if err != nil {
	//		return ctx.JSON(500, global.ReplyError(60021, ctx))
	//	}
	//	//文件格式校验
	//	keyCsr := strings.Split(sslKeyCsr.Filename, ".")
	//	if keyCsr[1] != "csr" {
	//		return ctx.JSON(200, global.ReplyError(60017, ctx))
	//	}
	//
	//	srcKeyCsr, err := global.ReadByte(sslKeyCsr)
	//	domainSet.SslCsr = string(srcKeyCsr)
	//	//domainSet.FileName = map[string]string{
	//	//	"ssl_key": sslKeyFile.Filename,
	//	//	"ssl_csr": sslKeyCsr.Filename,
	//	//}
	//} else {
	//	domainSet.SslKey = ""
	//	domainSet.SslCsr = ""
	//}
	_, err = siteDomainBean.Edit(domainSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//软删除域名配置
func (*SetupController) SoftDeleteDomainSetup(ctx echo.Context) error {
	var err error
	siteDomain := new(input.DelDomainSet) //请求数据的结构体
	code := global.ValidRequest(siteDomain, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	domainSet := new(schema.SiteDomain)
	domainSet.Id = siteDomain.Id

	//判断存不存在
	_, err = siteDomainBean.Delete(domainSet)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.NoContent(204)
}

//获取域名配置列表
func (si *SetupController) GetAllListDomainSetup(ctx echo.Context) error {
	inputList := new(input.DomainSiteList)
	code := global.ValidRequest(inputList, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	listparam := new(global.ListParams)
	//获取listparam的数据
	si.GetParam(listparam, ctx)

	list, count, err := siteDomainBean.DomainSiteList(inputList, listparam)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyPagination(listparam, list, int64(len(list)), count, ctx))
}

//域名是否已经使用的校验
func domainIsAlreadyCheck(domain string) (bool, error) {
	domainSetInfo := new(schema.SiteDomain)
	domainSetInfo.Domain = domain
	flag, err := siteDomainBean.ExistDomain(domainSetInfo)
	return flag, err
}

//查看单个域名配置的详情
func (sc *SetupController) GetSingleDomainSeetup(ctx echo.Context) error {
	var flag bool
	var err error
	siteInfo := new(input.DelDomainSet) //请求数据的结构体

	if err = ctx.Bind(siteInfo); err != nil {
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	}
	//数据校验
	validtor := validation.Validation{}
	flag, err = validtor.Valid(siteInfo)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(10000, ctx))
	}
	if !flag {
		for _, err := range validtor.Errors {
			return ctx.JSON(200, global.ReplyError(err.Code(), ctx))
		}
	}

	siteDomain := new(schema.SiteDomain)
	siteDomain.Id = siteInfo.Id
	//获取
	singleDomainInfo, flag, err := siteDomainBean.GetOneSiteDomain(siteDomain)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if !flag {
		return ctx.NoContent(204)
	}
	return ctx.JSON(200, global.ReplyItem(singleDomainInfo))
}
