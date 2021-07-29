package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
}

var pool *sync.Pool

func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new Person")
			return new(Person)
		},
	}
}

func main() {
	p1 := pool.Get().(*Person)
	fmt.Printf("p type: %#v\n", p1)

	var (
		p2 = &Person{Name: "p2"}
		p3 = &Person{Name: "p3"}
	)

	pool.Put(p2)
	pool.Put(p3)
	fmt.Printf("p name: %v\n", p2.Name)
	fmt.Printf("p name: %v\n", pool.Get().(*Person).Name)
	fmt.Printf("p name: %v\n", pool.Get().(*Person).Name)

}
