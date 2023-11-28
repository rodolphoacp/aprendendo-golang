package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := soma(0, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)
}

func soma(a, b int) (int, error) {
	if a == 0 {
		return 0, errors.New("valor não pode ser igual a 0")
	}

	return a + b, nil
}
