package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/todos", GetTodos)
	r.POST("/todos", CreateTodo)
	r.GET("/todos/:id", GetTodoByID)
	r.PUT("/todos/:id", UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)
	return r
}

func TestGetTodos(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateTodo(t *testing.T) {
	router := setupRouter()

	newTodo := map[string]interface{}{
		"title": "Test Todo",
		"done":  false,
	}

	jsonValue, _ := json.Marshal(newTodo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestGetTodoByID(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/todos/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateTodo(t *testing.T) {
	router := setupRouter()

	updatedTodo := map[string]interface{}{
		"title": "Updated Todo",
		"done":  true,
	}

	jsonValue, _ := json.Marshal(updatedTodo)
	req, _ := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteTodo(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/todos/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
