package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int, 10) // buffer of 10 > 9 in sync channels
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i) // a channel per iteration
	}
	wg.Wait()
	close(fooVal)
	for item := range fooVal {
		fmt.Println(item)
	}
}
