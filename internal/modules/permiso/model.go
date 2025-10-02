package permiso

import "time"

type Permiso struct {
	IdPermiso         int       `json:"id_permiso" db:"id_permiso" gorm:"primaryKey;autoIncrement"`
	Nombre            string    `json:"nombre" db:"nombre" gorm:"type:varchar(100);not null"`
	Accion            string    `json:"accion" db:"accion" gorm:"type:varchar(100);not null"`
	Descripcion       *string   `json:"descripcion" db:"descripcion" gorm:"type:text"`
	Estado            bool      `json:"estado" db:"estado" gorm:"default:true"`
	FechaCreacion     time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
	FechaModificacion time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
}
