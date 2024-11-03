package repositories

import (
	"go-wishlist-api/entities"
	"gorm.io/gorm"
)

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	return &wishlistRepository{db: db}
}

func (r *wishlistRepository) FindAll() ([]entities.Wishlist, error) {
	var wishlists []entities.Wishlist
	if err := r.db.Find(&wishlists).Error; err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (r *wishlistRepository) Create(wishlist *entities.Wishlist) error {
	return r.db.Create(wishlist).Error
}
