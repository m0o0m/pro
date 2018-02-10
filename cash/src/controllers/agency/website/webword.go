//文案编辑
package website

import (
	"controllers"

	"github.com/labstack/echo"
	"global"
	"models/back"
	"models/input"
	"strings"
)

type WebwordController struct {
	controllers.BaseController
}

//首页文案查询
func (*WebwordController) IwordList(ctx echo.Context) error {
	iwordIndexs := new(input.SiteIWodList)
	code := global.ValidRequest(iwordIndexs, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	iwordIndexs.SiteId = user.SiteId
	data, err := siteIwordBean.IWordList(iwordIndexs)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//文案单条查询（根据id查询）
func (*WebwordController) IwordInfor(ctx echo.Context) error {
	iwordInfo := new(input.SiteCopyInfo)
	code := global.ValidRequest(iwordInfo, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)

	ok, data, err := siteIwordBean.IwordInfo(iwordInfo, user.SiteId)

	if err != nil {
		return ctx.JSON(500, global.ReplyError(10001, ctx))
	}
	if !ok {
		return ctx.JSON(200, global.ReplyError(90715, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

////文案添加
//func (*WebwordController) IwordAdd(ctx echo.Context) error{
//	siteIwordAdd := new(input.SiteIwordAdd)
//	code := global.ValidRequest(siteIwordAdd, ctx)
//	if code != 0 {
//		return ctx.JSON(200, global.ReplyError(code, ctx))
//	}
//	user := ctx.Get("user").(*global.RedisStruct)
//	if len(siteIwordAdd.SiteId) == 0 {
//		siteIwordAdd.SiteId = user.SiteId
//	}
//	if len(siteIwordAdd.SiteIndexId) == 0 {
//		siteIwordAdd.SiteIndexId = "a"
//	}
//
//	count, err := siteIwordBean.IwordAdd(siteIwordAdd)
//	if err != nil {
//		global.GlobalLogger.Error("error:%s", err.Error())
//		return ctx.JSON(500, global.ReplyError(60000, ctx))
//	}
//	if count != 1 {
//		return ctx.JSON(200, global.ReplyError(90701, ctx))
//	}
//
//	return ctx.NoContent(204)
//}

//文案修改

func (*WebwordController) IwordEidt(ctx echo.Context) error {
	copyUpdate := new(input.IwordUpdate)
	code := global.ValidRequest(copyUpdate, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	code, count, err := siteIwordBean.IwordEidt(copyUpdate, user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(90713, ctx))
	}

	return ctx.NoContent(204)
}

//优惠文案查询
func (*WebwordController) ActivityList(ctx echo.Context) error {
	iwordIndex := new(input.SiteActivityCopyList)
	code := global.ValidRequest(iwordIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	data, err := siteIwordBean.IwordActivityList(iwordIndex, user.SiteId)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	return ctx.JSON(200, global.ReplyItem(data))
}

//优惠文案内容查询
func (*WebwordController) ActivityInfo(ctx echo.Context) error {
	iwordIndex := new(input.SiteActivityCopyInfo)
	code := global.ValidRequest(iwordIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	if iwordIndex.TopId == 0 && len(iwordIndex.TypeName) != 0 { //优惠分类
		data, err := siteIwordBean.IwordActivityType(iwordIndex, user.SiteId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(data))

	} else { //优惠详情
		ok, data, err := siteIwordBean.IwordActivityInfo(iwordIndex, user.SiteId)

		if err != nil {
			return ctx.JSON(500, global.ReplyError(10001, ctx))
		}
		if !ok {
			return ctx.JSON(200, global.ReplyError(90715, ctx))
		}
		return ctx.JSON(200, global.ReplyItem(data))
	}
}

//
////优惠文案添加
//func (*WebwordController) IwordActivityAdd(ctx echo.Context) error{
//	siteIwordActivityAdd := new(input.SiteIwordActivityAdd)
//	code := global.ValidRequest(siteIwordActivityAdd, ctx)
//	if code != 0 {
//		return ctx.JSON(200, global.ReplyError(code, ctx))
//	}
//
//	count, err := siteIwordBean.IwordActivityAdd(siteIwordActivityAdd)
//	if err != nil {
//		global.GlobalLogger.Error("error:%s", err.Error())
//		return ctx.JSON(500, global.ReplyError(60000, ctx))
//	}
//	if count != 1 {
//		return ctx.JSON(200, global.ReplyError(90701, ctx))
//	}
//
//	return ctx.NoContent(204)
//}

//优惠修改/添加

func (*WebwordController) ActivityEdite(ctx echo.Context) error {
	activityEditeTitle := new(input.ActivityEditeTitle)
	code := global.ValidRequest(activityEditeTitle, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	if activityEditeTitle.Id != 0 {
		code, count, err := siteIwordBean.ActivityEditeTitle(activityEditeTitle, user.SiteId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if code != 0 {
			return ctx.JSON(200, global.ReplyError(code, ctx))
		}
		if count != 1 {
			return ctx.JSON(200, global.ReplyError(90713, ctx))
		}

	} else {
		count, err := siteIwordBean.IwordActivityAdd(activityEditeTitle, user.SiteId)
		if err != nil {
			return ctx.JSON(500, global.ReplyError(60000, ctx))
		}
		if count != 1 {
			return ctx.JSON(200, global.ReplyError(90701, ctx))
		}
	}
	return ctx.NoContent(204)
}

//优惠内容修改
func (*WebwordController) ActivityEditeContent(ctx echo.Context) error {
	iwordIndex := new(input.ActivityEditeContent)
	code := global.ValidRequest(iwordIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//如果图片路径和内容都为空，返回
	if iwordIndex.Img == "" && iwordIndex.Content == "" {
		return ctx.JSON(200, global.ReplyError(50157, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	code, count, err := siteIwordBean.ActivityEditeContent(iwordIndex, user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(90713, ctx))
	}
	return ctx.NoContent(204)
}

//优惠删除

func (*WebwordController) ActivityDel(ctx echo.Context) error {
	iwordIndex := new(input.ActivityDel)
	code := global.ValidRequest(iwordIndex, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)

	code, count, err := siteIwordBean.ActivityDel(iwordIndex.Id, user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(90713, ctx))
	}

	return ctx.NoContent(204)
}

//站点线路检测数据查询

func (*WebwordController) SiteDetectList(ctx echo.Context) error {
	siteDetect := new(input.SiteDetect)
	code := global.ValidRequest(siteDetect, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	//查询数据
	sdata := []back.SiteDetect{}
	data, err := siteIwordBean.SiteDetectList(user.SiteId, siteDetect.SiteIndexId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//处理域名
	for _, v := range data {
		info := back.SiteDetect{}
		domainStr := strings.Split(v.Domain, "//")
		info.Id = v.Id
		info.SiteIndexId = v.SiteIndexId
		info.Domain = v.Domain
		info.Content = domainStr[1]
		if domainStr[0] == "http:" {
			info.Protocol = 1
		} else if domainStr[0] == "https:" {
			info.Protocol = 2
		}
		sdata = append(sdata, info)
	}
	return ctx.JSON(200, global.ReplyItem(sdata))
}

//站点线路检测数据修改

func (*WebwordController) SiteDetectEdit(ctx echo.Context) error {
	siteDetectEdit := new(input.SiteDetectEdit)
	code := global.ValidRequest(siteDetectEdit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	//存入的域名数据
	domainStr := ""
	if siteDetectEdit.Protocol == 1 {

		domainStr = "http://" + siteDetectEdit.Content
	} else {
		domainStr = "https://" + siteDetectEdit.Content

		domainStr = "http://" + siteDetectEdit.Content
	}
	user := ctx.Get("user").(*global.RedisStruct)
	code, count, err := siteIwordBean.SiteDetectEdit(siteDetectEdit.Id, user.SiteId, siteDetectEdit.SiteIndexId, domainStr)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(90719, ctx))
	}
	return ctx.NoContent(204)
}

//站点线路检测数据删除

func (*WebwordController) SiteDetectDel(ctx echo.Context) error {
	siteDetectDel := new(input.SiteDetectDel)
	code := global.ValidRequest(siteDetectDel, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	user := ctx.Get("user").(*global.RedisStruct)
	code, count, err := siteIwordBean.SiteDetectDel(siteDetectDel, user.SiteId)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(90719, ctx))
	}
	return ctx.NoContent(204)
}

//站点线路检测数据添加

func (*WebwordController) SiteDetectAdd(ctx echo.Context) error {
	siteDetectEdit := new(input.SiteDetectAdd)
	code := global.ValidRequest(siteDetectEdit, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	domainStr := ""
	if siteDetectEdit.Protocol == 1 {

		domainStr = "http://" + siteDetectEdit.Content
	} else {
		domainStr = "https://" + siteDetectEdit.Content
	}
	user := ctx.Get("user").(*global.RedisStruct)
	code, count, err := siteIwordBean.SiteDetectAdd(user.SiteId, siteDetectEdit.SiteIndexId, domainStr)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	if count != 1 {
		return ctx.JSON(200, global.ReplyError(90719, ctx))
	}
	return ctx.NoContent(204)
}
