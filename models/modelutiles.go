package models

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var (
	SessionFailError error = errors.New("get session fail")
)

const (
	STATUS_NOMARL = iota
	STATUS_DELETE
)

func GetFailLog(db *mgo.Session) *mgo.Collection {
	return db.DB("orderflume").C("fail_info")
}

func GetOrderDetail(db *mgo.Session) *mgo.Collection {
	return db.DB("orderflume").C("orderDetail")
}
