package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------CONTADOR DE NUMEROS IMPARES --------")

	fmt.Println("Ingrese el primer numero: ")
	var numero1 int
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el segundo numero: ")
	var numero2 int
	fmt.Scanln(&numero2)

	var contador int

	for a := numero1; a <= numero2; a++ {
		if a%2 != 0 {
			contador++

			fmt.Println("Los numeros impares son: ", a)
		}
	}

	fmt.Println("--------- La cantidad de numeros impares son: --------- ")
	fmt.Println(contador)

}
