package factura

import "time"

type Factura struct {
	IdFactura         uint       `gorm:"primaryKey;autoIncrement;column:id_factura" json:"id_factura"`
	Numero            string     `gorm:"type:varchar(50);uniqueIndex;not null;column:numero" json:"numero"`
	Serie             string     `gorm:"type:varchar(10);not null;column:serie" json:"serie"`
	RegimenFiscal     string     `gorm:"type:varchar(100);not null;column:regimen_fiscal" json:"regimen_fiscal"`
	UUID              string     `gorm:"type:varchar(100);uniqueIndex;not null;column:uuid" json:"uuid"`
	FechaEmision      time.Time  `gorm:"type:date;not null;column:fecha_emision" json:"fecha_emision"`
	DireccionFiscal   string     `gorm:"type:varchar(255);not null;column:direccion_fiscal" json:"direccion_fiscal"`
	RazonSocial       string     `gorm:"type:varchar(255);not null;column:razon_social" json:"razon_social"`
	Regimen           string     `gorm:"type:varchar(100);not null;column:regimen" json:"regimen"`
	Tipo              string     `gorm:"type:varchar(50);not null;column:tipo" json:"tipo"` // Ej: factura, nota crédito, etc.
	Descuento         float64    `gorm:"type:decimal(10,2);column:descuento" json:"descuento"`
	Subtotal          float64    `gorm:"type:decimal(10,2);not null;column:subtotal" json:"subtotal"`
	Impuestos         float64    `gorm:"type:decimal(10,2);not null;column:impuestos" json:"impuestos"`
	IdMoneda          int        `gorm:"not null;column:id_moneda" json:"id_moneda"`
	Total             float64    `gorm:"type:decimal(10,2);not null;column:total" json:"total"`
	IdPedido          int        `gorm:"not null;column:id_pedido" json:"id_pedido"`
	PDFUrl            *string    `gorm:"type:varchar(255);column:pdf_url" json:"pdf_url"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
