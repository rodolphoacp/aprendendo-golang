package main

import "fmt"

func main() {
	fmt.Println(soma(2, 5, 6, 7, 8, 9))
	fmt.Println(soma(2))
	fmt.Println(soma(2, 5))

}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
