package main

import (
	"fmt"
)

// Definir Estructura Persona
type Persona struct {
	Nombre string
	edad   int
}

// recibe 2 personas como parametro.  Devuelve una persona y un entero
func Older(p1 Persona, p2 Persona) (Persona, int) {
	if p1.edad > p2.edad { //Si la edad de la Persona es mayor
		return p1, p1.edad - p2.edad //Recibe a la persona y la diferencia de edad
	}
	return p2, p2.edad - p1.edad
}

func main() {

	var p Persona
	p.Nombre = "Ezequiel"
	p.edad = 20

	fmt.Println("Nombre: ", p.Nombre)
	fmt.Println("Edad: ", p.edad)

	p2 := Persona{Nombre: "Iara", edad: 21}
	fmt.Println("Nombre p2: ", p2.Nombre)
	fmt.Println("Edad p2: ", p2.edad)

	p3 := Persona{"KALA", 25}
	fmt.Println("Nombre p3: ", p3.Nombre)
	fmt.Println("Edad p3: ", p3.edad)

	//Ejercicio Older
	tom := Persona{"Tomas", 60}
	paul := Persona{"Paul", 20}
	rick := Persona{"Rick", 80}

	mayorA, DifereciaA := Older(tom, paul)
	mayorB, DeiferenciaB := Older(tom, rick)
	mayorC, DiferenciaC := Older(paul, rick)

	fmt.Printf("Entre %s y %s, %s es mayor con una diferencia de %d \n", tom.Nombre, paul.Nombre, mayorA.Nombre, DifereciaA)
	fmt.Printf("Entre %s y %s, %s es mayor con una diferencia de %d \n", tom.Nombre, rick.Nombre, mayorB.Nombre, DeiferenciaB)
	fmt.Printf("Entre %s y %s, %s es mayor con una diferencia de %d", paul.Nombre, rick.Nombre, mayorC.Nombre, DiferenciaC)
}
