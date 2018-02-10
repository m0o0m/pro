package lutils

import (
	"bytes"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//编码格式互转

func GBKDecode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
func GBKDecode_(str string) (string, error) {
	s := []byte(str)
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func GBKEncode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
func GBKEncode_(str string) (string, error) {
	s := []byte(str)
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return "", e
	}
	return string(d), nil
}
