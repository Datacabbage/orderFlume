package models

import (
	"order_flume/common/mongo"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type OrderDetail struct {
	ObjectId_  bson.ObjectId `bson:"_id"`
	ID         string        `bson:"id"`
	Type       uint8         `bson:"type"`       //任务类型
	Msg        string        `bson:"msg"`        //内容
	CreateTime time.Time     `bson:"createTime"` //创建时间
	UpdateTime time.Time     `bson:"endtime"`    //结束时间
}

func NewOrderDetail(id string) *OrderDetail {
	col := new(OrderDetail)
	col.ID = id
	return col
}

//添加一条信息
func (operate *OrderDetail) Add() error {
	var (
		err error
	)
	db := mongo.MSessionGet("operate")
	if db == nil {
		beego.Debug(SessionFailError)
		return SessionFailError
	}
	defer db.Close()
	tm := time.Now().UTC()
	operate.ObjectId_ = bson.NewObjectId()
	operate.UpdateTime = tm
	operate.CreateTime = tm
	if err = GetOrderDetail(db).Insert(operate); err == nil {
		return nil
	}
	return err
}

//查找信息
func (operate *OrderDetail) Get() (*OrderDetail, error) {

	db := mongo.MSessionGet("operate")
	if db == nil {
		beego.Debug(SessionFailError)
		return nil, SessionFailError
	}
	defer db.Close()
	var one *OrderDetail
	if err := GetOrderDetail(db).Find(bson.M{"id": operate.ID}).One(&one); err != nil {
		return nil, err
	}
	return one, nil
}

//查找信息
func (faillog *OrderDetail) IsExist() bool {
	failTmp := NewOrderDetail(faillog.ID)
	failTmp, err := failTmp.Get()

	if err == nil && failTmp != nil {
		return true
	}

	return false
}
