package main

// 사용하지 않는 경우 밑줄 -> 임시 import
import (
	"fmt"
	_ "section3/package/lib"
	checkUp "section3/package/lib"
)

func main() {
	// 패키지 접근제어
	// alias -> 별칭 사용
	// 빈 식별자 사용
	fmt.Println("10보다 큰 수? :", checkUp.CheckNum(11))
}
