package main

import "fmt"

func main() {

	// if ~ else

	var a int = 100
	b := 150

	if a >= 100 {
		fmt.Println("100 이상")
	} else {
		fmt.Println("100 이하")
	}

	if b >= 150 {
		fmt.Println("150 이상")
	} else {
		fmt.Println("150 이하")
	}

	// 아래의 주석과 같이 사용시 에러 -> else 괄호 개행 허용되지 않음.
	/*
	    if b >= 150 {
	       fmt.Println("150 이상")
	   } else
	   {
	       fmt.Println("150 이하")
	   }
	*/
}
