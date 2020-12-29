package main

import (
	"fmt"
	"sync"
)

func dataProducer(ch chan int, wg *sync.WaitGroup)  {
	// 创建一个协程来产生数据
	go func() {
		for i:=0; i<10; i++ {
			// 将i 放入channel
			ch <- i
		}
		wg.Done()
		// 关闭channel
		close(ch)
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup)  {
	// 这种方式存在一个问题：我们很有可能不知道channel 的个数
	/*go func() {
		for i:=0; i<10; i++ {
			// 从channel 中取出数据
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()*/

	// 当获取通道的次数大于通道中的数量时，会默认返回对应类型的零值
	go func() {
		// 第十一个值是 0
		for i:=0; i<11; i++ {
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
	
	/*go func() {
		for  {
			// true 表示正常接收，false 表示通道关闭
			if data, ok := <-ch; ok {
				fmt.Println(data)
			}else{
				break
			}
		}
		wg.Done()
	}()*/
}

func main()  {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Wait()
}