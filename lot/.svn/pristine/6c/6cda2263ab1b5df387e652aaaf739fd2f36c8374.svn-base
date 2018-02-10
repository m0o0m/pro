package lutils

import (
	"net/url"
	"reflect"
	"fmt"
	"errors"
)

func GetValuesGET(st interface{}, args ...string) (string, error) {
	var tagKey = "send"
	if len(args) > 0 {
		tagKey = args[0]
	}
	val := reflect.ValueOf(st)
	ind := reflect.Indirect(val)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("The struct paramer must be use ptr"))
	}
	ty := ind.Type()
	if ty.NumField() > 0 {
		var values string
		for i := 0; i < ty.NumField(); i++ {
			name := ty.Field(i).Tag.Get(tagKey)
			value := ind.Field(i).String()
			if name != "" && value != "" {
				if values != "" {
					values += "&"
				}
				values += name + "=" + value
			}
		}
		if len(values) > 0 {
			return values, nil
		}
	}
	return "", errors.New("Not found field in struct")
}

func GetValuesPOST(st interface{}, args ...string) (url.Values, error) {
	var tagKey = "send"
	if len(args) > 0 {
		tagKey = args[0]
	}
	val := reflect.ValueOf(st)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("The struct paramer must be use ptr"))
	}
	ind := reflect.Indirect(val)

	ty := ind.Type()
	if ty.NumField() > 0 {
		values := url.Values{}
		for i := 0; i < ty.NumField(); i++ {
			name := ty.Field(i).Tag.Get(tagKey)
			value := ind.Field(i).String()
			if name != "" && value != "" {
				values.Add(name, value)
			}
		}
		if len(values) > 0 {
			return values, nil
		}
	}
	return nil, errors.New("Not found field in struct")
}
