package main

import "fmt"

func main() {
	// 맵(Map)
	// 맵 조회 할 경우에 주의 할 점

	// ex1
	map1 := map[string]int{
		"apple":  5,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	// 존재 하지 않는 Key의 값을 가져오는 경우, 선언한 Value의 data type의 초기값으로 할당된다.
	// int : 0, string : "", float: 0.0
	value1 := map1["lemon"]
	// a, b는 선언되지 않았지만 할당하는 내용이 없다고 생각 할 수 있지만 맵의 2번째 값인 해당 키의 존재 여부 bool 값이 할당된다.
	value2, a := map1["lemon"]
	value3, b := map1["kiwi"] // 초기값이 들어가기 때문에 예외가 발생하지 않음

	fmt.Println("ex1 : ", value1)
	fmt.Println("ex1 : ", value2, a)
	fmt.Println("ex1 : ", value3, b)

	// ex2
	// ok는 key의 존재 여부.
	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : banana is not exist!")
	}

	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist!")
	}

	// ex3
	// key의 존재 여부만으로 다음과 같이 표현 가능.
	if _, ok := map1["banana"]; !ok {
		fmt.Println("ex3 : banana is not exist!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex3 : kiwi is not exist!")
	}
}
