package main

import (
	"fmt"
	"sync"
)

type MultiFields struct {
	a string
	b string
	c string
	d string
}


func main() {
	for i:=0; i < 10000; i++ {
		v := &MultiFields{}
		var wg sync.WaitGroup
		wg.Add(4)
		go func() {
			defer wg.Done()
			v.a ="a"
		} ()
		go func() {
			defer wg.Done()
			v.b ="b"
		} ()
		go func() {
			defer wg.Done()
			v.c ="c"
		} ()
		go func() {
			defer wg.Done()
			v.d ="d"
		} ()
		wg.Wait()
		if v.a != "a" || v.b != "b" || v.c != "c" || v.d != "d" || i % 1000 == 0{
			fmt.Println(v)
		}
	}
}
