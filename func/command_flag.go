package main

import (
	"flag"
	"fmt"
)

func main()  {
	var p = fmt.Println
	// Go 语言提供了一个 flag 包，支持命令行标志的解析
	// go build command_flag.go
	// ./command_flag -word=opt -numb=7 -fork -svar=flag

	// 声明一个默认值为 foo 的字符串标志 word
	// 这里的flag.String 返回一个字符串指针
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	// 用程序中已有的参数来声明一个标志
	flag.StringVar(&svar, "svar", "bar", "a string var")
	// 所有标志都声明完成之后，使用Parse 来解析
	flag.Parse()

	p("word: ", *wordPtr)
	p("numb: ", *numPtr)
	p("boolPtr: ", *boolPtr)
	p("svar: ", svar)
	p("tail: ", flag.Args())
}
