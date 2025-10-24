package inventario

import "time"

type Inventario struct {
	IdInventario      uint       `json:"id_inventario" db:"id_inventario" gorm:"primaryKey;autoIncrement"`
	Codigo            string     `json:"codigo" db:"codigo" gorm:"type:varchar(50);not null;unique"`
	TipoItem          string     `json:"tipo_item" db:"tipo_item" gorm:"type:varchar(20);not null"` // "producto" o "articulo"
	IdProducto        *uint      `json:"id_producto" db:"id_producto" gorm:"column:id_producto"`
	IdArticulo        *uint      `json:"id_articulo" db:"id_articulo" gorm:"column:id_articulo"`
	IdAlmacen         uint       `json:"id_almacen" db:"id_almacen" gorm:"not null"`
	Cantidad          int        `json:"cantidad" db:"cantidad" gorm:"default:1"`
	Ubicacion         string     `json:"ubicacion" db:"ubicacion" gorm:"type:varchar(100)"`                                     // Pasillo, estante, etc.
	EstadoInventario  string     `json:"estado_inventario" db:"estado_inventario" gorm:"type:varchar(50);default:'disponible'"` // disponible, vendido, empeñado, reservado
	FechaIngreso      *time.Time `json:"fecha_ingreso" db:"fecha_ingreso" gorm:"not null"`
	FechaSalida       *time.Time `json:"fecha_salida" db:"fecha_salida"`
	Observaciones     string     `json:"observaciones" db:"observaciones" gorm:"type:text"`
	Estado            *bool      `json:"estado" db:"estado" gorm:"default:true"`
	FechaModificacion *time.Time `json:"fecha_modificacion" db:"fecha_modificacion" gorm:"autoUpdateTime"`
	FechaCreacion     *time.Time `json:"fecha_creacion" db:"fecha_creacion" gorm:"autoCreateTime"`
}
