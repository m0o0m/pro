//[主函数] 前台接口
package main

import (
	_ "cmd/godaemon"
	"config"
	"framework"
	"framework/render"
	"global"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
		configFile := filepath.Join(filepath.Dir(path), "etc/front.yaml")
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
		cfg, err = config.ParseFromMongoDB(mongo_url, "/conf/p_config/p_front.yaml")
		if err != nil {
			log.Fatalf("parse config file error:%v\n", err.Error())
			return
		}
	}

	//全局初始化
	global.InitLog(cfg, "front")
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

	if cfg.TemplateConfig.MongoCache == "on" {
		//初始化mongo连接
		err = global.InitMgoPool(&cfg.MongoDb)
		if err != nil {
			//mongo连接失败
			global.GlobalLogger.Error("Mongo connection failed:%d\n", err.Error())
			return
		}
		// TODO 将html件全部读取到mongo中
		global.GlobalLogger.Debug(cfg.TemplateConfig.SourcePath)
		htmlNames := getHtmlNames(cfg.TemplateConfig.SourcePath)
		var htmls []global.Html
		for _, htmlName := range htmlNames {
			bs, err := ioutil.ReadFile(htmlName)
			if err != nil {
				global.GlobalLogger.Error("read template err")
				panic("read template err")
			}
			html := global.Html{
				Name: htmlName,
				Body: bs,
			}
			htmls = append(htmls, html)
		}
		err = global.SaveHTMLS(htmls)
		if err != nil {
			global.GlobalLogger.Error("err:%s", err.Error())
			panic("insert template mongo err")
		}
	}

	//启动web服务
	app, err := framework.NewFront(cfg)
	if err != nil {
		//错误日志
		global.GlobalLogger.Error("NewApp error:%v\n", err.Error())
		return
	}
	//初始化render
	render.InitRootPath(&cfg.TemplateConfig)
	//初始化静态目录
	if ENVIRONMENT == "development" {
		app.WebServer.Static("/template", render.ViewPath)
	} else if ENVIRONMENT == "testing" || ENVIRONMENT == "production" {
	}

	//启动
	err = app.Run()
	if err != nil {
		global.GlobalLogger.Error("app start error:%v\n", err.Error())
		return
	}
}

func getHtmlNames(src string) (srcFileNames []string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			if strings.Count(src, "/") >= 20 {
				global.GlobalLogger.Error("The folder level is too deep")
				return
			}
			childSrcFileNames := getHtmlNames(src + "/" + file.Name())
			srcFileNames = append(srcFileNames, childSrcFileNames...)
		} else if strings.HasSuffix(file.Name(), ".html") {
			srcFileNames = append(srcFileNames, src+"/"+file.Name())
		}
	}
	return
}
