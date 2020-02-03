package main

import "fmt"

func main() {
	a := 100

	// if ~ else if
	if a >= 120 {
		fmt.Println("120 이상")
	} else if a >= 100 && a < 120 {
		fmt.Println("100 이상 120 미만")
	} else if a < 100 && a >= 50 {
		fmt.Println("50 이상 100 미만")
	} else {
		fmt.Println("50 미만")
	}
}
