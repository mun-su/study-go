package main

import "fmt"

func main() {
	// 다양한 선언 방법
	// &struct, struct : &struct 포인터를 받아오며 역참조를 또 하기 때문에 속도는 조금느리다.
	// Go 인터페이스 메소드를 선언만 해둔 후 -> 오버라이딩 해서 메소드에 포인터 리시버를 사용할 경우 반드시 &struct 로 만들어야한다.
	// structType does not implement Interface (method has pointer receiver)

	// 선언 방법1
	var seo *Account1 = new(Account1)
	seo.number = "245-901"
	seo.balance = 10000000
	seo.interest = 0.015

	// 선언 방법2
	hong := &Account1{number: "245-902", balance: 150000, interest: 0.03}

	// 선언 방법3
	lee := new(Account1)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	// ex1
	fmt.Println("ex1 :", seo)
	fmt.Println("ex1 :", hong)
	fmt.Println("ex1 :", lee)
	fmt.Println()

	// ex2
	fmt.Printf("ex2 : %#v\n", seo)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)
	fmt.Println()

	// ex3
	fmt.Println("ex1 :", int(seo.Calculate()))
	fmt.Println("ex1 :", int(hong.Calculate()))
	fmt.Println("ex1 :", int(lee.Calculate()))
}

func (a *Account1) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

type Account1 struct {
	number   string
	balance  float64
	interest float64
}
