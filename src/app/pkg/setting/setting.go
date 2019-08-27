package setting

import (
	"github.com/go-ini/ini"
	log "github.com/shengkehua/xlog4go"
	"time"
)

type Service struct {
	Host  string
	Level string // production debug dev
}

var ServiceSetting = &Service{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     string
	MaxActive   string
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {
	log.Info("init settings")
	var err error
	confPath := "conf/app.ini"
	cfg, err = ini.Load(confPath)
	if err != nil {
		log.Fatal("setting setup: error to parse %s.", confPath)
	}

	mapTo("service", ServiceSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Error("Cfg.MapTo %s err: %v", section, err)
	}
}
