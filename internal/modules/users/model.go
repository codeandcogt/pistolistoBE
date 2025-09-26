package users

import "time"

type Usuario struct {
	IdUsuario         string
	Nombre            string
	Apellido          string
	CUI               string
	NIT               string
	IdRol             int64
	Email             string
	Telefono          string
	NombreUsuario     string
	IdSucursal        int64
	FechaNacimiento   time.Time
	Genero            string
	PrimerLogIn       bool
	Estado            bool
	FechaModificacion time.Time
}
