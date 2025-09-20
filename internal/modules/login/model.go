package login

import "time"

type IntentoLogin struct {
	IdIntentoLogIn    int
	Dispositivo       string
	MotivoFallo       string
	IdCliente         int
	IdUsuario         int
	IP                string
	Exitoso           bool
	Estado            bool
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

type LogLoginUsuario struct {
	IdLogInUsuario    int
	IdUsuario         int
	IP                string
	Dispositivo       string
	TiempoSesion      int
	Estado            bool
	FechaCreacion     time.Time
	FechaModificacion time.Time
}
