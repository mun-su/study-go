package main

import "fmt"

func main() {
	// ex1
	i := 7
	p := &i

	// &, *, 역참조에 대해 정확하게 이해해야함.
	fmt.Println("ex1 : ", i, *p, &i, p)

	// 주소값을 연산한것이 아니기 때문에 정상적으로 수행됨.
	*p++ // 1증가

	fmt.Println("ex1 : ", i, *p, &i, p)

	*p = 777 // 포인터 변수 역 참조 값 변경

	fmt.Println("ex1 : ", i, *p, &i, p)

	i = 77

	fmt.Println("ex1 : ", i, *p, &i, p)
	// 참조 값을 바꾸더라도 메모리 주소가 바뀌진 않는다.
}
