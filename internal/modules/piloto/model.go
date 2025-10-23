package piloto

import "time"

type Piloto struct {
	IdPiloto                 uint       `gorm:"primaryKey;autoIncrement;column:id_piloto" json:"idPiloto"`
	Nombres                  *string    `gorm:"type:varchar(100);column:nombres" json:"nombres"`
	Apellidos                *string    `gorm:"type:varchar(100);column:apellidos" json:"apellidos"`
	Telefono                 *string    `gorm:"type:varchar(20);column:telefono" json:"telefono"`
	Email                    *string    `gorm:"type:varchar(100);column:email" json:"email"`
	NumeroLicencia           *string    `gorm:"type:varchar(50);column:numero_licencia" json:"numeroLicencia"`
	TipoLicencia             *string    `gorm:"type:varchar(20);column:tipo_licencia" json:"tipoLicencia"`
	FechaVencimientoLicencia *time.Time `gorm:"type:date;column:fecha_vencimiento_licencia" json:"fechaVencimientoLicencia"`
	AdministrativoId         *uint      `gorm:"column:administrativo_id" json:"administrativoId"`
	Estado                   *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion        *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion            *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
