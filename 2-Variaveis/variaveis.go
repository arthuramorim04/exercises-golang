package main

import (
	"errors"
	"fmt"
)

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string
		variavel4 string
	)

	fmt.Println(variavel3 + " " + variavel4)

	variavel1, variavel2 = variavel2, variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// numeros

	//inteiros
	var numero1 int64 = -100000000000000000
	fmt.Println(numero1)

	var numero2 uint32 = 1000000000
	fmt.Println(numero2)

	// alternativa para o int32
	// INT32 = RUNE

	var numero3 rune = 123123
	fmt.Println(numero3)

	// byte uint = byte

	var numero4 byte = 123
	fmt.Println(numero4)

	// real

	var numero5 float64 = 1233232323232.45
	fmt.Println(numero5)

	numero6 := 123321.43
	fmt.Println(numero6)

	// String

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	//devolve o valor na tabela ASC
	char := 'B'
	fmt.Println(char)

	//devolve o valor do texto
	char2 := "B"
	fmt.Println(char2)

	//declarar sem inicializar
	var texto string
	fmt.Println(texto)

	var inteiro int
	fmt.Println(inteiro)

	var numFloat float64
	fmt.Println(numFloat)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("erro2")
	fmt.Println(erro2)

}
