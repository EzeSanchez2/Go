package main

import (
	"fmt"
)

type Persona struct {
	Nombre   string
	Apellido string
	edad     int
}

type Estudiante struct {
	Persona
	Carrera string
}
type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	ezequiel := Estudiante{
		Persona{
			Nombre:   "Ezequiel",
			Apellido: "Sanchez",
			edad:     23,
		},
		"Informatica",
	}

	fmt.Println("Nombre: ", ezequiel.Nombre)
	fmt.Println("Edad: ", ezequiel.edad)
	fmt.Println("Carrera: ", ezequiel.Carrera)

	profesor := Profesor{
		Estudiante{
			Persona{
				"Pedro",
				"Almonte",
				20,
			},
			"Informatica",
		},
		"Biotecnologa",
	}

	fmt.Println("Profesor: ", profesor)
	fmt.Println("Nombre: ", profesor.Nombre)
	fmt.Println("Apellido: ", profesor.Apellido)
	fmt.Println("Edad: ", profesor.edad)
	fmt.Println("Carrera: ", profesor.Carrera)
	fmt.Println("Carrera Estudiante: ", profesor.Estudiante.Carrera)

}
