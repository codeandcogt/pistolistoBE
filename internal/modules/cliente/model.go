package cliente

import "time"

type Cliente struct {
	IdCliente         uint       `gorm:"primaryKey;autoIncrement;column:id_cliente" json:"idCliente"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Apellido          string     `gorm:"type:varchar(100);not null;column:apellido" json:"apellido"`
	CUI               *string    `gorm:"type:varchar(20);uniqueIndex;column:cui" json:"cui"`
	Nit               *string    `gorm:"type:varchar(20);column:nit" json:"nit"`
	Email             string     `gorm:"type:varchar(255);uniqueIndex;column:email" json:"email"`
	Telefono          *string    `gorm:"type:varchar(20);column:telefono" json:"telefono"`
	NombreUsuario     string     `gorm:"type:varchar(50);uniqueIndex;not null;column:nombreUsuario" json:"nombreUsuario"`
	Contrasena        string     `gorm:"type:varchar(255);not null;column:contrasena" json:"contrasena"`
	FechaNacimiento   *time.Time `gorm:"type:date;column:fecha_nacimiento" json:"fecha_nacimiento"`
	Genero            *string    `gorm:"type:varchar(20);column:genero" json:"genero"`
	TipoCliente       *string    `gorm:"type:varchar(50);column:tipo_cliente" json:"tipo_cliente"`
	PrimerLogin       *time.Time `gorm:"type:timestamp;column:primer_login" json:"primer_login"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
