package main

import "fmt"

func main() {
	// 함수 고급
	// 익명 함수(Anonymous Function)
	// 즉시 실행 (선언과 동시에)

	// ex1
	// 기존에 func를 작성하던 비슷한 형태로 하면 실행되지않고 에러 발생 -> func literal evaluated but not used
	// func() {
	// 	fmt.Println("ex1 : Anonymous Function!")
	// }

	// 기존에 func를 호출하던 방법
	// test()

	// 함수 생성과 호출을 합쳐서 바로 실행.
	func() {
		fmt.Println("ex1 : Anonymous Function!")
	}()

	// ex2
	msg := "Hello Golang!"
	func(m string) {
		fmt.Println("ex2 :", m)
	}(msg)

	func(m string) {
		fmt.Println("ex2 :", m)
	}("Hello Golang!")

	// ex3
	func(x, y int) {
		fmt.Println("ex3 :", x+y)
	}(10, 20)

	// ex4
	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println("ex4 :", r)
}
