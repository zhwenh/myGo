package main

import (
	"fmt"
	// go fmt puts these in alphbetical order.
)

func main() {
	m := map[string]float64{}
	m["1"] = 1
	m["pi"] = 3.1415
	m2 := m
	m2["pi"] = 12
	fmt.Println(m, m2)	
}
