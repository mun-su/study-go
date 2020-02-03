package main

import (
	"fmt"
	"section3/init/lib"
)

var num int32

func init() {
	num = 30
	fmt.Println("Init Start!")
}

func main() {
	fmt.Println("100보다 작은수? :", lib.CheckSum(num))
	/*
		실행된 순서.
		Lib Package Init Start!
		Init Start!
		100보다 작은수? : true
	*/
}
