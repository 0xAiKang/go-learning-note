package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func InitLogConf() error {
	// 获取年月日
	y, m, d := time.Now().Date()
	// 拼接文件名
	fileName := fmt.Sprintf("%d_%d_%d_award.log", y, m, d)
	// 文件操作
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Errorf("file open error, %v", err)
	}

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(f)

	return nil
}