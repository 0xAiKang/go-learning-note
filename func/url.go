package main

import (
	"fmt"
	"net"
	"net/url"
)

func main()  {
	var p = fmt.Println

	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// 解析URL
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u.Scheme)

	// User 包含了所有认证信息
	p(u.User)

	// Host
	p(u.Host)

	// 解析Host、Port
	host, port, _ := net.SplitHostPort(u.Host)
	p(host)
	p(port)

	// 获取查询参数
	p(u.RawQuery)

	// 将查询参数解析成一个map
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])
}