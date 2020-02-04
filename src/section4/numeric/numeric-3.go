package main

import "fmt"

func main() {
	// 실수(부동소수점)
	// float32(7자리), float64(15자리)

	// ex 1
	var num1 float32 = 0.14
	var num2 float32 = .75647
	var num3 float32 = 442.0378373
	var num4 float32 = 10.0

	// 지속 표기법
	var num5 float32 = 14e6
	var num6 float64 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)
	fmt.Println("num5 : ", num5)
	fmt.Println("num6 : ", num6)
	fmt.Println("num7 : ", num7)

	fmt.Println("float32 num4 - 0.1 : ", float32(num4-0.1))
	fmt.Println("float64 num4 - 0.1 : ", float64(num4-0.1)) // 소수점 연산은 항상 주의가 필요함

}
