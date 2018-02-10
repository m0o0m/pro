package global

import (
	"config"
	"framework/influxdb/v2"
	"github.com/go-redis/redis"
	"time"
)

type RedisExp struct {
	Pc      time.Duration `json:"pc"`
	Wap     time.Duration `json:"wap"`
	Android time.Duration `json:"android"`
	Ios     time.Duration `json:"ios"`
}

//agency 存储进入redis的数据的struct
type RedisStruct struct {
	Id          int64    `json:"id"`            //agency id
	Account     string   `json:"account"`       //账号
	SiteId      string   `json:"site_id"`       //站点id
	SiteIndexId string   `json:"site_index_id"` //站点前台id
	Username    string   `json:"username"`      //用户名
	Level       int8     `json:"level"`         //等级
	IsSub       int8     `json:"is_sub"`        //是否子账号
	RoleId      int64    `json:"role_id"`       //角色id
	ExpirTime   int64    `json:"expir_time"`    //过期时间
	Type        string   `json:"type"`          //类型[agency]
	AccessSID   []string `json:"access_sid"`    //可控制站点
}

//admin　登录的时候存储进去的数据
type AdminRedisStruct struct {
	Id        int64  `json:"id"`         //id
	Account   string `json:"account"`    // 登录账号
	Status    int8   `json:"status"`     // 账号状态
	Type      string `json:"type"`       //类型
	ExpirTime int64  `json:"expir_time"` //过期时间
	RoleId    int64  `json:"role_id"`    // 角色
}

//member 登录的时候存储的数据
type MemberRedisToken struct {
	Id        int64  `json:"id"`         //会员id
	Account   string `json:"account"`    //账号
	Status    int8   `json:"status"`     //状态
	ExpirTime int64  `json:"expir_time"` //token过期时间
	Site      string `json:"site"`       //会员所属站点
	SiteIndex string `json:"site_index"` //会员所属前台站点
	LevelId   string `xorm:"level_id"`   //会员所属层级
	Type      string `json:"type"`       //会员登录的客户端类型
}

var (
	DefaultRedis     int           = 0
	AgencyRedisExp   time.Duration //token过期时间,外面配置文件过期时间以分钟为单位
	UpdateRankingExp time.Duration //排行榜更新时间
	Goroutinenum     int           //可以开的协程数量
	RedisClientMap   map[string]*redis.Client
	Influxdb         client.Client //统计使用的时序数据库
	DefaultRedisExp  = RedisExp{
		Pc:      time.Hour * 24,
		Wap:     time.Hour * 24,
		Android: time.Hour * 24,
		Ios:     time.Hour * 24,
	}
)

func InitRedis(configRedis []config.RedisConfig) error {
	RedisClientMap = make(map[string]*redis.Client)
	length := len(configRedis)
	for i := 0; i < length; i++ {
		node := configRedis[i].Name
		if len(node) == 0 {
			return errInvalidMysqlNode
		}
		host := configRedis[i].Host
		password := configRedis[i].Password
		db := configRedis[i].DB

		AgencyRedisExp = time.Minute * time.Duration(configRedis[i].Expirtime)
		UpdateRankingExp = time.Minute * time.Duration(configRedis[i].ExpUpdateRanking)
		Goroutinenum = configRedis[i].Goroutinenum

		RedisClientMap[node] = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		})

		_, err := RedisClientMap[node].Ping().Result()
		if err != nil {
			return err
		}
	}
	return nil
}

//获取默认redis
func GetRedis() *redis.Client {
	return RedisClientMap[DEFAULT]
}

func InitInfluxdb(influxConfig config.InfluxConfig) error {
	if Influxdb == nil {
		/*
			c, err := client.NewUDPClient(client.UDPConfig{
				Addr:        influxConfig.Host,
				PayloadSize: influxConfig.PayloadSize,
			})*/
		c, err := client.NewHTTPClient(client.HTTPConfig{
			Addr: influxConfig.Host,
		})
		if err != nil {
			return err
		}
		Influxdb = c
	}
	return nil
}
