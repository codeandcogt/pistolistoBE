package departamento

import "time"

type Departamento struct {
	IdDepartamento    uint       `json:"id_departamento" db:"id_Departamento" gorm:"primaryKey;autoIncrement"`
	Nombre            string     `json:"nombre" db:"nombre" gorm:"type:varchar(80);not null;column:nombre"`
	Codigo            string     `json:"codigo" db:"codigo" gorm:"type:varchar(20);uniqueIndex"`
	Descripcion       *string    `json:"descripcion,omitempty" db:"descripcion" gorm:"type:text"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true;"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
