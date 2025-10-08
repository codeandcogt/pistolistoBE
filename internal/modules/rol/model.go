package rol

import "time"

type Rol struct {
	IdRol             uint       `json:"id_rol" db:"id_rol" gorm:"primaryKey;autoIncrement"`
	Nombre            string     `json:"nombre" db:"nombre" gorm:"not null; varchar(80); uniqueIndex"`
	Descripcion       *string    `json:"descripcion,omitempty" db:"descripcion"`
	Nivel             int        `json:"nivel" db:"nivel" gorm:"not null"`
	Estado            bool       `json:"estado" db:"estado" gorm:"default:true"`
	FechaCreacion     *time.Time `json:"fecha_creacion,omitempty" gorm:"column:fecha_creacion;autoCreateTime"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
}
