package convert

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GBKtoUTF8(gbkData []byte) (utf8Data []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(gbkData), simplifiedchinese.GBK.NewDecoder())
	utf8Data, err = ioutil.ReadAll(reader)
	return
}

func UTF8toGBK(utf8Data []byte) (gbkData []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(utf8Data), simplifiedchinese.GBK.NewEncoder())
	gbkData, err = ioutil.ReadAll(reader)
	return
}
