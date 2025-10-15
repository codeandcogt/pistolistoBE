package producto

import (
	"time"
)

type Producto struct {
	IdProducto        uint       `gorm:"primaryKey;autoIncrement;column:id_producto" json:"id_producto"`
	SKU               string     `gorm:"type:varchar(50);uniqueIndex;not null;column:sku" json:"sku"`
	Costo             float64    `gorm:"type:decimal(10,2);not null;column:costo" json:"costo"`
	Precio            float64    `gorm:"type:decimal(10,2);not null;column:precio" json:"precio"`
	UnidadMedida      string     `gorm:"type:varchar(20);not null;column:unidad_medida" json:"unidad_medida"`
	IdArticulo        int        `gorm:"not null;column:id_articulo" json:"id_articulo"`
	IdDescuento       *int       `gorm:"column:id_descuento" json:"id_descuento"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
