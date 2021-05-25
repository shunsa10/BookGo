package config

import (
	"log"
	"todo/utils"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList //グローバルに使うため

func init()  { //main関数より先に読み込ませたい
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig()  {
	cfg, err := ini.Load("config.ini") //iniファイルの読み込み
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{ //読み込んだiniをConfigListに入れていく
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static: cfg.Section("web").Key("static").String(),
		//js css の読み込み
	}
}