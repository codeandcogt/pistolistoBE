package pedido

import (
	"time"
)

type Pedido struct {
	IdPedido          uint       `gorm:"primaryKey;autoIncrement;column:id_pedido" json:"id_pedido"`
	NumeroPedido      string     `gorm:"type:varchar(50);uniqueIndex;not null;column:numero_pedido" json:"numero_pedido"`
	IdCarrito         int        `gorm:"not null;column:id_carrito" json:"id_carrito"`
	IdRuta            *int       `gorm:"column:id_ruta" json:"id_ruta"`
	Subtotal          float64    `gorm:"type:decimal(10,2);not null;column:subtotal" json:"subtotal"`
	Impuesto          float64    `gorm:"type:decimal(10,2);not null;column:impuesto" json:"impuesto"`
	CostoEnvio        float64    `gorm:"type:decimal(10,2);not null;column:costo_envio" json:"costo_envio"`
	Total             float64    `gorm:"type:decimal(10,2);not null;column:total" json:"total"`
	IdCliente         int        `gorm:"not null;column:id_cliente" json:"id_cliente"`
	IdCupon           *int       `gorm:"column:id_cupon" json:"id_cupon"`
	IdDireccion       int        `gorm:"not null;column:id_direccion" json:"id_direccion"`
	Observaciones     *string    `gorm:"type:text;column:observaciones" json:"observaciones"`
	FechaEnvio        *time.Time `gorm:"type:timestamp;column:fecha_envio" json:"fecha_envio"`
	IdEstadoPedido    int        `gorm:"not null;column:id_estado_pedido" json:"id_estado_pedido"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
