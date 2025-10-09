package wishlist

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type WishlistRepository interface {
	Create(wishlist *Wishlist) error
	GetByID(id uint) (*Wishlist, error)
	GetAllByCliente(clienteId uint) ([]*Wishlist, error)
	DeleteWishlist(id uint) (string, error)
	UpdateWishlist(id uint, updated *Wishlist) (*Wishlist, error)
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlist(db *gorm.DB) WishlistRepository {
	return &wishlistRepository{db}
}

func (r *wishlistRepository) Create(wishlist *Wishlist) error {
	return r.db.Create(wishlist).Error
}

func (r *wishlistRepository) GetByID(id uint) (*Wishlist, error) {
	var wishlist Wishlist
	err := r.db.Where("estado = ?", true).First(&wishlist, id).Error
	if err != nil {
		return nil, err
	}
	return &wishlist, nil
}

func (r *wishlistRepository) GetAllByCliente(clienteId uint) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	err := r.db.Where("cliente_id = ? AND estado = ?", clienteId, true).Find(&wishlists).Error
	if err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (r *wishlistRepository) DeleteWishlist(id uint) (string, error) {
	result := r.db.Model(&Wishlist{}).Where("id_wishlist = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}

func (r *wishlistRepository) UpdateWishlist(id uint, updated *Wishlist) (*Wishlist, error) {
	var wishlist Wishlist
	err := r.db.First(&wishlist, id).Error
	if err != nil {
		return nil, err
	}

	// Actualiza los campos deseados
	wishlist.Nombre = updated.Nombre
	wishlist.Estado = updated.Estado
	wishlist.IdCliente = updated.IdCliente
	wishlist.FechaModificacion = updated.FechaModificacion

	err = r.db.Save(&wishlist).Error
	if err != nil {
		return nil, err
	}

	return &wishlist, nil
}
