package formulario

import "time"

type Formulario struct {
	IdFormulario      uint       `gorm:"primaryKey;autoIncrement;column:id_formulario" json:"idFormulario"`
	Nombres           *string    `gorm:"type:varchar(100);column:nombres" json:"nombres"`
	Apellidos         *string    `gorm:"type:varchar(100);column:apellidos" json:"apellidos"`
	Email             *string    `gorm:"type:varchar(100);column:email" json:"email"`
	Telefono          *string    `gorm:"type:varchar(20);column:telefono" json:"telefono"`
	Asunto            *string    `gorm:"type:varchar(200);column:asunto" json:"asunto"`
	Mensaje           *string    `gorm:"type:text;column:mensaje" json:"mensaje"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
