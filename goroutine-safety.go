package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	count int = 0
	count2 int = 0
	mu sync.Mutex
)

func increment() {
	count++
}
func incrementWithMutex() {
	mu.Lock()
	count2++
	mu.Unlock()
}
func main() {
	for i := 0; i < 1000000; i++ {
		go increment()
		go incrementWithMutex()
	}
	for i := 0; i < 30; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(count, count2)
	}
}
