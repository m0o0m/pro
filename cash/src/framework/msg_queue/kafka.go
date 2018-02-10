package msg_queue

import (
	"github.com/Shopify/sarama"
	"strings"
)

type KafkaMsgQueue struct {
	Producer sarama.AsyncProducer
}

func NewKafkaMsgQueue(brockerList string) (*KafkaMsgQueue, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll //sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	kmq := new(KafkaMsgQueue)
	ap, err := sarama.NewAsyncProducer(strings.Split(brockerList, ","), config)
	kmq.Producer = ap
	return kmq, err
}

func (kmq *KafkaMsgQueue) PutMsg(msg *Message) error {
	kmq.Producer.Input() <- &sarama.ProducerMessage{
		Topic: msg.Key,
		Key:   sarama.ByteEncoder([]byte(msg.Key)),
		Value: msg.Data,
	}

	return nil
}

func (kmq *KafkaMsgQueue) Close() {
	kmq.Producer.AsyncClose()
}
