package main

import "fmt"

func main() {
	// 맵(Map)
	// 맵 : 해시테이블, 딕셔너리(파이썬), Key - Value 로 자료 저장
	// 레퍼런스 타입(배열, 슬라이스와 같은 참조 값 전달)
	// 비교 연산자 사용 불가능(참조 타입이므로)
	// 특징 : 참조타입(key)로 사용 불가능, 값(Value) 모든 타입 사용 가능
	// make 함수 및 축약(리터럴)로 초기화 가능
	// 순서가 유지 되지 않으므로 iterator 사용시 순서대로 나오지 않음.

	// ex1
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)
	map3 := make(map[string]int) // 주로 사용됨

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)

	// ex2
	map4 := map[string]int{} // Json 형태 map4 = {"a" : 1, "b": 2}
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 73

	// json 형태로 넣기 가능.
	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}

	map6 := make(map[string]int, 10) // 기본 10개 할당
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 73

	// map이기 때문에 실행할때 마다 순서가 다를 수 있음.
	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["orange"])
	fmt.Println("ex2 : ", map6["apple"])
}
