package main

import "fmt"

func main() {
	// 함수 고급
	// 가변 인자 실습(매개변수 개수가 동적으로 변할 때 - 정해져 있지 않음)

	// ex1
	x := multiply(5, 6, 7, 8, 9, 10)

	fmt.Println("ex1 :", x)
	fmt.Println()

	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("ex1 :", y)
	fmt.Println()

	// ex2
	prtWord("a", "apple", "test", "new", "golang")
	fmt.Println()

	// ex3 -> 배열
	a := []int{5, 6, 7, 8, 9, 10}

	// throw : cannot use a (type []int) as type int in argument to multiply
	// m := multiply(a)
	m := multiply(a...)
	fmt.Println("ex3 :", m)
	fmt.Println()

	// throw : cannot use a (type []int) as type int in argument to multiply
	// n := sum(a)
	n := sum(a...)
	fmt.Println("ex3 :", n)
	fmt.Println()
}

// ex1
func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}
	return tot
}

// ex2
func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println(value)
	}
}

// ex3
