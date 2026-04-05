package main

import (
	"fmt"
)

// Declaramos tipo
type Dinero int

func (d Dinero) String() string {
	return fmt.Sprintf("$%d", d)
}

func main() {
	var sueldo Dinero
	sueldo = 25000
	fmt.Println("Sueldo: ", sueldo)

	aumento := 1000
	sueldo += Dinero(aumento)
	fmt.Println("Aumento: ", sueldo)
}
