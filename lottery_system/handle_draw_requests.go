package main

import (
	"fmt"
	"net/http"
)

/**
	该回调函数接收两个参数，分别是：
	rsp：响应对象
	req：请求对象
 */
func handleDrawRequests(rsp http.ResponseWriter, req *http.Request)  {
	// 获取请求数据
	params := req.URL.Query()
	var username string
	usernames, ok := params["username"]
	if ok {
		username = usernames[0]
	} else {
		fmt.Println("username is empty")
		rsp.Write([]byte("username cannot be empty"))
		return
	}
	awardBatch := GetAward(username)
	if awardBatch == nil {
		rsp.Write([]byte("sorry you don't win any prize"))
	}else{
		words := fmt.Sprintf("congratutions ! you won a %s", awardBatch.GetName())
		rsp.Write([]byte(words))
	}
}