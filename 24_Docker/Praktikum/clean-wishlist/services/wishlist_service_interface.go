package services

import "go-wishlist-api/entities"

type WishlistService interface {
	GetAllWishlists() ([]entities.Wishlist, error)
	CreateWishlist(wishlist *entities.Wishlist) error
}
