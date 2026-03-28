package main

import "fmt"

type Regla func(int) bool

func crearFiltroDeEdad(minima int) Regla {
	return func(edad int) bool {
		if edad > minima {
			return true
		} else {
			return false
		}
	}
}

func validarAcceso(valor int, reglas ...Regla) bool {
	for _, reglaActual := range reglas {
		if !reglaActual(valor) {
			return false
		}
	}
	return true
}

func monitorDeSeguridad() func(bool) int {
	rechazados := 0
	return func(b bool) int {
		if !b {
			rechazados++
		}
		return rechazados
	}
}

func main() {

	esMayor := crearFiltroDeEdad(18)
	esVip := crearFiltroDeEdad(21)

	//Monitor de errores
	contarErrores := monitorDeSeguridad()

	//Cliente
	edadCliente := []int{10, 15, 30, 11, 50}

	fmt.Println("--SISTEMA DE ACCESO AL CASINO--")
	for _, edad := range edadCliente {
		accesoPermitido:= validarAcceso(edad,esMayor)

		totalRechazados:= contarErrores(accesoPermitido)

		if accesoPermitido{
			fmt.Printf("EL CLIENTE TIENE %d [ENTRA]",edad )
		}else {
			fmt.Printf("ES MENOR NO ENTRA. LA CANTIDAD DE RECHAZADOS SON %d", totalRechazados)
		}

		fmt.Println("--CONTROL SECTOR VIP--")
		vip:=21
		if validarAcceso(vip,esMayor,esVip){
			fmt.Println("EL CLIENTE TIENE ACCESO VIP")
		}else {
			fmt.Println("NO ES ACCESO VIP")
		}

	}

}
