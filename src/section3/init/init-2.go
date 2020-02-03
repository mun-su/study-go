package main

import (
	"fmt"
	"section3/init/lib"
)

// init method 는 개수에 상관이 없으나 같은 함수를 여러개를 사용하는것은 좋지 않음.
func init() {
	fmt.Println("Init 1!")
}

func init() {
	fmt.Println("Init 2!")
}

func init() {
	fmt.Println("Init 3!")
}

func init() {
	fmt.Println("Init 4!")
}

func main() {
	fmt.Println("Main Source Code!")
	fmt.Println("100 보다 작은가? :", lib.CheckSum(10))

	/*
		결과 값, 순서는 다음과 같음.
		import된 순서대로 포함된 파일 중 init이 있다면 먼저 실행
		init 순서대로 실행됨
		main 실행

		Lib Package Init Start!
		Init 1!
		Init 2!
		Init 3!
		Init 4!
		Main Source Code!
		100 보다 작은가? : true
	*/
}
