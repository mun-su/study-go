package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	// 에러 처리 고급
	// error 타입이 아닌 경우 에러 처리 방법
	// Error 메소드를 구현해서 사용자 정의 에러 처리 예제 심화
	// 구조체를 사용해서 세부적인 정보 출력

	// ex1
	v, err := Pow(10, 3) // 정상 처리
	if err != nil {
		log.Fatal(err)
		// log.Fatal(err.Error())
	}
	fmt.Println("ex1 : ", v)

	// ex2
	t, err := Pow(0, 3)
	if err != nil {
		log.Fatal(err)
		// log.Fatal(err.Error())
		// fmt.Println(err.(PowError).value)
		// fmt.Println(err.(PowError).message)
		// fmt.Println(err.(PowError).time)
	}
	fmt.Println("ex2 : ", t)
}

// 예외(에러) 처리 구조체
type PowError struct {
	time    time.Time // 에러 발생 시간
	value   float64   // 파라미터
	message string    // 에러 메시지
}

func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g) - %s", e.time, e.value, e.message)
}

func Pow(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time.Now(), f, "0은 사용할 수 없습니다."}
	}

	// 실제로는 func Pow 의 parameter 가 숫자만 받기 때문에 발생하진 않으나 예제를 위하여 사용.
	if math.IsNaN(f) {
		return 0, PowError{time.Now(), f, "숫자가 아닙니다."}
	}
	if math.IsNaN(i) {
		return 0, PowError{time.Now(), i, "숫자가 아닙니다."}
	}
	return math.Pow(f, i), nil
}
