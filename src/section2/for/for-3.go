package main

import "fmt"

func main() {

	// ex 1
	// break 는 j loop 를 벗어나며 2, 3을 출력하고 2, 4를 패스 후 3,0 을 출력하게됨.
	// ex1 : 2 3
	// ex1 : 3 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				// break
				break
			}
			fmt.Println("ex1 : ", i, j)
		}
	}

	// ex 2 -> label 사용한 break
	// label 다음에는 반드시 for 를 사용해야함.
	// break Loop1 은 j loop 를 벗어나는 것이 아닌
	// 지정한 곳의 loop를 벗어나게 되므로 2, 3을 출력하고 종료됨.
	// ex2 : 2 3
	// 종료
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				// break and go to Loop1
				break Loop1
			}
			fmt.Println("ex2 : ", i, j)
		}
	}

	// ex 3
	// continue 를 사용하여 2로 나눠지는 수를 건너뛰고 출력
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex3 : ", i)
	}

	// ex 4 -> label 사용시 error
	// label 다음에는 반드시 for 를 사용해야함.
	/*
	   Loop2:
	   if i := 1; i < 2 {
	       fmt.Println("ex4")
	   }
	*/

	// ex 5 -> label 사용한 continue
	// label 다음에는 반드시 for 를 사용해야함.
	// 1, 2 일때 continue 로 i loop 로 돌아갔기 때문에 다음과 같이 출력됨
	// 1 0
	// 1 1
	// 2 0
	// 2 1
Loop3:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop3
			}
			fmt.Println("ex5 : ", i, j)
		}
	}
}
