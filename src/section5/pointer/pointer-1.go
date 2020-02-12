package main

import "fmt"

func main() {
	// 포인터
	// Golang : 포인터 지원(C/C++)
	/*
		숙련된 개발자가 사용하는 경우 정말 좋은 기능이며 적재 적소에 잘 사용 할수 있으나
		숙달된 프로그래머라고 할지라 포인터는 어려우며 실수는 치명적인 에러를 발생시킵니다
		포인터가 단점이 있다라기 보다 위험성이 내포하고 있다고 생각하는게 좋다.
	*/
	// 변수의 지역성, 연속된 메모리 참조..., 힙, 스택...
	// 파이썬, 자바 포인터를 지원하지 않지만 컴파일러, 인터프리터 내부적으로 동작합니다
	// 포인터 지원 X -> 파이썬, C#, Java 등
	// Golang은 리스크가 있는 포인터의 연산은 할수 없도록 되어있음.
	// 주소의 값은 직접 변경 불가능 (잘못된 코딩으로 인한 버그 방지)
	// 역참조의 값을 바꾸거나 주소값 자체를 입력하는 행위 불가.
	// * (아스터리스크)로 사용
	// nil 로 초기화 (nil == 0)

	// ex1
	// 주소값을 할당 해야하며 일반 값을 할당시 에러
	// var a *int = 5

	var a *int            // 방법 1
	var b *int = new(int) // 방법 2

	// a = 8 // cannot use 8 (type int) as type *int in assignment

	fmt.Println(a) // nil
	fmt.Println(b) // 주소 값.

	i := 7
	fmt.Println("ex1 : ", i, &i) // & 붙일시 i의 주소값.

	fmt.Println()

	a = &i // i의 주소값 할당
	b = &i

	fmt.Println("ex1 a : ", a)  // i의 주소값
	fmt.Println("ex1 a : ", &a) // i의 주소값을 저장하기 위한 a 자체의 주소값
	fmt.Println("ex1 a : ", *a) // 역참조 -> 최종적으로 저장되어있는 위치까지 가서 참조

	fmt.Println()

	fmt.Println("ex1 b : ", b)
	fmt.Println("ex1 b : ", &b)
	fmt.Println("ex1 b : ", *b)

	var c = &i
	d := &i

	*d = 10 // 역참조된 곳의 값을 변경

	fmt.Println("ex1 c : ", c)
	fmt.Println("ex1 c : ", &c)
	fmt.Println("ex1 c : ", *c)

	fmt.Println()

	fmt.Println("ex1 d : ", d)
	fmt.Println("ex1 d : ", &d)
	fmt.Println("ex1 d : ", *d)
}
