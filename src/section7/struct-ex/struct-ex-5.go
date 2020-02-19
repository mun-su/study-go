package main

import "fmt"

func main() {
	// 구조체 임베디드 메소드 오버라이딩 패턴
	//
	// ex1
	// 직원
	ep1 := Employee1{"KIM", 20000000, 150000}
	ep2 := Employee1{"PARK", 15000000, 200000}

	// 임원
	ex := Executives1{
		Employee1{"LEE", 50000000, 15000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate1()))
	fmt.Println("ex1 : ", int(ep2.Calculate1()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("ex2 : ", int(ex.Calculate1())) // 오버라이딩 정확한 값 반환
	fmt.Println("ex2 : ", int(ex.Employee1.Calculate1()+ex.specialBonus))

	// fmt.Println("ex3 : ", int(ex.Calculate1()+ex.specialBonus)) // 오버라이딩  : 잘못 된 값 반환
}

// super clas
type Employee1 struct {
	name   string
	salary float64
	bonus  float64
}

// sub class
type Executives1 struct {
	Employee1    // is a 관계
	specialBonus float64
}

func (e Employee1) Calculate1() float64 {
	return e.salary + e.bonus
}

func (e Executives1) Calculate1() float64 {
	return e.salary + e.bonus + e.specialBonus
}
