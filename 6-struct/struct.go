package main

import "fmt"

type user struct {
	name     string   // name
	lastname string   // last name of the user
	age      int      // age of the user
	endereco endereco // endereco of the user
}

type endereco struct {
	rua    string // name of the user
	numero uint   // number ofero
}

func main() {
	fmt.Println("Arquivo de struct")

	var usuario user
	usuario.name = "Arthur"
	usuario.lastname = "Amorim"
	usuario.age = 22
	fmt.Println(usuario)

	endereco1 := endereco{rua: "Rua Independencia", numero: 120}

	usuario2 := user{
		name: "Gabriela", lastname: "da Rocha", age: 24, endereco: endereco1,
	}

	fmt.Println(usuario2)
}
