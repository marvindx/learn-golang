package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func demo() {
	runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	t1 := time.Now()

	go func() {
		defer wg.Done()

		for i := 0; i < 3000; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 3000; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	wg.Wait()

	t2 := time.Now()
	fmt.Println()
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("time use:", t2.Sub(t1))

}

func RaceCanditionOne() {
	var (
		counter int
		wg      sync.WaitGroup
	)
	runtime.GOMAXPROCS(2)
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 30; i++ {
			value := counter
			value++
			counter = value
			fmt.Println("counter of func1:", counter)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 30; i++ {
			value := counter
			value++
			counter = value
			fmt.Println("counter of func2:", counter)
		}
	}()

	wg.Wait()

	fmt.Println("counter of end:", counter)

}

func RaceCandition() {
	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(2)

	runtime.GOMAXPROCS(1)

	incCounter := func(id int) {
		defer wg.Done()

		for count := 0; count < 2; count++ {
			//value := counter
			//runtime.Gosched()
			//value++
			//counter = value
			atomic.AddInt64(&counter, 1)
		}
	}

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

//func main() {
//	//demo()
//	//RaceCanditionOne()
//	RaceCandition()
//}
