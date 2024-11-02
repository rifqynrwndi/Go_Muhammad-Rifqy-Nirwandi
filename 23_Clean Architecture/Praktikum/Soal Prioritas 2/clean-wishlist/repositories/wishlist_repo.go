package repositories

import (
	"go-wishlist-api/entities"
	"gorm.io/gorm"
)

type WishlistRepository interface {
	FindAll() ([]entities.Wishlist, error)
	Create(wishlist entities.Wishlist) error
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	return &wishlistRepository{db: db}
}

func (r *wishlistRepository) FindAll() ([]entities.Wishlist, error) {
	var wishlists []entities.Wishlist
	err := r.db.Find(&wishlists).Error
	return wishlists, err
}

func (r *wishlistRepository) Create(wishlist entities.Wishlist) error {
	return r.db.Create(&wishlist).Error
}
