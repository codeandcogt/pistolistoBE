package moneda

import "time"

type Moneda struct {
	IdMoneda          uint       `gorm:"primaryKey;autoIncrement;column:id_moneda" json:"idMoneda"`
	Codigo            string     `gorm:"type:varchar(3);uniqueIndex;not null;column:codigo" json:"codigo"`
	Nombre            string     `gorm:"type:varchar(50);not null;column:nombre" json:"nombre"`
	Simbolo           string     `gorm:"type:varchar(5);not null;column:simbolo" json:"simbolo"`
	TipoCambio        float64    `gorm:"type:decimal(10,4);not null;column:tipo_cambio;default:1.0000" json:"tipo_cambio"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fecha_creacion"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fecha_modificacion"`
}
