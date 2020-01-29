package main

import "fmt"

func main() {
	// 밑줄과 iota 사용
	// 밑줄을 사용하는 경우 생략됨.
	const (
		_ = iota
		One
		_
		Three
		Four
		_
		Seven
	)

	fmt.Println(One)   // 1
	fmt.Println(Three) // 3
	fmt.Println(Four)  // 4
	fmt.Println(Seven) // 6
}
