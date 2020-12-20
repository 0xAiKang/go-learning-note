package main

import "fmt"

// defer 语句
func main()  {
	for i:=0; i<5; i++ {
		// 主调函数退出时，defer后的函数才会被调用。
		defer fmt.Println(i)   // 43310
		fmt.Println(i)	 	   // 01234

		// 通常在进行数据库连接、文件操作时，会使用 defer 语句进行数据库连接的释放，释放文件句柄和锁的释放。
	}
}