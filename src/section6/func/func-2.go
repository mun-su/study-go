package main

import "fmt"

func main() {
	// 함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)
	// ex1
	// 함수 전달
	sum1(100, add)

	// ex2
	// call by value
	a := 100
	multi_value(a)
	fmt.Println("ex2 :", a)

	// ex3
	// reference by value
	b := 100
	multi_reference(&b)
	fmt.Println("ex3 :", b)
}

// ex1
func sum1(i int, f func(int, int)) {
	f(i, 10)
}

// ex1
// 같은 data type인 경우 생략 가능.
func add(a, b int) {
	fmt.Println("ex :1", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i *= 10 // *i = *i * 10
}
