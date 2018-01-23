package conf

import (
	"github.com/astaxie/beego"
)

var Conf *Config

type Config struct {
	KafkaAddr         string
	KafkaTopic        string
	Sid               string
	AID               string
	Key               string
	Tokenurl          string
	Refreshurl        string
	Serviceurl        string
	ICODE             string
	MysqlDsn          string
	MysqlDriver       string
	MaoyanKey         string
	MeituanSecretSign string
	MeituanUrl        string
}

func InitConf() {
	var confs Config
	confs.KafkaAddr = beego.AppConfig.String("kafka::addr")
	confs.KafkaTopic = beego.AppConfig.String("kafka::topic")
	confs.AID = beego.AppConfig.String("xiechen::AID")
	confs.Sid = beego.AppConfig.String("xiechen::Sid")
	confs.Key = beego.AppConfig.String("xiechen::Key")
	confs.Tokenurl = beego.AppConfig.String("xiechen::tokenurl")
	confs.Refreshurl = beego.AppConfig.String("xiechen::refreshurl")
	confs.Serviceurl = beego.AppConfig.String("xiechen::serviceurl")
	confs.ICODE = beego.AppConfig.String("xiechen::Icode")
	confs.MysqlDriver = beego.AppConfig.String("driver")
	confs.MysqlDsn = beego.AppConfig.String("dsn")
	confs.MaoyanKey = beego.AppConfig.String("maoyan::key")
	confs.MeituanSecretSign = beego.AppConfig.String("meituan::secret_sign")
	confs.MeituanUrl = beego.AppConfig.String("meituan::geturl")
	Conf = &confs
}
