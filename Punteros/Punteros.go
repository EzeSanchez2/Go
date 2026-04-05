package main

import (
	"fmt"
)

func Incrementar(numero *int) {
	*numero++
	fmt.Println("Valor de numero en la funcion incrementar: ", *numero)
}

func main() {
	a := 25
	fmt.Println("Direccion de memoria : ", &a)

	b := &a

	fmt.Println(b)

	//Imprime el valor que contiene la variable "a" osea 25
	fmt.Println(*b)

	//b=21 Esto te da un error porque no es tipo int es tipo *int

	*b = 15
	fmt.Println("Valor de a : ", a)
	fmt.Println("Direccion de a: ", &a)

	a++
	fmt.Println("Valor de b: ", *b)

	//Funciones para punteros
	numeros := 10
	fmt.Println("Numero antes de ingresar en la funcion incrementar: ", numeros)
	Incrementar(&numeros)
	fmt.Println("Numero despues de ingresar en la funcion incrementar:  ", numeros)

}
