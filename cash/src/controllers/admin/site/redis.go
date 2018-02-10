package site

import (
	"controllers"
	"github.com/labstack/echo"
	"global"
	"models/input"
)

// 缓存管理
type RedisController struct {
	controllers.BaseController
}

func (c *RedisController) Search(ctx echo.Context) error {
	list := c.GetRedis(input.SearchRedis{})

	return ctx.JSON(200, global.ReplyItem(list))
}

// 获取缓存
func (c *RedisController) GetRedis(request input.SearchRedis) (keys []string) {
	var key string
	if request.Keyword != "" {
		key = "*" + request.Keyword + "*"
	}
	keys, _ = global.GetRedis().Keys(key).Result()

	return
}
