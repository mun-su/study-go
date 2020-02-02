package main

import "fmt"

func main() {
	// enumeration 열거형
	// 주로 상수를 사용하는 일정한 규칙에 따라 숫자를 계산 또는 증가
	const (
		Zero   = 0
		Ten    = 10
		Twenty = 20
	)

	fmt.Println(Zero)
	fmt.Println(Ten)
	fmt.Println(Twenty)
}
