package main

import (
	"fmt"
)

type user struct {
	name     string // name
	lastname string // last name of the user
	age      int    // age of the user
	endereco
}

type endereco struct {
	rua    string // name of the user
	numero uint   // number ofero
}

type estudante struct {
	user
	curso string // name of the curse
}

func main() {
	fmt.Println("Herança")

	endereco1 := endereco{"Rua independencia", 120}

	user := user{"Arthur", "Amorim", 22, endereco1}

	estudante1 := estudante{user, "ADS"}

	fmt.Println(estudante1)

}
