package main

import (
	"app/pkg/setting"
	"app/router"
	"fmt"
	log "github.com/shengkehua/xlog4go"
)

func initLogger(filename string) error {
	return log.SetupLogWithConf(filename)
}

func init() {
	err := initLogger("conf/log.json")
	if err != nil {
		fmt.Println("Fail to init Logger")
		return
	}
	err = setting.Setup()
	if err != nil {
		log.Fatal("Fail to init settings")
		return
	}
}

func main() {
	r := router.RouterRegister()
	// Listen and Server in 0.0.0.0:8080
	r.Run(setting.ServiceSetting.Host)
}
