package main

import "fmt"

func main() {
	// 슬라이스(슬라이스 참조 타입 증명)
	// ex1 (배열)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int
	arr2 = arr1
	arr2[0] = 7

	fmt.Println("ex1 : ", arr1)
	fmt.Println("ex1 : ", arr2)
	// ex1 :  [1 2 3]
	// ex1 :  [7 2 3]

	// ex2 (슬라이스)
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7
	fmt.Println("ex2 : ", slice1)
	fmt.Println("ex2 : ", slice2)
	// ex2 :  [7 2 3]
	// ex2 :  [7 2 3]

	// ex3 (슬라이스 예외 상황)
	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4])
	// 용량만큼 초기화 되는것이 아닌 길이 만큼 초기화 되기 때문에 예외발생.
	// fmt.Println("ex3 : ", slice3[50]) // 길이 index 예

	// ex4 (슬라이스 순환)
	for i, v := range slice1 {
		fmt.Println("ex4 : ", i, v)
	}

}
