package msg_queue

/**
//使用
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

*/

import (
	"bytes"
	"compress/zlib"
	"io"
)

type MsgQueue interface {
	PutMsg(msg *Message) error
	//GetMsg() (*Message, error)
}

type Message struct {
	Key  string
	Data Encoder
}

type Encoder interface {
	Encode() ([]byte, error)
	Length() int
}

type StringEncoder string

func (s StringEncoder) Encode() ([]byte, error) {
	return ZipStr([]byte(s))
}

func (s StringEncoder) Length() int {
	return len(s)
}

type ByteEncoder []byte

func (b ByteEncoder) Encode() ([]byte, error) {
	return ZipStr(b)
}

func (b ByteEncoder) Length() int {
	return len(b)
}

func ZipStr(str []byte) ([]byte, error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err := w.Write(str)
	if err != nil {
		return str, err
	}
	err = w.Close()
	if err != nil {
		return str, err
	}
	return in.Bytes(), nil
}

func UnZipToStr(data []byte) (string, error) {
	var in bytes.Buffer
	_, err := in.Write(data)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	r, err := zlib.NewReader(&in)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
