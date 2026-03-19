package main

import (
	"fmt"
)

func main() {
	var x [3]int // Un array x que tiene 3 numeros
	fmt.Println(x)

	var k [3][2]float64 //Un array de 3 dimensiones con dos valores adentro
	fmt.Println(k)

	x[1] = 25 // En el indice 1 se inserta el valor 25
	fmt.Println(x)

	fmt.Println(x[1]) // Accedo al indice del array muestro lo que hay

	y := [5]int{1, 2, 3, 4, 5} // Inicializo la variable y que va a tener 5 lugares y el {} lo llena con los valores
	fmt.Println(y)

	j := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // Omitimos la cantidad de elementos que va a tener el array , el compilador analiza cuantos elementos estamos inicializando
	fmt.Println(j)

	fmt.Println(len(j)) // Longitud del array

	fmt.Println(j[len(j)-1]) // Ultimo elemento del array

}
