package conf

import (
	"fmt"

	"order_flume/lib/goconfig"
)

type Config struct {
	MongoTable   string `goconf:"mongodb:table"`
	MongoURL     string `goconf:"mongodb:url"`
	MongoMaxconn int    `goconf:"mongodb:maxconn"`
	MysqlDsn     string `goconf:"mysql:dsn"`
	MysqlDriver  string `goconf:"mysql:driver"`

	//kafka
	KafkaAddrs     []string `goconf:"kafka:addrs:,"`
	KafkaGroupID   string   `goconf:"kafka:groupid"`
	KafkatopicList []string `goconf:"kafka:topiclist:,"`

	//base
	Version string `goconf:"base:version"`
	YMUrl   string `goconf:"base:url"`
	Type    string `goconf:"base:type"`
}

var Conf *Config

func InitConfig() {
	Conf = &Config{}
	if err := goconfig.InitConfig(Conf); err != nil {
		panic(fmt.Sprintf("init config err:%v", err))
	}
}
