package app

import (
	"config"
	"github.com/labstack/echo"
)

const (
	Offline = iota
	Online
	Unknown
)

type App struct {
	Cfg       *config.Config
	WebServer *echo.Echo
	State     int32
}

func (a *App) Status() string {
	var status string
	switch a.State {
	case Online:
		status = "online"
	case Offline:
		status = "offline"
	case Unknown:
		status = "unknown"
	default:
		status = "unknown"
	}
	return status
}

func (a *App) Stop() {
	//TODO 关闭
}

//初始化
func (a *App) Run() (err error) {
	//启动服务
	err = a.WebServer.Start(a.Cfg.Addr)
	if err != nil {
		return
	}
	return
}
