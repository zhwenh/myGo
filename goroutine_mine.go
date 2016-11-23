package main

import (
	"fmt"
)


func main() {
	go func(a, b int) {
		fmt.Println(a+b)
	} (3, 4)
	var input string
	fmt.Scanln(&input)
}
