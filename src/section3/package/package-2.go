package main

// gopath가 project root로 지정되어야 사용가능
import (
	"fmt"
	"section3/package/lib"
)

func main() {
	// lib 하위에 존재하는 func 실행.
	fmt.Println("10보다 큰 수? : ", lib.CheckNum(15))

	// checkNum 은 네이밍 규칙에 따라 private이므로 호출 불가
	// fmt.Println("10보다 큰 수? : ", lib.checkNum(15))
}
