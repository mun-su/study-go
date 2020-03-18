// 사용자 패키 작성 및 문서화 2 (godoc 확인 가능)
package arithmetic

// 제곱의 합산
func (o *Numbers) SquaredAddition() int {
	return (o.X * o.X) + (o.Y * o.Y)
}

// 제곱의 차
func (o *Numbers) SquaredSubtraction() int {
	return (o.X * o.X) - (o.Y * o.Y)
}
