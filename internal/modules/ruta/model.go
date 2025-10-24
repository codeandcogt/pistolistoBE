package ruta

import (
	"time"
)

type Ruta struct {
	IdRuta            uint       `gorm:"primaryKey;autoIncrement;column:id_ruta" json:"id_ruta"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Descripcion       *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	TiempoEstimado    *int       `gorm:"type:int;column:tiempo_estimado" json:"tiempo_estimado"`
	IdVehiculo        *int       `gorm:"column:id_vehiculo" json:"id_vehiculo"`
	IdEstadoRuta      *int       `gorm:"column:id_estado_ruta" json:"id_estado_ruta"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
