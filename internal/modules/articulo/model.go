package articulo

import (
	"time"
)

type Articulo struct {
	IdArticulo        uint       `gorm:"primaryKey;autoIncrement;column:id_articulo" json:"idArticulo"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Descripcion       *string    `gorm:"type:varchar(255);column:descripcion" json:"descripcion"`
	Peso              *float64   `gorm:"type:decimal(10,2);column:peso" json:"peso"`
	Dimension         *string    `gorm:"type:varchar(50);column:dimension" json:"dimension"`
	Color             *string    `gorm:"type:varchar(50);column:color" json:"color"`
	IdAlmacen         int        `gorm:"column:id_almacen;not null" json:"idAlmacen"`
	IdSubCategoria    int        `gorm:"column:id_sub_categoria;not null" json:"idSubCategoria"`
	Imagen            *string    `gorm:"type:varchar(255);column:imagen" json:"imagen"` // 🆕 Nueva propiedad
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
