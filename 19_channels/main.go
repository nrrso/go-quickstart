package main

import (
	"fmt"
)

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)
	go foo(fooVal, 5)
	go foo(fooVal, 3)
	a := <-fooVal
	b := <-fooVal
	fmt.Println(a, b)
}
