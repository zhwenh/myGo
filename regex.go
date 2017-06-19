package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1 := regexp.MustCompile(`([a-zA-Z0-9_.-/]{0,253}/)?[a-zA-Z0-9]([a-zA-Z0-9_.-]{0,61}[a-zA-Z0-9])?`)
	fmt.Println(re.MatchString("k8s_ns/lab.el/kubernetes-admin.caicloud.io/partition", 1))
	

	re := regexp.MustCompile("a.")	
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", 3))
	fmt.Println(re.FindAllString("none", -1))

}
