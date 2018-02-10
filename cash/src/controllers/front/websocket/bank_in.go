package websocket

import (
	"framework/logger"
	"github.com/bluele/gcache"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"global"
)

type BankIn struct {
}

//入款成功的通知
func (*BankIn) DepositSuccess(ctx echo.Context) error {
	ws, err := upGrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	for {
		// TODO 读取订单号
		_, msg, err := ws.ReadMessage()
		if err != nil {
			logger.Global.Error("err:%", err.Error())
			return err
		}
		// TODO 得到订单号后,一直轮询,直到订单成功或者前台页面断开
		if len(msg) > 0 {
			no := string(msg)
			for {
				//fmt.Println("订单号", no)
				_, err := global.DepositCache.Get(no)
				if err != nil {
					if err != gcache.KeyNotFoundError {
						global.GlobalLogger.Error("error:%s", err.Error())
						return err
					}
				} else {
					global.DepositCache.Remove(no)
					err := ws.WriteMessage(websocket.TextMessage, []byte("成功存款了"))
					if err != nil {
						logger.Global.Error("err:%", err.Error())
					}
					return err
				}
			}
		}
	}
}

//出款成功的通知
func (*BankIn) WithdrawSuccess(ctx echo.Context) error {
	ws, err := upGrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	for {
		// TODO 读取订单号
		_, msg, err := ws.ReadMessage()
		if err != nil {
			logger.Global.Error("err:%", err.Error())
			return err
		}
		// TODO 得到订单号后,一直轮询,直到订单成功或者前台页面断开
		if len(msg) > 0 {
			no := string(msg)
			for {
				//fmt.Println("订单号", no)
				_, err := global.WithdrawalCache.Get(no)
				if err != nil {
					if err != gcache.KeyNotFoundError {
						global.GlobalLogger.Error("error:%s", err.Error())
						return err
					}
				} else {
					global.WithdrawalCache.Remove(no)
					err := ws.WriteMessage(websocket.TextMessage, []byte("成功取款了"))
					if err != nil {
						logger.Global.Error("err:%", err.Error())
					}
					return err
				}
			}
		}
	}
}
