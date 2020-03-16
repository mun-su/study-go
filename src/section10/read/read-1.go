package main

import (
	"fmt"
	"os"
)

func main() {
	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의 (오류 메시지 확인)
	// 예외 처리 정말 중요!
	// 탐색 Seek 중요!
	// https://golang.org/pkg/os

	// 파일 읽기 ex
	// 파일 열기
	file, err := os.Open("sample.txt")
	errCheck1(err)

	// 읽기 ex1
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 흭득
	errCheck2(err)
	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용을 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력(1) :", fileInfo)
	fmt.Println("파일 이름(1) :", fileInfo.Name())
	fmt.Println("파일 크기(1) :", fileInfo.Size())
	fmt.Println("파일 수정 시간(1) :", fileInfo.ModTime())
	fmt.Printf("읽기 작업 완료(1) (%d bytes)\n\n", ct1)
	// fmt.Println(fd1) // 타입 변환 없을 경우
	fmt.Println(string(fd1)) // 타입 변환 있는 경우

	fmt.Println("======================================================")

	// 읽기 ex2 (탐색 : Seek(offset))
	o1, err := file.Seek(20, 0) // 0: 처음 위치, 1: 현재 위치, 2: 마지막 위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)
	fmt.Printf("읽기 작업 완료(2) (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2))
	fmt.Println("======================================================")

	// 읽기 ex3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치부터 읽어온다.
	errCheck1(err)

	fmt.Printf("읽기 작업 완료(3) (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3))
	fmt.Println("======================================================")

	defer file.Close()
}

func errCheck1(err error) {
	if err != nil {
		panic(err)
	}
}

func errCheck2(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
