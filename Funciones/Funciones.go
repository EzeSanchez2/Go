package main

import (
	"fmt"
)

func ImprimirNombre(nombre string) {
	fmt.Println("FUERA DEL MAIN")
	fmt.Println("EL NOMBRE ES: ", nombre)
}

func sumar(num1 int, num2 int) int { // int
	return num1 + num2 // devuekve int
}

func restar(n1, n2 int) (r int) { //int
	r = n1 - n2
	return r //devuelve int
}

func main() {
	fmt.Println("ADENTRO DEL MAIN")
	ImprimirNombre("Ezequiel")

	fmt.Println(sumar(10, 20))
	fmt.Println(restar(20, 5))

	fmt.Printf("FUNCION TIPO %T", sumar)
	fmt.Printf("FUNCION TIPO %T", restar)
}
