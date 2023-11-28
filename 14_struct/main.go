package main

import "fmt"

type Aluno struct {
	Nome      string
	NotaFinal int
}

func (a Aluno) Imprimir() {
	a.Nome = "Alterou o nome"
	fmt.Printf("Aluno %v com nota final %v\n", a.Nome, a.NotaFinal)
}

func main() {

	aluno := Aluno{
		Nome:      "João",
		NotaFinal: 8,
	}
	aluno.Imprimir()
	fmt.Printf("\nAluno %v\n", aluno)
}
