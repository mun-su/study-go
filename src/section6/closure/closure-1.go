package main

import "fmt"

func main() {
	// Closure
	// 익명함수를 사용할 경우 함수를 변수에 할당해서 사용가능.
	// 함수 안에서 함수를 선언 및 정의 가능 -> 외부 함수에 선언된 변수에 접근 가능한 함수.
	// 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 함수를 호출할 때 이전에 존재했던 값을 유지하기 위해서. -> 비동기, 누적 카운트 -> 무분별한 전역 변수 남발...
	// 남발 -> 객체들이 메모리에 존재하므로 -> 메모리 부족, 오버플로우 현상이 발생할 수 있다.
	// Closure를 정확하게 이해하고 사용해야 한다.

	// Javascript 에서 주로 사용, Golang, Java 8에서도 사용가능.
	// Closure 기능을 염두하고 사용하면 퍼포먼스적으로 좋을수 있다.

	// ex1 -> closure를 사용하지 않은 경우
	multiply := func(x, y int) int { // 익명 함수
		return x * y
	}
	r1 := multiply(5, 10)
	fmt.Println("ex1 :", r1)

	// ex2 -> closure 사용.
	m, n := 5, 10            // 변수가 캡처됨.
	sum := func(c int) int { // 익명함수 변수 할당
		return m + n + c // 지역 변수 소멸되지 않는다. (함수 호출시 마다 사용가능)
	}
	r2 := sum(10)
	fmt.Println(r2)
}
