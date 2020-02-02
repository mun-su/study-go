package main

import "fmt"

func main() {
	// 코드 서식 지정
	// 한 사람이 코딩한것과 동일한 일관성을 유지하게 해줌.
	// 다른 언어를 사용시 일부 IDE에 지정하여 사용할수 있으나 golang은 별다른 플러그인이 아닌 gofmt 강제 됨.
	// 터미널에서 실행 가능함
	// gofmt -h : 사용법
	// gofmt -w : 원본 파일에 반영

	// ex 1
	// for i:=0;i<=100;i++{
	// 	fmt.Println("ex : ",i)
	// }
	//
	// 위와 같이 작성 후 터미널에서 "gofmt -w ./feature-4.go" 실행하면 자동으로 다음과 같이 정렬됨.
	for i := 0; i <= 100; i++ {
		fmt.Println("ex : ", i)
	}

	// Jetbrain Goland 사용시에는 다음과 같이 사용 할 수 있음.
	// Preferences > Tools > File Watchers > go fmt 추가.
}
