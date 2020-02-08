package main

import "fmt"

func main() {
	// golang 의 문자열 표현
	// 문자열 : 기본적으로 UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의!

	// * index 접근 가능하며, index 로 접근시 ascii 정수로 출력이 되며 실제 문자열로 출력하기 위해서는 Printf 사용.
	// ex 1 (하나의 글자를 배열로 생각하는게 이해하기 좋다)
	var str1 string = "Golang" // 6 byte
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	// ascii code 값이 출력됨.
	// 71 111 108 97 110 103
	fmt.Println("ex 1 : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])

	// 87 111 114 108 100
	fmt.Println("ex 1 : ", str2[0], str2[1], str2[2], str2[3], str2[4])

	// 234 179 160 32 237 148 132 235 161
	fmt.Println("ex 1 : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5], str3[6], str3[7], str3[8])

	// ex 2 -> format 형태로 출력시 한글은 깨짐.
	// G o l a n g
	fmt.Printf("ex 2 : %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	// W o r l d
	fmt.Printf("ex 2 : %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
	// 문자열 깨짐 -> ê ³     í
	fmt.Printf("ex 3 : %c %c %c %c %c %c\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// 한글을 정상적으로 출력
	// * 한글이 깨질 경우 rune 형태 생성자로 문자열로 넣어서 하는 방법
	str4 := []rune(str3)
	fmt.Printf("ex 3 : %c %c %c %c %c %c\n", str4[0], str4[1], str4[2], str4[3], str4[4], str4[5])

	// * 문자열은 배열로 취급되기 때문에 반복으로 순회 가능
	// ex 3
	for i, char := range str1 {
		fmt.Printf("ex3 : %c(%d)\t", char, i)
	}
	fmt.Println("")

	for i, char := range str2 {
		fmt.Printf("ex3 : %c(%d)\t", char, i)
	}
	fmt.Println("")

}
