package wishlistitem

import (
	"pistolistoBE/internal/common"

	"gorm.io/gorm"
)

type WishListItemRepository interface {
	Create(wishListItem *WishListItem) error
	GetByID(id uint) (*WishListItem, error)
	GetByWishlistID(wishlistId uint) ([]*WishListItem, error)
	UpdateWishListItem(id uint, updated *WishListItem) (*WishListItem, error)
	DeleteWishListItem(id uint) (string, error)
}

type wishListItemRepository struct {
	db *gorm.DB
}

func NewWishListItemRepository(db *gorm.DB) WishListItemRepository {
	return &wishListItemRepository{db}
}

func (r *wishListItemRepository) Create(wishListItem *WishListItem) error {
	return r.db.Create(wishListItem).Error
}

func (r *wishListItemRepository) GetByID(id uint) (*WishListItem, error) {
	var item WishListItem
	err := r.db.Where("estado = ?", true).First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *wishListItemRepository) GetByWishlistID(wishlistId uint) ([]*WishListItem, error) {
	var items []*WishListItem
	err := r.db.Where("wishlist_id = ? AND estado = ?", wishlistId, true).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *wishListItemRepository) UpdateWishListItem(id uint, updated *WishListItem) (*WishListItem, error) {
	var item WishListItem
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}

	item.IdProducto = updated.IdProducto
	item.Estado = updated.Estado
	item.FechaAgregado = updated.FechaAgregado
	item.FechaModificacion = updated.FechaModificacion

	err = r.db.Save(&item).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *wishListItemRepository) DeleteWishListItem(id uint) (string, error) {
	result := r.db.Model(&WishListItem{}).Where("id_wish_list_item = ?", id).Update("estado", false)

	if result.Error != nil {
		return common.ERR_DATABASE_ERROR, result.Error
	}

	if result.RowsAffected == 0 {
		return common.ERR_NOT_FOUND, gorm.ErrRecordNotFound
	}

	return common.SUCCESS_DELETED, nil
}
