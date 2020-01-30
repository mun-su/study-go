package main

import "fmt"

func main() {
	/*
	   반드시 Boolean 을 사용해야 하며 1, 0은 허용 되지않음.
	   다른 언어에서는 1, 0을 true, false 로 형변환이 되지만 golang 에서는 자동 형 변환 불가
	   소괄호 사용 하지 않음.
	*/
	// if
	var a int = 10
	b := 20

	if a >= 5 {
		fmt.Println("10 이상")
	}

	// 25 이하 이므로 print 되지 않음.
	if b >= 25 {
		fmt.Println("25 이상")
	}

	// 아래의 주석과 같이 사용시 에러 -> 괄호 개행시
	/*
	   if b >= 15
	   {
	       fmt.Println("25 이상")
	   }
	*/

	// 아래와 주석과 같이 사용시 에러 -> 괄호 생략시
	/*
	   if b >= 15
	       fmt.Println("25 이상")
	*/

	// 아래와 주석과 같이 사용시 에러 -> 1, 0 자동 형변환 되지 않고 Boolean type 이 아니기 때문

	/*
	   if c := 1; c {
	       fmt.Println("true")
	   }
	*/

	// 아래와 주석과 같이 사용시 에러 -> 짧은 선언은 if 내에서만 사용할수 있으므로.
	/*
	   if c := 40; c >= 30 {
	       fmt.Println("30 이상")
	   }
	   c += 10
	*/
}
