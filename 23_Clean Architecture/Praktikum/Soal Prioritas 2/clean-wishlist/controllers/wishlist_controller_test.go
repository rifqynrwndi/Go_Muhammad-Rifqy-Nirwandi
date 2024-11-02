package controllers

import (
	"errors"
	"go-wishlist-api/entities"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Service
type MockWishlistService struct {
	mock.Mock
}

func (m *MockWishlistService) GetAllWishlists() ([]entities.Wishlist, error) {
	args := m.Called()
	return args.Get(0).([]entities.Wishlist), args.Error(1)
}

func (m *MockWishlistService) CreateWishlist(wishlist *entities.Wishlist) error {
	args := m.Called(wishlist)
	return args.Error(0)
}

// Unit tests
func TestGetAllWishlistsController(t *testing.T) {
	mockService := new(MockWishlistService)
	controller := NewWishlistController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/wishlists", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.On("GetAllWishlists").Return([]entities.Wishlist{{ID: 1, Title: "Test Wishlist"}}, nil)

	err := controller.GetAll(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Test Wishlist")
	mockService.AssertExpectations(t)
}

func TestCreateWishlistController(t *testing.T) {
	mockService := new(MockWishlistService)
	controller := NewWishlistController(mockService)

	e := echo.New()
	wishlistJSON := `{"title":"New Wishlist"}`
	req := httptest.NewRequest(http.MethodPost, "/wishlists", strings.NewReader(wishlistJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.On("CreateWishlist", mock.AnythingOfType("*entities.Wishlist")).Return(nil)

	err := controller.Create(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "New Wishlist")
	mockService.AssertExpectations(t)
}

func TestCreateWishlistControllerError(t *testing.T) {
	mockService := new(MockWishlistService)
	controller := NewWishlistController(mockService)

	e := echo.New()
	wishlistJSON := `{"title":"New Wishlist"}`
	req := httptest.NewRequest(http.MethodPost, "/wishlists", strings.NewReader(wishlistJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.On("CreateWishlist", mock.AnythingOfType("*entities.Wishlist")).Return(errors.New("creation failed"))

	err := controller.Create(c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "creation failed")
	mockService.AssertExpectations(t)
}
