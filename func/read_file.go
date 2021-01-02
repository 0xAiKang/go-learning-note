package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件需要经常进行错误检查，创建一个方法来精简程序
func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func main()  {
	// echo "hello" > /tmp/dat
	// echo "go" >> /tmp/dat
	var p = fmt.Println

	// 读取文件
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	p(string(dat))

	// 如果希望对文件进行更多的操作，可以先打开一个文件，获取 os.File
	f, err := os.Open("/tmp/dat")
	check(err)
	p(f)

	// 关闭文件
//	defer f.Close()

	// 从文件开始位置读取一些字节
	b1 := make([]byte, 5)
	p(b1)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s \n", n1, string(b1[:n1]))

	// Seek 到文件中已知的位置，并从这个位置开始读取
	o2, err := f.Seek(6, 0)
	p(o2)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ",n2, o2)
	fmt.Printf("%v \n", string(b2[:n2]))

	// io包提供了更健壮的ReadAtLeast，用于读取上面那种文件
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s \n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s \n", string(b4))
}
