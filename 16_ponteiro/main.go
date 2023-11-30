package main

import "fmt"

func entregar(correspondencia *string) {
	fmt.Printf("Entregou a correspondência '%v' no endereço '%v'\n", *correspondencia, correspondencia)
	*correspondencia = "Carta"
}

func main() {
	// memória -> endereço -> valor

	//Casa	| - |	&0xc00008e020
	//Casa	| - |	&0xc00008e070
	//Casa	| - |	&0xc00008e080

	correspondencia := "Boleto"
	entregar(&correspondencia)
	fmt.Printf("Entregou a correspondência '%v 'no endereço '%v'\n", correspondencia, &correspondencia)

}
