package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 파일 읽기, 버퍼 사용 쓰기 -> bufio 패키지 활용
	// ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, Writer, Read 메소드를 사용 가능.

	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/
	// 즉, bufio 의 NewReader, NewWriter 를 통하여 객체 반환 후 메소드 사용.

	// bufio(Buffered io) 패키지
	// https://golang.org/pkg/bufio

	// 파일 열기
	// 두 번째 매개변수 확인
	// https://golang.org/pkg/os/#pkg-constant

	/*
		상태
		a ----> a
		b ----> ab
		c ----> abc
		d ----> abcd
		e ----> e      ----> abcd
		f ----> ef     ----> abcd
		g ----> efg    ----> abcd
		h ----> efgh   ----> abcd
		i ----> i      ----> abcdefgh
	*/

	// 없으면 생성하고, 있으면 Read, Write 가능하게
	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))

	// bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) // Writer 반환(버퍼 사용)
	wt.WriteString("Hello Golang!\n File Write Test1!\n")
	wt.Write([]byte("Hello Golang!\n File Write Test2!\n"))

	// 에러 체크
	errCheck1(err)

	fmt.Printf("사용한 Buffer Size : (%d bytes)\n", wt.Buffered())
	fmt.Printf("남은 Buffer Size : (%d bytes)\n", wt.Available())
	fmt.Printf("전체 Buffer Size : (%d bytes)\n", wt.Size())

	wt.Flush() // 버퍼를 비우고 반영 (버퍼의 내용을 디스크에 기록)
	fmt.Println("쓰기 작업 완료")
	fmt.Println("=========================================================")

	rt := bufio.NewReader(file) // Reader 반환
	fi, err := file.Stat()
	errCheck1(err)

	b := make([]byte, fi.Size())

	fmt.Println("파일 정보 출력 :", fi)
	fmt.Println("파일 이름 :", fi.Name())
	fmt.Println("파일 크기 :", fi.Size())
	fmt.Println("파일 수정 시간 :", fi.ModTime())
	fmt.Println("=========================================================")

	// constant 에 os.SEEK_SET 등이 있으나 deprecated 예정이므로 가이드 되어있는 io pkg 를 사용
	file.Seek(0, io.SeekStart)
	data, _ := rt.Read(b) // 읽기 (ReadLine, ReadByte, ReadBytes 등)
	// rt.Read(b)

	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)
	fmt.Println("=========================================================")
	fmt.Println(string(b))
}

func errCheck1(err error) {
	if err != nil {
		panic(err)
	}
}
