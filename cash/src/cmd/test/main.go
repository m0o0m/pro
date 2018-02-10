package main

import (
	"fmt"
	"framework/msg_queue"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

var msgQueue msg_queue.MsgQueue

func main() {
	//agent()
	//mgo_test()
	//mgo_av()
	//fmt.Println(mgo_test())
	bt, err := msg_queue.ZipStr([]byte("sajlfjlajflajfljafljalfjalsjflajfljalfjlajflasjfafsf1"))
	fmt.Println("压缩前", []byte("sajlfjlajflajfljafljalfjalsjflajfljalfjlajflasjfafsf1"))
	fmt.Println("压缩后", bt, err)
	fmt.Println(msg_queue.UnZipToStr(bt))

	msgQueue = msg_queue.NewRedisMsgQueue(redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1",
		Password: "",
		DB:       0,
	}))
	msg := new(msg_queue.Message)
	msg.Key = "queue_test"
	msg.Data = msg_queue.ByteEncoder([]byte("test"))
	msgQueue.PutMsg(msg)

}

type BetRecordInfo struct {
	ID           int64   `xorm:"id" json:"-" bson:"id"`
	Line_id      string  `xorm:"line_id" json:"line_id"`
	At_id        int64   `xorm:"at_id" json:"at_id"`
	Ua_id        int64   `xorm:"'ua_id'" json:"ua_id"`
	Sh_id        int64   `xorm:"'sh_id'" json:"sh_id"`
	At_username  string  `xorm:"'at_username'" json:"at_username"`
	Uid          int64   `xorm:"'uid'" json:"uid"`
	Uname        string  `xorm:"'uname'" json:"uname"`
	Order_num    string  `xorm:"'order_num'" json:"order_num"`
	Bet          float64 `xorm:"'bet'" json:"bet"`
	Valid_bet    float64 `xorm:"'valid_bet'" json:"valid_bet"`
	Assets       float64 `xorm:"'assets'" json:"assets"`
	Fc_type      string  `xorm:"'fc_type'" json:"fc_type"`
	Odds         float64 `xorm:"'odds'" json:"odds"`
	Periods      int64   `xorm:"'periods'" json:"periods"`
	Win          float64 `xorm:"'win'" json:"win"`
	Result       float64 `xorm:"'result'" json:"result"`
	Handicap     int64   `xorm:"'handicap'" json:"handicap"`
	Addtime      int64   `xorm:"'addtime'" json:"addtime"`
	Addday       string  `xorm:"'addday'" json:"addday"`
	Updatetime   int64   `xorm:"'updatetime'" json:"updatetime"`
	Updateday    int64   `xorm:"'updateday'" json:"updateday"`
	Bet_info     string  `xorm:"'bet_info'" json:"bet_info"`
	Ptype        int64   `xorm:"'ptype'" json:"ptype"`
	Js           int64   `xorm:"'js'" json:"js"`
	Status       int64   `xorm:"'status'" json:"status"`
	Return_water float64 `xorm:"'return_water'" json:"return_water"`
	Bet_ip       string  `xorm:"'bet_ip'" json:"bet_ip"`
	Play_id      int64   `xorm:"'play_id'" json:"play_id"`
}

func (*BetRecordInfo) TableName() string {
	return "my_bet_record"
}

//SELECT * FROM `my_bet_record` ORDER BY `id` LIMIT 10000,500;
//插入
func mgo_test() error {
	engine, err := xorm.NewEngine("mysql", "lottery:tjVBd&RfWX0Y@(113.10.246.111:3306)/lottery?charset=utf8")
	if err != nil {
		return err
	}
	engine.ShowSQL(true)

	engine_local, err := xorm.NewEngine("mysql", "root:19861261@(127.0.0.1:3306)/lot?charset=utf8")
	if err != nil {
		return err
	}

	var lastId int64 = 0

	for {
		pEveryOne := make([]*BetRecordInfo, 0)
		err = engine.Where("id > " + strconv.FormatInt(lastId, 10)).Limit(2000).Asc("id").Find(&pEveryOne)
		if err != nil {
			return err
		}
		fmt.Println(len(pEveryOne), lastId)
		if len(pEveryOne) != 2000 {
			break
		}
		for _, v := range pEveryOne {
			lastId_ := v.ID
			if lastId_ > lastId {
				lastId = lastId_
			}
		}
		//写入数据库
		fmt.Println(engine_local.Insert(pEveryOne))
	}

	fmt.Println("---end")

	/*
		//数据库读取
		session, err := mgo.Dial("mongodb://127.0.0.1:27017/caipiaodb")
		if err != nil {
			return err
		}
		defer session.Close()
		for i := 1; i < 10000; i++ {
			// Optional. Switch the session to a monotonic behavior.
			session.SetMode(mgo.Monotonic, true)
			c := session.DB("").C("my_bet_record")
			//插入
			result := BetRecordInfo{}
			for _, v := range pEveryOne {
				err = c.Insert(v)
				if err != nil {
					return err
				}
			}
			err = c.Find(bson.M{"filename": "p_config.yaml"}).One(&result)
		}*/
	return err
}

//统计
func agent() {
	catIds := make(map[interface{}]interface{})
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/yapi")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//分类
	c := session.DB("").C("interface_cat")
	var result []bson.M
	err = c.Find(bson.M{"project_id": 53}).All(&result)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range result {
		catIds[v["_id"]] = v["name"]
		fmt.Println(v["name"])
	}

	c = session.DB("").C("interface")
	//var result []bson.M
	err = c.Find(bson.M{"project_id": 53}).Sort("catid").All(&result)
	if err != nil {
		fmt.Println(err)
	}
	cat_id := interface{}("12")
	for _, v := range result {
		tempid := v["catid"]
		if cat_id != tempid {
			fmt.Printf("//%s\n", catIds[v["catid"]])
			cat_id = tempid
		}
		fmt.Printf("e.%s(\"%s\", publicController.Todo) //%s\n", v["method"], v["path"], v["title"])
	}

	//fmt.Println(result, len(result))
	//fmt.Println(c.Count())
}

//统计
func mgo_av() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/config")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("").C("bet")

	var result []bson.M
	/*
		var pipe = []bson.M{
			bson.M{"$match": bson.M{"url": bson.M{"$ne": ""}}},
			bson.M{"$group": bson.M{
				"_id":   bson.M{"url": "$url"},
				"dups":  bson.M{"$addToSet": "$_id"},
				"count": bson.M{"$sum": 1},
			}},
			bson.M{"$match": bson.M{
				"count": bson.M{"$gt": 1},
			}},
		}*/
	var pipe = []bson.M{
		bson.M{"$group": bson.M{
			"_id":    "$platform",
			"betall": bson.M{"$sum": "$betall"},
			"betyx":  bson.M{"$sum": "$betyx"},
			"wins":   bson.M{"$sum": "$win"},
		}},
	}
	err = c.Pipe(pipe).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result, len(result))
	//fmt.Println(c.Count())
}
