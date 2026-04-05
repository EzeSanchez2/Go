package main

import (
	"fmt"
)

func imprimir() {
	fmt.Println("Hola Eze")
	//panic("error")

	defer func() {
		cadena := recover()
		fmt.Println(cadena)
	}()
	panic("Error")
}
func main() {
	imprimir()
	fmt.Println("Hola main")
}
