package main

import (
	"fmt"
)

type Base struct {
	ID int
}

type Producto struct {
	Base
	ID string
}

func main() {
	producto := Producto{
		Base{
			ID: 1,
		},
		"PROD-100",
	}
	fmt.Println(producto.ID)
	fmt.Println(producto.Base.ID)
}
