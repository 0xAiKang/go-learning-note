package main

import (
	"fmt"
	"time"
)

func main()  {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.ANSIC))

	fmt.Printf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}