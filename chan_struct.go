package main

import (
	"fmt"
	// go fmt puts these in alphbetical order.
)

type A struct {
	b int
}

func pump(ch chan *A, a *A) {
	ch <- a
}

func suck(ch chan *A) {
	for {
		a := <-ch
		a.b = 456
		fmt.Println("suck", a)
	}
}
func main() {
	ch1 := make(chan *A)
	a := A{b:123}
	go pump(ch1, &a)

	go suck(ch1)
	fmt.Println("main", a)
	var input string
	fmt.Scanln(&input)
	fmt.Println("main", a)
}
