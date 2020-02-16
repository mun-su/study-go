package main

import "fmt"

func main() {
	// 리시버(구조체 메소드) 전달(값, 참조) 형식
	// 함수는 기볹적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달
	// 리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능 가능
	cart1 := cart{3, 5000}
	fmt.Println("ex1(totPrice) :", cart1.purchase())

	// 참조 전달 (원본 값 수정 O)
	cart1.rePurchase1(7, 5000)
	fmt.Println("ex1(totPrice) :", cart1.purchase())

	// 값 전달(원본 값 수정 X)
	cart1.rePurchase2(10, 0)
	fmt.Println("ex1(totPrice) :", cart1.purchase())
}

type cart struct{ quantity, price int }

func (b cart) purchase() int {
	return b.quantity * b.price
}

// 원본 수정 O (참조 전달 방식)
func (b *cart) rePurchase1(quantity, price int) {
	b.quantity += quantity
	b.price += price
}

// 원본 수정 X (값 전달 형식)
func (b cart) rePurchase2(quantity, price int) {
	b.quantity += quantity
	b.price += price
}
