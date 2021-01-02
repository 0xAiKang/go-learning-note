package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	// 创建一个通道来接收通知
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify 注册给定的通道，用户接收特定的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTRAP)

	// 这个协程执行一个阻塞的信号接收操作，当它接收到一个值时，打印这个值，然后通知程序可以退出
	go func() {
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 程序在这里等待，直到它得到期望的信号，也就是上面的协程发送的done 值
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}