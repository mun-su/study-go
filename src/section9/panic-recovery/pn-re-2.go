package main

import "fmt"

func main() {
	// Go Recover 함수
	// 에러 복구 가능
	// Panic 으로 발생한 에러를 복구 후 코드 재 실행 (프로그램 종료 되지 않음)
	// 즉, 코드 흐름을 정상 상태로 복구하는 기능
	// Panic 에서 설정한 메시지를 받아올 수 있다.

	// ex1
	runFunc()

	fmt.Println("Hello Golang!")
}

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error Message :", s)
	}()
	panic("Error occurred!")
	fmt.Println("called?") // panic 이후의 코드는 실행 되지 않음.
}
