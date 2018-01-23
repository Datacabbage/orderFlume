package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

//处理生成成功数据
type ProduceSuccessHandle interface {
	HandleSuccess(*sarama.ProducerMessage) error
}

//处理生产失败数据
type ProduceErrorHandle interface {
	HandleError(*sarama.ProducerError) error
}

type ProduceHandler interface {
	ProduceSuccessHandle
	ProduceErrorHandle
}

type ConsumeHandler interface {
	MessageHandle(consumer *cluster.Consumer)
}

type DefaultProduceHandler struct {
}

func NewDefaultProduceHandler() *DefaultProduceHandler {
	return new(DefaultProduceHandler)
}

func (handler *DefaultProduceHandler) HandleSuccess(producer *sarama.ProducerMessage) error {
	return nil
}

func (handler *DefaultProduceHandler) HandleError(producer *sarama.ProducerError) error {
	return nil
}
