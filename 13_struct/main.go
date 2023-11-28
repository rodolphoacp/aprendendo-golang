package main

import "fmt"

type Aluno struct {
	Nome string
	Escola
}

type Escola struct {
	Nome string
}

func main() {
	aluno := Aluno{Nome: "João"}
	fmt.Printf("Nome %v da escola %v\n\n", aluno.Nome, aluno.Escola)

	aluno.Nome = "Maria"
	fmt.Printf("Nome %v da escola %v\n\n", aluno.Nome, aluno.Escola)
}
