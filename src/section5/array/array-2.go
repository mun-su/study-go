package main

import "fmt"

func main() {
	// 배열 순회
	// ex1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex1 : ", arr1[i])
	}
	fmt.Println("")

	// ex2
	// 가장 많이 사용됨.
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, v := range arr2 {
		fmt.Println("ex1 : ", i, v)
	}

	// ex3
	// 인덱스 생략1
	arr3 := [5]int{7, 77, 777, 7777, 77777}
	for _, v := range arr3 {
		fmt.Println("ex1 : ", v)
	}

	// ex4
	// 인덱스만 필요한 경우.
	arr4 := [5]int{7, 77, 777, 7777, 77777}
	for i := range arr4 {
		fmt.Println("ex1 : ", i)
	}
}
