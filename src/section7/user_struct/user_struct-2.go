package main

import "fmt"

func main() {
	// 기본 자료형 사용자 정의 타입
	// ex1
	a := cnt(5)
	fmt.Println("ex1 :", a)

	// ex2
	var b cnt = 15
	fmt.Println("exb :", b)

	// ex3
	// testConvert1(b) // 예외 발생 (중요), 사용자 정의 타입 <-> 기본타입 : 매개변서 전달시에 변환해야  사용가능 (int(변수))
	testConvert1(int(b))
	testConvert2(b)
}

type cnt int

func testConvert1(i int) {
	fmt.Println("ex3 : (Default Type)", i)
}

func testConvert2(i cnt) {
	fmt.Println("ex4 : (Custom Type)", i)
}
