package main

import "fmt"

func main() {
	// 문자 끝 세미콜론(;) 주의
	// 자동으로 끝에 세미콜론 삽입됨.
	// 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용 (권장하지 않음)
	// 반복문 및 제어문 (if, for) 사용할 경우 주의

	// ex 1
	// 두 문장을 사용하는 경우 -> 사용은 가능하지만..
	for i := 0; i <= 10; i++ {
		fmt.Println("ex1 : ", i)
		fmt.Println(i)
	}

	// ex 2
	// 두 문장을 사용하는 경우 -> 사용은 가능하지만..
	for i := 0; i <= 10; i++ {
		fmt.Println("ex2 : ", i)
		fmt.Println(i)
	}

	// ex 3
	/*
		for j := 10; j >= 0; j--
		{
			fmt.Println("ex3 : ", j)
		}
		위와 같이 작성시 에러가 발생하며 다음과 같이 해석하게 되어 실행 불가능하기 때문에 에러 발생.
		for j := 10; j >= 0; j--;{
			fmt.Println("ex3 : ", j)
		}
	*/
	for j := 10; j >= 0; j-- {
		fmt.Println("ex3 : ", j)
	}
}
