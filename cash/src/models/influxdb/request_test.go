package influxdb

import (
	"framework/influxdb/v2"
	"testing"
	"time"
)

var (
	host = "http://localhost:8086"
)

func InitDB(t *testing.T) client.Client { /*
		c, err := client.NewUDPClient(client.UDPConfig{
			Addr:        host,
			PayloadSize: 1000000,
		})
	*/
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: host,
	})
	if err != nil {
		t.Fatal("InitDB udp error %s %s", host, err.Error())
	}
	return c
}

func TestInsert(t *testing.T) {
	c := InitDB(t)
	rp := RequestPoint{}
	rp.Tag = RequestTag{}
	rp.Tag.IP = "127.0.0.1"
	rp.Tag.Name = "get"
	rp.Tag.SiteID = "aaa"
	rp.Fields = RequestFields{}
	rp.Fields.ExecutionTime = 30
	rp.Fields.StartTime = 1231231312312313
	rp.Fields.EndTime = 1231231312312313
	rp.TimePoint = time.Now()
	err := rp.Insert(c, "test")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(err)
	}
}
