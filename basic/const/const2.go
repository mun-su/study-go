package main

import "fmt"

func main() {
	const a, b int = 5, 10
	const c, d, e, f = true, 99.9, "e", 9

	const (
		i, j int16 = 72, 79
		x, y       = "i", 9999
	)

	fmt.Println("a : ", a, ", b : ", b)
	fmt.Println("c : ", c, ", d : ", d, ", e : ", e, ", f : ", f)
	fmt.Println("i : ", i, ", j : ", j)
	fmt.Println("x : ", x, ", y : ", y)
}
