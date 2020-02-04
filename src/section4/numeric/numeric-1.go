package main

import "fmt"

func main() {
	// 데이터 타입 : 숫자형
	// 정수, 실수, 복소수 지원
	// 32bit, 64bit, unsigned(양수)
	// 정수 : 8진수(0), 16진수(0x), 10진수

	/*
		int32, int64 가 복잡하다고 느껴지는 경우 단순한 int 를 사용하는게 좋다.
		데이터 타입에 따라 큰차이는 없음.
		사용할 범위가 명확한 경우에는 그것에 맞게 최적화를 하면 되겠지만
		요즘 언어에서는 외부적으로는 정확하게 지정하지 않지만
		내부적으로는 cast가 됨.
	*/
	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)

	//num1 :  17
	//num2 :  -68
	//num3 :  409
	//num4 :  855256181
}
