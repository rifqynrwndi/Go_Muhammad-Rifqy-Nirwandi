package repositories

import "go-wishlist-api/entities"

type WishlistRepository interface {
	FindAll() ([]entities.Wishlist, error)
	Create(wishlist *entities.Wishlist) error
}
