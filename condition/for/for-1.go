package main

import "fmt"

func main() {
	// Golang에서 유일하게 사용 가능한 반복문
	// 다양한 사용법 확인 필요.

	// ex 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	// error
	/*
	   for i := 0; i < 5; i++
	   {
	       fmt.Println("ex1 : ", i)
	   }
	*/

	// ex2 무한루프
	/*
	   for {
	       fmt.Println("ex2 : Infinite Loop!")
	   }
	*/

	// ex3 range
	// 다음과 같이 사용시 무조건 첫번째가 index이며 두번째가 값임.
	locations1 := []string{"Seoul", "Busan", "Gwangju"}
	for index, location := range locations1 {
		fmt.Println("ex3 : ", index, location)
	}
	// ex4
	// 값만 찍으려는 경우 다음과 같이 사용할 수 있음.
	// "_" 를 사용하여 index는 skip
	locations2 := []string{"Seoul", "Busan", "Gwangju"}
	for _, location := range locations2 {
		fmt.Println("ex3 : ", location)
	}
}
