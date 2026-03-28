package main

import (
	"fmt"
)

func sumarApuestas(apuestas ...int) (int, int) {
	sumador := 0
	contador := 0
	for _, a := range apuestas {
		sumador += a

		contador++
	}
	return sumador, contador
}

func auditorGrandesApuestas(limite int) func(int) (bool, int) {
	contadorPasarMonto := 0
	return func(monto int) (bool, int) {
		esSospechoso := false
		if monto > limite {
			contadorPasarMonto++
			esSospechoso = true
		}
		return esSospechoso, contadorPasarMonto
	}
}

func main() {
	apuestas := []int{500, 12427, 555, 3000, 400306721, 5000}

	fmt.Println("**SUMAR APUESTAS**")
	fmt.Println(sumarApuestas(apuestas...))

	auditoria := auditorGrandesApuestas(10000)
	fmt.Println("**AUDITAR APUESTAS GRANDES**")
	for _, apuesta := range apuestas {
		esMala, TotalMalas := auditoria(apuesta)
		if esMala {
			fmt.Printf("ALERTA!, apuesta de %d, cantidad %d\n", apuesta, TotalMalas)
		} else {
			fmt.Printf("Apuesta de %d. Hasta el momento la cantidad de sospechosos hasta ahora: %d\n", apuesta, TotalMalas)
		}
	}

	fmt.Println("**ES PAR O IMPAR**")
	for _, num := range apuestas {
		func(n int) {
			if n%2 == 0 {
				fmt.Println("Es par")
			} else {
				fmt.Println("Es impar")
			}
		}(num)
	}
}
