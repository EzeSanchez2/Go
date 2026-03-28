package main

import (
	"fmt"
)

func MultiplicarNumeros(numero int) (n1, n2, n3 int) {
	n1 = numero * 10
	n2 = numero * 10
	n3 = numero * 10
	return
}

func MultiplicarOtraManera(numero int) (int, int, int) {
	n1 := numero * 20
	n2 := numero * 30
	n3 := numero * 40
	return n1, n2, n3
}

func main() {

	fmt.Println(MultiplicarOtraManera(10))
}
