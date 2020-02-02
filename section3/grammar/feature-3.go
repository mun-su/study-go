package main

import "fmt"

func main() {
	// 후치(후위) 연산자 허용 (i++)
	// 전치(전위) 연산사 비허용 (++i) -> 사용 불가
	// 증감연산은 반환 값이 없음 sum : i++ -> 사용 불가.

	// ex 1
	// 문법적으로 후위 연산의 return 값은 없기 때문에 다음과 같이 작성 필요.
	sum, i := 0, 0
	for i <= 100 {
		// sum += i++
		sum += i
		i++
		// ++i error 발생
	}
	fmt.Println(sum)

}
