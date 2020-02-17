package main

import "fmt"

func main() {
	// 구조체
	// Go의 필드들의 집합체 또는 컨테이너
	// 필드는 갖지만 메소드는 갖지 않는다.
	// 객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	// 상속, 객체, 클래스 개념 없음.
	// 구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	// ex1
	seo := Account{number: "111-2222", balance: 10000000, interest: 0.015}
	lee := Account{number: "245-9122", balance: 2000000}
	park := Account{number: "125-2211", interest: 0.025}
	cho := Account{"521-1412", 150000, 0.03}

	fmt.Println("ex1 : ", seo)  // ex1 :  {111-2222 1e+07 0.015}
	fmt.Println("ex1 : ", lee)  // ex1 :  {245-9122 2e+06 0}
	fmt.Println("ex1 : ", park) // ex1 :  {125-2211 0 0.025}
	fmt.Println("ex1 : ", cho)  // ex1 :  {521-1412 150000 0.03}
	fmt.Println()

	// ex2
	// 보기편한 지수로 변경하기 위해 int로 치환.
	fmt.Println("ex2 : ", int(seo.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(park.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

type Account struct {
	number   string
	balance  float64
	interest float64
}
