package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main()  {
	// Go 标准库的  net/http  包为http 客户端提供了出色的支持

	resp, err := http.Get("http://0x2beace.com")
	if err != nil {
		panic(err)
	}
	// 关闭
	defer resp.Body.Close()

	// 打印响应状态
	fmt.Println("Response status: ", resp.Status)

	// 打印response body 的前五行
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i<5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// 整个过程是同步阻塞的，最后输出 ok
	fmt.Println("ok")
}