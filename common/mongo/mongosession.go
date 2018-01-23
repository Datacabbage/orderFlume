package mongo

import (
	"cloud_server/utils"
	"fmt"

	"gopkg.in/mgo.v2"
)

type mgoDict map[string]*mgo.Session

var mgoSessionPool mgoDict = make(mgoDict, 1)

//  [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
func MSessionAddPool(alias, url string, maxPoolSize int) error {
	if alias == "" || url == "" {
		return utils.ErrorParameter
	}

	fmt.Println(url)
	mgoSession, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err) //直接终止程序运行
		return err
	}

	//最大连接池默认为4096
	if maxPoolSize > 0 {
		mgoSession.SetPoolLimit(maxPoolSize)
	}

	mgoSessionPool[alias] = mgoSession

	return nil
}

func MSessionGet(alias string) *mgo.Session {
	mgoSession, ok := mgoSessionPool[alias]

	if !ok {
		return nil
	}

	return mgoSession.Clone()
}

func MSessionClose(session *mgo.Session) {
	session.Close()
}
