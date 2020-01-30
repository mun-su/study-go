package main

// 다건의 import 가능
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 아래와 같이 다양하게 사용 가능.
	// random 사용
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50이상 100미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25이상 50미만")
	default:
		fmt.Println("i -> ", i, " 기본값")
	}
}
