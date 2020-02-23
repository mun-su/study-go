package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널 (Channel)
	// 함수 등의 매개변수에 수신 및 발신 전용 채널 지정 가능
	// 전용 채널 설정 후 방향이 다를 경우 예외 발생
	// 발신 전용 channel <- 데이터형
	// 수신 전용 <- channel
	// 매개 변수를 통해서 전용 채널 확인할 수 있다.
	// 채널 또한 함수의 반환 값으로 사용 가능.

	// ex1
	c := make(chan int)
	go sendOnly(c, 10) // 발신전용
	go receiveOnly(c)  // 수신전용

	time.Sleep(2 * time.Second)
}

func sendOnly(c chan<- int, cnt int) {
	for i := 0; i < cnt; i++ {
		c <- i
	}
	c <- 777

	// 발신 전용에서 수신을 시도 하는 경우 다음과 같은 패닉 발생
	// invalid operation: <-c (receive from send-only type chan<- int)
	// fmt.Println(<-c)
}

func receiveOnly(c <-chan int) {
	// 채널이 종료될때 까지 받는 방법.
	for i := range c {
		fmt.Println("received : ", i)
	}
}
