package main

import "fmt"

func main() {
	// 구조체 인스턴스 값 변경시 -> 포인터 전달, 보통의 경우 -> 값 전달.
	kim := Account3{"245-901", 1000000, 0.015}
	lee := Account3{"245-902", 2000000, 0.015}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println()

	kim.CalculateA(100000)
	lee.CalculateB(100000)

	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))
}

type Account3 struct {
	number   string
	balance  float64
	interest float64
}

func (a Account3) CalculateA(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account3) CalculateB(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}
