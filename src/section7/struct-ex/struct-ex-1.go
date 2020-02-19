package main

import "fmt"

func main() {
	// 구조체 생성자 패턴 예제

	// ex1
	kim := Account{number: "245-901", balance: 100000, interest: 0.015}

	var lee *Account = new(Account)
	lee.number = "245-902" // getter, setter
	lee.balance = 130000
	lee.interest = 0.025

	fmt.Println("ex1 :", kim)
	fmt.Println("ex1 :", lee)

	// ex2
	park := NewAccount("245-903", 170000, 0.04)
	fmt.Println("ex2 :", park)
}

type Account struct {
	number   string
	balance  float64
	interest float64
}

// 생성자 패턴
func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환 아닌 경우 값 복사
	return &Account{number, balance, interest} // 구조체 인스턴스를 생성한 뒤 포인터 리턴
}
