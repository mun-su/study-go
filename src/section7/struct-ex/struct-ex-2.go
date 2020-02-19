package main

import "fmt"

func main() {
	kim := Account2{"245-901", 1000000, 0.015}
	lee := Account2{"245-902", 2000000, 0.015}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))
}

type Account2 struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account2) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account2) {
	a.balance = a.balance + (a.balance * a.interest)
}
