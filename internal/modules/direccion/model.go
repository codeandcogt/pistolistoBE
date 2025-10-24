package direccion

import (
	"time"
)

type Direccion struct {
	IdDireccion       uint       `gorm:"primaryKey;autoIncrement;column:id_direccion" json:"id_direccion"`
	Linea1            string     `gorm:"type:varchar(255);not null;column:linea1" json:"linea1"`
	Linea2            *string    `gorm:"type:varchar(255);column:linea2" json:"linea2"`
	Referencia        *string    `gorm:"type:varchar(255);column:referencia" json:"referencia"`
	Latitud           *float64   `gorm:"type:decimal(10,8);column:latitud" json:"latitud"`
	Longitud          *float64   `gorm:"type:decimal(11,8);column:longitud" json:"longitud"`
	IdMunicipio       uint       `gorm:"not null;column:id_municipio" json:"id_municipio"`
	IdCliente         *int       `gorm:"column:id_cliente" json:"id_cliente"`
	IdAdministrativo  *int       `gorm:"column:id_administrativo" json:"id_administrativo"`
	TipoDireccion     *string    `gorm:"type:varchar(50);column:tipo_direccion" json:"tipo_direccion"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
