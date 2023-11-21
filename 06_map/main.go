package main

import "fmt"

func main() {
	alunos := map[string]int{"Rodolpho": 0, "maria": 8}

	fmt.Println(alunos)

	valor, ok := alunos["João"]
	fmt.Printf("\nValor é %v se existe %v\n\n", valor, ok)

	valor, ok = alunos["Rodolpho"]
	fmt.Printf("\nValor é %v se existe %v\n\n", valor, ok)

	//
	alunos["José"] = 10
	fmt.Println(alunos)

	delete(alunos, "Rodolpho")
	fmt.Println(alunos)

	for key, aluno := range alunos {
		fmt.Printf("Nota do aluno %v é igual a %v\n", key, aluno)
	}

	clear(alunos)
	fmt.Println(alunos)
}
