package main

import (
	"fmt"
	"strings"
)

func main() {
	// 문자열 결합.
	// ex 1 -> 일반연산
	str1 := "The Go programming " +
		"language is an open source " +
		"project to make programmers more productive."
	str2 := "Instructions for downloading and installing the Go compilers"

	fmt.Println("ex 1 : ", str1+str2)

	// ex 2 -> Join
	// 문자열 결합은 strings.Join 을 사용하는것이 좋다.
	// java 에서 StringBuilder, StringBuffer 사용하듯이.
	strSet := []string{} // 슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)
	fmt.Println("ex2 : ", strings.Join(strSet, ""))
}
