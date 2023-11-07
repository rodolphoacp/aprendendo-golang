package main

import "fmt"

func main() {

	for contador := 1; contador <= 3; contador++ {
		fmt.Printf("Valor do contador é igual a %d\n", contador)
	}

	fmt.Printf("\n\n-------------------\n\n")
	contador := 1
	for contador <= 5 {
		fmt.Printf("Valor do contador é igual a %d\n", contador)
		contador++
	}
}
