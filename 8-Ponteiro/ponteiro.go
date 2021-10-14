package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// ponteiro

	var variavel3 int = 100
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	variavel3++
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

	*ponteiro = 99
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}
