package main

import (
	"fmt"
	"math"
)

func main() {
	// ex 1 (overflow error : 범위 초과)
	var n1 uint8 = math.MaxUint8 + 1
	var n2 uint16 = math.MaxUint16 + 1
	var n3 uint32 = math.MaxUint32 + 1
	var n4 uint64 = math.MaxUint64 + 1

	fmt.Println("ex : ", n1)
	fmt.Println("ex : ", n2)
	fmt.Println("ex : ", n3)
	fmt.Println("ex : ", n4)

	// ex 2 (overflow error : 범위 미만)
	// unsigned int
	var n5 uint8 = -1
	var n6 uint16 = -1
	var n7 uint32 = -1
	var n8 uint64 = -1

	fmt.Println("ex : ", n5)
	fmt.Println("ex : ", n6)
	fmt.Println("ex : ", n7)
	fmt.Println("ex : ", n8)
}
