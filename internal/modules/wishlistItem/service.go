package wishlistitem

type WishListItemService interface {
	AddWishListItem(wishListItem *WishListItem) error
	GetWishListItemByID(id uint) (*WishListItem, error)
	GetWishListItemByWishlist(IdWishlist uint) ([]*WishListItem, error)
	UpdateWishListItem(id uint, updated *WishListItem) (*WishListItem, error)
	DeleteWishListItem(id uint) (string, error)
}

type wishListItemService struct {
	repo WishListItemRepository
}

func NewWishListItemService(repo WishListItemRepository) WishListItemService {
	return &wishListItemService{repo}
}

func (s *wishListItemService) AddWishListItem(wishListItem *WishListItem) error {
	return s.repo.Create(wishListItem)
}

func (s *wishListItemService) GetWishListItemByID(id uint) (*WishListItem, error) {
	return s.repo.GetByID(id)
}

func (s *wishListItemService) GetWishListItemByWishlist(IdWishlist uint) ([]*WishListItem, error) {
	return s.repo.GetByWishlistID(IdWishlist)
}

func (s *wishListItemService) UpdateWishListItem(id uint, updated *WishListItem) (*WishListItem, error) {
	return s.repo.UpdateWishListItem(id, updated)
}

func (s *wishListItemService) DeleteWishListItem(id uint) (string, error) {
	return s.repo.DeleteWishListItem(id)
}
