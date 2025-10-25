package loan

import "time"

type Loan struct {
	IdPrestamo        uint       `gorm:"primaryKey;autoIncrement;column:id_prestamo" json:"idPrestamo"`
	MontoPrestado     float64    `gorm:"type:decimal(10,2);not null;column:monto_prestado" json:"montoPrestado"`
	IdCuota           *int       `gorm:"column:id_cuota" json:"idCuota"`
	IdContrato        int        `gorm:"not null;column:id_contrato" json:"idContrato"`
	CobroDespuesXDias *int       `gorm:"column:cobro_despues_x_dias" json:"cobroDespuesXDias"`
	Estado            *bool      `gorm:"type:boolean;default:true;column:estado" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;autoUpdateTime;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;autoCreateTime;column:fecha_creacion" json:"fechaCreacion"`
}
