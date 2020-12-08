
package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {
	c,err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("Cannot connect to Redis", err)
		return
	}
	fmt.Println("Redis connected ")
	defer c.Close()
}