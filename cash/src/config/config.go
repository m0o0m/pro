//[配置] 配置文件对应结构体
package config

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v2"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

//整个config文件对应的结构
type Config struct {
	Addr          string `yaml:"addr"`
	Language      string `yaml:"language"`
	Wesocketport  string `yaml:"websocketport"`
	RedisThirdSet string `yaml:"redisthirdset"` //同步第三方缓存到redis的key值

	Log            LogConfig      `yaml:"log"`
	Influx         InfluxConfig   `yaml:"influxdb"`
	Redis          []RedisConfig  `yaml:"redis"`
	Mysqls         []MysqlConfig  `yaml:"mysql"`
	Video          Video          `yaml:"video"`
	Version        VersionConfig  `yaml:"version"`
	Third          ThirdInterface `yaml:"third"`    //第三方接口请求路由
	TemplateConfig TemplateConfig `yaml:"template"` //模板配置
	ExePath        string         `yaml:"-"`
	MongoDb        MongoDbConfig  `yaml:"mongo"`
}

//版本配置
type VersionConfig struct {
	Test    string `yaml:"test"`
	Product string `yaml:"product"`
}

//log配置
type LogConfig struct {
	Level    string `yaml:"level"`
	Log_type string `yaml:"type"`
	Path     string `yaml:"path"`
}

//redis
type InfluxConfig struct {
	Name        string `yaml:"name"`
	Host        string `yaml:"host"`
	DbName      string `yaml:"dbname"`
	PayloadSize int    `yaml:"payloadsize"`
}

//redis
type RedisConfig struct {
	Name             string `yaml:"name"`
	Host             string `yaml:"host"`
	Password         string `yaml:"password"`
	DB               int    `yaml:"db"`
	Expirtime        int    `yaml:"expirtime"`          //token过期时间,外面配置文件过期时间以分钟为单位
	Runtime          int    `yaml:"runtime"`            //定时任务多长时间执行一次，以秒为单位
	Goroutinenum     int    `yaml:"goroutimenum"`       //可以开的协程数量
	ExpUpdateRanking int    `yaml:"exp_update_ranking"` //可以开的协程数量
}

//Mysql
type MysqlConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	ShowSql  bool   `yaml:"showsql"`
	Timeout  string `yaml:"timeout"`
}

//mongodb
type MongoDbConfig struct {
	Host     string `yaml:"host"`
	Account  string `yaml:"account" `
	Password string `yaml:"password" `
}

//第三方平台请求接口
type ThirdInterface struct {
	PaidType     string `yaml:"paidtype"`     //获取支付类型api
	GetThirdApi  string `yaml:"getthirdapi"`  //获取第三方配置api
	NewSetup     string `yaml:"newsetup"`     //增加或者修改三方设置api
	GetBank      string `yaml:"getbank"`      //获取某支付下面支持的银行卡的api
	GetSetup     string `yaml:"getsetup"`     //获取商户配置api
	ClientUserId string `yaml:"clientuserid"` //客户id
	ClientName   string `yaml:"clientname"`   //客户名称
	ClientSecret string `yaml:"clientsecret"` //客户授权证书
}

//第三方视讯
type Video struct {
	ApiUrl string `yaml:"api_url"`
	Md5Key string `yaml:"md5_key"`
	DesKey string `yaml:"des_key"`
}

//模板配置
type TemplateConfig struct {
	CacheSize   int    `yaml:"cache_size"`   //缓存页面数量
	CdnUrl      string `yaml:"cdn_url"`      //资源cdn地址
	CacheSwitch string `yaml:"cache_switch"` //缓存开关
	SourcePath  string `yaml:"source_path"`  //资源文件根路径
	MongoCache  string `yaml:"mongo_cache"`  //是否使用Mongo来缓存html
}

//从数据解析配置
func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	path, _ := GetExecPath()
	cfg.ExePath = filepath.Dir(path)
	cfg.TemplateConfig.SourcePath = filepath.Join(cfg.ExePath, cfg.TemplateConfig.SourcePath)

	if !strings.HasPrefix(cfg.Log.Path, "/") {
		cfg.Log.Path = filepath.Join(cfg.ExePath, cfg.Log.Path)
	}
	return &cfg, nil
}

//从文件解析配置
func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return ParseConfigData(data)
}

//写入文件
func WriteConfigFile(cfg *Config, filename string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	execPath, err := os.Getwd()
	if err != nil {
		return err
	}

	configPath := execPath + "/etc/" + filename
	err = ioutil.WriteFile(configPath, data, 0755)
	if err != nil {
		return err
	}

	return nil
}

//mongo数据存储结构
type ConfigMgo struct {
	Name     string `json:"name" bson:"name"`
	FullName string `json:"-" bson:"full_name"`
	FullPath string `json:"-" bson:"full_path"`
	Rights   string `json:"rights" bson:"rights"`
	Size     string `json:"size" bson:"size"`
	Ext      string `json:"-" bson:"ext"`
	Date     string `json:"date" bson:"date"`
	FileType string `json:"type" bson:"file_type"`
	Content  []byte `json:"-" bson:"content"`
	Status   int8   `json:"status" bson:"status"`
}

//将配置写入mongodb
func WriteToMongoDB(cfg *Config, source string, filename string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	session, err := mgo.Dial(source)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("").C("p_config")
	//_, err = c.RemoveAll(bson.M{"filename": "p_config.yaml"})
	//覆盖
	cm := ConfigMgo{}
	cm.Name = filename
	cm.FullName = "/conf/p_config/" + filename
	cm.FullPath = "/conf/p_config"
	cm.Rights = "drwxr-xr-x"
	cm.Size = strconv.Itoa(len(data))
	cm.Ext = ""
	cm.Date = time.Now().String()
	cm.FileType = "dir"
	cm.Status = 1
	_, err = c.Upsert(bson.M{"full_name": filename}, &cm)
	if err != nil {
		return err
	}
	return nil
}

//从mongodb中解析配置
func ParseFromMongoDB(source string, filename string) (*Config, error) {
	//source
	/*
		10.10.10.23 27017
		10.10.10.186 27017
		mongodb://10.10.10.23:27017,10.10.10.186:27017/config
	*/
	session, err := mgo.Dial(source)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("").C("p_config")

	result := ConfigMgo{}
	err = c.Find(bson.M{"full_name": filename}).One(&result)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(result.Content)
}

func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return p, nil
}
