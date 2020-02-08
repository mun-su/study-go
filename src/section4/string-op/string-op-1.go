package main

import "fmt"

func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// ex 1 (추출)
	var str1 string = "Golang"
	var str2 string = "World"

	// python 과 동일한 slice 사용
	// 0 부터 2번 index 까지
	// 0 번은 printf 가 아니기 때문에 해당되는 정수값 출력됨.
	fmt.Println("ex1 : ", str1[0:2], str1[0])

	// 3 부터 마지막 index 까지
	// 0 번은 printf 가 아니기 때문에 해당되는 정수값 출력됨.
	fmt.Println("ex1 : ", str2[3:], str2[0])

	fmt.Println("ex1 : ", str2[:4])
	fmt.Println("ex1 : ", str1[1:3])
}
