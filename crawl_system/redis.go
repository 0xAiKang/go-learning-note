package main

import "github.com/gomodule/redigo/redis"

/**存入Redis
	conn ：Redis 连接对象
	topic：文章id
	date：对应日期
 */
func save(conn redis.Conn, topic string, date string) error {
	_, err := conn.Do("SET", topic, date)
	return err
}

/** 检查是否存在
	conn ：Redis 连接对象
	topic：文章id
 */
func exists(conn redis.Conn, topic string) (int, error) {
	return redis.Int(conn.Do("EXISTS", topic))
}