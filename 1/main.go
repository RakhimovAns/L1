// L1.1
package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func (a Action) Run() {
	fmt.Printf("%s бежит!\n", a.Name)
}

func main() {
	a := Action{
		Human: Human{Name: "Ансар", Age: 17},
	}

	a.SayHello()

	a.Run()
}
