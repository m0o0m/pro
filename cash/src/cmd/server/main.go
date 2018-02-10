//[主函数] 代理后台接口
package main

import (
	_ "cmd/godaemon"
	"config"
	"controllers/agency/cash"
	"encoding/json"
	"framework"
	"github.com/go-redis/redis"
	"global"
	"log"
	"models/function"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	ENVIRONMENT = "development" //testing production development
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cfg *config.Config
	var err error
	if ENVIRONMENT == "development" {
		//加载config
		path, _ := config.GetExecPath()
		configFile := filepath.Join(filepath.Dir(path), "etc/conf.yaml")
		//加载config
		cfg, err = config.ParseConfigFile(configFile)
		if err != nil {
			log.Fatalf("parse config file error:%v\n", err.Error())
			return
		}
	} else if ENVIRONMENT == "testing" || ENVIRONMENT == "production" {
		//读取环境 MGO_URL
		//从 mgo读取文件
		//MGO_BOOT_URL = "mongodb://10.10.10.186:27017/config"
		mongo_url := os.Getenv("MGO_BOOT_URL")
		if len(mongo_url) == 0 {
			mongo_url = "mongodb://master.mongo.dev.com:27017/conf"
		}
		cfg, err = config.ParseFromMongoDB(mongo_url, "/conf/p_config/p_config.yaml")
		if err != nil {
			log.Fatalf("parse config file error:%v\n", err.Error())
			return
		}
	}
	//第三方
	cash.ConfigThird = cfg.Third
	//全局初始化
	global.InitLog(cfg, "server")
	defer global.GlobalLogger.Close()

	//初始化数据库
	err = global.InitMysql(cfg.Mysqls)
	if err != nil {
		//数据库连接错误
		global.GlobalLogger.Error("InitDb error:%v\n", err.Error())
		return
	}

	//初始化redis
	err = global.InitRedis(cfg.Redis)
	if err != nil {
		//redis连接失败
		global.GlobalLogger.Error("Redis connection failed:%v\n", err.Error())
		return
	}
	//初始化 Influxdb
	//err = global.InitInfluxdb(cfg.Influx)
	//if err != nil {
	//	//redis连接失败
	//	global.GlobalLogger.Error("Influxdb connection failed:%v\n", err.Error())
	//	//return
	//}
	//启动web服务
	app, err := framework.NewApp(cfg)
	if err != nil {
		//错误日志
		global.GlobalLogger.Error("NewApp error:%v\n", err.Error())
		return
	}

	//启动定时任务
	//go Crons(cfg.Redis[global.DefaultRedis])
	//启动
	err = app.Run()
	if err != nil {
		global.GlobalLogger.Error("app start error:%v\n", err.Error())
		return
	}
}

//定时任务
func Crons(cfg config.RedisConfig) {
	ticker := time.NewTicker(time.Second * time.Duration(cfg.Runtime))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for i := 0; i < cfg.Goroutinenum; i++ {
				go Runs()
			}
		}
	}
}

//执行判定流程
func Runs() {
	//弹出key
	keys, err := global.GetRedis().LPop("is_login").Result()
	if err == redis.Nil {
		return
	} else if err != nil {
		global.GlobalLogger.Error("error:%v\n", err.Error())
		return
	}
	if keys == "" {
		return
	}

	//获取序列化的数据
	result, err := global.GetRedis().Get(keys).Result()
	if err == redis.Nil {
		return
	} else if err != nil {
		global.GlobalLogger.Error("error:%v\n", err.Error())
		return
	}
	if result == "" {
		return
	}
	var s global.RedisStruct
	json.Unmarshal([]byte(result), &s)

	//过期
	if s.ExpirTime < time.Now().Unix() {
		//根据agencyId更新数据库在线状态,
		err = UpDataLoginStatus(s.Id)
		if err != nil {
			global.GlobalLogger.Error("error:%v\n", err.Error())
			return
		}
		//清除token
		err = global.GetRedis().Del(keys).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%v\n", err.Error())
			return
		}
	} else {
		//不过期
		err = global.GetRedis().RPush("is_login", keys).Err()
		if err != nil {
			global.GlobalLogger.Error("error:%v\n", err.Error())
			return
		}
	}
}

//更新??/
func UpDataLoginStatus(agencyId int64) (err error) {
	_, err = function.ChangeLoginStatus(agencyId)
	return
}
