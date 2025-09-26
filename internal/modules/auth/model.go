package auth

import (
	"time"
)

type LogLoginCliente struct {
	IdLoginCliente    int64     `gorm:"primaryKey;autoIncrement;column:id_login_cliente" json:"idLoginCliente"`
	IdCliente         int       `gorm:"column:id_cliente;not null" json:"idCliente"`
	Estado            *bool     `gorm:"column:estado;default:true" json:"estado"`
	Exito             bool      `gorm:"column:exitoso;" json:"exitoso"`
	FechaModificacion time.Time `gorm:"column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     time.Time `gorm:"column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
