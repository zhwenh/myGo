package main

import "fmt"

type T struct{
	name string
}
func main() {
	t := T{name:"aaa"}
	b := t
	b.name="bbb"
	fmt.Println(t, b)
	
}
