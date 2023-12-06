package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	MySqlUser     string
	MySqlPassword string
	MySqlDataBase string
	MysqlIP       string
	MysqlPort     string
)

var (
	RedisAddr     string
	RedisPassword string
	RedisDB       = 2 //redis的连接池
)

func LoadMysql() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	MySqlUser = cfg.Section("mysql").Key("User").String()
	MySqlPassword = cfg.Section("mysql").Key("Password").String()
	MySqlDataBase = cfg.Section("mysql").Key("DataBase").String()
	MysqlIP = cfg.Section("mysql").Key("IP").String()
	MysqlPort = cfg.Section("mysql").Key("Port").String()
}

func LoadRedis() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	RedisAddr = cfg.Section("redis").Key("Addr").String()
	RedisPassword = cfg.Section("redis").Key("Password").String()
	RedisDB, _ = cfg.Section("redis").Key("DB").Int()
}
