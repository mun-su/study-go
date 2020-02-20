package main

import "fmt"

func main() {
	// 인터페이스 활용(비어있는 인터페이스)
	// 함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다(만능) -> 모든 타입 지정가능

	// ex1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	printValue(dog) // ex1 : {poll 10}
	printValue(cat) // ex1 : {bob 5}

	fmt.Println()

	// 어떤 타입이든 받아서 처리가 가능함.
	printValue(15)
	printValue("Animal")
	printValue(true)
	printValue(25.5)
	printValue([]Dog{})
	printValue([5]Dog{})
}

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) { // 비어 있는 인터페이스
	fmt.Println("ex1 :", s)
}
