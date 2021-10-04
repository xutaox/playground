package main

import (
	"fmt"
	"io"
	"reflect"
	"sort"
)

func main() {
	var a sort.IntSlice
	fmt.Println(reflect.TypeOf(a).String())  // sort.IntSlice
	fmt.Println(reflect.TypeOf(a).Name())    // IntSlice
	fmt.Println(reflect.TypeOf(a).Kind())    // slice
	fmt.Println(reflect.TypeOf(a).Elem())    // int
	fmt.Println(reflect.TypeOf(a).Method(0)) // {Len  func(sort.IntSlice) int <func(sort.IntSlice) int Value> 0}
	//fmt.Println(reflect.TypeOf(a).Field(0))  // panic
	fmt.Println(reflect.ValueOf(a))             // []
	fmt.Println(reflect.ValueOf(a).Interface()) // []
	var b io.Reader
	fmt.Println(reflect.TypeOf(b))  // <nil>
	fmt.Println(reflect.ValueOf(b)) //<invalid reflect.Value>

	genNewVarByType()
}

type fooType struct {
	Content string
}

// 怎么根据给定的类型, 生成一个新的这个类型的变量?
// https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
func genNewVarByType() {
	foo := fooType{Content: "hello"}
	var IFoo interface{}
	IFoo = foo // 抹去了类型信息

	// 两种方式
	newIFoo := reflect.New(reflect.TypeOf(IFoo)).Interface()
	newFoo := newIFoo.(*fooType)
	fmt.Println("newFoo:", newFoo) // &{}

	newIFoo2 := reflect.New(reflect.TypeOf(IFoo)).Elem().Interface()
	newFoo2 := newIFoo2.(fooType)
	fmt.Println("newFoo2:", newFoo2) // {}

	// reflect.New(Type) -> Value // 根据Type生成一个这个类型的*指针*的value
	fmt.Println("reflect.New:", reflect.New(reflect.TypeOf(IFoo)).String()) // *main.fooType value

	// reflect.ValueOf(interface{}) -> 拿到这个interface的value
	fmt.Println("Value of:", reflect.ValueOf(IFoo).Interface()) // {hello}
	//fmt.Println(reflect.ValueOf(IFoo).Elem().Interface()) // panic 因为不能对struct取elem, valueOf已经是从interface取到了struct

}

/*
	official doc: https://go.dev/blog/laws-of-reflection

*/
