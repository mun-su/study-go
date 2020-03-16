package main

import (
	"fmt"
	"os"
)

func main() {
	// Go panic & Recover 최종 실습
	fileOpen("text.txt")
	fmt.Println("End Main")
}

func fileOpen(filename string) {
	// defer 함수 (Panic 호출 되면 실행)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error :", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1  :", f.Name())
	}

	defer f.Close()
}
