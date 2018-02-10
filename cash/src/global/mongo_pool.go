package global

import (
	"config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	TEMPLATE = "template"
	HTML     = "html"
)

var (
	host       string
	mgoSession *mgo.Session
)

func InitMgoPool(cfg *config.MongoDbConfig) error {
	var err error
	mgoSession, err = mgo.Dial(cfg.Host)
	if err != nil {
		return err
	}
	if cfg.Account != "" {
		GlobalLogger.Debug("connect mongo:%s-%s", cfg.Account, cfg.Password)
		err = mgoSession.DB(TEMPLATE).Login(cfg.Account, cfg.Password)
	}
	return err
}

//得到session,如果存在,则拷贝一份
func getMongo() *mgo.Session {
	return mgoSession.Clone()
}

type Html struct {
	Name string `bson:"_id"`
	Body []byte `bson:"body"`
}

//往mongo中写入多条html文件
func SaveHTMLS(htmls []Html) error {
	session := getMongo()
	defer session.Close()
	c := session.DB(TEMPLATE).C(HTML)
	c.DropCollection()
	ins := make([]interface{}, len(htmls))
	for i := range htmls {
		ins[i] = htmls[i]
	}
	return c.Insert(ins...)
}

//往mongo中写入html文件
func SaveHTML(name string, body []byte) error {
	session := getMongo()
	defer session.Close()
	c := session.DB(TEMPLATE).C(HTML)
	return c.Insert(Html{Name: name, Body: body})
}

//从mongo中读取html文件
func ReadHTML(name string) ([]byte, error) {
	session := getMongo()
	defer session.Close()
	c := session.DB(TEMPLATE).C(HTML)
	var html Html
	err := c.Find(bson.M{"_id": name}).One(&html)
	return html.Body, err
}
