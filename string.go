package main

import (
	"fmt"
	// "io/ioutil"
	// "strings"
	"strings"

)

func wordPattern(pattern string, str string) bool {
    arr := strings.Split(str, " ")
    m := make(map[string]string)
    m2 := make(map[string]bool)
    for i, word := range arr {
    	fmt.Println(string(pattern[i]))
        if val, ok := m[word]; !ok {
        	if m2[string(pattern[i])] {
        		return false
        	}
        	m2[string(pattern[i])] = true
            m[word] = string(pattern[i])
        } else {
            if val != string(pattern[i]) {
                return false
            }
        }
    fmt.Println(m, m2)
    }
    fmt.Println(m, m2)
    return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
