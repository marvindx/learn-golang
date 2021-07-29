package main

import (
	"fmt"
	"log"
	"os"
)

// init
func init() {
	log.SetOutput(os.Stdout)
}

func Array() {
	var arr [5]int
	fmt.Printf("%#v\n", arr)

	arr1 := [...]int{1, 2, 3, 4}
	fmt.Printf("%#v\n", arr1)

	arr2 := [...]int{1: 1, 2: 2}
	fmt.Printf("%#v\n", arr2)

	arr2[0] = 100
	fmt.Printf("%#v\n", arr2)

	a := 1
	fmt.Printf("a: %v\n", &a)

	arr3 := [3]*int{0: &a, 2: new(int)}
	fmt.Printf("%#v\n", arr3)

	*arr3[0] = 2
	fmt.Printf("%#v\n", arr3)
	fmt.Printf("a: %#v\n", a)

	arrStr := [...]string{"a", "b", "c", "d", "e"}
	fmt.Printf("arrStr: %#v\n", arrStr)

	var arrStr1 [5]string
	arrStr1 = arrStr // 值赋值
	fmt.Printf("arrStr1: %#v\n", arrStr1)

	fmt.Printf("arrStr: %p\n", &arrStr)   // 0xc0000b6050
	fmt.Printf("arrStr1: %p\n", &arrStr1) // 0xc0000b60f0

	var arrStr3 [3]*string
	arrStr4 := [3]*string{new(string), new(string), new(string)}
	fmt.Printf("arrStr3: %#v\n", arrStr3)
	fmt.Printf("arrStr4: %#v\n", arrStr4)

	*arrStr4[0] = "a"
	*arrStr4[1] = "b"
	*arrStr4[2] = "c"
	fmt.Printf("arrStr4: %#v\n", arrStr4)
	arrStr3 = arrStr4
	fmt.Printf("%p\n", &arrStr4)
	fmt.Printf("%p\n", &arrStr3)
	fmt.Printf("arrStr3: %#v\n", arrStr3)
	fmt.Println(*arrStr3[0])

	fmt.Printf("arrStr: %#v\n", arrStr)
	Swap(&arrStr)
	fmt.Printf("arrStr: %#v\n", arrStr)

	var arrStr5 [5]string
	fmt.Printf("%p\n", &arrStr5) // 0xc000066280
	arrStr6 := [5]string{}
	fmt.Printf("%p\n", &arrStr6) // 0xc0000662d0

}

func Swap(arr *[5]string) {
	fmt.Printf("%T\n", arr)
	fmt.Printf("%p\n", arr)
	//temp := arr[0]
	//arr[0] = arr[1]
	//arr[1] = temp

	//temp := arr
	//fmt.Printf("%p\n", temp)
	//temp[0] = "g"

	var temp *[5]string

	temp = arr
	fmt.Printf("%p\n", &temp)
	temp[0] = "g"
}

func Slice() {
	// make([]Type, len, cap)    cap >= len
	s := make([]string, 2, 5) // len=2 cap=5
	fmt.Println(len(s), cap(s))

	s1 := make([]string, 2) // len=cap=2
	fmt.Println(len(s1), cap(s1))

	s2 := [...]int{1, 2, 3}
	fmt.Println(len(s2), cap(s2)) // array: cap=len

	s3 := []int{1, 2, 3}
	fmt.Printf("%T\n", s2)
	fmt.Printf("%T\n", s3)

	s4 := []int{99: 99}
	fmt.Println(len(s4), cap(s4))

	var s5 []int           // nil slice
	fmt.Printf("%p\n", s5) // 0x0
	fmt.Println(len(s5), cap(s5))
	s5 = append(s5, 1)
	fmt.Printf("%p\n", s5) // 0xc0000160e8
	fmt.Println(len(s5), cap(s5))

	s6 := []int{}           // make([]int, 0)
	fmt.Printf("%p\n", &s6) //0x11923f0
	fmt.Println(len(s6), cap(s6))

	s7 := []int{1, 2, 4, 5, 7, 9}
	s8 := s7[1:3]
	fmt.Println(s8, len(s8), cap(s8)) // [2 4] 2 5
	s8[1] = 88
	fmt.Println(s8) // [2 88]
	fmt.Println(s7) // [1 2 88 5 7 9]

	s8 = append(s8, 99)
	fmt.Println(s8, len(s8), cap(s8)) // [2 88 99] 3 5
	fmt.Println(s7)                   // [1 2 88 99 7 9]

	s8 = append(s8, 100, 100)
	fmt.Println(s8, len(s8), cap(s8)) // [2 88 99 100 100] 5 5
	fmt.Println(s7)                   // [1 2 88 99 100 100]

	s8 = append(s8, 101)
	fmt.Println(s8, len(s8), cap(s8)) // [2 88 99 100 100 100] 6 10

	s9 := []int{1, 2, 3, 4}
	s10 := s9[1:3:3]
	fmt.Println(s10, len(s10), cap(s10)) // [2 3] 2 2
	s10 = append(s10, 5)
	fmt.Println(s10, len(s10), cap(s10)) // [2 3 5] 3 4
	fmt.Println(s9, len(s9), cap(s9))    // [1, 2, 3, 4] 4 4

	for index, value := range s10 {
		fmt.Printf("%d : %v\n", index, value)
	}

	for _, value := range s10 {
		fmt.Printf("%v\n", value)
	}

	for index := range s10 {
		fmt.Printf("%v\n", index)
	}

	s11 := make([]int, 1, 10)
	for index, value := range s11 {
		fmt.Println(index, value)
	}
}

func Map() {
	//m1 := new(map[int]int)
	//m2 := make(map[int]int)
	//fmt.Printf("%T\n", m1)
	//fmt.Printf("%T\n", m2)

	var m3 = make(map[int]int)
	m3[1], m3[2] = 1, 2
	fmt.Println(m3, len(m3))
	fmt.Printf("%#v\n", m3)

	value, exists := m3[5]
	fmt.Println(value, exists)

	for key, value := range m3 {
		fmt.Println(key, value)
	}

	delete(m3, 3)
	fmt.Println(m3, len(m3))
}

type user struct {
	name  string
	age   int
	email string
}

type Duration int64

func Type() {
	var dur Duration
	dur = Duration(1000)
	fmt.Printf("%d\n", dur)

	var u = user{
		name:  "marvin",
		age:   18,
		email: "gmail",
	}

	fmt.Printf("%#v\n", u)
	u.Haha()
	u.Update()
	fmt.Printf("%#v\n", u)

	u1 := &user{
		name:  "u1",
		age:   18,
		email: "gmail",
	}
	fmt.Printf("%#v\n", u1)
	u1.Haha()
	u1.Update()
	fmt.Printf("%#v\n", u1)

	m1 := myMap{}
	m1.TestMap()
	fmt.Printf("%#v\n", m1)
}

func (u user) Haha() {
	fmt.Println(u.name, "haha")
	//u.name = "zjg"
}

func (u *user) Update() {
	u.age = 23
}

func TestInt(i *int) {
	*i = 1
}

func TestString(s *string) {
	*s = "dsfd"
}

type myMap map[string]int

func (m myMap) TestMap() {
	m["a"] = 10
}

type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Println("user name:", u.name, " user email:", u.email)
	u.name = "new name"
}

func (ad *admin) notify() {
	fmt.Println("admin name:", ad.name, " admin email:", ad.email)
	//u.name = "new name"
}

func Interface() {
	var u = user{
		name:  "marvin",
		age:   18,
		email: "gmail",
	}
	sendN(&u)
	fmt.Println("user name:", u.name, " user email:", u.email)

	var ad = admin{
		user:  u,
		level: "high",
	}
	sendN(&ad)
	fmt.Printf("%#v\n", ad)

	var p = Person{
		user: &u,
	}
	p.Update()
	fmt.Printf("%#v\n", p.user)
	fmt.Printf("%#v\n", u)
	u.age = 1
	fmt.Printf("%#v\n", p.user)
	fmt.Printf("%#v\n", u)
}

func sendN(n notifier) {
	n.notify()
}

type admin struct {
	user
	level string
}

type Person struct {
	user *user
}

func (p *Person) Update() {
	p.user.age = 22
}

func main() {
	log.Println("starting!!!")

	//Slice()
	//Array()
	//Map()
	//Type()
	//Interface()

	//Info := mylog.FileLog()
	//
	//Info.Println("main unitTest info log...")
	//fmt.Sprintln("dfdfdf")
	//fmt.Println("dfddddfdf")

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3:3]
	fmt.Println(s2, len(s2), cap(s2))

}
