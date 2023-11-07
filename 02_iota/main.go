package main

import "fmt"

const (
	a = iota + 1
	_
	c
	d
)

const (
	Jan = (iota + 1) * 1000
	Fev
	Mar
)

func main() {

	fmt.Printf("Valor de a = %d\n", a)
	fmt.Printf("Valor de c = %d\n", c)
	fmt.Printf("Valor de d = %d\n", d)

	fmt.Printf("Valor de Jan = %d\n", Jan)
	fmt.Printf("Valor de Fev = %d\n", Fev)
	fmt.Printf("Valor de Mar = %d\n", Mar)
}
