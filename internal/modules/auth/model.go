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

type LogLoginAdmin struct {
	IdLoginAdmin      uint      `json:"id_login_admin" gorm:"primaryKey;autoIncrement;column:id_login_admin"`
	IdAdministrativo  uint      `json:"id_administrativo" gorm:"not null;column:id_administrativo;index"`
	Estado            *bool     `json:"estado,omitempty" gorm:"column:estado;default:true"`
	Exito             bool      `json:"exito" gorm:"column:exito;not null;default:false"`
	FechaCreacion     time.Time `json:"fecha_creacion" gorm:"column:fecha_creacion;autoCreateTime"`
	FechaModificacion time.Time `json:"fecha_modificacion" gorm:"column:fecha_modificacion;autoUpdateTime"`
}
