package pago

import (
	"time"
)

type Pago struct {
	IdPago            uint       `gorm:"primaryKey;autoIncrement;column:id_pago" json:"id_pago"`
	Monto             float64    `gorm:"type:decimal(10,2);not null;column:monto" json:"monto"`
	Metodo            string     `gorm:"type:varchar(50);not null;column:metodo" json:"metodo"`
	Observacion       *string    `gorm:"type:text;column:observacion" json:"observacion"`
	Autorizacion      *string    `gorm:"type:varchar(100);column:autorizacion" json:"autorizacion"`
	FechaAutorizacion *time.Time `gorm:"type:timestamp;column:fecha_autorizacion" json:"fecha_autorizacion"`
	IdMoneda          int        `gorm:"not null;column:id_moneda" json:"id_moneda"`
	IdPedido          int        `gorm:"not null;column:id_pedido" json:"id_pedido"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
