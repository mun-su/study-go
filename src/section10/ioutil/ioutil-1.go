package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 활용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기 가능
	// 아래 메소드 확인 가능
	// WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/pkg/io/ioutil/

	s := "Hello Golang\n File Write Test\n"

	// 파일 모드(chmod, chown, chgrp) -> 퍼미션
	// 읽기 : 4, 쓰기 : 2, 실행 : 1
	// 소유자, 그룹, 기타 사용자 순서 (777) -> 4+2+1 = 7 읽기 + 쓰기 + 실행 권한
	// https://golang.org/pkg/os/#FileMode

	// 0을 붙이는 이유는 bit 에서 전혀 다른값이 나오기 때문에 go 에서 퍼미션을 할당할때는 0을 붙인다.
	// 소유자 : 읽기/쓰기 가능, 그룹 : 읽기 가능, 기타 사용자 : 읽기 가능
	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	data, err := ioutil.ReadFile("sample.txt")
	errCheck(err)

	fmt.Println("=========================================================")
	fmt.Println(string(data))
	fmt.Println("=========================================================")
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
