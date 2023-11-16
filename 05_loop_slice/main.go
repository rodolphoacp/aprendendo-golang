package main

import "fmt"

func main() {
	array := [2]string{"Banana", "Maça"}
	fmt.Printf("Valores Array %v\n", array)

	slice := []string{"Banana", "Maça", "Melância"}
	slice = append(slice, "Uva")
	slice = append(slice, "Goiaba")
	fmt.Printf("Valores %v tamanho %d capacidade %d\n\n\n", slice, len(slice), cap(slice))

	sliceCarro := []string{"Fusca"}
	slice = append(slice, sliceCarro...)
	fmt.Printf("Valores %v tamanho %d capacidade %d tipo %T\n\n", slice, len(slice), cap(slice), sliceCarro)

	for _, valor := range slice {
		fmt.Printf("Valor: %v\n", valor)
	}

	fmt.Printf("\n\nPedaço do slice %v %v\n", slice[0:1], slice[3:4])

	sliceLinha := [][]string{
		[]string{"Banana", "uva"},
		[]string{"Goiaba"},
	}

	fmt.Printf("Multi %v\n", sliceLinha)
}
