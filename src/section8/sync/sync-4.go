package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 뮤텍스 : 상호 배제 -> Thread(고루틴)들이 서로 running time 에 서로 영향을 주지 않게 즉, 단독으로 실행되게 하는 기술.
	// 여러 고루틴에서 작업하는 공유 데이터 보호
	// RWMutex : 쓰기 Lock -> 쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 X, 읽기 락, 쓰기 락 전부 방어(방지)
	// RMutex : 읽기 Lock -> 읽기 시도 중에 값이 변경 방지 즉, 쓰기 락 방어(방지)

	// 동기화 사용
	// 쓰기 읽기 동작 순서가 일정

	// ex1

	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex) // var mutex = new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.Lock() // 쓰기 뮤텍스 잠금
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock() // 쓰기 뮤텍스 잠금 해제
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock() // 읽기 뮤텍스 잠금
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock() // 쓰기 뮤텍스 잠금 해제
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock() // 읽기 뮤텍스 잠금
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock() // 쓰기 뮤텍스 잠금 해제
		}
	}()

	time.Sleep(10 * time.Second)

	// 잠금과 해제는 반드시 쌍으로 구성되어야 함.
	// Lock <-> Unlock
	// RLock <-> RUnlock
}
