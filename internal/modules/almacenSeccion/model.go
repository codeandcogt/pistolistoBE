package almacenseccion

import "time"

type AlmacenSeccion struct {
	IdAlmacenSeccion  uint       `json:"id_almacen_seccion" db:"id_almacen_seccion" gorm:"primaryKey;autoIncrement"`
	IdAlmacen         uint       `json:"id_almacen" db:"id_almacen" gorm:"column:id_almacen; not null"`
	IdSeccion         uint       `json:"id_seccion" db:"id_seccion" gorm:"column:id_seccion; not null"`
	CapacidadSeccion  float64    `json:"capacidad_seccion" db:"capacidad_seccion" gorm:"column:capacidad_seccion"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
