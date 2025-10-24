package logUbicacionTiempoReal

import "time"

type LogUbicacionTiempoReal struct {
	IdLogTiempoReal   uint       `gorm:"primaryKey;autoIncrement;column:id_log_tiempo_real" json:"id_log_tiempo_real"`
	PilotoId          int        `gorm:"not null;uniqueIndex;column:piloto_id" json:"piloto_id"`
	Latitud           float64    `gorm:"type:decimal(10,8);not null;column:latitud" json:"latitud"`
	Longitud          float64    `gorm:"type:decimal(11,8);not null;column:longitud" json:"longitud"`
	Estado            *bool      `gorm:"type:boolean;default:true;column:estado" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;autoUpdateTime;column:fecha_modificacion" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;autoCreateTime;column:fecha_creacion" json:"fecha_creacion"`
}
