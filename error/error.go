package main

import (
	"errors"
	"fmt"
)

var (
	ErrorOne = errors.New("error one unitTest")
	ErrorTwo = errors.New("error two unitTest")
	ErrorFmt = fmt.Errorf("fmt error format %%")
)

// panicTest
func panicTest() {
	defer func() {
		fmt.Println("defer after panic")
		fmt.Println(recover()) // recover() 捕获异常
	}()

	if ErrorOne != nil {
		fmt.Printf("error: %v\n", ErrorOne)
		panic(ErrorOne) // raise `ErrorOne` exception
	}
}

// recoverTest
func recoverTest() {
	defer func() {
		//recover()
		fmt.Println(recover()) // recover() 捕获异常
	}()

	if ErrorTwo != nil {
		panic(ErrorTwo) // raise `ErrorOne` exception
	}
}

func main() {
	//panicTest()
	recoverTest()
}
