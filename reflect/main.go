package main

import (
	"fmt"
	"reflect"
)

type intAlias int

type Cat struct {
	Name string
	Age  int `json:"age"`
}

func reflectTypeDemo() {

	var (
		i  int      = 1
		ia intAlias = 2
	)

	typeOfI := reflect.TypeOf(i)
	typeOfIa := reflect.TypeOf(ia)

	// Name() 表层类型名
	// Kind() 类型实际种类
	// Align() 变量内存中占用字节数
	fmt.Println(typeOfI.Name(), typeOfI.Kind(), typeOfI.Align())    // int int 8
	fmt.Println(typeOfIa.Name(), typeOfIa.Kind(), typeOfIa.Align()) // intAlias int 8

	c := Cat{
		Name: "kitty",
		Age:  2,
	}
	typeOfC := reflect.TypeOf(c)
	// Cat type name:  Cat , Cat type kind:  struct
	fmt.Println("Cat type name: ", typeOfC.Name(), ", Cat type kind: ", typeOfC.Kind())

	cPtr := &Cat{}
	typeOfCptr := reflect.TypeOf(cPtr)
	// Cat Ptr type name: , Cat Ptr type kind: ptr
	fmt.Printf("Cat Ptr type name: %v, Cat Ptr type kind: %v\n", typeOfCptr.Name(), typeOfCptr.Kind())

	// 对指针类型进行类型反射, Elem()进行了取指针的操作
	typeOfCptr = typeOfCptr.Elem()
	// Cat Ptr type name: Cat, Cat Ptr type kind: struct
	fmt.Printf("Cat Ptr type name: %v, Cat Ptr type kind: %v\n", typeOfCptr.Name(), typeOfCptr.Kind())

	// 获取结构体内部的元素的类型信息
	/*
		name: Name, tag:
		name: Age, tag: json:"age"
	*/
	for i := 0; i < typeOfC.NumField(); i++ {
		// 获取内部元素的类型，返回 StructField
		fieldType := typeOfC.Field(i)
		fmt.Printf("name: %v, tag: %v\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名找到字段类型信息，返回 StructField
	// Cat age tag of json:  age
	if catAge, ok := typeOfC.FieldByName("Age"); ok {
		fmt.Println("Cat age tag of json: ", catAge.Tag.Get("json"))
	}

	// 另一种办法：获取结构体内部元素的kind
	typeOfCName := reflect.TypeOf(c.Name)
	// Cat's Name type:  string , Cat's Name kind:  string
	fmt.Println("Cat's Name type: ", typeOfCName.Name(), ", Cat's Name kind: ", typeOfCName.Kind())
}

func reflectValueDemo() {
	var (
		i int = 1
	)

	valueOfI := reflect.ValueOf(i)
	// 获取interface{} 类型的值，再进行类型断言
	fmt.Println(valueOfI.Interface().(int))
	// 转为 int64类型的值
	fmt.Println(valueOfI.Int())
	// int64强制转为int
	fmt.Println(int(valueOfI.Int()))

}

func main() {
	//reflectTypeDemo()
	reflectValueDemo()
}
