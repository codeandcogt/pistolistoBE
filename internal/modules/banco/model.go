package banco

import "time"

type Banco struct {
	IdBanco           uint       `gorm:"primaryKey;autoIncrement;column:id_banco" json:"idBanco"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Descripcion       *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	Direccion         *string    `gorm:"type:varchar(255);column:direccion" json:"direccion"`
	Telefono          *string    `gorm:"type:varchar(20);column:telefono" json:"telefono"`
	Email             *string    `gorm:"type:varchar(100);column:email" json:"email"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fechaModificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fechaCreacion"`
}
