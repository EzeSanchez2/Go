package main

import (
	"fmt"
)

func main() {
	var j []int //Declarar un slice
	fmt.Println("Slice j: ", j)

	x := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice x: ", x)

	i := make([]int, 5) //Declarar slice con Make , aca dice que el slice tiene 5 valor
	fmt.Println("Slice i: ", i)
	fmt.Println("Longitud slice i: ", len(i))
	fmt.Println("Capacidad slice i: ", cap(i)) // La capacidad del array

	e := make([]int, 5, 12) // tiene 5 valores pero la capacidad total es de 10
	fmt.Println("Slice e: ", e)
	fmt.Println("Longitud slice e: ", len(e))
	fmt.Println("Capacidad slice e: ", cap(e))

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Slice 10 elementos: ", ar)

	//Definimos dos Slice
	var a, b []int
	fmt.Println("Slice a: ", a)
	fmt.Println("Slice b: ", b)

	b = ar[0:8]
	b[3] = 25
	fmt.Println("Slice b[3] cambiamos a 25: ", b)

	//Imprime el primer numero pro no incluye el ultimo
	a = ar[2:5]
	fmt.Println("Slice a: ", a)

}
