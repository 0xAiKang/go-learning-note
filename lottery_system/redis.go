package main

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

/**
	获取Redis 连接
 */
/*func getConn() (redis.Conn, error) {
	// redis.Dial("tcp", "127.0.0.1:6379")
	conn, err := redis.Dial(Conf.Redis.Network, fmt.Sprintf("%s:%d", Conf.Redis.Ip, Conf.Redis.Port))

	if err != nil {
		fmt.Println("connect to redis error ", err)
		return nil, err
	}

	return conn, nil
}*/

// 由于 redis 连接太多没释放和网络超时导致，该用Redis 连接池
func getConn() (redis.Conn, error){
	p := getPool()
	if p == nil {
		return nil, errors.New("connect to redis error ")
	}
	return p.Get(), nil
}

var pool *redis.Pool

/**
	创建Redis 连接池
 */
func getPool() *redis.Pool {
	if pool == nil {
		// 创建连接池
		return &redis.Pool{
			// 设定最大连接数
			MaxIdle:     2000,
			IdleTimeout: 3 * time.Minute,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp",
					fmt.Sprintf("%s:%d", Conf.Redis.Ip, Conf.Redis.Port),
					// 分别设定Redis 连接、写、读超时时间
					redis.DialConnectTimeout(800 * time.Millisecond),
					redis.DialWriteTimeout(800 * time.Millisecond),
					redis.DialReadTimeout(800 * time.Millisecond),
				)
				if err != nil {
					log.Println("connect to redis error: ", err)
				}
				return conn, nil
			},
		}
	}
	return pool
}
