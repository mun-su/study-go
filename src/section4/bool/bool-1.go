package main

import "fmt"

func main() {
	// Bool (Boolean) : 참과 거짓
	// 조건부 논리 연산자와 주료 사용 : !, || (or), && (and)
	// 암묵적 형변환이 불가능 :  0, Nil -> false 변환 없음.

	// ex 1
	var b1 bool = true
	var b2 bool = false
	b3 := true
	//b4 := 1

	fmt.Println("b1 : ", b1)
	fmt.Println("b2 : ", b2)
	fmt.Println("b3 : ", b3)

	fmt.Println("b3 == b2 : ", b3 == b3)
	fmt.Println("b3 || b2 : ", b1 || b3)
	fmt.Println("b3 && b2 : ", b1 && b2)
	fmt.Println("!b1 : ", !b1)

	// 암묵적인 형변환이 되지 않음.
	// non-bool b4 (type int) used as if condition
	/*
		if b4 {
			fmt.Println("b4 is true")
		}
	*/
}
