package rolpermiso

import "time"

type RolPermiso struct {
	IdRolPermiso      uint       `json:"id_rol_permiso" gorm:"primaryKey;autoIncrement;column:id_rol_permiso"`
	IdRol             uint       `json:"id_rol" gorm:"not null;column:id_rol;index"`
	IdPermiso         uint       `json:"id_permiso" gorm:"not null;column:id_permiso;index"`
	Estado            *bool      `json:"estado,omitempty" gorm:"column:estado;default:true"`
	FechaCreacion     *time.Time `json:"fecha_creacion,omitempty" gorm:"column:fecha_creacion;autoCreateTime"`
	FechaModificacion *time.Time `json:"fecha_modificacion,omitempty" gorm:"column:fecha_modificacion;autoUpdateTime"`
}
