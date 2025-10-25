package prestamo

import "time"

type Prestamo struct {
	IdPrestamo        uint       `gorm:"primaryKey;autoIncrement;column:id_prestamo" json:"idPrestamo"`
	Monto             *float64   `gorm:"type:decimal(10,2);column:monto" json:"monto"`
	TasaInteres       *float64   `gorm:"type:decimal(5,2);column:tasa_interes" json:"tasaInteres"`
	PlazoMeses        *int       `gorm:"type:integer;column:plazo_meses" json:"plazoMeses"`
	FechaPrestamo     *time.Time `gorm:"type:date;column:fecha_prestamo" json:"fechaPrestamo"`
	FechaVencimiento  *time.Time `gorm:"type:date;column:fecha_vencimiento" json:"fechaVencimiento"`
	ClienteId         *uint      `gorm:"column:cliente_id" json:"clienteId"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
