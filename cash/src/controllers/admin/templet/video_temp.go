//[控制器] [平台]视讯模板管理
package templet

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

//视讯模板管理
type VideoTempController struct {
	controllers.BaseController
}

//视讯模板查询
func (c *VideoTempController) GetVideoTempList(ctx echo.Context) error {
	videolist, err := videoTempBean.GetVideoList()
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	//fmt.Println(videolist)
	//video_list := make(map[int]back.VideoStyle)
	//for k, v := range videolist {
	//	if v.Aid != 0 {
	//		vlist := back.VideoStyle{}
	//		for _, value := range videolist {
	//			if value.Pid == v.Aid {
	//				vlist.Pid = v.Pid
	//				vlist.Aid = v.Aid
	//				vlist.Id = v.Id
	//				vlist.Remark = v.Remark
	//				vlist.Name = v.Name
	//				vlist.Status = v.Status
	//				vlist.Style = v.Style
	//				vlist.List = append(vlist.List, value)
	//			}
	//		}
	//		video_list[k] = vlist
	//	}
	//
	//}
	return ctx.JSON(200, global.ReplyItem(videolist))
}

//添加
func (c *VideoTempController) PostVideoTempAdd(ctx echo.Context) error {
	add_data := new(input.VideoAdd)
	code := global.ValidRequest(add_data, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := videoTempBean.AddVideo(add_data)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90001, ctx))
	}
	return ctx.NoContent(204)
}

//修改
func (c *VideoTempController) PutVideoTempUpdate(ctx echo.Context) error {
	update_data := new(input.VideoUpdate)
	code := global.ValidRequest(update_data, ctx)
	if code != 0 {
		return ctx.JSON(200, global.ReplyError(code, ctx))
	}
	data, err := videoTempBean.UpdateVideo(update_data)
	if err != nil {
		return ctx.JSON(500, global.ReplyError(60000, ctx))
	}
	if data == 0 {
		return ctx.JSON(200, global.ReplyError(90001, ctx))
	}
	return ctx.NoContent(204)
}
