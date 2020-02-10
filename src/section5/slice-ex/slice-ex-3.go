package main

import "fmt"

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// 반드시 make로 공간을 할당 후 복사 해야함. []int{} x
	// 복사된 슬라이스 값 변경해도 원본에는 영향이 당현이 가지 않음.

	// ex1 (복사)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("ex1 : ", slice2) // slice2의 공간이 선언되지 않았으나 길이가 5이기 때문에 1,2,3,4,5 까지만 복사됨.
	fmt.Println("ex1 : ", slice3) // 공간 할당이 되지 않았기 때문에 slice3으로 복사되지 않음.
	// ex1 :  [1 2 3 4 5]
	// ex1 :  []

	// ex2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)
	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)
	// ex2 :  [1 2 3 4 5]
	// ex2 :  [7 2 3 4 10]

	// ex3 주의 할 내용
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 부분적으로 슬라이스 추출은 참조이므로 -> 원본 값 반영 됨.

	d[1] = 7
	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)

	// ex4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 3번째의 7은 용량 지정
	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f)
}
