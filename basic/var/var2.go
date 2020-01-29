package main

import "fmt"

func main() {

	// 한번에 여러개 선언
	var (
		zipCode       string = "123456"
		address       string = "서울시 강서구"
		addressDetail string = "마곡동"
		expose        bool   = false
	)

	zipCode = "654321"
	address = "서울시 강서구 마곡서로"
	addressDetail = "xxx호"
	expose = true

	fmt.Println("zipCode : ", zipCode, "address : ", address, "addressDetail : ", addressDetail, "expose : ", expose)

}
