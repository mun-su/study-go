package main

import (
	"fmt"
	"time"
)

func main() {
	// 고루틴(Gorouutine)
	exe("t1")
	fmt.Println("Main Routine Start : ", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second) // time.Second, Minute, Hour, Microsecond
	fmt.Println("Main Routine End : ", time.Now())
}

func exe(name string) {
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>", i)
	}
	fmt.Println(name, " end : ", time.Now())
}
