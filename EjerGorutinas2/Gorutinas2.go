package main

import(
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var TotalGastado float64 = 0
type Gastos struct{
	Categoria string
	Monto float64
}


func SumarGasto(g Gastos, wg * sync.WaitGroup ) {
	defer wg.Done() // Resta una gorutine
	sleep:=rand.Int63n(500)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	fmt.Println("Procesando:",g.Monto + TotalGastado)
	
}

func main()  {
	var wg sync.WaitGroup
	
	Gastos:=[]Gastos{
		{Categoria: "Supermercado", Monto: 1000},
		{Categoria: "Salud", Monto: 5000},
		{Categoria: "Verduleria", Monto: 200},
		{Categoria: "Carniceria", Monto: 100},
	}

	fmt.Println("-- Iniciando el calculo del total gastado --")

	for _ , g:= range Gastos{
		wg.Add(1)
		go SumarGasto(g,&wg)
	}

	wg.Wait()

	fmt.Println("El total general es de: ", TotalGastado)


}