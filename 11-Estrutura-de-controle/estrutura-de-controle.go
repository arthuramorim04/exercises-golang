package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 10 {
		fmt.Println("Menor que 10")
	} else {
		fmt.Println("Maior ou igual a 10")
	}

	if numero2 := 15; numero2 > 10 {
		fmt.Println("Maior igual a", numero2)
	} else {
		fmt.Println("Menor que 10")
	}
}
