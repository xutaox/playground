package main

import (
	"encoding/json"
	"fmt"
	pincode "playground/protobuf/servicepb"
)

func main()  {
	testIDL()
	fmt.Println("")
	testIDL()
}

// 测试序列化后的编码，然后用新的增加了字段的IDL去解析
func testIDL() {

	b := pincode.TestA{
		Nest: &pincode.TestB{
			Num:  13,
			Num2: 14,
		},
		Num: 15,
	}

	data, err := b.Marshal()
	if err != nil {
		panic(err)
	}

	fmt.Println("marshal old, len:", len(data))
	fmt.Printf("%b\n", data)

	p := pincode.TestC{}
	err = p.Unmarshal(data)
	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("unmarshal with new: %s\n", j)

	p = pincode.TestC{Nest:&pincode.TestD{}}
	p.Nest.Num = 12
	p.Nest.Num2 = 13
	p.Nest.Num3 = 14
	p.Num = 15

	data, err = p.Marshal()
	if err != nil {
		panic(err)
	}

	fmt.Println("marshal new, len:", len(data))
	fmt.Printf("%b\n", data)

	b = pincode.TestA{}
	err = b.Unmarshal(data)
	if err != nil {
		panic(err)
	}

	j, err = json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("unmarshal with old: %s\n", j)

}

/*
marshal old, len 8
[1|010 | 100 | 1|000 | 1101 | 10|000 | 1110 | 10|000 | 1111]
unmarshal with new: {"nest":{"num":13,"num2":14},"num":15}
marshal new, len: 10
[1|010 | 110 | 1|000 | 1100 | 10|000 | 1101 | 11|000 | 1110 | 10|000 | 1111]
unmarshal with old: {"nest":{"num":12,"num2":13},"num":15}
证明新增字段不影响protobuf解析
*/

// 测试oneof
func testOneof() {
	b := pincode.TestOneof{
		Num:   15,
		Union: &pincode.TestOneof_First{First: 15},
	}

	data, err := b.Marshal()
	if err != nil {
		panic(err)
	}

	fmt.Println("one of first, len:", len(data))
	fmt.Printf("%b\n", data)

	b = pincode.TestOneof{
		Num:   15,
		Union: &pincode.TestOneof_Second{Second: "abc"},
	}

	data, err = b.Marshal()
	if err != nil {
		panic(err)
	}

	fmt.Println("one of second, len:", len(data))
	fmt.Printf("%b\n", data)

	c := pincode.TestOneof{}
	err = c.Unmarshal(data)
	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("unmarshal second: %s\n", j)

	s, ok := c.Union.(*pincode.TestOneof_Second)
	if !ok {
		panic("not ok")
	}
	fmt.Println("type assertion second:", s)

}

/*
one of first, len: 4
[1|000 | 1111 | 10|000 | 1111]
one of second, len: 7
[1|000 | 1111 | 11|010 | 11 | 1100001 1100010 1100011]
unmarshal second: {"num":15,"Union":{"Second":"abc"}}
type assertion second: &{abc}

*/
