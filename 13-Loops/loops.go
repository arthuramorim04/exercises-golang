package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"arthur", "gabriela", "joao"}

	for indice, value := range nomes {
		fmt.Println(indice, " - ", value)
	}

	for indice, value := range "arthur" {
		fmt.Println(indice, " - ", string(value))
	}

	usuario := map[string]string{
		"nome":      "Arthur",
		"sobrenome": "Amorim",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, " - ", valor)
	}

	for {
		fmt.Println("executando infinitamente...")
		time.Sleep(time.Second)
	}
}
