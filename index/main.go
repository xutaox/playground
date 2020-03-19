package main

import (
	"fmt"
	"index/suffixarray"
	"regexp"
)

func main() {

	idx := suffixarray.New([]byte("orderchinesefoodtoday,somesuacegetheaway"))
	reg, err := regexp.Compile("ay")
	if err != nil {
		panic(err)
	}
	fmt.Println(idx.FindAllIndex(reg, -1))
	fmt.Println(idx.Lookup([]byte("ay"), -1))
}
/*
[[19 21] [38 40]]
[38 19]
 */
