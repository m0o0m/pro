package influxdb

import (
	"encoding/json"
	"framework/influxdb/v2"
	"time"
)

type RequestTag struct {
	Name   string `json:"Name"`
	IP     string `json:"IP"`
	SiteID string `json:"SiteID"`
}

type RequestFields struct {
	ExecutionTime int64 `json:"ExecutionTime"`
	StartTime     int64 `json:"StartTime"`
	EndTime       int64 `json:"EndTime"`
}

type RequestPoint struct {
	Tag       RequestTag
	Fields    RequestFields
	TimePoint time.Time
}

func (r *RequestPoint) TableName() string {
	return "game_api_request"
}

func (r *RequestPoint) Insert(influxdb client.Client, db string) error {
	buf, err := json.Marshal(r.Tag)
	if err != nil {
		return err
	}
	tags_ := make(map[string]string)
	err = json.Unmarshal(buf, &tags_)
	if err != nil {
		return err
	}

	buf, err = json.Marshal(r.Fields)
	if err != nil {
		return err
	}
	fields_ := make(map[string]interface{})
	err = json.Unmarshal(buf, &fields_)
	if err != nil {
		return err
	}

	pt, err := client.NewPoint(r.TableName(), tags_, fields_, r.TimePoint)
	if err != nil {
		return err
	}
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})
	bp.AddPoint(pt)
	err = influxdb.Write(bp)
	if err != nil {
		return err
	}
	return nil
}
