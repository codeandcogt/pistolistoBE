package almacen

import "time"

type Almacen struct {
	IdAlmacen         uint       `json:"id_almacen" db:"id_almacen" gorm:"primaryKey;autoIncrement"`
	Nombre            string     `json:"nombre" db:"nombre" gorm:"type:varchar(100);not null"`
	Codigo            string     `json:"codigo" db:"codigo" gorm:"type:varchar(50);not null"`
	TipoAlmacen       string     `json:"tipo_almacen" db:"tipo_almacen" gorm:"type:varchar(50);not null"`
	CapacidadMaxima   float64    `json:"capacidad_maxima" db:"capacidad_maxima" gorm:"column:capacidad_maxima"`
	Descripcion       string     `json:"descripcion" db:"descripcion" gorm:"type:text"`
	IdSucursal        uint       `json:"id_sucursal" db:"id_sucursal" gorm:"column:id_sucursal;not null"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
