package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 채널(Channel)
	// ex1 (비동기 : 버퍼 사용)
	// 버퍼 : 발신 -> 가득차면 대기, 비어있으면 작동.
	//       수신 -> 비어있으면 대기, 가득차있으면 작동.

	runtime.GOMAXPROCS(4)

	// 채널의 버퍼 용량.
	ch := make(chan bool, 4)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

	// 순서가 맞지 않는 경우는 단순히 콘솔이 먼저 찍힌경우에 불과함.
}
