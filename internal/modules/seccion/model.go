package seccion

import "time"

type Seccion struct {
	IdSeccion         uint       `json:"id_seccion" db:"id_seccion" gorm:"primaryKey;autoIncrement"`
	Nombre            string     `json:"nombre" db:"nombre" gorm:"type:varchar(100);not null"`
	Descripcion       string     `json:"descripcion" db:"descripcion" gorm:"type:text"`
	Codigo            string     `json:"codigo" db:"codigo" gorm:"type:varchar(50);not null"`
	TipoSeccion       string     `json:"tipo_seccion" db:"tipo_seccion" gorm:"type:varchar(50);not null"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
