package cupon

import "time"

type Cupon struct {
	IdCupon             uint       `gorm:"primaryKey;autoIncrement;column:id_cupon" json:"idCupon"`
	Nombre              string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Codigo              string     `gorm:"type:varchar(50);uniqueIndex;not null;column:codigo" json:"codigo"`
	Descripcion         *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	TipoCupon           string     `gorm:"type:varchar(50);not null;column:tipo_cupon" json:"tipoCupon"`
	MontoDescuento      *float64   `gorm:"type:decimal(10,2);column:monto_descuento" json:"montoDescuento"`
	PorcentajeDescuento *float64   `gorm:"type:decimal(5,2);column:porcentaje_descuento" json:"porcentajeDescuento"`
	MontoMaximo         *float64   `gorm:"type:decimal(10,2);column:monto_maximo" json:"montoMaximo"`
	MontoMinimo         *float64   `gorm:"type:decimal(10,2);column:monto_minimo" json:"montoMinimo"`
	FechaInicio         *time.Time `gorm:"type:date;column:fecha_inicio" json:"fechaInicio"`
	FechaVencimiento    *time.Time `gorm:"type:date;column:fecha_vencimiento" json:"fechaVencimiento"`
	CantidadUso         *int       `gorm:"type:integer;column:cantidad_uso" json:"cantidadUso"`
	Estado              *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion   *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion       *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
