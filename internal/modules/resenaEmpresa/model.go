package resenaEmpresa

import "time"

type ResenaEmpresa struct {
	IdResenaEmpresa   uint       `gorm:"primaryKey;autoIncrement;column:id_resena_empresa" json:"idResenaEmpresa"`
	Calificacion      *int       `gorm:"type:integer;column:calificacion" json:"calificacion"`
	ClienteId         *uint      `gorm:"column:cliente_id" json:"clienteId"`
	FacturaId         *uint      `gorm:"column:factura_id" json:"facturaId"`
	Comentario        *string    `gorm:"type:text;column:comentario" json:"comentario"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
