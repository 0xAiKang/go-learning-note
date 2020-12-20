package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程
func main()  {
	var wg sync.WaitGroup
	wg.Add(2)

	// 协程A
	go func() {
		fmt.Println("1")
		time.Sleep(100)
		fmt.Println("2")
		time.Sleep(100)
		fmt.Println("3")
		wg.Done()
	}()

	// 协程B
	go func() {
		fmt.Println("A")
		time.Sleep(100)
		fmt.Println("B")
		time.Sleep(100)
		fmt.Println("C")
		wg.Done()
	}()

	wg.Wait()
}