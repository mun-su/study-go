package main

import "fmt"

func main() {
	// 함수 고급
	// 함수를 변수에 할당

	// ex1 (슬라이스에 할당)
	f := []func(int, int) int{multiply1, sum1}
	a := f[0](10, 10) // f[0] -> multiply(10, 10)
	b := f[1](10, 10) // f[1] -> sum(10, 10)

	fmt.Println("ex1 :", a, f[0](10, 10))
	fmt.Println("ex1 :", b, f[1](10, 10))

	// ex2 (변수에 할당)
	var f1 func(int, int) int = multiply1
	f2 := sum1
	fmt.Println("ex2 :", f1(10, 10))
	fmt.Println("ex2 :", f2(10, 10))

	// ex3 (맵에 할당)
	m := map[string]func(int, int) int{
		"mul_func": multiply1,
		"sum_func": sum1,
	}

	fmt.Println("ex3 :", m["mul_func"](10, 10))
	fmt.Println("ex3 :", m["sum_func"](10, 10))
}

func multiply1(x, y int) (r int) {
	r = x * y
	return r
}

func sum1(x, y int) (r int) {
	r = x + y
	return r
}
