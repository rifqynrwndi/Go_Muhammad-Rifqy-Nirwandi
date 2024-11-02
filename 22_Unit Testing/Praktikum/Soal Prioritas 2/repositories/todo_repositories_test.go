package repositories

import (
	_ "errors"
	"go-todos-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) GetAll() ([]models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetByID(id string) (models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoRepository) Create(todoInput models.TodoInput) (models.Todo, error) {
	args := m.Called(todoInput)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoRepository) Update(todoInput models.TodoInput, id string) (models.Todo, error) {
	args := m.Called(todoInput, id)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	mockData := []models.Todo{
		{ID: 1, Title: "Test Todo", Description: "Test Description"},
	}
	mockRepo.On("GetAll").Return(mockData, nil)
	todos, err := mockRepo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, todos, 1)
	assert.Equal(t, "Test Todo", todos[0].Title)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	mockData := models.Todo{ID: 1, Title: "Sample Todo", Description: "Sample Description"}
	mockRepo.On("GetByID", "1").Return(mockData, nil)
	todo, err := mockRepo.GetByID("1")
	assert.NoError(t, err)
	assert.Equal(t, "Sample Todo", todo.Title)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoInput := models.TodoInput{Title: "New Todo", Description: "New Description"}
	mockData := models.Todo{ID: 1, Title: "New Todo", Description: "New Description"}
	mockRepo.On("Create", todoInput).Return(mockData, nil)
	todo, err := mockRepo.Create(todoInput)
	assert.NoError(t, err)
	assert.Equal(t, "New Todo", todo.Title)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	todoInput := models.TodoInput{Title: "Updated Title", Description: "Updated Description"}
	mockData := models.Todo{ID: 1, Title: "Updated Title", Description: "Updated Description"}
	mockRepo.On("Update", todoInput, "1").Return(mockData, nil)
	todo, err := mockRepo.Update(todoInput, "1")
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", todo.Title)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	mockRepo.On("Delete", "1").Return(nil)
	err := mockRepo.Delete("1")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
