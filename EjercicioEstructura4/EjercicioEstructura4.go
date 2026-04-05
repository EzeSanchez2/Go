package main

import "fmt"

type Rectangulo struct {
	Ancho float64
	Alto  float64
}

func (r Rectangulo) Area() float64 {
	return r.Alto * r.Ancho
}
func (r *Rectangulo) Escalar(factor float64) {
	r.Alto *= factor
	r.Ancho *= factor

}
func main() {
	r1 := Rectangulo{
		Ancho: 2,
		Alto:  5,
	}
	fmt.Println("Area: ", r1.Area())
	r1.Escalar(3)
	fmt.Println("Alto Actual: ", r1.Alto)
	fmt.Println("Ancho Actual: ", r1.Ancho)
	fmt.Println("Area actual: ", r1.Area())
}
