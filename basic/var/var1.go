package main

import "fmt"

func main() {

	// 정수 타입 : 0
	// 실수 (소수점) : 0.0
	// 문자열 : ""
	// 논리형 : true, false

	// 변수명 규칙 -> 숫자 첫글자(x), 대소문자 구분(O), 문자, 숫자, 밑줄, 특수기호 사용가능

	// 변수 및 상수 : 함수 내외 사용 가능

	// int 는 int32와 동일
	// int 는 초기화 하지 않을시 primitive type 으로서 0이 할당.
	var a int
	var b string
	var c, d, e int

	// f = 1, g = 2, h = 3
	var f, g, h int = 1, 2, 3

	var i float32 = 11.4
	var j string = "Hello Golang!"
	var k = 4.74
	var l = "Hello Korea"
	var m = true

	// 위 의 코드에서 선언 및 초기화된 변수를 사용하지 않을시 에러 발생함.

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)
}
