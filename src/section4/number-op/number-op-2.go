package main

import "fmt"

func main() {
	// ex 1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex1 : ", n1+n2)
	fmt.Println("ex1 : ", n1-n2)
	fmt.Println("ex1 : ", n1*n2)
	fmt.Println("ex1 : ", n1/n2)
	fmt.Println("ex1 : ", n1%n2)
	fmt.Println("ex1 : ", n1<<2)
	fmt.Println("ex1 : ", n1>>2)
	fmt.Println("ex1 : ", ^n1)

	// ex 2
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	// 형태가 같지 않기 때문에 예외발생
	// fmt.Println("ex2 : ", n3+n4)

	// 정수와 실수 값으로 연산하는 경우 실수로 변환하여 사용하는 것이 좋다. (하지만 정수로 반올림 및 의도하는 경우 상관없음)
	fmt.Println("ex2 : ", float32(n3)+n4)
	fmt.Println("ex2 : ", n3+int(n4)) // 주의

	// n6 는 120000 이지만 형변환시 최대값을 초과하기 때문에 잘못된 값이 들어올 수 있으므로 계산이 잘못됨.
	// unit16 의 최대값은 65535
	fmt.Println("ex2 : ", n5+uint16(n6))
}
