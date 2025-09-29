package banco

import "time"

type Banco struct {
	IdBanco           uint       `gorm:"primaryKey;autoIncrement;column:id_banco" json:"idBanco"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	CodigoBanco       string     `gorm:"type:varchar(10);uniqueIndex;not null;column:codigo_banco" json:"codigo_banco"`
	CodigoSwift       *string    `gorm:"type:varchar(11);column:codigo_swift" json:"codigo_swift"`
	Direccion         *string    `gorm:"type:text;column:direccion" json:"direccion"`
	Telefono          *string    `gorm:"type:varchar(20);column:telefono" json:"telefono"`
	Email             *string    `gorm:"type:varchar(255);column:email" json:"email"`
	SitioWeb          *string    `gorm:"type:varchar(255);column:sitio_web" json:"sitio_web"`
	TipoBanco         string     `gorm:"type:varchar(50);not null;column:tipo_banco;default:'COMERCIAL'" json:"tipo_banco"` // COMERCIAL, COOPERATIVO, PUBLICO
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion" json:"fecha_creacion"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion" json:"fecha_modificacion"`
}
