package import_member

import (
	"controllers"
	"github.com/labstack/echo"
	"io"
	"models/input"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"global"
	"models/function"
)

//代理管理
type UploadController struct {
	controllers.BaseController
}

var imb = new(function.ImportMemberBean)

//导入会员
func (*UploadController) Upload(ctx echo.Context) error {
	//获取表单数据
	siteId := ctx.FormValue("siteId") //站点id
	if siteId == "" {
		return ctx.JSON(200, global.ReplyError(30296, ctx))
	}
	siteIndexId := ctx.FormValue("siteIndexId") //站点前台id
	if siteIndexId == "" {
		return ctx.JSON(200, global.ReplyError(30296, ctx))
	}
	//查询站点是否存在
	has, err := imb.IsExistSite(siteId, siteIndexId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(60010, ctx))
	}
	levelId := ctx.FormValue("levelId") //层级id
	if levelId == "" {
		return ctx.JSON(200, global.ReplyError(30296, ctx))
	}
	//查询层级是否存在
	has, err = imb.IsExistMemberLevel(siteId, siteIndexId, levelId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30141, ctx))
	}
	firstAgencyId := ctx.FormValue("firstAgencyId") //股东id
	if firstAgencyId == "" {
		return ctx.JSON(200, global.ReplyError(30296, ctx))
	}
	//查询股东是否存在
	has, err = imb.IsExistFirstAgency(siteId, siteIndexId, firstAgencyId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30238, ctx))
	}
	secondAgencyId := ctx.FormValue("secondAgencyId") //总代id
	if secondAgencyId == "" {
		return ctx.JSON(200, global.ReplyError(30296, ctx))
	}
	//查询总代是否存在
	has, err = imb.IsExistSecondAgency(siteId, siteIndexId, firstAgencyId, secondAgencyId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30297, ctx))
	}
	thirdAgencyId := ctx.FormValue("thirdAgencyId") //代理id
	if thirdAgencyId == "" {
		return ctx.JSON(200, global.ReplyError(30296, ctx))
	}
	//查询代理是否存在
	has, err = imb.IsExistThirdAgency(siteId, siteIndexId, secondAgencyId, thirdAgencyId)
	if err != nil {
		return ctx.JSON(200, global.ReplyError(60000, ctx))
	}
	if !has {
		return ctx.JSON(200, global.ReplyError(30298, ctx))
	}
	// 上传文件
	file, err := ctx.FormFile("excel")
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	src, err := file.Open()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	defer src.Close()
	//文件路径
	path := "../src/controllers/admin/import_member/"
	//新建文件
	dst, err := os.Create(path + file.Filename)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	//读取excel
	excel, err := excelize.OpenFile(path + file.Filename)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}

	var im input.ImportMember
	members := make([]input.ImportMember, 0)
	rows := excel.GetRows("工作表1")
	//给结构体赋值
	for i := 0; i < len(rows)-1; i++ {
		im.SiteId = siteId
		im.SiteIndexId = siteIndexId
		im.LevelId = levelId
		im.FirstAgencyId, err = strconv.ParseInt(firstAgencyId, 10, 64)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		im.SecondAgencyId, err = strconv.ParseInt(secondAgencyId, 10, 64)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		im.ThirdAgencyId, err = strconv.ParseInt(thirdAgencyId, 10, 64)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
		im.Account = rows[i+1][0]
		im.UserName = rows[i+1][1]
		im.Money, _ = strconv.ParseFloat(rows[i+1][2], 10)
		im.PayCard, _ = strconv.ParseInt(rows[i+1][3], 10, 64)
		im.PayNum = rows[i+1][4]
		im.Password = rows[i+1][5]
		members = append(members, im)
	}
	//存在的账号
	var existAccount []string
	//做新增操作的会员
	insertMember := make([]input.ImportMember, 0)
	//判断会员账号是否存在
	for k := range members {
		has, err := imb.GetMemberAccount(members[k].SiteId, members[k].Account)
		if err != nil {
			return ctx.JSON(200, global.ReplyError(60000, ctx))
		}
		if has { //数据库存在就提出来
			existAccount = append(existAccount, members[k].Account)
		} else { //数据库不存在的就添加
			insertMember = append(insertMember, members[k])
		}
	}
	dst.Close()
	if len(insertMember) != 0 {
		count, err := imb.AddMember(insertMember)
		if err != nil {
			return ctx.JSON(200, global.ReplyError(60000, ctx))
		}
		if count == 0 {
			return ctx.JSON(200, global.ReplyError(50174, ctx))
		}
	}
	//功能完成后删除excel
	err = os.Remove("./" + path + file.Filename)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return err
	}
	if len(existAccount) != 0 {
		return ctx.JSON(200, existAccount)
	}
	return ctx.NoContent(204)
}

//层级下拉
func (*UploadController) GetLevelDrop(ctx echo.Context) error {
	importSite := new(input.ImportSite)
	code := global.ValidRequest(importSite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := imb.LevelDrop(importSite)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//股东下拉
func (*UploadController) GetFirstAgencyDrop(ctx echo.Context) error {
	importSite := new(input.ImportSite)
	code := global.ValidRequest(importSite, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := imb.FirstAgencyDrop(importSite)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//总代下拉
func (*UploadController) GetSecondAgencyDrop(ctx echo.Context) error {
	importAgency := new(input.ImportAgency)
	code := global.ValidRequest(importAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := imb.SecondAgencyDrop(importAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}

//代理下拉
func (*UploadController) GetThirdAgencyDrop(ctx echo.Context) error {
	importAgency := new(input.ImportAgency)
	code := global.ValidRequest(importAgency, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	list, err := imb.ThirdAgencyDrop(importAgency)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(list))
}
