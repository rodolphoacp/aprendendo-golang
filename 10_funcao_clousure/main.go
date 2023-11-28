package main

import "fmt"

func main() {
	soma := func() int {
		a := 5
		b := 5
		return a + b
	}()

	fmt.Println(soma)

	subtracao := func(a, b int) int {
		return a + b
	}(2, 1)

	fmt.Println(subtracao)
}
