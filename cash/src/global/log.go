package global

import (
	"config"
	"framework/exception"
	"framework/logger"
	"log"
	"os"
	"path"
)

const (
	sysLogName = "sys.log"
	MaxLogSize = 1024 * 1024 * 1024
)

var GlobalLogger *logger.Logger

//日志等级 "FNST", "FINE", "DEBG", "TRAC", "INFO", "WARN", "EROR", "CRIT"
func InitLog(cfg *config.Config, pathname string) {
	defer exception.Exception()
	//全局数据库，redis mogon
	//初始化
	switch cfg.Log.Log_type {
	case "file":
		err := os.MkdirAll(cfg.Log.Path+"/"+pathname, 0777)
		if err != nil {
			log.Println(cfg.Log.Path+"Folder create failure:", err.Error())
		}
		fileLog := logger.NewFileLogWriter(path.Join(cfg.Log.Path+"/"+pathname, sysLogName), true)
		fileLog = fileLog.SetRotateSize(MaxLogSize)

		GlobalLogger = &logger.Logger{
			"fileOut": &logger.Filter{logger.StringToLevel(cfg.Log.Level), fileLog},
		}
	default:
		GlobalLogger = &logger.Logger{
			"stdout": &logger.Filter{logger.DEBUG, logger.NewConsoleLogWriter()},
		}
	}
}
