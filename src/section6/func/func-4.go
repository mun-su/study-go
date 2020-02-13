package main

import "fmt"

func main() {
	// 리턴 값 변수 사용
	// ex1
	a, b := multiply1(10, 5)
	fmt.Println("ex1 :", a, b)

	// ex2
	c, d := multiply2(10, 5)
	fmt.Println("ex1 :", c, d)
}

// 같은 data type인 경우 생략 가능.
func multiply1(x, y int) (r1, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

// 리턴 값 변수 미사용
func multiply2(x, y int) (int, int) {
	return x * 10, y * 20
}
