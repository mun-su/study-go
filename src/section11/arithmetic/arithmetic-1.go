// 사용자 패키 작성 및 문서화 예제 2 (godoc 확인 가능)
package arithmetic

// 두 숫자의 연산 계산 제공 패키지
type Numbers struct {
	X int
	Y int
}

func (o *Numbers) Addition() int {
	return o.X + o.Y
}

func (o *Numbers) Subtraction() int {
	return o.X - o.Y
}

func (o *Numbers) Multiply() int {
	return o.X * o.Y
}

func (o *Numbers) Division() int {
	return o.X / o.Y
}
