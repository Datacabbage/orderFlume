package service

import (
	"encoding/json"

	"order_flume/lib/beelog"
	"order_flume/protobuf"
	"order_flume/thirdOrder/conf"
	"order_flume/utils"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
)

//处理kafka生产成功和错误消息
type OrderProduceHandler struct {
}

func handleRecover() {
	if r := recover(); r != nil {
		//fmt.Println("Recovered in sendProcessMessage", r)
		//check exactly what the panic was and create error.
		switch x := r.(type) {
		case string:
			beelog.Error("produce handleRecover Error : ", x)
		case error:
			beelog.Error("produce handleRecover Error : ", x)
		default:
			//err = errors.New("Unknow panic")
			beelog.Error("produce handleRecover Unknow: ", x)
		}
	}
}

func handleSync(ch chan<- int) {
	ch <- 1
}

func (handler *OrderProduceHandler) HandleSuccess(msg *sarama.ProducerMessage) error {
	return nil
}

func (handler *OrderProduceHandler) HandleError(perr *sarama.ProducerError) error {

	return nil
}

func SendToKafka(info interface{}, types string) interface{} {
	var (
		msg      []byte
		err      error
		oderlist []interface{}
	)
	ret := make(map[string]string, 0)
	collect := make(map[string]interface{})
	oderlist = append(oderlist, info)
	collect["order_list"] = oderlist
	collect["partner"] = types
	if msg, err = MessageSerialize(collect, types); err != nil {
		beelog.Error("MessageSerialize:", err)
		ret["errcode"] = "196610"
		ret["errmsg"] = err.Error()
		return ret
	}
	producer := ProducerInstance()
	if err = producer.PutMsg(msg, types, conf.Conf.KafkaTopic, collect); err != nil {
		beelog.Error("put message to kafka error:", err)
		ret["errcode"] = "196610"
		ret["errmsg"] = err.Error()
		return ret
	}
	//插入数据库
	if err != nil {
		beelog.Debug("insert failed")
		ret["errcode"] = "196610"
		ret["errmsg"] = "insert fail"
		return ret
	}
	ret["errcode"] = "0"
	ret["errmsg"] = "ok"
	return ret
}

func MessageSerialize(collect map[string]interface{}, types string) ([]byte, error) {
	msg := &protobuf.OrderList{}
	msg.MsgId = utils.CreateId()
	switch types {
	case "美团":
		msg.MsgType = 0
	case "携程酒店":
		msg.MsgType = 1
	case "猫眼":
		msg.MsgType = 2
	case "美团外卖":
		msg.MsgType = 3
	default:
		msg.MsgType = 0
	}

	if msgdata, err := json.Marshal(collect); err != nil {
		return nil, err
	} else {
		msg.Data = string(msgdata)
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return data, nil
}
