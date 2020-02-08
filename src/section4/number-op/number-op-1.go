package main

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 연산(산술, 비교)
	// 타입이 같아야 가능. (정수, 실수) 엄격함.
	// 다른 타입끼리는 반드시 형 변환 후 연산
	// 형 변환 없을 경우 예외(에러) 발생
	// +, -, *, %, /, <<, >>, &, ^

	// ex 1 unsigned int
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex 1 : ", n1)
	fmt.Println("ex 1 : ", n2)
	fmt.Println("ex 1 : ", n3)
	fmt.Println("ex 1 : ", n4)

	// 일반 int는 unit에 비해 절반
	fmt.Println("ex 1 : ", math.MaxInt8)
	fmt.Println("ex 1 : ", math.MaxInt16)
	fmt.Println("ex 1 : ", math.MaxInt32)
	fmt.Println("ex 1 : ", math.MaxInt64)

	// Float
	fmt.Println("ex 1 : ", math.MaxFloat32)
	fmt.Println("ex 1 : ", math.MaxFloat64)

	n5 := 100000       // int
	n6 := int16(10000) // 형변환
	n7 := uint8(100)

	// 사용 형태가 틀리기 때문에 예외 발생
	// fmt.Println("ex2 :  ", n5+n6)
	fmt.Println("ex2 : ", n5+int(n6))

	// 사용 형태가 틀리기 때문에 예외 발생
	// fmt.Println("ex2 : ", n6+n7)
	fmt.Println("ex2 : ", n6+int16(n7))

	fmt.Println("ex2 : ", n6 > int16(n7))
	fmt.Println("ex2 : ", n6-int16(n7) > 5000)
}
