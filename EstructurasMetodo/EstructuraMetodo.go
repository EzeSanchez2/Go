package main

import (
	"fmt"
	"math"
)

type Rectangulo struct {
	ancho, alto float64
}

type Circulo struct {
	radio float64
}

func (c Circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}

func (r Rectangulo) area() float64 {
	return r.ancho * r.alto
}
func (a Rectangulo) incr(n float64) Rectangulo {
	return Rectangulo{
		a.ancho * n,
		a.alto * n,
	}
}
func (r *Rectangulo) inc(n float64) {
	r.alto *= n
	r.ancho *= n
}

func main() {
	r1 := Rectangulo{
		10,
		12,
	}

	r2 := Rectangulo{
		5,
		2,
	}

	c1 := Circulo{
		2,
	}

	fmt.Println("Area del r1: ", r1.area())
	fmt.Println("Area de r2: ", r2.area())
	fmt.Println("Area del circulo: ", c1.area())

	//Puntero

	r5 := Rectangulo{10, 5}
	fmt.Println(r5)
	r5.inc(5)
	fmt.Println(r5)
}
