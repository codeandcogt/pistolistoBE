package warehouse

import (
	"time"
)

type Warehouse struct {
	IdAlmacen         int        `gorm:"primaryKey;autoIncrement" json:"id_sucursal"`
	Nombre            string     `gorm:"type:varchar(100);not null" json:"nombre"`
	Codigo            string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"codigo"`
	TipoAlmacen       string     `gorm:"type:varchar(50);not null" json:"tipo_almacen"`
	CantidadMaxima    int        `gorm:"type:int;not null" json:"cantidad_maxima"`
	Descripcion       *string    `gorm:"type:varchar(255)" json:"descripcion"`
	Estado            *bool      `gorm:"type:boolean;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp" json:"fecha_creacion"`
}
