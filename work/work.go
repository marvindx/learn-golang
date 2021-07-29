package main

import (
	"fmt"
	"sync"
)

type Person struct{}

func (p Person) Task() {
	for i := 0; i < 200; i++ {
		fmt.Println(i)
	}
}

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := &Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			defer p.wg.Done()
			for w := range p.work {
				w.Task()
			}
		}()
	}
	return p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

func main() {
	p := New(1)
	p.Run(Person{})
	//fmt.Println("run 1")
	p.Run(Person{})
	//fmt.Println("run 2")
	close(p.work)	// 关闭通道才会跳出通道阻塞，走到wg.Done()
	p.wg.Wait()

}
