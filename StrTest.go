package main

import (
	"fmt"
)

func main() {
	var str string
	str += "abc"
	fmt.Println(str)
	arr2d := make([][]string, 3)
	arr := []string{"a", "b"}
	arr2d[0] = arr
	arr[1] = "c"
	arr2d[1] = arr
	fmt.Println(arr2d)
}
