package main

import "fmt"

func main() {
	// 다중 값 반환
	// ex1
	a, b := multiply(10, 5)

	// 아래와 같이 사용시 예외 발생
	// c := multiply(10, 5)

	// 밑줄을 사용하여 특정 할당만 사용가능
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)

	fmt.Println("ex 1 :", a, b)
	fmt.Println("ex 1 :", c)
	fmt.Println("ex 1 :", d)

	// ex2
	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, y3, _, y5 := arrayMultiply(1, 2, 3, 4, 5)
	fmt.Println("ex 2 :", x1, x2, x3, x4, x5)
	fmt.Println("ex 2 :", y1, y3, y5)
}

// 같은 data type인 경우 생략 가능.
func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}
