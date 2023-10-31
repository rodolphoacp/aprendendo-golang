package main

import "fmt"

type nome string

const UF string = "SC"

var idade int
var razao nome
var cidade = "Criciuma"
var ativo = false
var inativo bool
var tempo int
var sobrenome string

func main() {
	razao = "12"
	idade = 1

	ativo = true
	fmt.Printf("Helo world %T %T %T %T %T\n ", razao, idade, UF, cidade, ativo)

	estado := "SC"
	fmt.Printf("Helo world %T %v\n ", estado, estado)

	fmt.Printf("%T %v\n ", inativo, inativo)
	fmt.Printf("%T %v\n ", tempo, tempo)
	fmt.Printf("%T %v\n ", sobrenome, sobrenome)
}
