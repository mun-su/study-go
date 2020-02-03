package main

// 방법 1
/*
import "fmt"
import "os"
*/

// 방법 2 -> 아래의 방법을 대체로 많이 사용
import (
	"fmt"
	"os"
)

func main() {
	// Package : 코드 구조화(모듈) 및 재사용
	// 응집도, 결합도
	// Golang : 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그램을 작성 할 것을 권고.
	// 패키지 이름 == 디렉토리 이름
	// 같은 패키지 내 소스파일들은 디렉토리 명을 패키지 명으로 사용.
	// 네이밍 규칙 : 소문자 -> private 이며 지역변수, 대문자 -> public 이며 전역함수
	// main 패키지는 특별하게 인식 -> 컴파일러는 공유 라이브러리 X, 프로그램의 시작점으로 인식.

	var name string

	fmt.Println("이름 : ")

	// %s 는 문자열로 입력 받겠다는 의미
	fmt.Scanf("%s", &name)

	// 아래와 같이 name을 바로사용하여 포인트로 쓸 수 있지만 이번에는 생략
	// fmt.Scanf("%s", name)

	// command line 에 출력하기 위해 os 를 사용.
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)

	// terminal 에서 go run package-1.go
}
