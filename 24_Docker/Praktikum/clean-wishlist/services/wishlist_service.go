package services

import (
	"go-wishlist-api/entities"
	"go-wishlist-api/repositories"
)

type wishlistService struct {
	repo repositories.WishlistRepository
}

func NewWishlistService(repo repositories.WishlistRepository) WishlistService {
	return &wishlistService{repo: repo}
}

func (s *wishlistService) GetAllWishlists() ([]entities.Wishlist, error) {
	return s.repo.FindAll()
}

func (s *wishlistService) CreateWishlist(wishlist *entities.Wishlist) error {
	return s.repo.Create(wishlist)
}
