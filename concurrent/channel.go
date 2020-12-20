package main

import "fmt"

// 信道
func sum(a []int, c chan int)  {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// 将合计传入channel
	c <- sum
}

func main()  {
	a := []int{7, 2, 8, -9, 4, 0}

	// 创建一个通道
	c := make(chan int)

	// go 并发
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// 从 channel 中获取
	x, y := <-c, <-c
	fmt.Printf("x: %d \n y: %d \n x+y: %d \n", x, y, x +y)
}