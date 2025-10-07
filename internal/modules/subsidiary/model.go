package subsidiary

import (
	"time"
)

type Subsidiary struct {
	IdSucursal        uint       `gorm:"primaryKey;autoIncrement;column:id_sucursal" json:"idSucursal"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Descripcion       *string    `gorm:"type:varchar(255);column:descripcion" json:"descripcion"`
	Codigo            string     `gorm:"type:varchar(50);uniqueIndex;not null;column:codigo" json:"codigo"`
	Telefono          *string    `gorm:"type:varchar(20);column:telefono" json:"telefono"`
	Email             *string    `gorm:"type:varchar(255);uniqueIndex;column:email" json:"email"`
	Direccion         *string    `gorm:"type:varchar(255);column:direccion" json:"direccion"`
	HoraApertura      *time.Time `gorm:"type:timestamp;column:hora_apertura" json:"hora_apertura"`
	HoraCierre        *time.Time `gorm:"type:timestamp;column:hora_cierre" json:"hora_cierre"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
