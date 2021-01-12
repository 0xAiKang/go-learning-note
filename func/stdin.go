package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	// echo hello go | go run stdin.go
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(bytes)
	}
	fmt.Println(string(bytes))
}