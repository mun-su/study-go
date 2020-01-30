package main

import "fmt"

func main() {
	// switch
	// switch 뒤 표현식(Expression) 생략 가능
	// case 뒤 표현식(Expression) 사용 가능
	// 자동 break 때문에 fallthrough 존재
	// Type 분기 -> 값이 아닌 변수 Type으로 분기 가능

	// ex 1
	a := -7

	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// ex 2
	// 위의 선언 후 사용보다 아래의 방법을 golang 레퍼런스에서 추천 -> scope 가 switch 내로 제한하수 있으므로.
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// ex 3
	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치 하는것이 없음.")
	}

	// ex 4
	// 연산자로 계산이 가능.
	switch d := "go"; d + "lang" {
	case "golang":
		fmt.Println("Golang!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치 하는것이 없음.")
	}

	// ex 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j")
	case i == j:
		fmt.Println("i == j")
	case i > j:
		fmt.Println("i > j")
	}
}
