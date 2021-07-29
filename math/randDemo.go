package main

import (
	"fmt"
	"math/rand"
)

func randDemo() {
	// 0 <= n <= 100 的随机整数
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int())
}

func main() {
	randDemo()
}