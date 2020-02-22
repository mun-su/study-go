package main

import "fmt"

func main() {
	// 채널(Channel)
	// Close : 채널 닫기

	// ex1
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good!"
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("ex1 : ", val1, ok1)

	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)

	val3, ok3 := <-ch
	fmt.Println("ex1 : ", val3, ok3)

	close(ch) // 채널 닫기.

	val4, ok4 := <-ch
	fmt.Println("ex1 : ", val4, ok4) // string 기본값 공백과 성공여부 false
}
