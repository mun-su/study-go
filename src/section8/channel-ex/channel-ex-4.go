package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널(Channel)
	// 채널 Select 구문 -> 채널에 값이 수신되면 해당 case 문 실행
	// 일회성 구문이므로, for(반복문) 안에서 수행
	// default 구문 처리 주의.

	// ex1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 주로 이런 패턴은 채팅에서 이미지, 일반 문자, 비밀문자 와 같은 것을 받을때 구분하여 처리 할수 있도록 사용.
	go func() {
		for {
			select {
			case num := <-ch1: // 수신
				fmt.Println("ch1 : ", num)
			case str := <-ch2: // 수신
				fmt.Println("ch2 : ", str)
				// 채널 select 수신용도로 사용하는 경우 default 를 사용하지 않아야함.
				// 값이 도착하기 전에 채널을 받아오기 전에 실행되면 무조건 default 이기 때문에
				// default:
				// 	fmt.Println("default test")
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
