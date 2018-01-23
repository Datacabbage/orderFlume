package main

import (
	_ "order_flume/lib/beelog"
	"order_flume/thirdOrder/conf"
	_ "order_flume/thirdOrder/routers"
	"order_flume/thirdOrder/service"
	"order_flume/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/redis.v5"
)

func InitKafka() {
	service.InitProduce(conf.Conf.KafkaAddr)
}
func main() {
	conf.InitConf()
	InitKafka()
	orm.RegisterDataBase("default", conf.Conf.MysqlDriver, conf.Conf.MysqlDsn)
	var err error
	redisIp := beego.AppConfig.String("redisIp")
	redisPort := beego.AppConfig.String("redisPort")
	connectInfo := redisIp + redisPort
	utils.Cache = redis.NewClient(&redis.Options{
		Addr:     connectInfo,
		PoolSize: 100,
	})
	err = utils.Cache.Ping().Err()
	if err != nil {
		panic(err)
	}
	beego.Run()
}
