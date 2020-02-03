package main

import (
	"fmt"
)

// init 은 main 보다 먼저 실행됨.
func init() {
	fmt.Println("Init Method Start!")
}

func main() {
	// init : 패키지 로드시에 가장 먼저 호출
	// 가장 먼저 초기화 되는 작업 작성시 유용.
	fmt.Println("Main Method Start!")
}
