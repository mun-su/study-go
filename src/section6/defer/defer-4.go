package main

import "fmt"

func main() {
	a()
	/*
		start : b
		in a
		end : b
	*/
}

func a() {
	defer end(start("b")) // 중첩 함수 주의
	fmt.Println("in a")
}

func start(t string) string {
	fmt.Println("start :", t)
	return t
}

func end(t string) {
	fmt.Println("end :", t)
}
