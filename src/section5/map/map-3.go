package main

import "fmt"

func main() {
	// 맵(Map)
	// 맵 값 변경 및 삭제

	// ex1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"me":     "http://mun-su.github.io",
	}
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	map1["home"] = "http://github.io/mun-su" // 추가
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	map1["home"] = "http://mun-su.github.io" // 수정
	fmt.Println("ex1 : ", map1)
	fmt.Println()

	// ex2 (삭제)
	delete(map1, "home")
	delete(map1, "google")

	fmt.Println("ex2 : ", map1)
	fmt.Println()
}
