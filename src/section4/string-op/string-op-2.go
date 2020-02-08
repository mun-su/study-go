package main

import "fmt"

func main() {
	// 문자열 비교
	// ex 1
	str1 := "Golang" // 6 byte
	str2 := "World"  // 5 byte

	fmt.Println("ex 1 : ", str1 == str2)
	fmt.Println("ex 1 : ", str1 != str2)

	// str2 가 더 크다고 판단하여 false
	// 이유는 Go 문자열 -> ascii 코드 정수 값을 비교하기 때문임.
	// abc, ABC 비교시 대문자 ABC가 더크게 나옴.
	fmt.Println("ex 1 : ", str1 > str2)
	fmt.Println("ex 1 : ", str1 < str2)
}
