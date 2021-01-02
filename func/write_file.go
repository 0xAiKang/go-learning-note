package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func main()  {
	// var p = fmt.Println

	d1 := []byte("hello go")
	// 写入文件，注意这里并不是直接将字符串写入到文件中
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 对于更细粒度的写入，先打开一个文件
	f, err := os.Create("/tmp/dat2")
	fmt.Println(f)
	check(err)
	// 打开文件后，习惯性的使用 defer 调用文件 Close
	defer f.Close()

	// 创建一个字节切片
	d2 := []byte{155, 111, 109, 101, 10}
	// 向文件中写入字节切片
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	// WriteString 也是可用的
	n3, err := f.WriteString("writes \n")
	fmt.Printf("wrote %d bytes \n", n3)
	// 调用Sync 将缓冲区的数据写入磁盘
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered \n")
	fmt.Printf("wrote %d bytes \n ", n4)
	// 使用flush 来确保，将所有缓冲操作应用于底层writer
	w.Flush()
}