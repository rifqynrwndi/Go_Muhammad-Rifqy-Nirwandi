package services

import (
	"go-wishlist-api/entities"
	"go-wishlist-api/repositories"
)

type WishlistService interface {
	GetAll() ([]entities.Wishlist, error)
	Create(wishlist entities.Wishlist) (entities.Wishlist, error)
}

type wishlistService struct {
	repo repositories.WishlistRepository
}

func NewWishlistService(repo repositories.WishlistRepository) WishlistService {
	return &wishlistService{repo: repo}
}

func (s *wishlistService) GetAll() ([]entities.Wishlist, error) {
	return s.repo.FindAll()
}

func (s *wishlistService) Create(wishlist entities.Wishlist) (entities.Wishlist, error) {
	err := s.repo.Create(wishlist)
	return wishlist, err
}
