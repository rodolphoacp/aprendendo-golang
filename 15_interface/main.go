package main

import "fmt"

type Pessoa interface {
	string()
}

type Aluno struct {
	Nome      string
	NotaFinal int
}

type Professor struct {
	Nome string
}

func (a Aluno) string() {
	fmt.Printf("Nome do aluno é %v\n", a.Nome)
}

func (p Professor) string() {
	fmt.Printf("Nome do professor é %v\n", p.Nome)
}

func ToString(p Pessoa) {

	aluno, resultado := p.(Aluno)
	if !resultado {
		fmt.Printf("Não foi possível converter a interface do tipo %T\n\n", p)
		return
	}

	// switch p.(type)
	fmt.Printf("O tipo do objeto é %T\n", aluno)
	fmt.Printf("A nota final é %v\n", aluno.NotaFinal)

	// Type Assertion
	p.string()
}

func main() {
	aluno := Aluno{
		Nome:      "João",
		NotaFinal: 9,
	}

	ToString(aluno)

	professor := Professor{
		Nome: "Maria",
	}

	ToString(professor)

}
