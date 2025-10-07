package municipality

import (
	"time"
)

type Municipality struct {
	IdMunicipio       uint       `gorm:"primaryKey;autoIncrement;column:id_municipio" json:"idMunicipio"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Codigo            string     `gorm:"type:varchar(50);uniqueIndex;not null;column:codigo" json:"codigo"`
	Descripcion       *string    `gorm:"type:varchar(255);column:descripcion" json:"descripcion"`
	IdDepartamento    int        `gorm:"column:id_departamento;not null" json:"idDepartamento"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
