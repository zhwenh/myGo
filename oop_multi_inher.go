package main 

import ( 
        "fmt" 
) 

type A interface { 
	helloA()
} 
type B interface {
	helloB()
}

type commonSon struct {
	A
	B
}

func (c *commonSon) helloA() {
	fmt.Println("A says hello!")
}

func (c *commonSon) helloB() {
	fmt.Println("B says hello!")
}
func main() { 
        c := &commonSon{}
	    c.helloA()
	    c.helloB()
}