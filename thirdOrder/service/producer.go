package service

import (
	"order_flume/lib/kafka"

	"strings"
)

var produce *kafka.Producer

func ProducerInstance() *kafka.Producer {
	return produce
}

func InitProduce(kafkaAddr string) *kafka.Producer {
	if produce == nil {
		addrs := strings.Split(kafkaAddr, ",")
		produce = kafka.NewDefaultProducer(addrs, &OrderProduceHandler{})
	}
	return produce
}
