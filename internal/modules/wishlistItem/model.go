package wishlistitem

import "time"

type WishListItem struct {
	IdWishListItem    uint       `json:"id_wishlist_item" db:"id_wishlist_item" gorm:"primaryKey;autoIncrement"`
	IdWishlist        uint       `json:"id_wishlist" db:"id_wishlist" gorm:"column:id_wishlist;not null"`
	IdProducto        uint       `json:"id_producto" db:"id_producto" gorm:"column:id_producto;not null"`
	FechaAgregado     *time.Time `json:"fecha_agregado" db:"fecha_agregado" gorm:"autoUpdateTime"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
