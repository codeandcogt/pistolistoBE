package moneda

import "time"

type Moneda struct {
	IdMoneda          uint       `gorm:"primaryKey;autoIncrement;column:id_moneda" json:"idMoneda"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Codigo            string     `gorm:"type:varchar(10);uniqueIndex;not null;column:codigo" json:"codigo"`
	Simbolo           string     `gorm:"type:varchar(10);not null;column:simbolo" json:"simbolo"`
	TasaCambio        float64    `gorm:"type:decimal(10,4);not null;column:tasa_cambio" json:"tasaCambio"`
	Descripcion       *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
