package main

import (
	"errors"
	"fmt"
	"strconv"
)

func GetFibonacci(n int) ([]int, error) {
	// 及早失败，避免嵌套
	if n<2 || n>100 {
		return nil, errors.New("n should be in 2-100")
	}
	// 声明一个切片
	fibList := []int{1, 1}

	for i:=2; i<n; i++ {
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList, nil
}

func GetFibonacci2(str string)  {
	var (
		i int
		err error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}

	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(list)
}

func main()  {
	if v, err := GetFibonacci(-2); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(v)
	}

	GetFibonacci2("s")
}