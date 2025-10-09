package administrativo

import (
	"pistolistoBE/internal/modules/rol"
	"pistolistoBE/internal/modules/subsidiary"
	"time"
)

type Administrativo struct {
	IdAdministrativo  uint                   `json:"id_administrativo" gorm:"primaryKey;autoIncrement;column:id_administrativo"`
	Nombre            string                 `json:"nombre" gorm:"type:varchar(100);not null;column:nombre"`
	Apellido          string                 `json:"apellido" gorm:"type:varchar(100);not null;column:apellido"`
	CUI               string                 `json:"cui" gorm:"type:varchar(13);unique;not null;column:cui"`
	NIT               string                 `json:"nit" gorm:"type:varchar(15);column:nit"`
	IdRol             uint                   `json:"id_rol" gorm:"not null;column:id_rol;index"`
	Rol               *rol.Rol               `json:"rol,omitempty" gorm:"foreignKey:IdRol;references:IdRol"`
	Email             string                 `json:"email" gorm:"type:varchar(100);unique;not null;column:email"`
	Telefono          string                 `json:"telefono" gorm:"type:varchar(20);column:telefono"`
	Contrasenia       string                 `json:"contrasenia" gorm:"type:varchar(255);not null;column:contrasenia"`
	NombreUsuario     string                 `json:"nombre_usuario" gorm:"type:varchar(50);unique;not null;column:nombre_usuario"`
	IdSucursal        uint                   `json:"id_sucursal" gorm:"not null;column:id_sucursal;index"`
	Sucursal          *subsidiary.Subsidiary `json:"sucursal,omitempty" gorm:"foreignKey:IdSucursal;references:IdSucursal"`
	FechaNacimiento   time.Time              `json:"fecha_nacimiento" gorm:"type:date;not null;column:fecha_nacimiento"`
	Genero            string                 `json:"genero" gorm:"type:varchar(1);column:genero"`
	PrimerLogin       *bool                  `json:"primer_login,omitempty" gorm:"column:primer_login;default:false"`
	Estado            *bool                  `json:"estado,omitempty" gorm:"column:estado;default:true"`
	FechaCreacion     *time.Time             `json:"fecha_creacion,omitempty" gorm:"column:fecha_creacion;autoCreateTime"`
	FechaModificacion *time.Time             `json:"fecha_modificacion,omitempty" gorm:"column:fecha_modificacion;autoUpdateTime"`
}
