package cuota

import "time"

type Cuota struct {
	IdCuota           uint       `gorm:"primaryKey;autoIncrement;column:id_cuota" json:"idCuota"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Descripcion       *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	NumeroCuotas      int        `gorm:"not null;column:numero_cuotas" json:"numeroCuotas"`
	Estado            *bool      `gorm:"type:boolean;default:true;column:estado" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;autoUpdateTime;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;autoCreateTime;column:fecha_creacion" json:"fechaCreacion"`
}
