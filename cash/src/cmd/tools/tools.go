//[主函数] 更新mongodb配置
package main

import (
	"flag"
	"fmt"
	"os"
	//"strconv"
	//"strings"

	"config"
	//"github.com/go-xorm/xorm"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

//const mongodb_master_url = "mongodb://master.mongo.dev.com:27017/config"

func main() {
	var configType *string = flag.String("type", "conf", "config type: conf|admin|front|wap")
	var configFile *string = flag.String("config", "../bin/etc/conf.yaml", "config file")
	flag.Parse()

	mongodb_master_url := os.Getenv("MONGODB_MASTER_URL")
	if mongodb_master_url == "" {
		mongodb_master_url = "mongodb://master.mongo.dev.com:27017/config"
	}

	fmt.Printf("config file type: %v, path: %v\n", *configType, *configFile)
	cfg, err := config.ParseConfigFile(*configFile)

	if err != nil {
		fmt.Printf("parse config file error:%v\n", err.Error())
		return
	}

	fmt.Println("Addr: ", cfg.Addr)
	fmt.Println("Language: ", cfg.Language)
	fmt.Println("Wesocketport", cfg.Wesocketport)
	fmt.Println("Log Configuration: ")
	fmt.Println("	log level: ", cfg.Log.Level)
	fmt.Println("	log type: ", cfg.Log.Log_type)
	fmt.Println("	log path: ", cfg.Log.Path)
	fmt.Println("Test Version: ", cfg.Version.Test)
	fmt.Println("Production Version: ", cfg.Version.Product)

	fmt.Println("Redis Configuration: ", cfg.Redis)
	fmt.Println("Mysql Configuration: ", cfg.Mysqls)
	fmt.Println("MongoDB Configuration: ", cfg.MongoDb.Host)

	filename := "p_config.yaml"
	switch *configType {
	case "conf":
		filename = "p_config.yaml"
	case "admin":
		filename = "p_admin.yaml"
	case "front":
		filename = "p_front.yaml"
	case "wap":
		filename = "p_wap.yaml"
	default:
		fmt.Printf("Config type invalid, %v", configType)
		return
	}

	fmt.Println("Master MongoDB URL: ", mongodb_master_url)
	fmt.Println("File name in mongodb is: ", filename)
	//fmt.Println("Configuration: ", *cfg)
	// 写更新后的配置文件到本地
	config.WriteConfigFile(cfg, filename)

	err = config.WriteToMongoDB(cfg, mongodb_master_url, filename)
	if err != nil {
		fmt.Printf("WriteToMongoDB error: %v-%v\n", filename, err.Error())
		return
	}
	cfg2, err := config.ParseFromMongoDB(mongodb_master_url, filename)
	if err != nil {
		fmt.Printf("ParseFromMongoDB error: %v-%v\n", filename, err.Error())
		return
	}
	fmt.Printf("Read Front from Mongodb success: %v-%v\n", filename, cfg2)
}

//type BetRecordInfo struct {
//	ID             int64   `xorm:"'id'" json:"-" bson:"_id"`
//	OrderId        string  `xorm:"pk 'order_id'" json:"order_id"`
//	Platform       string  `xorm:"pk 'platform'" json:"platform"`
//	UserName       string  `xorm:"'username'" json:"username"`
//	GUserName      string  `xorm:"'g_username'" json:"g_username"`
//	Currency       string  `xorm:"'currency'" json:"currency"`
//	SiteId         string  `xorm:"'site_id'" json:"site_id"`
//	IndexId        string  `xorm:"'index_id'" json:"index_id"`
//	AgentId        int64   `xorm:"'agent_id'" json:"agent_id"`
//	UaId           int64   `xorm:"'ua_id'" json:"ua_id"`
//	ShId           int64   `xorm:"'sh_id'" json:"sh_id"`
//	BetAll         float64 `xorm:"'bet_all'" json:"bet_all"`
//	BetYx          float64 `xorm:"'bet_yx'" json:"bet_yx"`
//	Win            float64 `xorm:"'win'" json:"win"`
//	OtherBet       float64 `xorm:"'other_bet'" json:"other_bet"` //彩金
//	OtherWin       float64 `xorm:"'other_win'" json:"other_win"` //彩金
//	BetTimeLine    int64   `xorm:"'bet_timeline'" json:"bet_timeline"`
//	BetTime        string  `xorm:"'bet_time'" json:"bet_time"`
//	SettleTimeLine int64   `xorm:"'settle_timeline'" json:"settle_timeline"`
//	SettleTime     string  `xorm:"'settle_time'" json:"settle_time"`
//	GameType       int     `xorm:"'game_type'" json:"game_type"`     //游戏类型，1 电子，2 视讯，3 捕鱼 4 彩票 5 体育 6 红包 7 小费
//	GameName       string  `xorm:"'game_name'" json:"game_name"`     //游戏名字
//	GameResult     string  `xorm:"'game_result'" json:"game_result"` //游戏结构
//	UpdateTime     int64   `xorm:"'update_time'" json:"update_time"`
//
//	Extra  string `xorm:"'extra'" json:"extra"`   //附加信息
//	Status int8   `xorm:"'status'" json:"status"` //状态 1正常的 0 是注销的注单
//}
//
//func (*BetRecordInfo) TableName() string {
//	return "bet_record_info"
//}
//
////插入
//func mgo_test() error {
//	engine, err := xorm.NewEngine("mysql", "root:19861261@/shixun?charset=utf8")
//	if err != nil {
//		return err
//	}
//	pEveryOne := make([]*BetRecordInfo, 0)
//	err = engine.Find(&pEveryOne)
//	if err != nil {
//		return err
//	}
//	//数据库读取
//	session, err := mgo.Dial("mongodb://127.0.0.1:27017/config")
//	if err != nil {
//		return err
//	}
//	defer session.Close()
//
//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//
//	c := session.DB("").C("bet")
//	//插入
//	result := BetRecordInfo{}
//	for _, v := range pEveryOne {
//		err = c.Insert(v)
//		if err != nil {
//			return err
//		}
//	}
//
//	err = c.Find(bson.M{"filename": "p_config.yaml"}).One(&result)
//
//	return err
//}
//
////统计
//func mgo_av() {
//	session, err := mgo.Dial("mongodb://127.0.0.1:27017/config")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//
//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//
//	c := session.DB("").C("bet")
//
//	var result []bson.M
//	/*
//		var pipe = []bson.M{
//			bson.M{"$match": bson.M{"url": bson.M{"$ne": ""}}},
//			bson.M{"$group": bson.M{
//				"_id":   bson.M{"url": "$url"},
//				"dups":  bson.M{"$addToSet": "$_id"},
//				"count": bson.M{"$sum": 1},
//			}},
//			bson.M{"$match": bson.M{
//				"count": bson.M{"$gt": 1},
//			}},
//		}*/
//	var pipe = []bson.M{
//		bson.M{"$group": bson.M{
//			"_id":    "$platform",
//			"betall": bson.M{"$sum": "$betall"},
//			"betyx":  bson.M{"$sum": "$betyx"},
//			"wins":   bson.M{"$sum": "$win"},
//		}},
//	}
//	err = c.Pipe(pipe).All(&result)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(result, len(result))
//	//fmt.Println(c.Count())
//}
