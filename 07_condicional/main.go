package main

import "fmt"

func main() {
	alunos := map[string]int{"Rodolpho": 0, "maria": 8}

	if alunos["Rodolpho"] > 0 {
		fmt.Println("Passou")
	} else if alunos["Rodolpho"] == 0 {
		fmt.Println("Não passou")
	}

	cor := 1
	switch cor {
	case 1:
		fmt.Println("Amarelo")
	case 2:
		fmt.Println("Vermelho")
	default:
		fmt.Println("Sem cor")
	}

	for aluno, nota := range alunos {
		if nota == 0 {
			break
		}
		fmt.Printf("Aluno %v com nota %v\n", aluno, nota)
	}
}
