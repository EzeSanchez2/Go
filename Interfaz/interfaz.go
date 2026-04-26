package main

import "fmt"

type Usuario struct {
	nombre string
	email  string
}

type UsuarioInstagram struct {
	Usuario
	tipo string
}

type UsuarioFacebook struct {
	Usuario
	frase string
}

func (u Usuario) LoginNombre() {
	fmt.Println("Nombre: ", u.nombre)
}

func (u Usuario) LoginEmail() {
	fmt.Println("Email: ", u.email)
}

func (ui UsuarioInstagram) DecirTipo() {
	fmt.Println("Tipo de: ", ui.tipo)
}

func (uf UsuarioFacebook) Descripcion() {
	fmt.Println("Descripcion: ", uf.frase)
}

func Presentarse(a IniciarSesion) {
	a.LoginNombre()
	a.LoginEmail()
	a.DecirTipo()
}

type IniciarSesion interface {
	LoginNombre()
	LoginEmail()
	DecirTipo()
}

func main() {

	u2 := UsuarioInstagram{
		Usuario{
			nombre: "Ale",
			email:  "Ale@gmail.com",
		},
		"Instagram",
	}

	Presentarse(u2)

}
