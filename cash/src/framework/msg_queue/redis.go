package msg_queue

import (
	"github.com/go-redis/redis"
)

type RedisMsgQueue struct {
	Client *redis.Client
}

func NewRedisMsgQueue(client *redis.Client) *RedisMsgQueue {
	return &RedisMsgQueue{Client: client}
}

func (rmq *RedisMsgQueue) PutMsg(msg *Message) error {
	rmq.Client.Ping()
	data, err := msg.Data.Encode()
	if err != nil {
		return err
	}
	ic := rmq.Client.LPush(msg.Key, data)
	return ic.Err()
}

/*
func (rmq *RedisMsgQueue) GetMsg(key string) (*Message, error) {
	data, err := rmq.Client.RPop(key).Bytes()
	msg := new(Message)
	msg.Key = key
	msg.Data = ByteEncoder(data)
	return msg, err
}
*/
