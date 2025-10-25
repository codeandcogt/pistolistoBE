package avaluo

import "time"

type Avaluo struct {
	IdAvaluo          uint       `gorm:"primaryKey;autoIncrement;column:id_avaluo" json:"idAvaluo"`
	IdUsuario         int        `gorm:"not null;column:id_usuario" json:"idUsuario"`
	IdInteres         int        `gorm:"not null;column:id_interes" json:"idInteres"`
	MontoAvaluado     float64    `gorm:"type:decimal(10,2);not null;column:monto_avaluado" json:"montoAvaluado"`
	Observacion       *string    `gorm:"type:text;column:observacion" json:"observacion"`
	FechaAvaluo       time.Time  `gorm:"type:date;not null;column:fecha_avaluo" json:"fechaAvaluo"`
	IdFormulario      int        `gorm:"not null;column:id_formulario" json:"idFormulario"`
	IdRuta            *int       `gorm:"column:id_ruta" json:"idRuta"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fechaCreacion"`
}
