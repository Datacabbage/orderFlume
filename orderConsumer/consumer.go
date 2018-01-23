package main

import (
	"order_flume/lib/beelog"
	"order_flume/lib/kafka"
	"order_flume/orderConsumer/conf"
	"order_flume/protobuf"
	"order_flume/services"
	"order_flume/utils"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/protobuf/proto"
)

type OrderServer struct {
	KafkaAddrs     []string
	KafkaGroupID   string
	KafkatopicList []string
}

type OrderConsumeHandle struct {
	BatchSize    int64
	BatchTimeout int64
}

func (handle *OrderConsumeHandle) PutMsg(msg *protobuf.OrderList) error {
	if msg == nil {
		return utils.ErrorParameter
	}
	if msg.MsgType == 0 {
		services.AddMeituanOrder(msg)
	} else if msg.MsgType == 1 {
		services.AddXiechenOrder(msg)
	} else if msg.MsgType == 2 {
		services.AddMaoyanOrder(msg)
	} else {
		services.AddMeituanWaiMaiOrder(msg)
	}
	services.SendToYM(msg, conf.Conf.YMUrl)
	return nil
}

func checkTopic(topicList []string, topic string) bool {
	for _, t := range topicList {
		if t == topic {
			return true
		}
	}

	return false
}

func (handle *OrderConsumeHandle) MessageHandle(consumer *cluster.Consumer) {
	for msg := range consumer.Messages() {
		beelog.Debug(msg)
		syncMsg := &protobuf.OrderList{}
		err := proto.Unmarshal(msg.Value, syncMsg)
		if err != nil {
			beelog.Error("unmarshaling error: ", err)
		}

		//检查topic
		if checkTopic(conf.Conf.KafkatopicList, msg.Topic) == false {
			beelog.Warning("wrong topic:", msg.Topic, " current topic is:", conf.Conf.KafkatopicList, "msg:", syncMsg)
			goto MARKOFFSET
		}

		if err := handle.PutMsg(syncMsg); err != nil {
			//TODO 处理失败消息， 放入日志数据库， 由定时任务进行处理。
			beelog.Debug("handle failed:")
		} else {
			beelog.Debug("handle success:")
		}

	MARKOFFSET:
		consumer.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

func NewCloudServer(addrs []string, groupID string, topicList []string) *OrderServer {
	serv := new(OrderServer)
	serv.KafkaAddrs = addrs
	serv.KafkaGroupID = groupID
	serv.KafkatopicList = topicList

	return serv
}

func (serv *OrderServer) Start() error {
	var handle = OrderConsumeHandle{}

	consumer := kafka.NewConsumerDefault(serv.KafkaGroupID, serv.KafkaAddrs, serv.KafkatopicList, &handle)
	consumer.Start()
	beelog.Debug("connect to kafka:", serv.KafkaAddrs, " topic:", serv.KafkatopicList, "group:", serv.KafkaGroupID)

	return nil
}
