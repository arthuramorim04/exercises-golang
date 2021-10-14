package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"name":     "Arthur",
		"lastname": "Amorim",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["name"])
	fmt.Println(usuario["lastname"])

	usuario2 := map[string]map[string]string{
		"user": {
			"name":     "Arthur",
			"lastname": "Amorim",
		},
		"user2": {
			"name":     "Gabriela",
			"lastname": "da Rocha",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "user2")
	fmt.Println(usuario2)
	usuario2["user2"] = map[string]string{
		"name":     "Gabriela",
		"lastname": "da Rocha",
	}
	fmt.Println(usuario2)
}
