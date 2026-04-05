package main

import (
	"fmt"
)

var dbBoloqueada = false

var saldo = map[string]int{
	"Ezequiel": 2000,
	"Iara":     1000,
}

func gestionarBaseDeDatos(accion string) {
	if accion == "BLOQUEAR" {
		dbBoloqueada = true
		fmt.Println("BASE DE DATOS CERRADA POR MANTENIMIENTO")
	}
	if accion == "DESBLOQUEAR" {
		dbBoloqueada = false
		fmt.Println("BASE DE DATOS ABIERTA NUEVAMENTE")
	}
}

func realizarBackup() {
	gestionarBaseDeDatos("BLOQUEAR")

	//Este defer aunque la base de datos se prenda fuego se ejecuta igual y desbloquea la base de datos
	defer gestionarBaseDeDatos("DESBLOQUEAR")

	fmt.Println("-> Copiando datos críticos del Casino...")

	panic("¡ERROR CRÍTICO! El servidor de backup se prendió fuego")
}

func main() {
	realizarBackup()
	fmt.Println("ESTADO DE LA FUNCION ¿BLOQUEADA?:   ", dbBoloqueada)
}
