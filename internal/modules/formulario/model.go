package formulario

import "time"

type Formulario struct {
	IdFormulario        uint       `gorm:"primaryKey;autoIncrement;column:id_formulario" json:"idFormulario"`
	IdCliente           int        `gorm:"not null;column:id_cliente" json:"idCliente"`
	IdArticulo          int        `gorm:"not null;column:id_articulo" json:"idArticulo"`
	NumeroFormulario    string     `gorm:"type:varchar(50);uniqueIndex;not null;column:numero_formulario" json:"numeroFormulario"`
	Descripcion         *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	MontoSolicitado     float64    `gorm:"type:decimal(10,2);not null;column:monto_solicitado" json:"montoSolicitado"`
	MontoAprobado       *float64   `gorm:"type:decimal(10,2);column:monto_aprobado" json:"montoAprobado"`
	TipoOperacion       string     `gorm:"type:varchar(50);not null;column:tipo_operacion" json:"tipoOperacion"`
	DetallesAdicionales *string    `gorm:"type:text;column:detalles_adicionales" json:"detallesAdicionales"`
	Estado              *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion   *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fechaModificacion"`
	FechaCreacion       *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fechaCreacion"`
}
