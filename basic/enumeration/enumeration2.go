package main

import "fmt"

func main() {
	// iota 사용
	// iota 는 0으로 시작 됨.
	// iota 를 활용하여 다음과 같이 규칙 적용
	const (
		Zero = iota * 10
		Ten
		Twenty
	)

	fmt.Println(Zero)   // 0
	fmt.Println(Ten)    // 10
	fmt.Println(Twenty) // 10
}
