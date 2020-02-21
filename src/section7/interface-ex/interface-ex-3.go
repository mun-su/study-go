package main

import "fmt"

func main() {
	// 실제 타입 검사 switch 사용
	// 타입검사시 switch를 주로 사용함
	// 비어있는 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크를 통해 형 변환 후 사용 가능.

	// ex1
	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("Hello Golang!")
}

// 아래와 같은 타입과 포인터인지 일반타입인지 switch 사용가능.
func checkType(arg interface{}) {
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int", arg)
	case float32, float64:
		fmt.Println("This is a float", arg)
	case string:
		fmt.Println("This is a string", arg)
	case nil:
		fmt.Println("This is a nil", arg)
	default:
		fmt.Println("What is this type?", arg)
	}
}
