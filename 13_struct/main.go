package main

import "fmt"

type Aluno struct {
	Nome string
	Escola
}

type Escola struct {
	Nome      string
	Descricao string
}

func main() {
	aluno := Aluno{Nome: "João"}
	fmt.Printf("Nome %v da escola %v\n\n", aluno.Nome, aluno.Escola)

	aluno.Nome = "Maria"
	aluno.Escola.Nome = "Escola de Criciúma"
	fmt.Printf("Nome %v da escola %v\n\n", aluno.Nome, aluno.Escola)

	aluno.Descricao = "Escola de Criciúma"
	fmt.Printf("Nome %v da escola %v\n\n", aluno.Nome, aluno.Escola)

	escola := Escola{Nome: "Escola de outra cidade", Descricao: "Sem"}
	fmt.Printf("Escola %v\n", escola)

	ImplicitaStruct()
	InicializarStructComComposicao()
	AnonimoStruct()

}

func ImplicitaStruct() {
	escola := Escola{"Escola", ""}
	fmt.Printf("\n\nEscola %v\n", escola)
}

func InicializarStructComComposicao() {
	aluno := Aluno{
		Nome: "Maria",
		Escola: Escola{
			Nome:      "Teste",
			Descricao: "",
		},
	}

	fmt.Printf("\n\nAluno %v\n\n", aluno)
}

func AnonimoStruct() {
	aluno := struct {
		Nome string
	}{
		Nome: "Sem nome",
	}

	fmt.Printf("Aluno %v\n\n", aluno)
	aluno.Nome = ""

	aluno = struct {
		Nome string
	}{
		Nome: "Sem nome",
	}
}
