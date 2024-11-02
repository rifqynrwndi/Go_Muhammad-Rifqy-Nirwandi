package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestBaseResponse(t *testing.T) {

	response := BaseResponse[string]{
		Status:  "success",
		Message: "Data retrieved successfully",
		Data:    "Test Data",
	}

	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "Data retrieved successfully", response.Message)
	assert.Equal(t, "Test Data", response.Data)
}

func TestTodoInput(t *testing.T) {

	todoInput := TodoInput{
		Title:       "Test Title",
		Description: "Test Description",
	}

	assert.Equal(t, "Test Title", todoInput.Title)
	assert.Equal(t, "Test Description", todoInput.Description)
}

func TestTodo(t *testing.T) {

	now := time.Now()

	todo := Todo{
		ID:          1,
		Title:       "Sample Todo",
		Description: "This is a sample todo item.",
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   gorm.DeletedAt{Time: now, Valid: true},
	}

	assert.Equal(t, uint(1), todo.ID)
	assert.Equal(t, "Sample Todo", todo.Title)
	assert.Equal(t, "This is a sample todo item.", todo.Description)
	assert.Equal(t, now, todo.CreatedAt)
	assert.Equal(t, now, todo.UpdatedAt)
	assert.True(t, todo.DeletedAt.Valid)
	assert.Equal(t, now, todo.DeletedAt.Time)
}
