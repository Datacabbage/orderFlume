package models

import (
	"order_flume/common/mongo"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type FailLog struct {
	ObjectId_  bson.ObjectId `bson:"_id"`
	ID         string        `bson:"id"`
	Type       uint8         `bson:"type"`       //任务类型
	Count      int           `bson:"count"`      //任务执行次数
	Msg        string        `bson:"msg"`        //内容
	CreateTime time.Time     `bson:"createTime"` //创建时间
	UpdateTime time.Time     `bson:"endtime"`    //结束时间
}

func NewFailLog(types uint8) *FailLog {
	col := new(FailLog)
	col.Type = types
	return col
}

//添加一条信息
func (operate *FailLog) Add() error {
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
	if err = GetFailLog(db).Insert(operate); err == nil {
		return nil
	}
	return err
}

//查找信息
func (operate *FailLog) Get() (*FailLog, error) {

	db := mongo.MSessionGet("operate")
	if db == nil {
		beego.Debug(SessionFailError)
		return nil, SessionFailError
	}
	defer db.Close()
	var one *FailLog
	if err := GetFailLog(db).Find(bson.M{"id": operate.ID}).One(&one); err != nil {
		return nil, err
	}
	return one, nil
}

//查找信息
func (faillog *FailLog) IsExist() bool {
	failTmp := NewFailLog(faillog.Type)
	failTmp.ID = faillog.ID
	failTmp, err := failTmp.Get()

	if err == nil && failTmp != nil {
		return true
	}

	return false
}

//更改状态
//func (operate *FailLog) SetStatus(status uint8, errmsg string) error {

//	db := mongo.MSessionGet("operate")
//	if db == nil {
//		beego.Debug(SessionFailError)
//		return SessionFailError
//	}
//	defer db.Close()
//	endtime := time.Now()

//	return GetFailLog(db).Update(bson.M{"type": operate.Type, "uid": operate.UID, "id": operate.ID}, bson.M{"$set": bson.M{"status": status, "endtime": endtime, "errmsg": errmsg}})
//}
