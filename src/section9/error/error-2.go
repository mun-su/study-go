package main

import (
	"fmt"
	"log"
)

func main() {
	// 에러 처리
	// Errorf 를 이용한 에러 처리

	a, err := notZero(11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ex1 : ", a)

	b, err := notZero(0)
	if err != nil {
		log.Fatal(err)
	}
	// Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("ex2 : ", b)

	fmt.Println("End Error Handling!")
}

func notZero(n int) (string, error) { // 메소드 리턴 값 error 타입 중요!
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}
	return "", fmt.Errorf("%d를 입력했습니다. 에러발생!", n)
}
