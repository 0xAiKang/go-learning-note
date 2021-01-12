package main

import (
	"fmt"
	"os"
)

func main () {
	// 获取进程 id
	fmt.Println(os.Getpid())
	// 获取父进程 id
	fmt.Println(os.Getppid())
}
