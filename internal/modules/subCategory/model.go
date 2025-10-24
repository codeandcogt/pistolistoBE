package subCategory

import "time"

type SubCategory struct {
	IdSubCategoria    uint       `gorm:"primaryKey;autoIncrement;column:id_subcategoria" json:"idSubCategoria"`
	Nombre            string     `gorm:"type:varchar(80);not null;column:nombre" json:"nombre"`
	Descripcion       *string    `gorm:"type:varchar(255);column:descripcion" json:"descripcion"`
	IdCategoria       uint       `gorm:"column:id_categoria;not null" json:"idCategoria"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
