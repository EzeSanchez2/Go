package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Moto struct {
	Patente string
	ID      int
}

func ProcesarIngresos(m Moto, wg *sync.WaitGroup) {
	defer wg.Done()
	sleep := rand.Int63n(2000)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	fmt.Printf("Moto %q procesada correctamente", m.Patente)
}

func main() {
	var wg sync.WaitGroup
	motos := []Moto{
		{Patente: "ABC-123", ID: 1},
		{Patente: "XYZ-789", ID: 2},
		{Patente: "DEF-456", ID: 3},
		{Patente: "GHI-012", ID: 4},
		{Patente: "JKL-345", ID: 5},
	}
	fmt.Println("--- Iniciando escaneo de seguridad ---")
	for _, m := range motos {
		wg.Add(1)
		go ProcesarIngresos(m, &wg)
	}

	wg.Wait()

	fmt.Println("--- Terminando el escaneo de las motos ---")
}
