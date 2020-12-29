package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// 通过结构体定义全局变量
var Conf = struct {
	// Award info
	Award struct{
		StartTime string `toml:"startTime"`
		EndTime string `toml:"endTime"`
	}

	AwardMap map[string]int64

	// 全局Mysql 配置
	Mysql struct{
		Dsn string `toml:"Dsn"`
	}

	// 全局Redis 配置
	Redis struct{
		Ip string `toml:"ip"`
		Port uint32 `toml:"port"`  			// 无符号
		Network string `toml:"network"`
	}

	// Go 语言对数据类型的要求非常严格，所以数据类型一定需要保持一致
}{}

// 读取配置文件
func pareConf() error {
	//if _, err := toml.DecodeFile("./conf/config.toml", &Conf); err!= nil {
	if _, err := toml.DecodeFile("/Users/boo/Project/GitHub/go_learning_note/lottery_system/conf/config.toml", &Conf); err!= nil {
		fmt.Println("decode config file error", err)
	}

	fmt.Println("awardMap", Conf.AwardMap)
	return nil
}