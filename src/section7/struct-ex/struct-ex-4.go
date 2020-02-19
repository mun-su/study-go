package main

import "fmt"

func main() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재사용 하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	// ex1
	// 직원
	ep1 := Employee{"KIM", 20000000, 150000}
	ep2 := Employee{"PARK", 15000000, 200000}

	// 임원
	ex := Executives{
		Employee{"LEE", 50000000, 15000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	// ex.Employee.Calculate() 로 접근하지 않아도 구조체가 알아서 접근하여 사용 가능.
	// 이런 형식을 Golang에서 메소드 재사용이라 한다.
	fmt.Println("ex2 : ", int(ex.Calculate()+ex.specialBonus))
}

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

type Executives struct {
	Employee
	specialBonus float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}
