//[主函数] websocket服务
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"config"
	"framework/logger"
	"global"
	"models/function"
	"models/schema"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} //不使用默认设置，如果线上环境可能需要使用默认配置
var chananel = make(chan schema.Listening)                                                 //数据chan

var configFile *string = flag.String("config", "./bin/etc/conf.yaml", "agency config file")
var agentSlice []map[string]*websocket.Conn //socket对应关系存储

//发送消息结构体
type Message struct {
	Message     interface{} `json:"message"`
	SiteId      string      `json:"site_id"`
	SiteIndexId string      `json:"site_index_id"`
	Count       int64       `json:"count"`
}

//测试用[正式修改之后可以删除]
func hu(w http.ResponseWriter, r *http.Request) {
	siteid := r.FormValue("site_id")
	siteIndexId := r.FormValue("site_index_id")
	fmt.Println(siteIndexId, siteid)
	s := schema.Listening{"zym", "b", 1}
	chananel <- s
}

func main() {
	//数据库初始化
	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		log.Fatalf("parse config file error:%v\n", err.Error())
		return
	}
	//初始化日志
	global.InitLog(cfg)
	defer global.GlobalLogger.Close()
	//初始化数据库
	err = global.InitMysql(cfg.Mysqls)
	if err != nil {
		//数据库连接错误
		global.GlobalLogger.Error("InitDb error:%v\n", err.Error())
		return
	}
	//TODO 这里需要channel放在公司入款,线上入款，出款管理申请部分，传递消息过来这边，监听之后从channel中间传递过来的是[models/schema]schema.Listening
	http.HandleFunc("/o", hu) //这个路由是不走线上入款，公司入款和出款管理的channel部分，用来测试用
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	err = http.ListenAndServe(cfg.Wesocketport, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
func handleConnections(w http.ResponseWriter, r *http.Request) {
	//如果限制连接就可以使用ip+port限制，根据ip区分客户端，其他的可以根据r.Request提交的数据查找相应的内容
	siteId := r.FormValue("site_id")
	siteIndexId := r.FormValue("site_index_id")
	if siteId == "" || siteIndexId == "" {
		http.Error(w, "site_id and site_index_id must not empty", 403)
	}

	////token验证,权限验证
	//token := r.FormValue("token")    //token
	//roleId := r.FormValue("role_id") //角色id
	//err := parseToken(token, roleId)
	//if err != nil {
	//	global.GlobalLogger.Error("error:%s", err.Error())
	//	return
	//}
	s := fmt.Sprintf("%s%s", siteId, siteIndexId)

	//注册成为websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return
	}
	defer ws.Close()
	//存储连接[todo 这里可能还要考虑map并发读写问题]
	agent := make(map[string]*websocket.Conn)
	agent[s] = ws
	agentSlice = append(agentSlice, agent)
	clients[ws] = true
	//监听接收一个[models/schema]schema.Listening，
	for {
		var msg Message
		//todo 这里检测全局变量chan从公司入款，线上入款，出款管理传递过来的channel,如果没有值过来会堵塞
		s := <-chananel
		if s.Types == 1 {
			//todo 这里解析取出来的数据可能还需要加工
			//获取最新的没有确认得公司入款
			newincome := new(function.MemberCompanyIncomeBean)
			info, count, err := newincome.GetNotConfirm(s.SiteId, s.SiteIndexId)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return
			}
			msg = Message{SiteIndexId: s.SiteIndexId, SiteId: s.SiteId, Message: info, Count: count}
		} else if s.Types == 2 {
			//获取最新的线上入款
			onLineBean := new(function.OnlineEntryRecordBean)
			info, count, err := onLineBean.GetNotConfirm(s.SiteId, s.SiteIndexId)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return
			}
			msg = Message{SiteIndexId: s.SiteIndexId, SiteId: s.SiteId, Message: info, Count: count}
		} else {
			//获取没有确认得最新的出款管理
			makeMoney := new(function.MakeMoneyBean)
			info, count, err := makeMoney.GetOperateRecord(s.SiteId, s.SiteIndexId)
			if err != nil {
				global.GlobalLogger.Error("error:%s", err.Error())
				return
			}
			msg = Message{SiteIndexId: s.SiteIndexId, SiteId: s.SiteId, Count: count, Message: info}
		}
		broadcast <- msg
	}
}

//单点推送
func handleMessages() {
	for {
		msg := <-broadcast
		var pushClient []*websocket.Conn
		newS := fmt.Sprintf("%s%s", msg.SiteId, msg.SiteIndexId)
		lenAgent := len(agentSlice)
		for i := 0; i < lenAgent; i++ {
			for k, v := range agentSlice[i] {
				if newS == k {
					pushClient = append(pushClient, v)
				}
			}
		}
		for i := 0; i < len(pushClient); i++ {
			for client := range clients {
				if pushClient[i] == client {
					err := client.WriteJSON(msg)
					if err != nil {
						global.GlobalLogger.Error("error:%s", err.Error())
						client.Close()
						delete(clients, client)
					}
				}
			}
		}
	}
}

//token解析校验
func parseToken(token string, roleId string) error {
	//从redis获取，获取不到就是过期
	result, err := global.GetRedis().Get(token).Result()
	if err == redis.Nil {
		return echo.ErrUnauthorized
	} else if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return echo.ErrUnauthorized
	}
	if result == "" {
		return echo.ErrUnauthorized
	}
	//从数据库获取获取不到就是过期
	admin_bean := new(function.AdminBean)
	_, flag, err := admin_bean.GetInfoByToken(token)
	if err != nil {
		global.GlobalLogger.Error(logger.ERROR, err.Error())
		return err
	}
	if !flag {
		return echo.ErrUnauthorized
	}
	var b []byte
	//解析
	if roleId == "5" { //操作的是管理员及其子账号
		results := new(global.AdminRedisStruct)
		err = json.Unmarshal([]byte(result), &results)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		//刷新时间
		results.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
		b, err = json.Marshal(results)
		if err != nil {
			return err
		}
	} else if roleId == "1" { //操作这里的是开户人及其子账号
		s := new(global.RedisStruct)
		err = json.Unmarshal([]byte(result), &s)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}

		s.ExpirTime = time.Now().Add(global.AgencyRedisExp).Unix()
		b, err = json.Marshal(s)
		if err != nil {
			global.GlobalLogger.Error("error:%s", err.Error())
			return err
		}
	}
	//重新set
	err = global.GetRedis().Set(token, b, 0).Err()
	if err != nil {
		global.GlobalLogger.Error("error:%s", err.Error())
		return echo.ErrUnauthorized
	}
	return nil
}
