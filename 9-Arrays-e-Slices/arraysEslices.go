package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int

	// Array (Lista Fixa)
	fmt.Println(array1)

	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5

	fmt.Println(array1)
	fmt.Println("===============")
	array2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array2)

	fmt.Println("--------------------------")
	// slice (Lista dinamica)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)
	fmt.Println("===============")
	slice = append(slice, 9, 11, 12, 13, 14)
	last := slice[len(slice)-1]
	fmt.Println(slice)
	fmt.Println(last)
	fmt.Println("===============")
	slice2 := array2[2:6]
	fmt.Println(slice2)

	fmt.Println("--------------------------")
	// Arrays internos

	slice3 := make([]int32, 10, 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("===============")
	slice4 := make([]int32, 5)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println(slice4)
}
