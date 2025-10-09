package wishlist

type WishlistService interface {
	CreateWishlist(wishlist *Wishlist) error
	GetWishlistByID(id uint) (*Wishlist, error)
	GetAllByCliente(IdCliente uint) ([]*Wishlist, error)
	UpdateWishlist(id uint, updated *Wishlist) (*Wishlist, error)
	DeleteWishlist(id uint) (string, error)
}

type wishlistService struct {
	repo WishlistRepository
}

func NewWishlistService(repo WishlistRepository) WishlistService {
	return &wishlistService{repo}
}

func (s *wishlistService) CreateWishlist(wishlist *Wishlist) error {
	return s.repo.Create(wishlist)
}

func (s *wishlistService) GetWishlistByID(id uint) (*Wishlist, error) {
	return s.repo.GetByID(id)
}

func (s *wishlistService) GetAllByCliente(IdCliente uint) ([]*Wishlist, error) {
	return s.repo.GetAllByCliente(IdCliente)
}

func (s *wishlistService) UpdateWishlist(id uint, updated *Wishlist) (*Wishlist, error) {
	return s.repo.UpdateWishlist(id, updated)
}

func (s *wishlistService) DeleteWishlist(id uint) (string, error) {
	return s.repo.DeleteWishlist(id)
}
