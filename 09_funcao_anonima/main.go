package main

import "fmt"

func main() {

	frase := "Função anônima"
	func(frase string) {
		fmt.Println(frase)
	}(frase)

	func() {
		frase := "Teste"
		fmt.Println(frase)
	}()
}
