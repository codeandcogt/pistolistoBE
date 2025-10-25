package contrato

import "time"

type Contrato struct {
	IdContrato            uint       `gorm:"primaryKey;autoIncrement;column:id_contrato" json:"idContrato"`
	TipoContrato          string     `gorm:"type:varchar(50);not null;column:tipo_contrato" json:"tipoContrato"`
	CondicionesEspeciales *string    `gorm:"type:text;column:condiciones_especiales" json:"condicionesEspeciales"`
	IdAvaluo              int        `gorm:"not null;column:id_avaluo" json:"idAvaluo"`
	Estado                *bool      `gorm:"type:boolean;default:true;column:estado" json:"estado"`
	FechaModificacion     *time.Time `gorm:"type:timestamp;autoUpdateTime;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion         *time.Time `gorm:"type:timestamp;autoCreateTime;column:fecha_creacion" json:"fechaCreacion"`
}
