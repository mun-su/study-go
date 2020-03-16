package main

import "fmt"

func main() {
	// Go Recover 함수
	// 에러 복구 가능

	// ex1
	runFunc1()
	fmt.Println("Hello Golang!")
}

func runFunc1() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message :", s)
		}
		// s := recover()
		// fmt.Println("Error Message :", s)
	}()

	a := [3]int{1, 2, 3}
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 :", a[i])
	}
}
