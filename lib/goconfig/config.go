package goconfig

import (
	"flag"
	"fmt"

	"github.com/Terry-Mao/goconf"
	"github.com/cihub/seelog"
)

var (
	confFile string
)

func init() {
	flag.StringVar(&confFile, "c", "./app.conf", " set api  config file path")
}

//InitConfig  初始化配置文件
func InitConfig(conf interface{}) error {
	return InitConfigWithfile(conf, confFile)
}

//InitConfigWithfile  初始化配置文件
func InitConfigWithfile(conf interface{}, confFile string) error {
	gconf := goconf.New()
	if err := gconf.Parse(confFile); err != nil {
		return fmt.Errorf("goconf parse config err :%v", err)
	}
	if err := gconf.Unmarshal(conf); err != nil {
		seelog.Errorf("unmarshal config err :%v", err)
		return err
	}
	return nil
}
