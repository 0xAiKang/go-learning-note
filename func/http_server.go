package main

import (
	"fmt"
	"net/http"
)

/**
	handlers 是 net/http 服务器里面的一个基本概念
	handler 对象实现了 http.Handler 接口
	handler 有两个参数，分别是 http.ResponseWriter、http.Request
*/
func hello(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "hello go \n")
}

/**
	这个handler 稍微复杂一些，需要读取 http 请求的所有内容，然后将他们输出至response body
 */
func handlers(w http.ResponseWriter, req *http.Request)  {
	for name, heanders := range req.Header{
		for _, h := range heanders{
			fmt.Fprintf(w, "%v: %v \n", name, h)
		}
	}
}

func main()  {
	// 使用net/http 包，可以轻松实现一个简单的 http 服务器
	// 使用 http.HandleFunc 函数，可以将我们的handler 注册到服务器路由
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/handlers", handlers)

	// 设置监听端口
	http.ListenAndServe(":8090", nil)
}