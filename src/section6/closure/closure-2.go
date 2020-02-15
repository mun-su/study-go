package main

import "fmt"

func main() {
	// ex1
	cnt := increaseCnt()
	fmt.Println("ex1 :", cnt()) // 1
	fmt.Println("ex1 :", cnt()) // 2
	fmt.Println("ex1 :", cnt()) // 3
	fmt.Println("ex1 :", cnt()) // 4
	// cnt() 지역변수가 캡쳐된 값이 지속적으로 유지되므로 주의 해야함.
	// 사이드 이펙트 주의.
}

// ex1 -> 함수를 return
func increaseCnt() func() int {
	n := 0 // 지역변수 캡쳐됨
	return func() int {
		n += 1
		return n
	}
}
