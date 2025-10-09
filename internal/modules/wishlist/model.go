package wishlist

import "time"

type Wishlist struct {
	IdWishlist        uint       `json:"id_wishlist" db:"id_wishlist" gorm:"primaryKey;autoIncrement"`
	Nombre            string     `json:"nombre" db:"nombre" gorm:"varchar(80);column:nombre"`
	IdCliente         uint       `json:"id_cliente" db:"id_cliente" gorm:"column:id_cliente; not null"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true;"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
