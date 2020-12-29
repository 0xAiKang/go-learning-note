package main

import (
	"fmt"
	"net/http"
)

//   初始化
func init()  {
	// 解析配置文件
	pareConf()

	// 初始化日志文件
	InitLogConf()

	// 初始化Redis
	InitAwardRedis()
}

// 建立Http 请求
func Start()  {
	// 定义路由
	http.HandleFunc("/draw", handleDrawRequests)
	// 监听端口
	http.ListenAndServe(":8080", nil)
}

func main()  {
	fmt.Println("ok")
	Start()
}
