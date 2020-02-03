package lib

import "fmt"

// main method가 존재하지 않지만 다른 파일에 import 될때 실행됨.
func init() {
	fmt.Println("Lib Package Init Start!")
}

func CheckSum(c int32) bool {
	return c < 100
}
