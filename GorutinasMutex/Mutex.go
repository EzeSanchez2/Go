package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	KiTotal int
	mu      sync.Mutex
	wg      sync.WaitGroup
)

type Guerreros struct {
	Nombre   string
	KiCargas int
}

func CargarPoder(g Guerreros, wg *sync.WaitGroup) {
	defer wg.Done() // Restamos una Gorutina
	sleep := rand.Int63n(1000)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	// Se entra de a uno en la varibale KiTotal para sumar
	mu.Lock() // Cierro la puerta , nadie entra
	KiTotal += g.KiCargas
	fmt.Printf("¡%s ha terminado de cargar %d de Ki!\n", g.Nombre, g.KiCargas)
	mu.Unlock() // Abro la puerta para que entre otro

}

func main() {
	guerreros := []Guerreros{
		{Nombre: "Goku", KiCargas: 500},
		{Nombre: "Vegeta", KiCargas: 450},
		{Nombre: "Gohan", KiCargas: 300},
		{Nombre: "Piccolo", KiCargas: 200},
	}

	fmt.Println("-- Iniciando la carga de poderes --")
	for _, e := range guerreros {
		wg.Add(1)
		go CargarPoder(e, &wg)

	}

	wg.Wait()
	fmt.Printf("\n --- ¡Poder total reunido: %d --- \n!", KiTotal)

}
