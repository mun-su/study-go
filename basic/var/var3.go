package main

import "fmt"

func main() {
	// 짧은 선언
	// 기존 var 변수명 변수형
	// 함수 안에서만 사용(전역 X), 선운 후 할당 예외 발생
	// 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있다.
	var a int = 1

	// 주로 Oracle 의 프로시저의 값 할당과 비슷
	b := 2
	c := "c"
	d := false

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)

	// 다음과 같이 재할당 하는 경우 에러가 발생
	// d := true

	if i := 7; i < 10 {
		fmt.Println("i is variable")
	}

}
