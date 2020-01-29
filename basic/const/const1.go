package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한번 선언 후 값 변경 불가, 고정 된 값 관리 용
	const a string = "a"
	const b = "b"
	const c int32 = 15 * 15

	// 다음과 같이 func return 값을 사용하게 되면 func 는 항상 같은 값을 보장하지 않기 때문에 에러가 발생함.
	// const d = getData()

	const e = 99.9
	const f = false

	// 에러 발생
	// const g string
	// g = "g"

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
}
