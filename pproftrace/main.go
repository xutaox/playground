package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

var ch = make(chan int)
var wg = sync.WaitGroup{}

func main() {

	runtime.GOMAXPROCS(1)

	err := trace.Start(os.Stderr)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("cap", cap(ch))

	wg.Add(1)
	go myTest()
	wg.Wait()

}

func myTest() {
	go send()
	recv()
	wg.Done()
}

func send() {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

func recv() {
	for i := range ch {
		fmt.Println("recv", i)
	}
}
