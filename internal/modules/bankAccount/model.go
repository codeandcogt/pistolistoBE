package bankAccount

import (
	"time"
)

type BankAccount struct {
	IdCuentaBancaria  uint       `gorm:"primaryKey;autoIncrement;column:id_cuenta_bancaria" json:"idCuentaBancaria"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Numero            string     `gorm:"type:varchar(50);uniqueIndex;not null;column:numero" json:"numero"`
	Tipo              string     `gorm:"type:varchar(50);not null;column:tipo" json:"tipo"`
	Descripcion       *string    `gorm:"type:varchar(255);column:descripcion" json:"descripcion"`
	IdCliente         int        `gorm:"column:id_cliente;not null" json:"idCliente"`
	IdMoneda          int        `gorm:"column:id_moneda;not null" json:"idMoneda"`
	IdBanco           int        `gorm:"column:id_banco;not null" json:"idBanco"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
