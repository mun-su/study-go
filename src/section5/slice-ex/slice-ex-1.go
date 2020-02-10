package main

import "fmt"

func main() {
	// 슬라이스 추가 및 병합
	// 용량초과시 최초 생성시 용량의 2배에 해당하는 용량을 가진 임시 temp 에 할당을 한 후 s1에 재할당됨.
	// 생각보다 temp 생성 후 재할당은 퍼포먼스에 좋지 않으므로 용량에 신경써야함.
	// java 기준으로 볼때 HashMap 의 default capacity 의 1.7배 만큼 사용하게 되엇을때 일때 capacity 를 2배 확장하게 되는데 그와 비슷한것으로 보임
	// ex1
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      // 슬라이스를 삽입할 경우 ... 사용
	s3 = append(s3, s3[0:3]...) // 추출 후 병합.

	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)

	// ex2
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len : %d, cap : %d, value %v\n", len(s4), cap(s4), s4) // 길이 및 용량 자동 증가
	}
	// ex2 -> len : 1, cap : 5, value [0]
	// ex2 -> len : 2, cap : 5, value [0 1]
	// ex2 -> len : 3, cap : 5, value [0 1 2]
	// ex2 -> len : 4, cap : 5, value [0 1 2 3]
	// ex2 -> len : 5, cap : 5, value [0 1 2 3 4]
	// ex2 -> len : 6, cap : 10, value [0 1 2 3 4 5]
	// ex2 -> len : 7, cap : 10, value [0 1 2 3 4 5 6]
	// ex2 -> len : 8, cap : 10, value [0 1 2 3 4 5 6 7]
	// ex2 -> len : 9, cap : 10, value [0 1 2 3 4 5 6 7 8]
	// ex2 -> len : 10, cap : 10, value [0 1 2 3 4 5 6 7 8 9]
	// ex2 -> len : 11, cap : 20, value [0 1 2 3 4 5 6 7 8 9 10]
	// ex2 -> len : 12, cap : 20, value [0 1 2 3 4 5 6 7 8 9 10 11]
	// ex2 -> len : 13, cap : 20, value [0 1 2 3 4 5 6 7 8 9 10 11 12]
	// ex2 -> len : 14, cap : 20, value [0 1 2 3 4 5 6 7 8 9 10 11 12 13]
	// ex2 -> len : 15, cap : 20, value [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
}
