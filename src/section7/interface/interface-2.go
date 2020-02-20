package main

import "fmt"

func main() {
	// 인터페이스 구현 예제
	// ex1
	dog1 := Dog{"poll", 10}
	var interface1 Behavior
	interface1 = dog1
	interface1.bite()
	// dog1.bite() // 단순한 struct 를 사용하는 경우

	// ex2
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2) // 코드에서 이런형식의 사용이 더 많이 쓰임
	inter2.bite()

	// ex3 slice behavior
	inters := []Behavior{dog1, dog2}

	// 인덱스 형태로 실행
	for idx, _ := range inters {
		inters[idx].bite()
	}

	// 값 형태로 실행(인터페이스)
	for _, val := range inters {
		val.bite()
	}

}

type Dog struct {
	name   string
	weight int
}

// bite 메소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}
