package main

import (
	"fmt"
)

func main() {
	var Numero int
	fmt.Println("INGRESE UN NUMERO: ") //5
	fmt.Scanln(&Numero)
	fmt.Println("***EMPIEZA LA TABLA DE MULTIPLICAR***")
	for i := 10; i > 0; i-- {
		fmt.Println(Numero * i)
	}

}
