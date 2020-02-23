package main

import "fmt"

func main() {
	// 채널(Channel)

	// ex1
	c := receiveOnly1(101) // 채널 반환
	output := total(c)     // 채널 전달 후 반환
	// output <- 777 // 예외 -> invalid operation: output <- 777 (send to receive-only type <-chan int)
	fmt.Println("ex1 : ", <-output)
}

// 수신 전용
func receiveOnly1(cnt int) <-chan int {
	sum := 1
	tot := make(chan int)
	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

// 채널을 받아서 처리.
func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()
	return tot
}
