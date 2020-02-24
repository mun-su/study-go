package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 고루틴 동기화 예제
	// 실행 흐름 제어 및 변수 동기화 가능
	// 공유 데이터 보호가 가장 중요.
	// 뮤텍스(Mutex)
	// sync.Mutex 선언 후 Lock, Unlock 사용.

	// ex1
	// 동기화 사용 예제
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := count{number: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // cpu 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()
}

// 구조체 선언(공유 데이터)
type count struct {
	number int
	mutex  sync.Mutex
}

func (c *count) increment() {
	c.mutex.Lock() // 공유 데이터 수정 전 뮤텍스로 보호
	c.number += 1
	c.mutex.Unlock() // 공유 데이터 수정 후 보호 해제
}

func (c *count) result() {
	fmt.Println(c.number)
}
