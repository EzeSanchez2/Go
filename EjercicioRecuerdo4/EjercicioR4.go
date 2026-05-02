package main

import (
	"errors"
	"fmt"
)

var (
	ErrPresupuestoInsuficiente = errors.New("El Barça no tiene palancas suficientes para este fichaje")
	ErrEdadInvalida            = errors.New("El jugador es demasiado joven o ya esta para el retiro")
	ErrCupoExtracomunitario    = errors.New("No hay mas lugar para jugadores fuera de la UE")
)

func FicharJugador(presupesto float64, precio float64, edad int, esUE bool) (float64, error) {
	if precio > presupesto {
		return presupesto, ErrPresupuestoInsuficiente
	}
	if edad < 16 || edad > 40 {
		return presupesto, ErrEdadInvalida
	}
	if esUE == false {
		return presupesto, ErrCupoExtracomunitario
	}

	return presupesto, nil
}

func ProcesarFichaje(presupesto float64, precio float64, edad int, esUE bool) {
	NuevoSaldo, err := FicharJugador(presupesto, precio, edad, esUE)
	if err != nil {
		if errors.Is(err, ErrPresupuestoInsuficiente) {
			fmt.Println("¡ERROR!: Llamen a Laporta para activar una palanca.")
		} else if errors.Is(err, ErrEdadInvalida) {
			fmt.Println("¡ERROR!: Demasiado joven o muy veterano. Busquemos en La Masía.")
		} else if errors.Is(err, ErrCupoExtracomunitario) {
			fmt.Println("¡ERROR!: No podemos ficharlo, el cupo de extranjeros está lleno.")
		} else {
			fmt.Println("Ocurrió un error desconocido en el mercado.")
		}
		return
	}
	fmt.Println("Fichaje estrella concretado su salario es de ", NuevoSaldo)
}

func main() {
	fmt.Println("\n--- Intento de fichaje 1 ---")
	ProcesarFichaje(100000000, 20000000, 17, true)

	fmt.Println("\n--- Intento de Fichaje 2 ---")
	ProcesarFichaje(50000000, 200000000, 25, true)

	fmt.Println("\n--- Intento de Fichaje 3 ---")
	ProcesarFichaje(100000000, 5000000, 12, true)

	fmt.Println("\n--- Intento de Fichaje 4 ---")
	ProcesarFichaje(100000000, 30000000, 20, false)

}
