package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	Handler  ProduceHandler
	producer sarama.AsyncProducer
}

func NewProducer(config *sarama.Config, addrs []string, handler ProduceHandler) *Producer {
	produce, err := sarama.NewAsyncProducer(addrs, config)
	if err != nil {
		panic("kafka producer start err")
	}
	p := &Producer{
		producer: produce,
		Handler:  handler,
	}
	go p.handleSuccess(produce)
	go p.handleError(produce)
	return p
}

func NewDefaultProducer(addrs []string, handler ProduceHandler) *Producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = time.Second * 10

	return NewProducer(config, addrs, handler)
}

func (p *Producer) handleSuccess(producer sarama.AsyncProducer) {
	var (
		pm *sarama.ProducerMessage
	)
	for {
		pm = <-producer.Successes()
		if pm != nil {
			if e := p.Handler.HandleSuccess(pm); e != nil {
				fmt.Errorf("%s\n", e.Error())
			}
		}
	}
}

func (p *Producer) handleError(producer sarama.AsyncProducer) {
	var (
		err *sarama.ProducerError
	)
	for {
		err = <-producer.Errors()
		if e := p.Handler.HandleError(err); e != nil {
			fmt.Errorf("%s\n", err.Error())
		}
	}
}

//PutMsg put msg to kafka
func (p *Producer) PutMsg(msg []byte, key, topic string, metadata interface{}) error {
	p.producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.ByteEncoder(msg), Metadata: metadata}
	return nil
}
