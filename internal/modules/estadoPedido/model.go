package estadoPedido

import (
	"time"
)

type EstadoPedido struct {
	IdEstadoPedido    uint       `gorm:"primaryKey;autoIncrement;column:id_estado_pedido" json:"id_estado_pedido"`
	Nombre            string     `gorm:"type:varchar(100);not null;column:nombre" json:"nombre"`
	Descripcion       *string    `gorm:"type:text;column:descripcion" json:"descripcion"`
	Color             *string    `gorm:"type:varchar(20);column:color" json:"color"`
	EstadoFinal       *bool      `gorm:"type:boolean;column:estado_final;default:false" json:"estado_final"`
	Estado            *bool      `gorm:"type:boolean;column:estado;default:true" json:"estado"`
	FechaModificacion *time.Time `gorm:"type:timestamp;column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
	FechaCreacion     *time.Time `gorm:"type:timestamp;column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
}
