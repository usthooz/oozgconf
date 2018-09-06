package main

import (
	"github.com/usthooz/oozgconf"
	"github.com/usthooz/oozlog/go"
)

type Config struct {
	Author string
	Mysql  struct {
		User     string
		Password string
	}
}

func main() {
	var (
		conf Config
	)
	ozconf := oozgconf.NewConf(&oozgconf.OozGconf{
		ConfPath: "./config.json",
		Subffix:  "",
	})
	err := ozconf.GetConf(&conf)
	if err != nil {
		uoozg.Errorf("GetConf Err: %v", err.Error())
	}
	uoozg.Infof("Res: %v", conf)
}
