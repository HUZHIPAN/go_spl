package main

import (
	"fmt"
	"sync"
)

func escape01() *int {
	i := 100

	return &i

}

func escape02() int {
	return 0
}

func demo003(m map[int]int) int {
	defer func () int {
		err := recover()
		if err != nil {
			fmt.Println("捕获异常：", err)
		}
		fmt.Println("defer err")
		return 23
	}()

	panic("abc")

	m[34] = 34
	return 123
}

func main() {
	sync.Mutex{}

	tm := map[int]int{1: 1, 2: 2}
	fmt.Println(tm)
	iir := demo003(tm)

	fmt.Println(iir)

	fmt.Println(tm)
	// escape01()
	escape02()
}
