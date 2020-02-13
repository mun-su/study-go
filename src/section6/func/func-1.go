package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 함수
	// 선언 : func 키워드로 선언
	// func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
	// func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재
	// func 함수명() : 매개변수 없음, 반환 값 없음
	// 타 언어와 달리 반환 값(return value) 다수 가능

	// ex1
	helloGolang()

	// ex2
	sayOne("Hello World!")

	// ex3
	result := sum(5, 5)
	fmt.Println("ex3 :", result)
	fmt.Println("ex3 :", sum(5, 5))
	// 숫자를 문자열로 변경. -> strconv.Itoa
	fmt.Println("ex3 :", strconv.Itoa(sum(5, 5)))

}

// 함수 선언 위치는 어느 곳이든 가능, main 위 아래 모두 가능

// ex1
func helloGolang() {
	fmt.Println("ex1 : Hello Golang")
}

// ex2
func sayOne(m string) {
	fmt.Println("ex2 :", m)
}

// ex3
func sum(x int, y int) int {
	return x + y
}
