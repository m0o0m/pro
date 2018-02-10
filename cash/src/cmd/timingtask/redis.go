//[主函数] redis定时任务,处理登录状态
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"config"
	"global"
	"models/function"

	"github.com/go-redis/redis"
)

//配置文件路径
var configFile *string = flag.String("config", "../bin/etc/conf.yaml", "kingshard config file")

var (
	DefaultRedis   int = 0 //默认redis库
	TimeTickerRun  int     //定时任务多长时间执行一次，以秒为单位
	Goroutinenum   int     //可以开的协程数量
	RedisClientMap map[int]*redis.Client
)

//主函数
func main() {
	Cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		fmt.Printf("parse config file error:%v\n", err.Error())
		return
	}
	//初始化数据库
	err = global.InitMysql(Cfg.Mysqls)
	if err != nil {
		//数据库连接错误
		log.Fatalf("InitDb error : %v", err)
	}
	err = InitRedis(Cfg.Redis)
	if err != nil {
		//redis连接失败
		log.Fatalf("Redis connection failed : %v", err)
	}
	//定时任务
	go Crons()
	select {}
}

//定时任务
func Crons() {
	ticker := time.NewTicker(time.Second * time.Duration(TimeTickerRun))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for i := 0; i < Goroutinenum; i++ {
				go Runs()
			}
		}
	}
}

//初始化链接
func InitRedis(configRedis []config.RedisConfig) error {
	var redisConfig config.RedisConfig
	RedisClientMap = make(map[int]*redis.Client)
	length := len(configRedis)
	for i := 0; i < length; i++ {
		if configRedis[i] == redisConfig {
			continue
		}
		host := configRedis[i].Host
		password := configRedis[i].Password
		name := configRedis[i].Name
		TimeTickerRun = configRedis[i].Runtime
		Goroutinenum = configRedis[i].Goroutinenum
		RedisClientMap[name] = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       name})
		_, err := RedisClientMap[name].Ping().Result()
		if err != nil {
			return err
		}
	}
	return nil
}

//执行判定流程
func Runs() {
	//弹出key
	keys, err := RedisClientMap[DefaultRedis].LPop("is_login").Result()
	if keys == "" {
		return
	}

	//获取序列化的数据
	result, err := RedisClientMap[DefaultRedis].Get(keys).Result()
	if err != nil {
		return
	}
	var s global.RedisStruct
	json.Unmarshal([]byte(result), &s)

	//过期
	if s.ExpirTime < time.Now().Unix() {
		//根据agencyId更新数据库在线状态,
		UpDataLoginStatus(s.Id)
		//清除token
		err = RedisClientMap[DefaultRedis].Del(keys).Err()
		if err != nil {
			return
		}
	} else {
		//不过期
		err = RedisClientMap[DefaultRedis].RPush("is_login", keys).Err()
		if err != nil {
			return
		}
	}
}

//更新
func UpDataLoginStatus(agencyId int64) {
	_, err := function.ChangeLoginStatus(agencyId)
	if err != nil {
		return
	}
}
