package main

import (
	"fmt"
)

type JugadorInscripto interface {
	Info() string
	JugarPartido()
}

type Atleta struct {
	Nombre  string
	Energia int
}

type Futbolista struct {
	Atleta
	Posicion string
}

func (f Futbolista) Info() string {
	return fmt.Sprintf("%q - %s", f.Nombre, f.Posicion)
}

func (f *Futbolista) JugarPartido() {
	f.Energia = f.Energia - 20
	fmt.Println("La energia del jugador es de ", f.Energia)
}

func (f *Futbolista) Recuperar() {
	f.Energia = f.Energia + 100
	fmt.Println("Se recupero la energia del jugador, ", f.Energia)
}

func PrepararFicha(j JugadorInscripto) {
	j.Info()
	j.JugarPartido()
	fmt.Println("¡Partido disputado!.El jugador ha agotado energia")
}

func main() {
	RL := Futbolista{
		Atleta: Atleta{
			Nombre:  "Robert Lewandowski",
			Energia: 100,
		},
	}

	PrepararFicha(&RL)
	RL.Recuperar()
}
