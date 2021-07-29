package main

import "fmt"

// threePointOne 接受不定个数的参数 等同于切片
func threePointOne(s ...string) {
	fmt.Printf("...string type: %T\n", s) // []string
	fmt.Println(s)
}

func main() {
	threePointOne()

	s := []string{"1", "2", "3"}
	// 切片打散进行传递
	threePointOne(s...)

	// 标识未知元素数组的个数
	arr := [...]int{1, 2, 3}
	arr1 := [3]int{1, 2, 3}
	fmt.Println("arr length", len(arr))
	fmt.Println("arr1 length", len(arr1))
}
