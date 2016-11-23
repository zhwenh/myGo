package main

import "fmt"

type Config struct {
	Num    int                // config number
	Shards [10]int64          // gid
	Groups map[int64][]string // gid -> servers[]
	dd []int
}

func Do(arr *[]int) {
	*arr = make([]int, 4)
	(*arr)[0] = 123
}
func main() {
	var config Config
	config.Num = 100
	config.Groups = make(map[int64][]string)
	fmt.Println(len(config.Shards), "dd", len(config.dd))
	Do(&config.dd)
	fmt.Println(config)
	
}
