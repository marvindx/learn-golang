package main

import (
	"fmt"
	"sync"
)

// BufferedChan buffered channel demo
func BufferedChan() {
	var (
		chanl = make(chan int, 10)
		wg    sync.WaitGroup
	)
	wg.Add(2)

	get := func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			v := <-chanl
			fmt.Println("get from chanl", v)
		}
	}

	set := func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			chanl <- i
			fmt.Println("set into chanl", i)
		}
	}
	go get()
	go set()

	wg.Wait()
}

// NoBufferedChan no buffered channel demp
func NoBufferedChan() {
	var (
		chanl = make(chan int)
		wg    sync.WaitGroup
	)
	wg.Add(2)

	get := func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			v := <-chanl
			fmt.Println("get from chanl", v)
		}
	}

	set := func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			chanl <- i
			fmt.Println("set into chanl", i)
		}
	}
	go get()
	go set()

	wg.Wait()
}

// 通道选择器
func selectChanl() {
	var (
		wg sync.WaitGroup
	)
	wg.Add(2)

	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		defer wg.Done()
		for n := 0; n < 100; n++ {
			if n%2 == 0 {
				chan1 <- n
			} else {
				//chan2 <- n
			}
		}

	}()

	go func() {
		defer wg.Done()
		for n := 0; n < 100; n++ {
			if n%2 == 0 {
				//chan1 <- n
			} else {
				chan2 <- n
			}
		}

	}()

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		select {
		case n := <-chan1:
			fmt.Println("偶数： ", n)
		case n := <-chan2:
			fmt.Println("奇数： ", n)
		}
	}
	wg.Wait()

}

func main() {
	//BufferedChan()
	//NoBufferedChan()
	selectChanl()
}
