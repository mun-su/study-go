package main

import "fmt"

func main() {
	// 배열
	// 배열은 용량, 길이 항상 같다.

	// 배열 vs 슬라이스 차이점 중요
	// 길이고정 vs 길이 가변
	// 값 타입 vs 참조 타입
	// 복사 전달 vs 참조 값 전달
	// 전체 비교연산자 사용 가능 vs 비교 연산자 사용불가
	// 대부분 슬라이스를 사용함.

	// cap() : 배열, 슬라시으 용량
	// len() : 배열, 슬라이스 개수

	// ex1
	// 선언
	var arr1 [5]int
	// 선언과 초기화 -> 선언한 개수와 초기화되는 값의 개수가 같아야함, 들어가지 않은 값은 0으로 초기화됨.
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}         // 할당되지 않은 곳은 0으로 초기화됨
	arr6 := [...]int{1, 2, 3, 4, 5} // 배열크기 자동 맞춤
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, // 마지막에 콤마가 들어가야함
	}

	fmt.Printf("ex1 : %-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("ex1 : %-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("ex1 : %-5T %d %v\n", arr2, len(arr3), arr3)
	fmt.Printf("ex1 : %-5T %d %v\n", arr2, len(arr4), arr4)
	fmt.Printf("ex1 : %-5T %d %v\n", arr2, len(arr5), arr5)
	fmt.Printf("ex1 : %-5T %d %v\n", arr2, len(arr6), arr6)
	fmt.Printf("ex1 : %-5T %d %v\n", arr2, len(arr7), arr7)

	// ex2
	arr8 := [5]int{1, 2, 3, 4, 5}
	// 가독성을 위해 다음과 같이 사용하기도 함.
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"seo", "kim", "han"}
	fmt.Printf("ex1 : %-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("ex1 : %-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("ex1 : %-5T %d %v\n", arr10, len(arr10), arr10)

}
