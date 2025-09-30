package permiso

import "time"

type Permiso struct {
	IdPermiso         int       `json:"id_permiso" db:"id_permiso" gorm:"primaryKey;autoIncrement"`
	Nombre            string    ``
	Accion            string    ``
	Descripcion       *string   ``
	Estado            bool      ``
	FechaCreacion     time.Time ``
	FechaModificacion time.Time ``
}
