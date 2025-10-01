package categoria

import "time"

type Categoria struct {
	IdCategoria       uint       `json:"id_categoria" db:"id_categoria" gorm:"primaryKey;autoIncrement"`
	Nombre            string     `json:"nombre" db:"nombre" gorm:"type:varchar(80);column:nombre"`
	Descripcion       *string    `json:"descripcion,omitempty" db:"descripcion"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true;"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
