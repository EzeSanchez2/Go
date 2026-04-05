package main

import (
	"fmt"
)

// Si la estructura no tiene campo es que se le van a agregar metodos (estructura de tamaño 0)
type OpPunto struct {
}

type Punto struct {
	x, y int
}

type Punto3d struct {
	x, y int
	*Punto3d
}

func main() {
	p := Punto{}
	fmt.Println(p)

	p2 := Punto3d{
		2,
		3,
		&Punto3d{
			5,
			6,
			nil,
		},
	}
	fmt.Println(p2)

	//Comparando Estructuras
	a := Punto{2, 3}
	b := Punto{2, 3}
	fmt.Println("a == b: ", a == b)

	//Usar como indice en los mapas
	i := make(map[Punto]string)
	i[a] = "Hola como estas"
	fmt.Println(i[a])
}
