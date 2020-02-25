package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 고급
	// Once : 한번만 실행(주로 초기화에 사용)
	// Do로 실행

	// ex1
	runtime.GOMAXPROCS(runtime.NumCPU())
	once := new(sync.Once) // Once 객체 생성

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Goroutine : ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(3 * time.Second)
}

func onceTest() {
	// 이 부분에 한 번 실행 할 코드 작성
	fmt.Println("Once test Execute!")
}
