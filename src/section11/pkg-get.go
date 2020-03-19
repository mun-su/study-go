package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// 외부 저장소 패키지 설치
	// 2가지 방법
	// 1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	// 2. go get 패키지 주소 설치 -> 선언.

	// go get github.com/360EntSecGroup-Skylar/excelize

	// 선언 후 go get 설치 예제 (엑셀 파일 읽기)
	xfile := "sample.xlsx"
	xlfile, err := excelize.OpenFile(xfile)

	if err != nil {
		panic("Excel Loads Error!")
	}

	rows, err := xlfile.GetRows("Sheet1")
	for _, row := range rows {
		for _, cell := range row {
			fmt.Print(cell, "\t")
		}
		fmt.Println()
	}
}
