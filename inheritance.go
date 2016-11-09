package main

import "fmt"

type I interface {
	Talk()
	SameMethod()
}

type Person struct {
	I
  Name string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) SameMethod() {
  fmt.Println("just the same")
}

type Android struct {
  Person
  Model string
}

func getPerson(name string) I {
	return &Person {
		Name: name,
	}
}

func getAndroid(name string) I {
	return &Android{
		Person: Person {
			Name: name,
		},
	}
}
func (a *Android) Talk() {
  fmt.Println("Hi, Android my name is", a.Name)
}

func main() {
	p := Person{
		Name: "sbcd",	
	}
	p.Talk()
	a := Android{
		Person: p,
	}
	
	a.Talk()

	pi := getAndroid("adaf")
	pi.SameMethod()
	pi.Talk()
}