package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// 큰 따옴표 "", 백 스쿼트 ``
	// 문자 char 는 제공하지 않음 -> rune(int32)로 문자 코드 값으로 표현
	// 문자 : '' 로 작성
	// 자주 사용하는 escape : \\, \', \", \a (콘솔벨), \b(백스페이스), \f(쪽 바끔), \n (줄바끔), \r (복귀), \t(탭)

	// ex 1
	var str1 string = "d:\\study\\go\\src\\" // -> d:\study\go\src\
	str2 := `d:\study\go\src\`

	fmt.Println("ex 1 : ", str1)
	fmt.Println("ex 1 : ", str2)

	// ex 2
	var str3 string = "Hello, Golang!"
	var str4 string = "안녕하세요."
	// 유니코드 표현시 소문자 u 는 4바이트, 대문자 U 는 8 바이트
	var str5 string = "\ud55c\uae00"

	fmt.Println("ex2 : ", str3)
	fmt.Println("ex2 : ", str4)
	fmt.Println("ex2 : ", str5)

	// ex 3 -> 문자열 byte 길이
	// 다른 언어에서는 -> len, length
	// golang 에서는 길이와 크기를 구하는 것을 명확하게 지원한다
	// 길이 (바이트 수) -> 한글 하나는 3byte 이므로 주의.
	fmt.Println("ex 3 : ", len(str3)) // 14
	fmt.Println("ex 3 : ", len(str4)) // 16

	// ex 4 -> 문자열 길이
	// rune 의 배열형태로 선언해서 byte 를 이용해 구하는 방법
	fmt.Println("ex 4 : ", len([]rune(str3))) // 14
	fmt.Println("ex 4 : ", len([]rune(str4))) // 6

	// RuneCountInString 사용
	// 대다수가 rune 의 배열 형태보다 RuneCountInString 을 많이 사용.
	fmt.Println("ex 4 : ", utf8.RuneCountInString(str3)) // 14
	fmt.Println("ex 4 : ", utf8.RuneCountInString(str4)) // 6
}
