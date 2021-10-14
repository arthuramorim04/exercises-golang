package main

import "fmt"

//soma dois valores
func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func calculos(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}

func main() {
	resultado := somar(1, 2)

	var log = func(text string) {
		fmt.Println(text)
	}

	log("Soma anterior igual a: ")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculos(1, 2)
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSubtracao)

	resultadoSoma2, _ := calculos(5, 3)
	fmt.Println(resultadoSoma2)
}
