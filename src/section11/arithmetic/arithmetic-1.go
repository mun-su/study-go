// 사용자 패키 작성 및 문서화 1 (godoc 확인 가능)
package arithmetic

// 두 숫자의 연산 계산 제공 구조체
type Numbers struct {
	X int
	Y int
}

// 덧셈
func (o *Numbers) Addition() int {
	return o.X + o.Y
}

// 뺄셈
func (o *Numbers) Subtraction() int {
	return o.X - o.Y
}

// 곱셈
func (o *Numbers) Multiply() int {
	return o.X * o.Y
}

// 나눗셈
func (o *Numbers) Division() int {
	return o.X / o.Y
}
