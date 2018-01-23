package main

import (
	"os"

	"order_flume/common/mongo"
	"order_flume/lib/beelog"
	"order_flume/lib/sign"
	_ "order_flume/models"
	"order_flume/orderConsumer/conf"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitMongo(model string) {
	var err error
	url := conf.Conf.MongoURL
	maxNum := conf.Conf.MongoMaxconn
	if err = mongo.MSessionAddPool(model, url, maxNum); err != nil {
		beelog.Error("error resolving address:", err)
		os.Exit(1)
	}

}

func main() {
	conf.InitConfig()
	beelog.Debug(conf.Conf)
	orm.RegisterDataBase("default", conf.Conf.MysqlDriver, conf.Conf.MysqlDsn)
	InitMongo(conf.Conf.MongoTable)
	serv := NewCloudServer(conf.Conf.KafkaAddrs, conf.Conf.KafkaGroupID, conf.Conf.KafkatopicList)
	serv.Start()

	beelog.Trace("===start===")
	sign.InitSignal(conf.Conf.Version)
}
