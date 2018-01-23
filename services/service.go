package services

import (
	"encoding/json"
	//	"strconv"
	"time"

	"order_flume/lib/beelog"
	"order_flume/models"
	"order_flume/protobuf"
	"order_flume/utils"
)

func AddMeituanOrder(msg *protobuf.OrderList) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Data), &data); err != nil {
		beelog.Error(err)
	}
	orderDetail := models.NewOrderDetail("")
	orderDetail.ID = msg.MsgId
	orderDetail.Msg = msg.Data
	orderDetail.Type = uint8(msg.MsgType)
	if err := orderDetail.Add(); err != nil {
		beelog.Error(err)
	}
	beelog.Debug(data)
	if content, ok := data["order_list"].([]interface{}); ok {
		for i := 0; i < len(content); i++ {
			if order, ok := content[i].(map[string]interface{}); ok {
				var meituan models.TMeituan
				beelog.Debug(order, meituan)
				meituan.Orderid = order["orderid"].(string)
				meituan.Quantity = order["quantity"].(string)
				meituan.Dealid = order["dealid"].(string)
				meituan.Direct = order["direct"].(string)
				meituan.Modtime = order["modtime"].(string)
				meituan.Smstitle = order["smstitle"].(string)
				meituan.Paytime = order["paytime"].(string)
				meituan.Sid = order["sid"].(string)
				meituan.Total = order["total"].(string)
				meituan.Type = order["type"].(string)
				meituan.UID = order["uid"].(string)
				meituan.CreateTime = time.Now().UTC()
				//		meituan.Orderid = content["orderid"]
				//		meituan.Quantity = content["quantity"]
				//		meituan.Dealid = content["dealid"]
				//		meituan.Direct = content["direct"]
				//		meituan.Modtime = content["modtime"]
				//		meituan.Smstitle = content["smstitle"]
				//		meituan.Paytime = content["paytime"]
				//		meituan.Sid = content["sid"]
				//		meituan.Total = content["total"]
				//		meituan.UID = content["uid"]
				if _, err := models.AddTMeituan(&meituan); err != nil {
					beelog.Error(err)
				}
			}
		}

	}

}
func AddXiechenOrder(msg *protobuf.OrderList) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Data), &data); err != nil {
		beelog.Error(err)
	}
	orderDetail := models.NewOrderDetail("")
	orderDetail.ID = msg.MsgId
	orderDetail.Msg = msg.Data
	orderDetail.Type = uint8(msg.MsgType)
	if err := orderDetail.Add(); err != nil {
		beelog.Error(err)
	}
	//	timeLayout := "2006-01-02 15:04:05"
	//	dateLayout := "2006-01-02"
	//	loc, _ := time.LoadLocation("Local") //重要：获取时区
	//	beelog.Debug(data)
	//	if context, ok := data["order_list"].([]interface{}); ok {
	//		beelog.Debug("test2222")
	//		for i := 0; i < len(context); i++ {
	//			if order, ok := context[i].(map[string]interface{}); ok {
	//				var xiechen models.TXiechen
	//				xiechen.Orderid = order["orderid"].(string)
	//				xiechen.Ouid = order["ouid"].(string)
	//				userstatus := order["usestatus"].(string)
	//				xiechen.Usestatus, _ = strconv.Atoi(userstatus)
	//				xiechen.Orderstatus = order["orderstatus"].(string)
	//				xiechen.Orderstatusname = order["orderstatusname"].(string)
	//				xiechen.Orderamount = order["orderamount"].(string)
	//				orderdate := order["orderdate"].(string)
	//				xiechen.Orderdate, _ = time.ParseInLocation(dateLayout, orderdate, loc)
	//				xiechen.Sid = order["sid"].(string)
	//				xiechen.Ordertype = order["ordertype"].(string)
	//				xiechen.Ordername = order["ordername"].(string)
	//				Startdatetime := order["startdatetime"].(string)
	//				Pushdate := order["pushdate"].(string)
	//				xiechen.Startdatetime, _ = time.ParseInLocation(dateLayout, Startdatetime, loc)
	//				xiechen.Pushdate, _ = time.ParseInLocation(timeLayout, Pushdate, loc)
	//				xiechen.Guid = order["guid"].(string)
	//				xiechen.CreateTime = time.Now().UTC()
	//				beelog.Debug(&xiechen)
	//				if _, err := models.AddTXiechen(&xiechen); err != nil {
	//					beelog.Error(err)
	//				}
	//			}
	//		}
	//	}

}

func AddMaoyanOrder(msg *protobuf.OrderList) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Data), &data); err != nil {
		beelog.Error(err)
	}
	orderDetail := models.NewOrderDetail("")
	orderDetail.ID = msg.MsgId
	orderDetail.Msg = msg.Data
	orderDetail.Type = uint8(msg.MsgType)
	if err := orderDetail.Add(); err != nil {
		beelog.Error(err)
	}
}
func AddMeituanWaiMaiOrder(msg *protobuf.OrderList) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Data), &data); err != nil {
		beelog.Error(err)
	}
	orderDetail := models.NewOrderDetail("")
	orderDetail.ID = msg.MsgId
	orderDetail.Msg = msg.Data
	orderDetail.Type = uint8(msg.MsgType)
	if err := orderDetail.Add(); err != nil {
		beelog.Error(err)
	}
}
func SendToYM(msg *protobuf.OrderList, url string) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Data), &data); err != nil {
		beelog.Error(err)
	}
	if err := utils.SendToSZ([]byte(msg.Data), url); err != nil {
		faillog := models.NewFailLog(uint8(msg.MsgType))
		faillog.ID = msg.MsgId
		faillog.Msg = msg.Data
		faillog.Count = 1
		if err := faillog.Add(); err != nil {
			beelog.Error(err)
		}
	}
}
