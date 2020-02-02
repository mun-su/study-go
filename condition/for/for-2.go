package main

import "fmt"

func main() {

	// ex 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1
	}

	fmt.Println("ex1 : ", sum1)

	// ex 2
	// 해외에서는 다음과 같이 for 안에서 i++ 과 같이 사용하여 가독성을 더 높이려함.
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
	}

	fmt.Println("ex2 : ", sum2)

	// ex 3
	// while 형태와 비슷하게 사용
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)

	// ex 4
	// 최초 i, j 선언
	// i <= 10 조건
	// i, j 할당
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}

	// ex 5 (error)
	// 후칙 연산은 return 값이 없기 때문에 다음과 같이 사용시 에러발생
	/*
	   for i, j := 0, 0; i <= 10; i, j = i++, j+10 {
	       fmt.Println("ex5 : ", i, j)
	   }
	*/
}
