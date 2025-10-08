package descuento

import "time"

type Descuento struct {
	IdDescuento       uint       `json:"id_descuento" db:"id_descuento" gorm:"primaryKey;autoIncrement"`
	Nombre            string     `json:"nombre" db:"nombre" gorm:"type:varchar(80);not null"`
	Porcentaje        float64    `json:"porcentaje" db:"porcentaje" gorm:"type:decimal(5,2)"`
	TipoDescuento     *string    `json:"tipo_descuento" db:"tipo_descuento" gorm:"type:varchar(100)"`
	Monto             float64    `json:"monto" db:"monto" gorm:"type:decimal(10,2)"`
	Descripcion       *string    `json:"descripcion" db:"descripcion" gorm:"type:text"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
