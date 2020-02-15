package main

import "fmt"

func main() {
	// ex1
	stack()
	// defer가 10번 누적되기 때문에 가장 마지막 부터 처음까지 누적된 defer 실행됨.
}

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 :", i)
	}
}
