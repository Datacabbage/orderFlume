package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

//OperatorKafkaMsg  consumer and operator kafka msg
type OperatorKafkaMsg func(*sarama.ConsumerMessage) error

//Consumer  kido trace kafka Consumer
type Consumer struct {
	addrs           []string
	groupID         string
	topicList       []string
	ReadPerformance bool
	config          *cluster.Config
	consumer        *cluster.Consumer
	handler         ConsumeHandler
}

//NewConsumer 初始化消费者
func NewConsumer(groupid string, addrs, topicList []string, config *cluster.Config, handler ConsumeHandler) *Consumer {
	return &Consumer{
		addrs:     addrs,
		groupID:   groupid,
		topicList: topicList,
		config:    config,
		handler:   handler,
	}
}

//NewConsumerDefault defualt consumer
func NewConsumerDefault(groupid string, addrs, topicList []string, handler ConsumeHandler) *Consumer {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest //初始从最新的offset开始
	return NewConsumer(groupid, addrs, topicList, config, handler)
}

func (c *Consumer) ClusterConsumer() *cluster.Consumer {
	return c.consumer
}

//Start kafka consumer
func (c *Consumer) Start() {
	var err error
	c.consumer, err = cluster.NewConsumer(c.addrs, c.groupID, c.topicList, c.config)
	if err != nil {
		fmt.Errorf("Failed open consumer: %v", err)
		return
	}
	go func() {
		for err := range c.consumer.Errors() {
			fmt.Errorf("Error: %s\n", err.Error())
		}
	}()
	go func() {
		for note := range c.consumer.Notifications() {
			fmt.Printf("Rebalanced: %v", note)
		}
	}()

	go c.handler.MessageHandle(c.consumer)
}
