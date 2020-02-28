package main

import (
	"fmt"
	"math"
)

func main() {
	// 에러 처리 고급
	// Go error 패키지 New 메소드 사용 -> 사용자 에러 처리 생성.

	// ex1
	if f, err := Power1(7, 3); err != nil {
		fmt.Printf("Errror message : %s\n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}

	// ex2
	if f, err := Power1(0, 3); err != nil {
		fmt.Printf("Errror message : %s\n", err.Error())
	} else {
		fmt.Println("ex2 : ", f)
	}
}

// f의 i 제곱 구하는 함수
func Power1(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("(%g)는/은 사용 불가능 합니다.", f)
	}
	return math.Pow(f, i), nil // 제곱, nil 반환
}
