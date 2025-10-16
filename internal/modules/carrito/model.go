package carrito

import "time"

type Carrito struct {
	IdCarrito         uint          `gorm:"primaryKey;autoIncrement;column:id_carrito" json:"idCarrito"`
	ClienteId         *uint         `gorm:"column:cliente_id" json:"clienteId"`
	SubTotal          *float64      `gorm:"type:decimal(10,2);column:sub_total" json:"subTotal"`
	Descuento         *float64      `gorm:"type:decimal(10,2);column:descuento" json:"descuento"`
	Impuesto          *float64      `gorm:"type:decimal(10,2);column:impuesto" json:"impuesto"`
	Total             *float64      `gorm:"type:decimal(10,2);column:total" json:"total"`
	FechaExpiracion   *time.Time    `gorm:"type:timestamp;column:fecha_expiracion" json:"fechaExpiracion"`
	PedidoId          *uint         `gorm:"column:pedido_id" json:"pedidoId"`
	Estado            *bool         `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time    `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time    `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
	Items             []CarritoItem `gorm:"foreignKey:CarritoId" json:"items,omitempty"`
}

type CarritoItem struct {
	IdCarritoItem     uint       `gorm:"primaryKey;autoIncrement;column:id_carrito_item" json:"idCarritoItem"`
	ProductoId        *uint      `gorm:"column:producto_id" json:"productoId"`
	Cantidad          *int       `gorm:"type:integer;column:cantidad" json:"cantidad"`
	PrecioUnitario    *float64   `gorm:"type:decimal(10,2);column:precio_unitario" json:"precioUnitario"`
	DescuentoUnitario *float64   `gorm:"type:decimal(10,2);column:descuento_unitario" json:"descuentoUnitario"`
	Subtotal          *float64   `gorm:"type:decimal(10,2);column:subtotal" json:"subtotal"`
	FechaAgregado     *time.Time `gorm:"type:timestamp;column:fecha_agregado" json:"fechaAgregado"`
	CarritoId         *uint      `gorm:"column:carrito_id" json:"carritoId"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
