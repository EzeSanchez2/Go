package main

import (
	"fmt"
)

type Guerrero interface {
	Atacar() int
	obtenerNombres() string
}

type Saiyajin struct {
	Nombre     string
	NivelPoder int
}

type Androide struct {
	NombreAndroide string
	Bateria        int
}

func (s Saiyajin) Atacar() int {
	return s.NivelPoder * 10
}

func (a *Androide) Atacar() int {
	a.Bateria = a.Bateria - 10
	return 800
}

func (s *Saiyajin) Entrenar() int {
	return s.NivelPoder + 50
}

func (s Saiyajin) obtenerNombres() string {
	return s.Nombre
}

func (a Androide) obtenerNombres() string {
	return a.NombreAndroide
}
func PresentarPeleador(g Guerrero) {
	daño := g.Atacar()
	nombre := g.obtenerNombres()
	fmt.Printf("¡%d entra a la plataforma y lanza un ataque de %q de daño!", daño, nombre)
}
func main() {
	goku := Saiyajin{
		Nombre:     "Goku",
		NivelPoder: 100,
	}
	numero18 := Androide{
		NombreAndroide: "Numero18",
		Bateria:        50,
	}

	goku.Entrenar()
	PresentarPeleador(goku)
	PresentarPeleador(&numero18)
}
