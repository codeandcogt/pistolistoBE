package vehiculo

import "time"

type Vehiculo struct {
	IdVehiculo               uint       `gorm:"primaryKey;autoIncrement;column:id_vehiculo" json:"idVehiculo"`
	Modelo                   *string    `gorm:"type:varchar(100);column:modelo" json:"modelo"`
	Descripcion              *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	Marca                    *string    `gorm:"type:varchar(100);column:marca" json:"marca"`
	Placa                    *string    `gorm:"type:varchar(20);column:placa" json:"placa"`
	Anio                     *int       `gorm:"type:integer;column:anio" json:"anio"`
	Tipo                     *string    `gorm:"type:varchar(50);column:tipo" json:"tipo"`
	Capacidad                *int       `gorm:"type:integer;column:capacidad" json:"capacidad"`
	CapacidadPeso            *float64   `gorm:"type:decimal(10,2);column:capacidad_peso" json:"capacidadPeso"`
	Combustible              *string    `gorm:"type:varchar(50);column:combustible" json:"combustible"`
	Kilometraje              *int       `gorm:"type:integer;column:kilometraje" json:"kilometraje"`
	FechaUltimoMantenimiento *time.Time `gorm:"type:date;column:fecha_ultimo_mantenimiento" json:"fechaUltimoMantenimiento"`
	PilotoId                 *uint      `gorm:"column:piloto_id" json:"pilotoId"`
	Estado                   *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion        *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion            *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
