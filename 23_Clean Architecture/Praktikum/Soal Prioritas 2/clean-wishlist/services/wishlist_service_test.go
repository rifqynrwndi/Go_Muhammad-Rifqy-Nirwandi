package services

import (
	"errors"
	"go-wishlist-api/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockWishlistRepository struct {
	mock.Mock
}

func (m *MockWishlistRepository) FindAll() ([]entities.Wishlist, error) {
	args := m.Called()
	return args.Get(0).([]entities.Wishlist), args.Error(1)
}

func (m *MockWishlistRepository) Create(wishlist *entities.Wishlist) error {
	args := m.Called(wishlist)
	return args.Error(0)
}

// Unit tests
func TestGetAllWishlists(t *testing.T) {
	mockRepo := new(MockWishlistRepository)
	service := NewWishlistService(mockRepo)

	mockRepo.On("FindAll").Return([]entities.Wishlist{{ID: 1, Title: "Test Wishlist", IsAchieved: false}}, nil)

	result, err := service.GetAllWishlists()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "Test Wishlist", result[0].Title)
	mockRepo.AssertExpectations(t)
}

func TestCreateWishlist(t *testing.T) {
	mockRepo := new(MockWishlistRepository)
	service := NewWishlistService(mockRepo)

	wishlist := &entities.Wishlist{Title: "New Wishlist"}
	mockRepo.On("Create", wishlist).Return(nil)

	err := service.CreateWishlist(wishlist)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateWishlistError(t *testing.T) {
	mockRepo := new(MockWishlistRepository)
	service := NewWishlistService(mockRepo)

	wishlist := &entities.Wishlist{Title: "New Wishlist"}
	mockRepo.On("Create", wishlist).Return(errors.New("failed to create wishlist"))

	err := service.CreateWishlist(wishlist) // Memanggil metode CreateWishlist dengan wishlist yang sudah didefinisikan

	assert.NotNil(t, err)
	assert.Equal(t, "failed to create wishlist", err.Error())
	mockRepo.AssertExpectations(t)
}
