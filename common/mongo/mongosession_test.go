package mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestConn(t *testing.T) {
	url := "mongodb://10.0.12.104:27017/"
	MSessionAddPool("contact", url, 10)

	var ret map[string]interface{} = make(map[string]interface{}, 1)

	session := MSessionGet("contact")
	if session == nil {
		t.Log("get session error")
		t.Fail()
	}

	if err := session.DB("sync_contact").C("items_info").Find(bson.M{"uid": "456"}).One(ret); err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(ret)
}
