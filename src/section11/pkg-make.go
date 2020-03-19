package main

import (
	"fmt"
	oper "section11/arithmetic" // alias 사용 (패키지 중복 또는 약자로 사용
	// "section11/arithmetic" // alias 미사용
)

func main() {
	// 사용자 패키지 작성 & Go 문서화
	// 기준은 GOPATH/src
	// 패키지 폴더명(디렉토리 명) 명확하게 지정
	// 패키지 파일의 package 이름으로 사용한다. -> 길면 alias
	// package main 을 제외하고 package 문서에 등록
	// 지금까지 패키지 구조로 이미 사용해 왔다.
	// 기존적으로 GOROOT 의 패키지(pkg) 검색 -> GOPATH 의 패키지 (src/pkg) 검색
	// go install 명령어 실행의 경우 -> GOPATH/pkg 에 등록 (문서화)
	// godoc -http=:6060(임의의 포트) -> pkg 이동 -> 본인 패키지 메소드 및 주석 확인 (패키지, 타입, 메소드) 주석
	// golang 최신버전에는 기본적으로 godoc 가 설치 되지 않는 경우가 존재하므로 `go get golang.org/x/tools/cmd/godoc` 명령어를 사용하여 설

	// 패키지 사용 예제 (사칙연산)
	nums := oper.Numbers{100, 10}
	fmt.Println("Package Used(1) :", nums.Addition())
	fmt.Println("Package Used(2) :", nums.Subtraction())
	fmt.Println("Package Used(3) :", nums.Multiply())
	fmt.Println("Package Used(4) :", nums.Division())
	fmt.Println("Package Used(5) :", nums.SquaredAddition())
	fmt.Println("Package Used(6) :", nums.SquaredSubtraction())

	// 다음과 같이 하여 주석등 문서화를 볼 수 있다.
	// go install
	// godoc -http=:6060
	// 브라우저에서 localhost:6060 으로 접속
	// 아래 Third Party 에 작성한 코드 및 주석들이 보이는 것을 알 수 있다.
	// 파일의 제일 상단에 작성 된 주석은 package 설명
	// struct 및 func 상단에 주석하는 경우 해당 항목에 대한 설명
}
