//[wap后台主函数]
package main

import (
	"config"
	"log"
	"os"
	"runtime"

	"controllers/wap/wdeposit"
	"framework"
	"global"
	"path/filepath"
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
		configFile := filepath.Join(filepath.Dir(path), "etc/wap.yaml")
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
		cfg, err = config.ParseFromMongoDB(mongo_url, "/conf/p_config/p_wap.yaml")
		if err != nil {
			log.Fatalf("parse config file error:%v\n", err.Error())
			return
		}
	}
	//全局初始化
	global.InitLog(cfg, "wap")
	defer global.GlobalLogger.Close()
	//配置文件全局变量
	//配置文件全局变量
	wdeposit.TotalConfig = cfg
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
	/*err = global.InitInfluxdb(cfg.Influx)
	if err != nil {
		global.GlobalLogger.Error("Influxdb connection failed:%v\n", err.Error())
		//return
	}*/
	//启动web服务
	app, err := framework.NewWap(cfg)
	if err != nil {
		//错误日志
		global.GlobalLogger.Error("NewApp error:%v\n", err.Error())
		return
	}
	//初始化render
	//render.InitRootPath(&cfg.TemplateConfig)
	////初始化静态目录
	//if ENVIRONMENT == "development" {
	//	app.WebServer.Static("/template", render.ViewPath)
	//} else if ENVIRONMENT == "testing" || ENVIRONMENT == "production" {
	//}
	//启动
	err = app.Run()
	if err != nil {
		global.GlobalLogger.Error("app start error:%v\n", err.Error())
		return
	}
}
