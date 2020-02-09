package main

import "fmt"

func main() {
	// 배열 복사
	// 값 복사 확인. 중요

	// ex1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Println("ex1 : ", arr1, &arr1)
	fmt.Println("ex1 : ", arr2, &arr2)
	// ex1-1 :  [1 10 100 1000 10000] &[1 10 100 1000 10000]
	// ex1-2 :  [1 10 100 1000 10000] &[1 10 100 1000 10000]

	// %p 는 포인터
	fmt.Printf("ex1 : %p %v\n", &arr1, arr1)
	fmt.Printf("ex1 : %p %v\n", &arr2, arr2)
	// ex1 : 0xc000098030 [1 10 100 1000 10000]
	// ex1 : 0xc000098060 [1 10 100 1000 10000]
	// pointer 주소 확인시 주소가 다르므로 값이 복사 되었다는 것을 확인할 수 있다.

}
