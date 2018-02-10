package global

//数据库初始化
import (
	"config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	EngineMap map[string]*xorm.Engine
)

func InitMysql(configSql []config.MysqlConfig) error {
	var err error
	EngineMap = make(map[string]*xorm.Engine)
	length := len(configSql)
	for i := 0; i < length; i++ {
		node := configSql[i].Name
		if len(node) == 0 {
			return errInvalidMysqlNode
		}
		host := configSql[i].Host
		dbName := configSql[i].DbName
		password := configSql[i].Password
		username := configSql[i].Username
		timeout := configSql[i].Timeout
		data := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&timeout=%s", username, password, host, dbName, timeout)
		EngineMap[node], err = xorm.NewEngine("mysql", data)
		if err != nil {
			return err
		}
		err = EngineMap[node].Ping()
		if err != nil {
			return err
		}
		EngineMap[node].ShowSQL(configSql[i].ShowSql)
	}
	return nil
}

//获取默认xorm
func GetXorm() *xorm.Engine {
	return EngineMap[DEFAULT]
}

//获取xormSession
func GetXormSession() *xorm.Session {
	return EngineMap[DEFAULT].NewSession()
}

//获取视讯库
func GetVideo() *xorm.Engine {
	return EngineMap[VIDEO]
}
