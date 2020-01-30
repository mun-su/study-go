package main

import "fmt"

func main() {

	// ex 1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	}

	// ex 2
	switch b := 30; b / 15 {
	case 2, 4, 6:
		fmt.Println("a -> ", b, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", b, "는 홀수")
	}

	// ex 3
	/*
	   fallthrough -> 통과
	   조건에 해당되는 경우 자동으로 break 되지만
	   fallthrough 가 존재하면 조건인 "go" 다음 case 도 실행이됨.
	   go, python, c 에 해당하는 print가 모두 실행됨.
	   하지만 마지막의 조건에서는 다음 case 가 존재하지 않기 때문에 fallthrough 이 존재하는 경우 에러발생.
	   cannot fallthrough final case in switch
	*/
	switch c := "go"; c {
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "go":
		fmt.Println("Golang!")
		fallthrough
	case "python":
		fmt.Println("Python!")
		fallthrough
	case "c":
		fmt.Println("C!")
		// 조건의 마지막에는 fallthrough 를 사용할수 없음.
		// fallthrough
	}
}
